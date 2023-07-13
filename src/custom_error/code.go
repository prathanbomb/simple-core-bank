package custom_error

const errorCodeBase = 10000

const (
	UnknownError         uint64 = errorCodeBase + 1
	InvalidJSONString    uint64 = errorCodeBase + 2
	InputValidationError uint64 = errorCodeBase + 3
	Unauthorized         uint64 = errorCodeBase + 4
	DBError              uint64 = errorCodeBase + 5
	AccountNoNotFound    uint64 = errorCodeBase + 6
	BalanceNotEnough     uint64 = errorCodeBase + 7
)
