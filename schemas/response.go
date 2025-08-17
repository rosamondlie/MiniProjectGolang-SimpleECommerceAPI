package schemas

type ListUserResponse struct{
	ID	   uint   `json:"id"`
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	NoHP   *string `json:"no_hp"`
	Status  string   `json:"status"`
}

type ListProductResponse struct {
	ID     int     `json:"id"`
	Nama   string  `json:"nama"`
	Harga  float64 `json:"harga"`
	Stok   int     `json:"stok"`
	UserID int     `json:"user_id"`
	Photo  string  `json:"photo"`
}

type DashboardResponse struct {
	TotalUsers int64 `json:"total_users"`
	ActiveUsers int64 `json:"active_users"`
	TotalProducts int64 `json:"total_products"`
	TotalAvailProducts int64 `json:"total_available_products"`
	LatestProducts []ListLatestProductResponse `json:"latest_products"`
}

type ListLatestProductResponse struct {
	ID    int     `json:"id"`
	Photo string  `json:"photo"`
	Nama  string  `json:"nama"`
	Date string `json:"date"`
	Harga int `json:"harga"`
}

type ListProductLandingResponse struct {
	Photo string  `json:"photo"`
	Nama  string  `json:"nama"`
	Harga int `json:"harga"`
}