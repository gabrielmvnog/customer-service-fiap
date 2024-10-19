package dtos

type FindCustomerResponse struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	DocumentNumber string `json:"document_number"`
}
