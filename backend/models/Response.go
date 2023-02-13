package models

type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Detail  map[string]interface{} `json:"detail"`
}
