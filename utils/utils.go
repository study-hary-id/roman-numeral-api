package utils

import (
	"encoding/json"
	"net/http"

	"github.com/study-hary-id/roman-numeral-api/models"
)

// ConvertToRoman return roman numeral effectively until 3999.
func ConvertToRoman(num int) (romanNum string) {

	var (
		keyNumerals = [13]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
		index       = len(keyNumerals) - 1
	)
	for num != 0 {
		maxChance := num / keyNumerals[index]
		num %= keyNumerals[index]

		for maxChance != 0 {
			romanNum += Numerals[keyNumerals[index]]
			maxChance -= 1
		}
		index -= 1
	}
	return romanNum

}

// ResponseJSON marshal the payload and send back.
func ResponseJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		InternalServerError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(response)
	if err != nil {
		InternalServerError(w, err)
	}
}

// ResponseError wraps payload using models.ErrorPayload.
func ResponseError(w http.ResponseWriter, status int, payload models.Errors) {
	ResponseJSON(w, status, models.ErrorPayload{Errors: payload})
}

// ResponseSuccess wraps payload using models.Payload.
func ResponseSuccess(w http.ResponseWriter, status int, payload models.Numeral) {
	ResponseJSON(w, status, models.Payload{Data: payload})
}

// InternalServerError handles server error.
func InternalServerError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
