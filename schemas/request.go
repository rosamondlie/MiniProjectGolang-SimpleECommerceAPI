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

type CreateProductRequest struct {
	Nama   string `form:"nama" binding:"required"`
	Harga  int    `form:"harga" binding:"required,min=0"`
	Stok   *int    `form:"stok" binding:"min=0"`
	UserID uint   `form:"user_id" binding:"required"`
}

type LoginRequest struct {
    Login    string `json:"login" binding:"required"`   // email / no hp
    Password string `json:"password" binding:"required"`
}