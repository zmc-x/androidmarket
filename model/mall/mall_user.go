package mall

type User struct {
	Uid       string `json:"uid"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Privilege int    `json:"privilege"`
}
