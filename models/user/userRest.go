package user

// Defines the structure of our user in our DB
type UserRest struct {
	ID        string `json:"_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
