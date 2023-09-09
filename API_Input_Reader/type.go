package api_input_reader

type SDC struct {
	ConnectionKey     string          `json:"connection_key"`
	Result            bool            `json:"result"`
	RedisKey          string          `json:"redis_key"`
	Filepath          string          `json:"filepath"`
	APIStatusCode     int             `json:"api_status_code"`
	RuntimeSessionID  string          `json:"runtime_session_id"`
	BusinessPartnerID *int            `json:"business_partner"`
	ServiceLabel      string          `json:"service_label"`
	APIType           string          `json:"APIType"`
	InputParameters   InputParameters `json:"InputParameters"`
	Header            Header          `json:"ProductionOrderConfirmation"`
	APISchema         string          `json:"api_schema"`
	Accepter          []string        `json:"accepter"`
	Deleted           bool            `json:"deleted"`
}

type InputParameters struct {
	MRPArea                                 *[]*string `json:"MRPArea"`
	MRPAreaTo                               *string    `json:"MRPAreaTo"`
	MRPAreaFrom                             *string    `json:"MRPAreaFrom"`
	OwnerProductionPlantBusinessPartner     *[]*int    `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlantBusinessPartnerTo   *int       `json:"OwnerProductionPlantBusinessPartnerTo"`
	OwnerProductionPlantBusinessPartnerFrom *int       `json:"OwnerProductionPlantBusinessPartnerFrom"`
	OwnerProductionPlant                    *[]*string `json:"OwnerProductionPlant"`
	OwnerProductionPlantTo                  *string    `json:"OwnerProductionPlantTo"`
	OwnerProductionPlantFrom                *string    `json:"OwnerProductionPlantFrom"`
	ProductInHeader                         *[]*string `json:"ProductInHeader"`
	ProductInHeaderTo                       *string    `json:"ProductInHeaderTo"`
	ProductInHeaderFrom                     *string    `json:"ProductInHeaderFrom"`
	ProductInItem                           *[]*string `json:"ProductInItem"`
	ProductInItemTo                         *string    `json:"ProductInItemTo"`
	ProductInItemFrom                       *string    `json:"ProductInItemFrom"`
	ProductionPlantBusinessPartner          *[]*int    `json:"ProductionPlantBusinessPartner"`
	ProductionPlantBusinessPartnerTo        *int       `json:"ProductionPlantBusinessPartnerTo"`
	ProductionPlantBusinessPartnerFrom      *int       `json:"ProductionPlantBusinessPartnerFrom"`
	ProductionPlant                         *[]*string `json:"ProductionPlant"`
	ProductionPlantTo                       *string    `json:"ProductionPlantTo"`
	ProductionPlantFrom                     *string    `json:"ProductionPlantFrom"`
	ProductionPlantStorageLocation          *[]*string `json:"ProductionPlantStorageLocation"`
	ProductionPlantStorageLocationTo        *string    `json:"ProductionPlantStorageLocationTo"`
	ProductionPlantStorageLocationFrom      *string    `json:"ProductionPlantStorageLocationFrom"`
	ExecuteProductAvailabilityCheck         *bool      `json:"ExecuteProductAvailabilityCheck"`
	ReleaseProductionOrder                  *bool      `json:"ReleaseProductionOrder"`
}

type Header struct {
	ProductionOrder                          int      `json:"ProductionOrder"`
	ProductionOrderItem                      int      `json:"ProductionOrderItem"`
	Operations                               int      `json:"Operations"`
	OperationsItem                           int      `json:"OperationsItem"`
	OperationID                              int      `json:"OperationID"`
	ConfirmationCountingID                   int      `json:"ConfirmationCountingID"`
	OperationPlannedQuantityInBaseUnit       float32  `json:"OperationPlannedQuantityInBaseUnit"`
	OperationPlannedQuantityInProductionUnit float32  `json:"OperationPlannedQuantityInProductionUnit"`
	OperationPlannedQuantityInOperationUnit  float32  `json:"OperationPlannedQuantityInOperationUnit"`
	ProductBaseUnit                          string   `json:"ProductBaseUnit"`
	ProductProductionUnit                    string   `json:"ProductProductionUnit"`
	ProductOperationUnit                     string   `json:"ProductOperationUnit"`
	OperationPlannedScrapInPercent           *float32 `json:"OperationPlannedScrapInPercent"`
	ConfirmationEntryDate                    *string  `json:"ConfirmationEntryDate"`
	ConfirmationEntryTime                    *string  `json:"ConfirmationEntryTime"`
	ConfirmationText                         *string  `json:"ConfirmationText"`
	IsFinalConfirmation                      *string  `json:"IsFinalConfirmation"`
	WorkCenter                               int      `json:"WorkCenter"`
	EmployeeIDWhoConfirmed                   int      `json:"EmployeeIDWhoConfirmed"`
	ConfirmedExecutionStartDate              *string  `json:"ConfirmedExecutionStartDate"`
	ConfirmedExecutionStartTime              *string  `json:"ConfirmedExecutionStartTime"`
	ConfirmedSetupStartDate                  *string  `json:"ConfirmedSetupStartDate"`
	ConfirmedSetupStartTime                  *string  `json:"ConfirmedSetupStartTime"`
	ConfirmedProcessingStartDate             *string  `json:"ConfirmedProcessingStartDate"`
	ConfirmedProcessingStartTime             *string  `json:"ConfirmedProcessingStartTime"`
	ConfirmedExecutionEndDate                *string  `json:"ConfirmedExecutionEndDate"`
	ConfirmedExecutionEndTime                *string  `json:"ConfirmedExecutionEndTime"`
	ConfirmedSetupEndDate                    *string  `json:"ConfirmedSetupEndDate"`
	ConfirmedSetupEndTime                    *string  `json:"ConfirmedSetupEndTime"`
	ConfirmedProcessingEndDate               *string  `json:"ConfirmedProcessingEndDate"`
	ConfirmedProcessingEndTime               *string  `json:"ConfirmedProcessingEndTime"`
	ConfirmedWaitDuration                    *float32 `json:"ConfirmedWaitDuration"`
	WaitDurationUnit                         *string  `json:"WaitDurationUnit"`
	ConfirmedQueueDuration                   *float32 `json:"ConfirmedQueueDuration"`
	QueueDurationUnit                        *string  `json:"QueueDurationUnit"`
	ConfirmedMoveDuration                    *float32 `json:"ConfirmedMoveDuration"`
	MoveDurationUnit                         *string  `json:"MoveDurationUnit"`
	ConfirmedYieldQuantity                   *float32 `json:"ConfirmedYieldQuantity"`
	ConfirmedScrapQuantity                   *float32 `json:"ConfirmedScrapQuantity"`
	OperationVarianceReason                  *string  `json:"OperationVarianceReason"`
	CreationDate                             string   `json:"CreationDate"`
	CreationTime                             string   `json:"CreationTime"`
	LastChangeDate                           string   `json:"LastChangeDate"`
	LastChangeTime                           string   `json:"LastChangeTime"`
	IsCancelled                              *bool    `json:"IsCancelled"`
}
