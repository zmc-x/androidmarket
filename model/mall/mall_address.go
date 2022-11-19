package mall

type Address struct {
	Id             int    `json:"id"`
	Uid            string `json:"uid"`
	Name           string `json:"name"`
	Location       string `json:"location"`
	Tel            int    `json:"tel"`
	Defaultaddress int    `json:"defaultaddress"`
}
