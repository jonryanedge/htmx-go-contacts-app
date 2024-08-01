package data

type Contact struct {
	ID     string            `json:"id"`
	First  string            `json:"first"`
	Last   string            `json:"last"`
	Phone  string            `json:"phone"`
	Email  string            `json:"email"`
	Errors map[string]string `json:"errors"`
}
