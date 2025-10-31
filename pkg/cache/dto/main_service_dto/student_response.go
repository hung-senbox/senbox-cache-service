package mainservicedto

type StudentResponse struct {
	StudentID      string `json:"id"`
	OrganizationID string `json:"organization_id"`
	StudentName    string `json:"name"`
	Avatar         Avatar `json:"avatar"`
	Code           string `json:"code"`
}
