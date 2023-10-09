package constant

type ResponseStatus int
type Headers int
type General int

const (
	Success ResponseStatus = iota + 1
	DataNotFound
	UnknownError
	InvalidRequest
	Unauthorized
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "INVALID_REQUEST", "UNKNOWN_ERROR", "UNAUTHORIZED"}[r-1]
}
func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data Not Found", "Invalid Request", "Unknown Error", "Unauthorized"}[r-1]
}
