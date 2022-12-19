package structs

type StatusTask struct {
	ID        int64  `json:"id"`
	Status    string `json:"status"`
	UserId    int64  `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}