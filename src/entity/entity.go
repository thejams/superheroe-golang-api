package entity

// Message is the health response struct
type Message struct {
	MSG string `json:"msg"`
}

// TradeMade struct
type TradeMade struct {
	Endpoint       string                   `json:"endpoint"`
	Quotes         []map[string]interface{} `json:"quotes"`
	Requested_time string                   `json:"requested_time"`
	Timestamp      int32                    `json:"timestamp"`
}
