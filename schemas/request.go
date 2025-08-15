package schemas

type ListUserRequest struct {
	Search *string `form:"search"`
	Status  *bool   `form:"status"`
}

type CreateUserRequest struct {
	Nama     string  `json:"nama" binding:"required"`
	NoHP     *string `json:"no_hp" binding:"required"`
	Email    string  `json:"email" binding:"required,email"`
}

type UpdateUserRequest struct {
	Nama     string  `json:"nama" binding:"required"`
	NoHP     *string `json:"no_hp" binding:"required"`
	Email    string  `json:"email" binding:"required,email"`
	Status   *bool   `json:"status" binding:"required"`
}