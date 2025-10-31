package mainservicedto

type TeacherResponse struct {
	TeacherID      string `json:"id"`
	OrganizationID string `json:"organization_id"`
	TeacherName    string `json:"name"`
	Avatar         Avatar `json:"avatar"`
	Code           string `json:"code"`
}
