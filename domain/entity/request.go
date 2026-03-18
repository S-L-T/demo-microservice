package entity

type UserRequest struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email,omitempty"`
	Country   string `json:"country,omitempty"`
}

type AddUserRequest struct {
	User UserRequest `json:"user,omitempty"`
}

type UpdateUserRequest struct {
	ID   string      `json:"id,omitempty"`
	User UserRequest `json:"user,omitempty"`
}

type DeleteUserRequest struct {
	ID string `json:"id,omitempty"`
}

type FilterRequest struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	Email     string `json:"email,omitempty"`
	Country   string `json:"country,omitempty"`
}

type GetPaginatedUsersRequest struct {
	Filter   FilterRequest `json:"filter,omitempty"`
	PageNum  uint64        `json:"page_num,omitempty"`
	PageSize uint64        `json:"page_size,omitempty"`
}
