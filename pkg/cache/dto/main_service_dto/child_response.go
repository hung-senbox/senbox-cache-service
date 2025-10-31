package mainservicedto

type ChildResponse struct {
	ChildID   string `json:"id"`
	ChildName string `json:"name"`
	Avatar    Avatar `json:"avatar"`
	Code      string `json:"code"`
}
