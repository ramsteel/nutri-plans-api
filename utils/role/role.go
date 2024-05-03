package role

import "nutri-plans-api/entities"

var roleTypes = []entities.RoleType{
	{
		ID:   1,
		Name: "user",
	},
	{
		ID:   2,
		Name: "admin",
	},
}

func GetRoleTypes() []entities.RoleType {
	return roleTypes
}
