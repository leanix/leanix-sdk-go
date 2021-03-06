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

type SavedQuery struct {

	Id string `json:"id,omitempty"`

	WorkspaceId string `json:"workspaceId"`

	CreatorId string `json:"creatorId"`

	Name string `json:"name"`

	QueryType string `json:"queryType"`

	GroupKey string `json:"groupKey,omitempty"`

	State map[string]interface{} `json:"state"`

	Rev int64 `json:"_rev,omitempty"`
}
