package api

import (
	"foundOriginalPoint/calcul"

	"log/slog"
	"net/http"
)

func (a *API) FindOriginalPoint(w http.ResponseWriter, r *http.Request) {
	var body calcul.FinalInput
	if err := DecodeJSON(r, &body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	resultat := calcul.Operation(body)
	slog.Info("Calculating")

	AnswerJson(w, resultat, 200)
}
