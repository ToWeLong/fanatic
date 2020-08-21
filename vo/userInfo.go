package vo

type UserInfo struct {
	ID uint `json:"id"`
	Account string `json:"account"`
	Permission []Permission `json:"permission"`
}

type Permission struct {
	ID uint `json:"id"`
	Name string `json:"name"`
}

