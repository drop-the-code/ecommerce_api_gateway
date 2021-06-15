package models

type Cart struct {
	Id          string   `json:"id"`
	UpdatedAt   string   `json:"updateAt"`
	ClientId    string   `json:"client_id"`
	Status      bool     `json:"status"`
	ProductList []string `json:"product_list"`
}
