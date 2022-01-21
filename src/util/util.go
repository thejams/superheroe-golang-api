//Package util provides all utilities functions to help with the logic
package util

import (
	"errors"
	"superheroe-api/superheroe-golang-api/src/entity"
)

//VerifySuperheroe verify if a field from a superheroe is already taken
func VerifySuperheroe(s []entity.Character, c entity.Character) error {
	for _, v := range s {
		if v.Name == c.Name {
			return errors.New("Name is already taken")
		}
	}
	return nil
}

//SuperheroeExists verify if a superheroe already exists
func SuperheroeExists(s []*entity.Character, name string) bool {
	for _, v := range s {
		if v.Name == name {
			return true
		}
	}
	return false
}
