package domain

type Otp struct {
	Value string `json:"value"`
	Sub   string `json:"sub"`
	Name  string `json:"name"`
}
