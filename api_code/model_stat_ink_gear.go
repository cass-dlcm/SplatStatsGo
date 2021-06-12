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

type StatInkGear struct {
	Key string `json:"key,omitempty"`

	Name *StatInkName `json:"name,omitempty"`

	Splatnet int32 `json:"splatnet,omitempty"`

	Type_ *StatInkKeyName `json:"type,omitempty"`

	Brand *StatInkKeyName `json:"brand,omitempty"`

	PrimaryAbility *StatInkKeyName `json:"primary_ability,omitempty"`
}
