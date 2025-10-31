package mainservicedto

type StaffResponse struct {
	StaffID        string `json:"id"`
	OrganizationID string `json:"organization_id"`
	StaffName      string `json:"name"`
	Avatar         Avatar `json:"avatar"`
	Code           string `json:"code"`
}
