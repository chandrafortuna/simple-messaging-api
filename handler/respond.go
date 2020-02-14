package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ReplyError would response errors to api request
func ReplyError(w *http.ResponseWriter, code int, args ...interface{}) {
	response := make(map[string]string)
	response["error"] = fmt.Sprint(args...)
	responseJSON, _ := json.Marshal(response)

	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(code)
	(*w).Write(responseJSON)
}

// ReplySuccess would response success to api request
func ReplySuccess(w *http.ResponseWriter, responseData interface{}) {
	responseJSON, _ := json.Marshal(responseData)

	(*w).Header().Set("Content-Type", "application/json")
	(*w).Write(responseJSON)
}

// func decodeBody(r *http.Request, v interface{}) error {
// 	defer r.Body.Close()
// 	return json.NewDecoder(r.Body).Decode(v)
// }

// func encodeBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
// 	return json.NewEncoder(w).Encode(v)
// }

// func respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
// 	w.WriteHeader(status)
// 	if data != nil {
// 		encodeBody(w, r, data)
// 	}
// }

// func respondErr(w http.ResponseWriter, r *http.Request, status int, args ...interface{}) {
// 	respond(w, r, status, map[string]interface{}{
// 		"error": map[string]interface{}{
// 			"message": fmt.Sprint(args...),
// 		},
// 	})
// }

// func respondHTTPErr(w http.ResponseWriter, r *http.Request, status int) {
// 	respondErr(w, r, status, http.StatusText(status))
// }
