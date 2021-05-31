package models

type Card struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Number       string `json:"number"`
	SecurityCode string `json:"securityCode"`
	ValidThru    string `json:"validThru"`
}
