package controller

import (
	"superheroe-api/superheroe-golang-api/src/entity"
	"superheroe-api/superheroe-golang-api/src/repository"
	"superheroe-api/superheroe-golang-api/src/util"

	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
)

//Service main interface for the service with the business logic
type Controller interface {
	GetAll() ([]*entity.Superhero, error)
	GetByID(id string) (*entity.Superhero, error)
	Add(c *entity.Superhero) (*entity.Superhero, error)
	Edit(c *entity.Superhero) (*entity.Superhero, error)
	Delete(id string) (string, error)
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
func (s *controller) GetAll() ([]*entity.Superhero, error) {
	log.WithFields(log.Fields{"package": "controller", "method": "GetAll"}).Info("ok")
	return s.repo.GetSuperheroes(), nil
}

//GetAll return a single superheroe
func (s *controller) GetByID(id string) (*entity.Superhero, error) {
	resp, err := s.repo.GetSuperheroeById(id)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "GetByID"}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{"package": "controller", "method": "GetByID"}).Info("ok")
	return resp, nil
}

//GetAll add a new superheroe
func (s *controller) Add(c *entity.Superhero) (*entity.Superhero, error) {
	resp := s.repo.GetSuperheroes()
	err := util.VerifySuperheroe(resp, *c)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "Add"}).Error(err.Error())
		return nil, &util.BadRequestError{Message: err.Error()}
	}

	uuid, _ := uuid.NewV4()
	c.ID = uuid.String()
	s.repo.AddSuperheroe(c)

	log.WithFields(log.Fields{"package": "controller", "method": "Add"}).Info("ok")
	return c, nil
}

//Edit a superheroe
func (s *controller) Edit(c *entity.Superhero) (*entity.Superhero, error) {
	heroe, err := s.repo.EditSuperheroe(c)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "Edit"}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{"package": "controller", "method": "Edit"}).Info("ok")
	return heroe, nil
}

//Delete delete a superheroe
func (s *controller) Delete(id string) (string, error) {
	response, err := s.repo.DeleteSuperheroe(id)
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "Delete"}).Error(err.Error())
		return "", err
	}

	log.WithFields(log.Fields{"package": "controller", "method": "Edit"}).Info("ok")
	return response, nil
}