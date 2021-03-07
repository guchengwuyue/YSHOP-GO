package vo

type ResultList struct {
	Content interface{} `json:"content"`
	TotalElements int `json:"totalElements"`
}
