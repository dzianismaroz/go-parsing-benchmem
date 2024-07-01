package hw3

//easyjson:json
type User struct {
	Browsers []string `json:"browsers"`
	Email    string   `json:"email"`
	Name     string   `json:"name"`
	// no need to use other fields during parsing unique browsers by users
}
