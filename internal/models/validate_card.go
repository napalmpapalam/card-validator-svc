package models

import "github.com/napalmpapalam/card-validator-svc/internal/problems"

type ValidateCardRequest struct {
	CardNumber      string `json:"card_number"`
	ExpirationMonth string `json:"expiration_month"`
	ExpirationYear  string `json:"expiration_year"`
}

type ValidateCardResponse struct {
	Valid bool            `json:"valid"`
	Error *problems.Error `json:"error,omitempty"`
}
