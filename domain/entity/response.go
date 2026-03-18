package entity

type UserResponse struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email,omitempty"`
	Country   string `json:"country,omitempty"`
	CreatedAt string  `json:"created_at,omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty"`
}

type AddUserResponse struct {
	ID    string `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

type UpdateUserResponse struct {
	User  UserResponse `json:"user,omitempty"`
	Error string       `json:"error,omitempty"`
}

type DeleteUserResponse struct {
	Error string `json:"error,omitempty"`
}

type GetPaginatedUsersResponse struct {
	Users []UserResponse `json:"users,omitempty"`
	Error string         `json:"error,omitempty"`
}
