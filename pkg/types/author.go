package types

type Author struct {
	ID         int64  `json:"id,omitempty"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}
