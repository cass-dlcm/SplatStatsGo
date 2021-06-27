package site_objects

type LoginAuth struct {
	Auth AuthInfo
}

type AuthInfo struct {
	Authenticated bool
	UserId        int64
	Username      string
}
