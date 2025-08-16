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