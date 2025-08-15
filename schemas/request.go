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
	Nama   string `json:"nama" binding:"required"`
	Harga  int    `json:"harga" binding:"required,min=0,numeric"`
	Stok   int    `json:"stok" binding:"required,min=0,numeric"`
	UserID uint   `json:"user_id" binding:"required"`
	Photo  *string `json:"photo"`
}