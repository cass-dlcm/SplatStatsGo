/*
 * SplatStats
 *
 * This is the API documentation for a future revision of SplatStats.
 *
 * API version: 0.4.0
 * Contact: splatstats@cass-dlcm.dev
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package api_code

import (
	"net/http"
)

func GetToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetUserByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
