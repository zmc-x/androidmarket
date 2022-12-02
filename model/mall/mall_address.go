package mall

type Address struct {
	Id   int    `json:"id"`
	Uid  string `json:"uid"`
	Name string `json:"name"`
	//Location       string `json:"location"`
	Province       string `json:"province"`
	City           string `json:"city"`
	County         string `json:"county"`
	Detaillocation string `json:"detaillocation"`
	Tel            int    `json:"tel"`
	Defaultaddress int    `json:"defaultaddress"`
}
