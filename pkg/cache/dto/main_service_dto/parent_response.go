package mainservicedto

type ParentResponse struct {
	ParentID       string `json:"id"`
	OrganizationID string `json:"organization_id"`
	ParentName     string `json:"name"`
	Avatar         Avatar `json:"avatar"`
	Code           string `json:"code"`
}
