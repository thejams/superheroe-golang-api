//Package repository provides all the methods to persit data in a slice
package repository

import (
	"context"
	"superheroe-api/superheroe-golang-api/src/entity"
)

//Repository main interface
type Repository interface {
	GetSuperheroes(ctx context.Context) ([]entity.Superhero, error)
	GetSuperheroeById(id string, ctx context.Context) (*entity.Superhero, error)
	AddSuperheroe(c *entity.Superhero, ctx context.Context) (*entity.Superhero, error)
	EditSuperheroe(id string, c *entity.Superhero, ctx context.Context) (*entity.Superhero, error)
	DeleteSuperheroe(id string, ctx context.Context) (string, error)
}
