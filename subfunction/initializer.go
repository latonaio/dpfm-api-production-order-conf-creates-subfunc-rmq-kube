package subfunction

import (
	"context"
	api_input_reader "data-platform-api-production-order-confirmation-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-production-order-confirmation-headers-creates-subfunc-rmq-kube/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-production-order-confirmation-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"data-platform-api-production-order-confirmation-headers-creates-subfunc-rmq-kube/config"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type SubFunction struct {
	ctx  context.Context
	db   *database.Mysql
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient
	l    *logger.Logger
}

func NewSubFunction(ctx context.Context, db *database.Mysql, conf *config.Conf, rmq *rabbitmq.RabbitmqClient, l *logger.Logger) *SubFunction {
	return &SubFunction{
		ctx:  ctx,
		db:   db,
		conf: conf,
		rmq:  rmq,
		l:    l,
	}
}

func (f *SubFunction) MetaData(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.MetaData {
	metaData := psdc.ConvertToMetaData(sdc)

	return metaData
}

func (f *SubFunction) ProcessType(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.ProcessType, error) {
	processType := psdc.ConvertToProcessType()

	if isBulkProcess(sdc, processType) {
		processType.BulkProcess = true
	}

	// if isIndividualProcess(sdc, processType, referenceType) {
	// 	processType.IndividualProcess = true
	// }

	//if !processType.BulkProcess && !processType.IndividualProcess {
	//	return nil, xerrors.Errorf("一括登録または個別登録に必要な入力パラメータが揃っていません。")
	//}

	return processType, nil
}

func isBulkProcess(
	sdc *api_input_reader.SDC,
	processType *api_processing_data_formatter.ProcessType,
) bool {
	inputParameters := sdc.InputParameters

	if inputParameters.MRPArea != nil && inputParameters.OwnerProductionPlantBusinessPartner != nil &&
		inputParameters.OwnerProductionPlant != nil && inputParameters.ProductInHeader != nil &&
		inputParameters.ProductInItem != nil && inputParameters.ProductionPlantBusinessPartner != nil &&
		inputParameters.ProductionPlant != nil && inputParameters.ProductionPlantStorageLocation != nil {
		if (*inputParameters.MRPArea)[0] != nil && (*inputParameters.OwnerProductionPlantBusinessPartner)[0] != nil &&
			(*inputParameters.OwnerProductionPlant)[0] != nil && (*inputParameters.ProductInHeader)[0] != nil &&
			(*inputParameters.ProductInItem)[0] != nil && (*inputParameters.ProductionPlantBusinessPartner)[0] != nil &&
			(*inputParameters.ProductionPlant)[0] != nil && (*inputParameters.ProductionPlantStorageLocation)[0] != nil {
			if len(*(*inputParameters.MRPArea)[0]) != 0 && len(*(*inputParameters.OwnerProductionPlant)[0]) != 0 &&
				len(*(*inputParameters.ProductInHeader)[0]) != 0 && len(*(*inputParameters.ProductInItem)[0]) != 0 &&
				len(*(*inputParameters.ProductionPlant)[0]) != 0 && len(*(*inputParameters.ProductionPlantStorageLocation)[0]) != 0 {
				processType.PluralitySpec = true
				return true
			}
		}
	}
	if inputParameters.MRPAreaTo != nil && inputParameters.MRPAreaFrom != nil &&
		inputParameters.OwnerProductionPlantBusinessPartnerTo != nil && inputParameters.OwnerProductionPlantBusinessPartnerFrom != nil &&
		inputParameters.OwnerProductionPlantTo != nil && inputParameters.OwnerProductionPlantFrom != nil &&
		inputParameters.ProductInHeaderTo != nil && inputParameters.ProductInHeaderFrom != nil &&
		inputParameters.ProductInItemTo != nil && inputParameters.ProductInItemFrom != nil &&
		inputParameters.ProductionPlantBusinessPartnerTo != nil && inputParameters.ProductionPlantBusinessPartnerFrom != nil &&
		inputParameters.ProductionPlantTo != nil && inputParameters.ProductionPlantFrom != nil &&
		inputParameters.ProductionPlantStorageLocationTo != nil && inputParameters.ProductionPlantStorageLocationFrom != nil {
		if len(*inputParameters.MRPAreaTo) != 0 && len(*inputParameters.MRPAreaFrom) != 0 &&
			len(*inputParameters.OwnerProductionPlantTo) != 0 && len(*inputParameters.OwnerProductionPlantFrom) != 0 &&
			len(*inputParameters.ProductInHeaderTo) != 0 && len(*inputParameters.ProductInHeaderFrom) != 0 &&
			len(*inputParameters.ProductInItemTo) != 0 && len(*inputParameters.ProductInItemFrom) != 0 &&
			len(*inputParameters.ProductionPlantTo) != 0 && len(*inputParameters.ProductionPlantFrom) != 0 &&
			len(*inputParameters.ProductionPlantStorageLocationTo) != 0 && len(*inputParameters.ProductionPlantStorageLocationFrom) != 0 {
			processType.RangeSpec = true
			return true
		}
	}

	return false
}

func (f *SubFunction) CreateSdc(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	//var e error

	wg := sync.WaitGroup{}

	psdc.MetaData = f.MetaData(sdc, psdc)
	//psdc.ProcessType, err = f.ProcessType(sdc, psdc)
	if err != nil {
		return err
	}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	f.l.Info(psdc)
	err = f.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}
