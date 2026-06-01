package booking

import (
	"encoding/json"
	"net/http"
)

type PayloadBooking struct {
	SeatId  string `json:"seatID"`
	MovieId string `json:"movieID"`
	UserId  string `json:"userID"`
}

func HandleBooking(w http.ResponseWriter, r *http.Request) {
	var payload PayloadBooking

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "incorrect payload", http.StatusBadRequest)
	}

}
