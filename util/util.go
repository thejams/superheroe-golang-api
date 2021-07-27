//Package util provides all utilities functions to help with the logic
package util

import (
	"errors"
	"superheroe-api/superheroe-golang-api/entity"
)

//VerifySuperheroe verify if a field from a superheroe is already taken
func VerifySuperheroe(s []*entity.Superheroe, c entity.Superheroe) error {
	for _, v := range s {
		if v.Name == c.Name {
			return errors.New("Name is already taken")
		}
	}
	return nil
}

//SuperheroeExists verify if a superheroe already exists
func SuperheroeExists(s []*entity.Superheroe, id string) bool {
	for _, v := range s {
		if v.ID == id {
			return true
		}
	}
	return false
}
