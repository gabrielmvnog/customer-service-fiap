package dtos

type CreateCustomerRequest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	DocumentNumber string `json:"document_number"`
}
