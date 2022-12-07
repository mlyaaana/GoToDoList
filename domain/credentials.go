package domain

type Credentials struct {
	UserId   string
	Username string
	Password string
}

func NewCredentials(username, password, userid string) *Credentials {
	return &Credentials{
		UserId:   userid,
		Username: username,
		Password: password,
	}
}
