package requests

type PostRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Likes  int    `json:"likes"`
	Draft  bool   `json:"draft"`
	Author string `json:"author"`
	// UserID uint `gorm:"user_id"`
}
