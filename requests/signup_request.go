package requests

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `gorm:"unique"`
	Password string `json:"password"`
}
