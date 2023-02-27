package data

type UserProfile struct {
	UserId   int64
	UserName string
}

type RequestLog struct {
	UserId   int64
	Requests []RequestId
}
