package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mohdjishin/sportsphere/internal/meta"
)

type infoResponse struct {
	Version   string `json:"version"`
	BuildTime string `json:"buildTime"`
	CommitID  string `json:"commitID"`
}

func Info(w http.ResponseWriter, r *http.Request) {
	response := infoResponse{
		Version:   meta.GetVersion(),
		BuildTime: meta.GetBuildTime(),
		CommitID:  meta.GetCommitID(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}
