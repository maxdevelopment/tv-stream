package api

import (
	"net/http"
	"github.com/maxdevelopment/tv-stream/utils"
	"time"
	"encoding/json"
)

type StreamClient struct {
	Id string `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	sc := &StreamClient{
		Id:utils.RandStringBytes(8),
		CreatedAt:time.Now(),
	}

	resp, err := json.Marshal(&sc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
