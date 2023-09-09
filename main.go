package main

import (
	"context"
	api_input_reader "data-platform-api-production-order-conf-creates-subfunc/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-production-order-conf-creates-subfunc/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-production-order-conf-creates-subfunc/API_Processing_Data_Formatter"
	"data-platform-api-production-order-conf-creates-subfunc/config"
	"data-platform-api-production-order-conf-creates-subfunc/subfunction"
	"fmt"
	"time"

	database "github.com/latonaio/golang-mysql-network-connector"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

func main() {
	ctx := context.Background()
	l := logger.NewLogger()
	c := config.NewConf()
	db, err := database.NewMySQL(c.DB)
	if err != nil {
		l.Error(err)
		return
	}
	defer db.Close()

	rmq, err := rabbitmq.NewRabbitmqClient(c.RMQ.URL(), c.RMQ.QueueFrom(), c.RMQ.SessionControlQueue(), c.RMQ.QueueTo(), -1)
	if err != nil {
		l.Fatal(err.Error())
	}
	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()
	for msg := range iter {
		start := time.Now()
		msg.Success()
		sdc, err := callProcess(ctx, db, rmq, msg, c)
		sdc.SubfuncResult = getBoolPtr(err == nil)
		if err != nil {
			sdc.SubfuncError = err.Error()
			l.Error(err)
		}
		l.JsonParseOut(sdc)
		err = msg.Respond(sdc)
		if err != nil {
			l.Error(err)
		}
		l.Info("process time %v\n", time.Since(start).Milliseconds())
	}
}

func getSessionID(data map[string]interface{}) string {
	id := fmt.Sprintf("%v", data["runtime_session_id"])
	return id
}

func callProcess(ctx context.Context, db *database.Mysql, rmq *rabbitmq.RabbitmqClient, msg rabbitmq.RabbitmqMessage, conf *config.Conf) (dpfm_api_output_formatter.SDC, error) {
	var err error
	l := logger.NewLogger()
	l.AddHeaderInfo(map[string]interface{}{"runtime_session_id": getSessionID(msg.Data())})

	subfunc := subfunction.NewSubFunction(ctx, db, conf, rmq, l)
	sdc := api_input_reader.ConvertToSDC(msg.Raw())
	psdc := api_processing_data_formatter.ConvertToSDC()
	osdc := dpfm_api_output_formatter.ConvertToSDC(msg.Raw())

	err = subfunc.CreateSdc(&sdc, &psdc, &osdc)
	if err != nil {
		return osdc, err
	}

	return osdc, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}
