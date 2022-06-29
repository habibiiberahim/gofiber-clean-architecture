package schemas

type SchemaDatabaseError struct {
	Type string
	Code int
}

type SchemaErrorResponse struct {
	Status string      `json:"statusCode"`
	Code   int         `json:"code"`
	Error  interface{} `json:"error"`
}

type SchemaUnauthorizedError struct {
	Status  string `json:"statusCode"`
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
}
