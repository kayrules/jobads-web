package model

// User struct
type User struct {
	// ID       bson.ObjectId `json:"id,omitempty" form:"id,omitempty"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
