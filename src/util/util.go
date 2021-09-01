//Package util provides all utilities functions to help with the logic
package util

import (
	"errors"
	"superheroe-api/superheroe-golang-api/src/entity"
)

//VerifySuperheroe verify if a field from a superheroe is already taken
func VerifySuperheroe(s []*entity.Superhero, c entity.Superhero) error {
	for _, v := range s {
		if v.Name == c.Name {
			return errors.New("Name is already taken")
		}
	}
	return nil
}

//SuperheroeExists verify if a superheroe already exists
func SuperheroeExists(s []*entity.Superhero, name string) bool {
	for _, v := range s {
		if v.Name == name {
			return true
		}
	}
	return false
}
