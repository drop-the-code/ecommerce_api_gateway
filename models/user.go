package models

type User struct {
	Id       string
	Name     string
	Email    string
	Cpf      string
	Address  string
	Card     Card
	Role     string
	Password string
	Token    string
}

type LoginRequest struct {
	Email    string
	Password string
}

// string name = 1;
// string email = 2;
// string password = 3;
// string cpf = 4;
// string endereco = 5;
// string cartao = 6;
// Role role = 7;
