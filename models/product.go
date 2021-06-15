package models

type Product struct {
	Id           int32   `json:"id"`
	Name         string  `json:"name" validate:"required"`
	Price        float64 `json:"price" validate:"required"`
	ProviderCnpj string  `json:"provider_cnpj" validate:"required"`
	Description  string  `json:"description"`
}

// message Product {
// 	int32 productID = 1;
// 	string productName = 2;
// 	double productPrice = 3;
// 	string providerCNPJ = 4;
// 	string description = 5;
//   }
