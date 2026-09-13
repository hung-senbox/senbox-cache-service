package keys

import "github.com/hung-senbox/senbox-cache-service/helper"

func GetOrganizationHierarchyDepartmentCacheKey(organizationID string) string {
	return string(helper.DepartmentPrefix) + string(helper.HierarchyDepartmentPrefix) + organizationID
}
