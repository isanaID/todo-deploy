package structs

type Task struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	User_id     int64  `json:"user_id"`
	Category_id int64  `json:"category_id"`
	Status_id   int64  `json:"status_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}