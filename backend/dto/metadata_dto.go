package dto

type Metadata struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

type Response struct {
	Metadata Metadata    `json:"metadata"`
	Data     interface{} `json:"data"`
}
