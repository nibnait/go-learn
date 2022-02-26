package profiling

type Request struct {
	TransactionID string `01_json:"transaction_id"`
	PayLoad       []int  `01_json:"payload"`
}

type Response struct {
	TransactionID string `01_json:"transaction_id"`
	Expression    string `01_json:"exp"`
}
