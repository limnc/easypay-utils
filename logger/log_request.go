package logger

type LogRequest struct {
	Service      string `json:"service"`
	LogLevel     string `json:"logLevel"`
	LogMessage   string `json:"logMessage"`
	RequestID    string `json:"requestId,omitempty"`
	HTTPStatus   string `json:"httpStatus,omitempty"`
	RequestBody  string `json:"requestBody,omitempty"`
	ResponseBody string `json:"responseBody,omitempty"`
	StackTrace   string `json:"stackTrace,omitempty"`
	CreatedBy    string `json:"createdBy,omitempty"`
	CreatedAt    string `json:"createdAt,omitempty"`
}
