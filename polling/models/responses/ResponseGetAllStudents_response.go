package responses

type ResponseLinkStudents struct {
	Self string `json:"self"`
}

type ResponseDataStudents struct {
	Type       string `json:"type"`
	Id         string  `json:"id"`
	Attributes struct {
		Name        string `json:"name"`
		Age         string  `json:"age"`
		PhoneNumber string `json:"phoneNumber"`
	} `json:"attributes"`
}

type ResponseGetAllStudents struct {
	Links ResponseLinkStudents   `json:"links"`
	Data  []ResponseDataStudents `json:"data"`
}