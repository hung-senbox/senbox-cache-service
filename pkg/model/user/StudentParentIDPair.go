package user

// StudentParentIDPair mô tả một quan hệ phụ huynh - học sinh.
type StudentParentIDPair struct {
	ParentID  string `json:"parentid"`
	StudentID string `json:"studentid"`
}
