package custom_error

const errorCodeBase = 10000

const (
	UnknownError         uint64 = errorCodeBase + 1
	InvalidJSONString    uint64 = errorCodeBase + 2
	InputValidationError uint64 = errorCodeBase + 3
	Unauthorized         uint64 = errorCodeBase + 4
)
