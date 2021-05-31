package entity

type UserDetail struct {
	ID        uint32 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
	Location  string `json:"location"`
	UserID    uint32 `json:"user_id"`
}
