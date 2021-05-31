package models

type User struct {
	Id       string `json:"id"`
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Cpf      string `validate:"required" json:"cpf"`
	Address  string `validate:"required" json:"address"`
	Card     Card   `validate:"dive" json:"card"`
	Role     string `validate:"required" json:"role"`
	Password string `validate:"required" json:"password"`
	Token    string `json:"token"`
}

type LoginRequest struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

// string name = 1;
// string email = 2;
// string password = 3;
// string cpf = 4;
// string endereco = 5;
// string cartao = 6;
// Role role = 7;
