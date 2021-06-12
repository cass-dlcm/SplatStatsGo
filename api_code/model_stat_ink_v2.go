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

type StatInkV2 struct {
	UpdatedAt *StatInkTime `json:"updated_at,omitempty"`

	Entire *StatInkEntire `json:"entire,omitempty"`

	Nawabari *StatInkNawabari `json:"nawabari,omitempty"`

	Gachi *StatInkGachi `json:"gachi,omitempty"`
}
