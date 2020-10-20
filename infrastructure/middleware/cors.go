package middleware

import "net/http"

func Cors(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	response.Header().Set("Access-Control-Allow-Origin", "Accept, Authorization, Content-Type")
	response.Header().Set("Content-Type", "appication/json")
	if request.Method == "OPTIONS" {
		return
	}
	next(response, request)
}
