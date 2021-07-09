package types

// chaincode invoke request
type InvokeRequest struct {
	ChaincodeID string   `json:"chaincodeID"`
	Fcn         string   `json:"fcn"`
	Args        []string `json:"args"`
	NeedSubmit  bool     `json:"needSubmit"`
	// only for admin client
	Endpoints []string `json:"endpoints"`
}

// chaincode invoke response
type InvokeResponse struct {
	Payload         []byte           `json:"payload"`
	TransactionInfo *TransactionInfo `json:"transactionInfo"`
	ChaincodeStatus int32            `json:"chaincodeStatus"`
}
