package api

import (
	"encoding/json"
	"net/http"
)

func ReleasesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, 200, map[string]any{"releases": listReleases()})
	case http.MethodPost:
		var v Release
		if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20)).Decode(&v); err != nil {
			writeError(w, 400, "invalid JSON request")
			return
		}
		saved, err := saveRelease(v)
		if err != nil {
			writeError(w, 400, err.Error())
			return
		}
		writeJSON(w, 201, map[string]any{"release": saved})
	default:
		w.Header().Set("Allow", "GET, POST")
		writeError(w, 405, "method not allowed")
	}
}
