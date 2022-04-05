//Package repository provides all the methods to persit data in a slice
package repository

import (
	"context"

	"superheroe-api/superheroe-golang-api/src/entity"
)

//Repository main interface
type Repository interface {
	Conn(context.Context, *entity.APPConfig) error
	Close(context.Context) error
	GetAll(ctx context.Context) ([]entity.Character, error)
	Get(id string, ctx context.Context) (*entity.Character, error)
	Add(c *entity.Character, ctx context.Context) (*entity.Character, error)
	Edit(id string, c *entity.Character, ctx context.Context) (*entity.Character, error)
	Delete(id string, ctx context.Context) (string, error)
}
