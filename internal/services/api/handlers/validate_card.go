package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/napalmpapalam/card-validator-svc/internal/models"
	"github.com/napalmpapalam/card-validator-svc/internal/problems"
	"github.com/pkg/errors"
)

func newValidateCardRequest(r *http.Request) (*models.ValidateCardRequest, error) {
	var req models.ValidateCardRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(err, "failed to decode body")
	}

	return &req, nil
}

func ValidateCard(w http.ResponseWriter, r *http.Request) {
	req, err := newValidateCardRequest(r)
	if err != nil {
		RenderErr(w, problems.BadRequest(err))
		return
	}

	validationError := validateValidateCardRequest(req)
	if validationError != nil {
		Render(w, models.ValidateCardResponse{Error: validationError})
		return
	}

	Render(w, models.ValidateCardResponse{Valid: true})
}

func validateValidateCardRequest(req *models.ValidateCardRequest) *problems.Error {
	if req == nil {
		return problems.BadRequest(errors.New("request is empty"))
	}

	if req.CardNumber == "" {
		return problems.BadRequest(errors.New("card number is empty"))
	}

	expirationMonth, err := strconv.Atoi(req.ExpirationMonth)
	if err != nil {
		return problems.BadRequest(errors.Wrap(err, "failed to parse expiration month"))
	}

	expirationYear, err := strconv.Atoi(req.ExpirationYear)
	if err != nil {
		return problems.BadRequest(errors.Wrap(err, "failed to parse expiration year"))
	}

	if expirationMonth == 0 {
		return problems.BadRequest(errors.New("expiration month is empty"))
	}

	if expirationYear == 0 {
		return problems.BadRequest(errors.New("expiration year is empty"))
	}

	if expirationMonth < 1 || expirationMonth > 12 {
		return problems.BadRequest(errors.New("expiration month is invalid, must be between 1 and 12"))
	}

	if expirationYear < 2024 || expirationYear > 2030 {
		return problems.BadRequest(errors.New("card is expired, expiration year must be greater than 2024, and less than 2030"))
	}

	return validateCardNumber(req.CardNumber)
}

func validateCardNumber(cardNumber string) *problems.Error {
	var sum int
	var alternate bool

	numberLen := len(cardNumber)

	if numberLen < 13 || numberLen > 19 {
		return problems.BadRequest(errors.New("card number is invalid, must be between 13 and 19 digits"))
	}

	for i := numberLen - 1; i > -1; i-- {
		mod, _ := strconv.Atoi(string(cardNumber[i]))
		if alternate {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}

		alternate = !alternate
		sum += mod
	}

	if sum%10 == 0 {
		return nil
	}

	return problems.BadRequest(errors.New("card number is invalid"))
}
