package types

// chaincode invoke request
type InvokeRequest struct {
	ChaincodeID string   `json:"chaincodeID" yaml:"chaincodeID"`
	Fcn         string   `json:"fcn" yaml:"fcn"`
	Args        []string `json:"args" yaml:"args"`
	NeedSubmit  bool     `json:"needSubmit" yaml:"needSubmit"`
	// only for admin client
	Endpoints []string `json:"endpoints" yaml:"endpoints"`
}

// chaincode invoke response
type InvokeResponse struct {
	Payload         []byte           `json:"payload" yaml:"payload"`
	TransactionInfo *TransactionInfo `json:"transactionInfo" yaml:"transactionInfo"`
	ChaincodeStatus int32            `json:"chaincodeStatus" yaml:"chaincodeStatus"`
}
