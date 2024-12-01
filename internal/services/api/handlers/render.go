package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/napalmpapalam/card-validator-svc/internal/problems"
	"github.com/pkg/errors"
)

const mediaType = "application/json"

func RenderErr(w http.ResponseWriter, err *problems.Error) {
	if err == nil {
		panic("expected non-empty error")
	}

	w.Header().Set("content-type", mediaType)
	w.WriteHeader(int(err.Code))
	json.NewEncoder(w).Encode(err)
}

func Render(w http.ResponseWriter, res interface{}) {
	w.Header().Set("content-type", mediaType)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(errors.Wrap(err, "failed to render response"))
	}
}
