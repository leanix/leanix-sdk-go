/* 
 * LeanIX Pathfinder REST API
 *
 * Core application for storage and analysis of IT landscape data
 *
 * OpenAPI spec version: 0.1.10-SNAPSHOT
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type JobResponse struct {

	ErrorMessage string `json:"errorMessage,omitempty"`

	Errors []ApiError `json:"errors,omitempty"`

	JobId string `json:"jobId,omitempty"`

	SynclogId string `json:"synclogId,omitempty"`
}
