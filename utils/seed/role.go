package seed

import "nutri-plans-api/entities"

func GetRoleTypes() *[]entities.RoleType {
	roleTypes := &[]entities.RoleType{
		{
			ID:   1,
			Name: "user",
		},
		{
			ID:   2,
			Name: "admin",
		},
	}
	return roleTypes
}
