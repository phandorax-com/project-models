package email

type Otp struct {
	Value    string `json:"value"`
	Sub      string `json:"sub"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Template string `json:"template"`
}
