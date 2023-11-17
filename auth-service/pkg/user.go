package pkg

type user struct{
	ID int `json:d`
	Username string `json:username`
	Password string `json:password`
}