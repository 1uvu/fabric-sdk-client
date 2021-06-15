package types

// chaincode invoke request
type InvokeRequest struct {
	ChaincodeID string
	Fcn         string
	Args        []string
	NeedSubmit  bool
	// for admin client
	Endpoints []string
}

// chaincode invoke response
type InvokeResponse struct {
	Payload         []byte
	TransactionInfo *TransactionInfo
	ChaincodeStatus int32
}
