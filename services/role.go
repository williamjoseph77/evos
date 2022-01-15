package services

import (
	"evos/be/db"
	"evos/be/domains"

	"github.com/go-pg/pg/v10"
)

// GetRoleList ...
func GetRoleList(dbi *pg.DB) ([]domains.Role, error) {
	roles, err := db.GetRoles(dbi)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
