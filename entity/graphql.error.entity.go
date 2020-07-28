package entity

type GrapqlError struct {
	Data   interface{}        `json:"data"`
	Errors []GraphqlErrorList `json:"errors"`
}

type GraphqlErrorList struct {
	Message  string        `json:"message"`
	Location []interface{} `json:"locations"`
}
