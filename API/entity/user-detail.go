package entity

type UserDetail struct {
	UserID    uint32 `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
	Location  string `json:"location"`
}
