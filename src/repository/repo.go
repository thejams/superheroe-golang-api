//Package repository provides all the methods to persit data in a slice
package repository

import (
	"context"
	"superheroe-api/superheroe-golang-api/src/entity"
)

var superheroesList []*entity.Superhero

//Repository main repository interface
type Repository interface {
	GetSuperheroes(ctx context.Context) ([]*entity.Superhero, error)
	GetSuperheroeById(id string, ctx context.Context) (*entity.Superhero, error)
	AddSuperheroe(c *entity.Superhero, ctx context.Context) (*entity.Superhero, error)
	EditSuperheroe(id string, c *entity.Superhero, ctx context.Context) (*entity.Superhero, error)
	DeleteSuperheroe(id string, ctx context.Context) (string, error)
}

/* type repository struct{}

//NewRepository initialice a new repository with clean data
func NewRepository() Repository {
	superheroesList = []*entity.Superhero{
		{Name: "Thor", Alias: "Thor Odinson", ID: "1"},
		{Name: "Batman", Alias: "Bruce Wayne", ID: "2"},
		{Name: "Iron Man", Alias: "Tony Stark", ID: "3"},
		{Name: "Superman", Alias: "Clark Kent", ID: "4"},
	}

	return &repository{}
} */

//GetSuperheroes returns all the superheroes in the slice
/* func (r *repository) GetSuperheroes() []*entity.Superhero {
	return superheroesList
} */

//GetSuperheroeById returns a single superheroe from the slice
/* func (r *repository) GetSuperheroeById(i string) (*entity.Superhero, error) {
	for _, value := range superheroesList {
		if value.ID == i {
			return value, nil
		}
	}
	return nil, &util.NotFoundError{Message: fmt.Sprintf("no superheroe with id %v found", i)}
} */

//AddSuperheroe add a new superheroe to the superheroes slice
/* func (r *repository) AddSuperheroe(c *entity.Superhero) *entity.Superhero {
	superheroesList = append(superheroesList, c)
	return c
} */

//EditCharacter edit a superheroe with new information
/* func (r *repository) EditSuperheroe(c *entity.Superhero) (*entity.Superhero, error) {
	for index, value := range superheroesList {
		if value.ID == c.ID {
			superheroesList = append(superheroesList[:index], superheroesList[index+1:]...)
			superheroesList = append(superheroesList, c)
			return c, nil
		}
	}
	return nil, &util.NotFoundError{Message: fmt.Sprintf("Superheroe with ID %v does not exist", c.ID)}
} */

//DeleteSuperheroe remove a superheroe from the superheroes slice
/* func (r *repository) DeleteSuperheroe(id string) (string, error) {
	for index, value := range superheroesList {
		if value.ID == id {
			superheroesList = append(superheroesList[:index], superheroesList[index+1:]...)
			return fmt.Sprintf("Character deleted %v", id), nil
		}
	}
	return "", &util.NotFoundError{Message: fmt.Sprintf("Superheroe with ID %v does not exist", id)}
} */
