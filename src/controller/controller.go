package controller

import (
	"context"
	"superheroe-api/superheroe-golang-api/src/entity"
	"superheroe-api/superheroe-golang-api/src/repository"
	"superheroe-api/superheroe-golang-api/src/util"

	log "github.com/sirupsen/logrus"
)

//Service main interface for the service with the business logic
type Controller interface {
	GetAll(ctx context.Context) ([]*entity.Superhero, error)
	GetByID(id string, ctx context.Context) (*entity.Superhero, error)
	Add(c *entity.Superhero, ctx context.Context) (*entity.Superhero, error)
	Edit(id string, c *entity.Superhero, ctx context.Context) (*entity.Superhero, error)
	Delete(id string, ctx context.Context) (string, error)
}

type controller struct {
	repo repository.Repository
}

//NewController initialice a new controller
func NewController(rep repository.Repository) Controller {
	log.SetFormatter(&log.JSONFormatter{})
	return &controller{
		repo: rep,
	}
}

//GetAll return all superheroes
func (s *controller) GetAll(ctx context.Context) ([]*entity.Superhero, error) {
	superheroes, err := s.repo.GetSuperheroes(ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "GetAll"}).Error(err.Error())
		return nil, err
	}
	log.WithFields(log.Fields{"package": "controller", "method": "GetAll"}).Info("ok")

	return superheroes, nil
}

//GetAll return a single superheroe
func (s *controller) GetByID(id string, ctx context.Context) (*entity.Superhero, error) {
	resp, err := s.repo.GetSuperheroeById(id, ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "GetByID"}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{"package": "controller", "method": "GetByID"}).Info("ok")
	return resp, nil
}

//GetAll add a new superheroe
func (s *controller) Add(c *entity.Superhero, ctx context.Context) (*entity.Superhero, error) {
	resp, err := s.repo.GetSuperheroes(ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "Add"}).Error(err.Error())
		return nil, err
	}

	err = util.VerifySuperheroe(resp, *c)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "Add"}).Error(err.Error())
		return nil, &util.BadRequestError{Message: err.Error()}
	}

	superhero, err := s.repo.AddSuperheroe(c, ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "Add"}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{"package": "controller", "method": "Add"}).Info("ok")
	return superhero, nil
}

//Edit a superheroe
func (s *controller) Edit(id string, c *entity.Superhero, ctx context.Context) (*entity.Superhero, error) {
	heroe, err := s.repo.EditSuperheroe(id, c, ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "Edit"}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{"package": "controller", "method": "Edit"}).Info("ok")
	return heroe, nil
}

//Delete delete a superheroe
func (s *controller) Delete(id string, ctx context.Context) (string, error) {
	response, err := s.repo.DeleteSuperheroe(id, ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "Delete"}).Error(err.Error())
		return "", err
	}

	log.WithFields(log.Fields{"package": "controller", "method": "Edit"}).Info("ok")
	return response, nil
}
