package errno

var (
	OK                  = &Errno{Code: 0, Message: "Ok"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	RouterNotFound      = &Errno{Code: 10005, Message: "router not found"}
)
