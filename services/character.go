package services

import (
	"github.com/williamjoseph77/evos/db"
	"github.com/williamjoseph77/evos/domains"
	"github.com/williamjoseph77/evos/objects"

	"github.com/go-pg/pg/v10"
)

// CreateCharacterNonSecure ...
func CreateCharacterNonSecure(dbi *pg.DB, request objects.CreateCharacterNonSecureRequest) (*domains.Character, error) {
	newCharacter, err := db.CreateCharacter(dbi, domains.Character{
		Name:   request.Name,
		RoleID: request.RoleID,
		Power:  request.Power,
		Wealth: request.Wealth,
	})

	if err != nil {
		return nil, err
	}

	return newCharacter, nil
}

// CreateCharacterSecure ...
func CreateCharacterSecure(dbi *pg.DB, request objects.CreateCharacterSecureRequest) (*domains.Character, error) {
	role, err := db.GetRoleByGUID(dbi, request.RoleGUID)
	if err != nil {
		return nil, err
	}

	newCharacter, err := db.CreateCharacter(dbi, domains.Character{
		Name:   request.Name,
		RoleID: role.ID,
		Power:  request.Power,
		Wealth: request.Wealth,
	})
	if err != nil {
		return nil, err
	}

	return newCharacter, nil
}

// GetCharacterList ...
func GetCharacterList(dbi *pg.DB) ([]domains.Character, error) {
	characters, err := db.GetCharacters(dbi)
	if err != nil {
		return nil, err
	}

	return characters, nil
}
