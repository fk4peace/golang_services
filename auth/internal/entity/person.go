package entity

type Person struct {
	Id       *int64  `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Role     *Role   `json:"role,omitempty"`
}
