package structs

type StatusTask struct {
	ID        int64  `json:"id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}