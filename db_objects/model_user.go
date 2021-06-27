package db_objects

type User struct {
	Pk            int64
	Username      string
	Password      string
	Email         string
	EmailVerified bool
	Locale        string
}

type SessionToken struct {
	Pk           int64
	Parent       int64
	SessionToken string
}
