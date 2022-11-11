package types

type Info struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Data struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

type Page struct {
	Page      int64       `json:"page"`
	Limit     int64       `json:"limit"`
	Count     int64       `json:"count"`
	PageCount int64       `json:"page_count"`
	Data      interface{} `json:"data"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
}
