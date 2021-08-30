package controller_test

import (
	"context"
	"fmt"
	"superheroe-api/superheroe-golang-api/src/controller"
	"superheroe-api/superheroe-golang-api/src/entity"
	mock "superheroe-api/superheroe-golang-api/src/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	batman entity.Superhero
	sh     []*entity.Superhero
)

func init() {
	batman = entity.Superhero{
		ID:    "1",
		Name:  "Batman",
		Alias: "World Greatest Detective",
	}
	sh = []*entity.Superhero{&batman}
}

func TestGetAll(t *testing.T) {
	t.Run("should return an array with 1 superheroe", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo)
		mockRepo.On("GetSuperheroes", context.TODO()).Return([]*entity.Superhero{&batman})
		result, _ := ctrl.GetAll(context.TODO())

		assert.Equal(t, "Batman", result[0].Name)
		assert.Equal(t, "World Greatest Detective", result[0].Alias)
		assert.Equal(t, "1", result[0].ID)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("should return error when no superheroe is found", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo)
		mockRepo.On("GetSuperheroeById", "1").Return(nil, fmt.Errorf("no superheroe with id %v found", 1))
		_, err := ctrl.GetByID("1")

		assert.NotNil(t, err)
		assert.Equal(t, "no superheroe with id 1 found", err.Error())
	})

	t.Run("should return a superheroe", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo)
		mockRepo.On("GetSuperheroeById", "1").Return(&batman, nil)
		result, _ := ctrl.GetByID("1")

		assert.Equal(t, "Batman", result.Name)
		assert.Equal(t, "World Greatest Detective", result.Alias)
		assert.Equal(t, "1", result.ID)
	})
}

func TestAdd(t *testing.T) {
	t.Run("should return when name of heroe already exists", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo)
		nh := entity.Superhero{
			Name:  "Batman",
			Alias: "The Dark Knight",
		}
		mockRepo.On("GetSuperheroes", context.TODO()).Return(sh)
		_, err := ctrl.Add(&nh, context.TODO())

		assert.NotNil(t, err)
		assert.Equal(t, "BadRequest: Name is already taken", err.Error())
	})

	t.Run("should add a new superheroe", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo)
		nh := entity.Superhero{
			Name:  "Superman",
			Alias: "The Man Of Steel",
		}
		mockRepo.On("GetSuperheroes", context.TODO()).Return(sh)
		mockRepo.On("AddSuperheroe", &nh).Return(&nh)
		result, _ := ctrl.Add(&nh, context.TODO())

		assert.Equal(t, "Superman", result.Name)
		assert.Equal(t, "The Man Of Steel", result.Alias)
	})
}

func TestEdit(t *testing.T) {
	t.Run("should return error when heroe does not exists", func(t *testing.T) {
		nh := entity.Superhero{
			ID:    "2",
			Name:  "Superman",
			Alias: "The Man Of Steel",
		}
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo)
		mockRepo.On("EditSuperheroe", &nh).Return(nil, fmt.Errorf("Superheroe with ID %v does not exist", 2))
		_, err := ctrl.Edit(&nh)

		assert.NotNil(t, err)
		assert.Equal(t, "Superheroe with ID 2 does not exist", err.Error())
	})

	t.Run("should edit a superheroe information", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo)
		nh := entity.Superhero{
			ID:    "1",
			Name:  "Superman",
			Alias: "The Last Son Of Krypton",
		}
		mockRepo.On("EditSuperheroe", &nh).Return(&nh, nil)
		result, _ := ctrl.Edit(&nh)

		assert.Equal(t, "Superman", result.Name)
		assert.Equal(t, "The Last Son Of Krypton", result.Alias)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should return error when heroe does not exists", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo)
		mockRepo.On("DeleteSuperheroe", "2").Return("", fmt.Errorf("Superheroe with ID %v does not exist", 2))
		_, err := ctrl.Delete("2")

		assert.NotNil(t, err)
		assert.Equal(t, "Superheroe with ID 2 does not exist", err.Error())
	})

	t.Run("should delete a superheroe", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo)
		mockRepo.On("DeleteSuperheroe", "1").Return("Character deleted 1", nil)
		result, _ := ctrl.Delete("1")

		assert.Equal(t, "Character deleted 1", result)
	})
}

func BenchmarkByID(b *testing.B) {
	mockRepo := new(mock.Repository)
	ctrl := controller.NewController(mockRepo)
	mockRepo.On("GetSuperheroes", context.TODO()).Return([]*entity.Superhero{&batman})

	for i := 0; i < b.N; i++ {
		ctrl.GetAll(context.TODO())
	}
}

func BenchmarkAdd(b *testing.B) {
	mockRepo := new(mock.Repository)
	ctrl := controller.NewController(mockRepo)
	nh := entity.Superhero{
		Name:  "Superman",
		Alias: "The Man Of Steel",
	}
	mockRepo.On("GetSuperheroes", context.TODO()).Return(sh)
	mockRepo.On("AddSuperheroe", &nh).Return(&nh)

	for i := 0; i < b.N; i++ {
		ctrl.Add(&nh, context.TODO())
	}
}

func BenchmarkEdit(b *testing.B) {
	mockRepo := new(mock.Repository)
	ctrl := controller.NewController(mockRepo)
	nh := entity.Superhero{
		ID:    "1",
		Name:  "Superman",
		Alias: "The Last Son Of Krypton",
	}
	mockRepo.On("EditSuperheroe", &nh).Return(&nh, nil)
	mockRepo.On("GetSuperheroes", context.TODO()).Return(sh)

	for i := 0; i < b.N; i++ {
		ctrl.Edit(&nh)
	}
}

func BenchmarkDelete(b *testing.B) {
	mockRepo := new(mock.Repository)
	ctrl := controller.NewController(mockRepo)
	mockRepo.On("GetSuperheroes", context.TODO()).Return(sh)
	mockRepo.On("DeleteSuperheroe", "1").Return("Character deleted 1", nil)

	for i := 0; i < b.N; i++ {
		ctrl.Delete("1")
	}
}
