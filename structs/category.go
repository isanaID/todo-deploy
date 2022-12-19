package structs

type Category struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	UserId    int64  `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}