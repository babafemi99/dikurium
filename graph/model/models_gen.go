// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type SignInResult struct {
	Success   bool    `json:"success"`
	Message   string  `json:"message"`
	User      *User   `json:"user,omitempty"`
	AuthToken *string `json:"authToken,omitempty"`
}

type SignUpResult struct {
	Success   bool    `json:"success"`
	Message   string  `json:"message"`
	User      *User   `json:"user,omitempty"`
	AuthToken *string `json:"authToken,omitempty"`
}

type Todo struct {
	ID    string `json:"id"`
	Text  string `json:"text"`
	Done  bool   `json:"done"`
	Email string `json:"Email"`
	User  *User  `json:"user" gorm:"embedded;foreignKey:userid"`
}

type User struct {
	Userid   string `json:"userid" gorm:"primarykey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
