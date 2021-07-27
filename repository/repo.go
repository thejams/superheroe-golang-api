//Package repository provides all the methods to persit data in a slice
package repository

import (
	"fmt"
	"superheroe-api/superheroe-golang-api/entity"
	"superheroe-api/superheroe-golang-api/util"
)

var superheroesList []*entity.Superheroe

//Repository main repository interface
type Repository interface {
	GetSuperheroes() []*entity.Superheroe
	GetSuperheroeById(id string) (*entity.Superheroe, error)
	AddSuperheroe(c *entity.Superheroe) *entity.Superheroe
	EditSuperheroe(c *entity.Superheroe) (*entity.Superheroe, error)
	DeleteSuperheroe(id string) (string, error)
}

type repository struct{}

//NewRepository initialice a new repository with clean data
func NewRepository() Repository {
	superheroesList = []*entity.Superheroe{
		{Name: "Thor", Alias: "Thor Odinson", ID: "1"},
		{Name: "Batman", Alias: "Bruce Wayne", ID: "2"},
		{Name: "Iron Man", Alias: "Tony Stark", ID: "3"},
		{Name: "Superman", Alias: "Clark Kent", ID: "4"},
	}

	return &repository{}
}

//GetSuperheroes returns all the superheroes in the slice
func (r *repository) GetSuperheroes() []*entity.Superheroe {
	return superheroesList
}

//GetSuperheroeById returns a single superheroe from the slice
func (r *repository) GetSuperheroeById(i string) (*entity.Superheroe, error) {
	for _, value := range superheroesList {
		if value.ID == i {
			return value, nil
		}
	}
	return nil, &util.NotFoundError{Message: fmt.Sprintf("no superheroe with id %v found", i)}
}

//AddSuperheroe add a new superheroe to the superheroes slice
func (r *repository) AddSuperheroe(c *entity.Superheroe) *entity.Superheroe {
	superheroesList = append(superheroesList, c)
	return c
}

//EditCharacter edit a superheroe with new information
func (r *repository) EditSuperheroe(c *entity.Superheroe) (*entity.Superheroe, error) {
	for index, value := range superheroesList {
		if value.ID == c.ID {
			superheroesList = append(superheroesList[:index], superheroesList[index+1:]...)
			superheroesList = append(superheroesList, c)
			return c, nil
		}
	}
	return nil, &util.NotFoundError{Message: fmt.Sprintf("Superheroe with ID %v does not exist", c.ID)}
}

//DeleteSuperheroe remove a superheroe from the superheroes slice
func (r *repository) DeleteSuperheroe(id string) (string, error) {
	for index, value := range superheroesList {
		if value.ID == id {
			superheroesList = append(superheroesList[:index], superheroesList[index+1:]...)
			return fmt.Sprint("Character deleted %v", id), nil
		}
	}
	return "", &util.NotFoundError{Message: fmt.Sprintf("Superheroe with ID %v does not exist", id)}
}
