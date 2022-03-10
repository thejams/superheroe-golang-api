package controller

import (
	"context"

	log "github.com/sirupsen/logrus"

	"superheroe-api/superheroe-golang-api/src/client"
	"superheroe-api/superheroe-golang-api/src/entity"
	"superheroe-api/superheroe-golang-api/src/repository"
	"superheroe-api/superheroe-golang-api/src/util"
)

//Service main interface for the service with the business logic
type Controller interface {
	GetAll(ctx context.Context) ([]entity.Character, error)
	GetByID(id string, ctx context.Context) (*entity.Character, error)
	Add(c *entity.Character, ctx context.Context) (*entity.Character, error)
	Edit(id string, c *entity.Character, ctx context.Context) (*entity.Character, error)
	Delete(id string, ctx context.Context) (string, error)
	GetHttpRequest() (interface{}, error)
}

type controller struct {
	repo   repository.Repository
	client client.Client
}

//NewController initialice a new controller
func NewController(rep repository.Repository, client client.Client) Controller {
	log.SetFormatter(&log.JSONFormatter{})
	return &controller{
		repo:   rep,
		client: client,
	}
}

//GetHttpRequest makes an http get request via client
func (s *controller) GetHttpRequest() (interface{}, error) {
	res, err := s.client.Get()
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "GetHttpRequest"}).Error(err.Error())
		return nil, err
	}
	log.WithFields(log.Fields{"package": "controller", "method": "GetHttpRequest"}).Info("ok")

	return res, nil
}

//GetAll return all superheroes
func (s *controller) GetAll(ctx context.Context) ([]entity.Character, error) {
	superheroes, err := s.repo.GetAll(ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "GetAll"}).Error(err.Error())
		return nil, err
	}
	log.WithFields(log.Fields{"package": "controller", "method": "GetAll"}).Info("ok")

	return superheroes, nil
}

//GetAll return a single superheroe
func (s *controller) GetByID(id string, ctx context.Context) (*entity.Character, error) {
	resp, err := s.repo.Get(id, ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "GetByID"}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{"package": "controller", "method": "GetByID"}).Info("ok")
	return resp, nil
}

//GetAll add a new superheroe
func (s *controller) Add(c *entity.Character, ctx context.Context) (*entity.Character, error) {
	resp, err := s.repo.GetAll(ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "Add"}).Error(err.Error())
		return nil, err
	}

	err = util.VerifySuperheroe(resp, *c)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "Add"}).Error(err.Error())
		return nil, &util.BadRequestError{Message: err.Error()}
	}

	superhero, err := s.repo.Add(c, ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "Add"}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{"package": "controller", "method": "Add"}).Info("ok")
	return superhero, nil
}

//Edit a superheroe
func (s *controller) Edit(id string, c *entity.Character, ctx context.Context) (*entity.Character, error) {
	heroe, err := s.repo.Edit(id, c, ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "Edit"}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{"package": "controller", "method": "Edit"}).Info("ok")
	return heroe, nil
}

//Delete delete a superheroe
func (s *controller) Delete(id string, ctx context.Context) (string, error) {
	response, err := s.repo.Delete(id, ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "Delete"}).Error(err.Error())
		return "", err
	}

	log.WithFields(log.Fields{"package": "controller", "method": "Edit"}).Info("ok")
	return response, nil
}
