package schemas

type ListUserResponse struct{
	ID	   uint   `json:"id"`
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	NoHP   *string `json:"no_hp"`
	Status  string   `json:"status"`
}