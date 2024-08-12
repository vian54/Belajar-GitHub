package models

type DbSearchObject struct {
	Mode         string                 `json:"-"`
	Limit        int64                  `json:"limit"`
	Page         int64                  `json:"page"`
	Order        []string               `json:"orders"`
	PayloadData  map[string]interface{} `json:"data"`
	ResponseData interface{}            `json:"response_data"`
	TotalData    int64                  `json:"total"`
}
