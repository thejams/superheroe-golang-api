package controller_test

import (
	"context"
	"fmt"
	"superheroe-api/superheroe-golang-api/src/controller"
	"superheroe-api/superheroe-golang-api/src/entity"
	mock "superheroe-api/superheroe-golang-api/src/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	batman entity.Superhero
	sh     []*entity.Superhero
)

func init() {
	oid, _ := primitive.ObjectIDFromHex("1")
	batman = entity.Superhero{
		ID:    oid,
		Name:  "Batman",
		Alias: "World Greatest Detective",
	}
	sh = []*entity.Superhero{&batman}
}

func TestGetAll(t *testing.T) {
	t.Run("should return an array with 1 superheroe", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo, nil)
		mockRepo.On("GetSuperheroes", context.TODO()).Return([]*entity.Superhero{&batman}, nil)
		result, _ := ctrl.GetAll(context.TODO())

		assert.Equal(t, "Batman", result[0].Name)
		assert.Equal(t, "World Greatest Detective", result[0].Alias)
	})
	t.Run("should return error when  database connection fails", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo, nil)
		mockRepo.On("GetSuperheroes", context.TODO()).Return(nil, fmt.Errorf("error in database connection"))
		_, err := ctrl.GetAll(context.TODO())

		assert.NotNil(t, err)
		assert.Equal(t, "error in database connection", err.Error())
	})
}

func TestGetByID(t *testing.T) {
	t.Run("should return error when no superheroe is found", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo, nil)
		mockRepo.On("GetSuperheroeById", "1", context.TODO()).Return(nil, fmt.Errorf("no superheroe with id %v found", 1))
		_, err := ctrl.GetByID("1", context.TODO())

		assert.NotNil(t, err)
		assert.Equal(t, "no superheroe with id 1 found", err.Error())
	})

	t.Run("should return a superheroe", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo, nil)
		mockRepo.On("GetSuperheroeById", "1", context.TODO()).Return(&batman, nil)
		result, _ := ctrl.GetByID("1", context.TODO())

		assert.Equal(t, "Batman", result.Name)
		assert.Equal(t, "World Greatest Detective", result.Alias)
	})
}

func TestAdd(t *testing.T) {
	t.Run("should return error when name of heroe already exists", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo, nil)
		nh := entity.Superhero{
			Name:  "Batman",
			Alias: "The Dark Knight",
		}
		mockRepo.On("GetSuperheroes", context.TODO()).Return(sh, nil)
		_, err := ctrl.Add(&nh, context.TODO())

		assert.NotNil(t, err)
		assert.Equal(t, "BadRequest: Name is already taken", err.Error())
	})

	t.Run("should add a new superheroe", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo, nil)
		nh := entity.Superhero{
			Name:  "Superman",
			Alias: "The Man Of Steel",
		}
		mockRepo.On("GetSuperheroes", context.TODO()).Return(sh, nil)
		mockRepo.On("AddSuperheroe", &nh, context.TODO()).Return(&nh, nil)
		result, _ := ctrl.Add(&nh, context.TODO())

		assert.Equal(t, "Superman", result.Name)
		assert.Equal(t, "The Man Of Steel", result.Alias)
	})

	t.Run("should error when data base insert operation fails", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo, nil)
		nh := entity.Superhero{
			Name:  "Superman",
			Alias: "The Man Of Steel",
		}
		mockRepo.On("GetSuperheroes", context.TODO()).Return(sh, nil)
		mockRepo.On("AddSuperheroe", &nh, context.TODO()).Return(nil, fmt.Errorf("Insert operation fail"))
		_, err := ctrl.Add(&nh, context.TODO())

		assert.NotNil(t, err)
		assert.Equal(t, "Insert operation fail", err.Error())
	})
}

func TestEdit(t *testing.T) {
	t.Run("should return error when heroe does not exists", func(t *testing.T) {
		oid, _ := primitive.ObjectIDFromHex("2")
		nh := entity.Superhero{
			ID:    oid,
			Name:  "Superman",
			Alias: "The Man Of Steel",
		}
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo, nil)
		mockRepo.On("EditSuperheroe", "1", &nh, context.TODO()).Return(nil, fmt.Errorf("Superheroe with ID %v does not exist", 2))
		_, err := ctrl.Edit("1", &nh, context.TODO())

		assert.NotNil(t, err)
		assert.Equal(t, "Superheroe with ID 2 does not exist", err.Error())
	})

	t.Run("should edit a superheroe information", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo, nil)
		oid, _ := primitive.ObjectIDFromHex("1")
		nh := entity.Superhero{
			ID:    oid,
			Name:  "Superman",
			Alias: "The Last Son Of Krypton",
		}
		mockRepo.On("EditSuperheroe", "1", &nh, context.TODO()).Return(&nh, nil)
		result, _ := ctrl.Edit("1", &nh, context.TODO())

		assert.Equal(t, "Superman", result.Name)
		assert.Equal(t, "The Last Son Of Krypton", result.Alias)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should return error when heroe does not exists", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo, nil)
		mockRepo.On("DeleteSuperheroe", "2", context.TODO()).Return("", fmt.Errorf("Superheroe with ID %v does not exist", 2))
		_, err := ctrl.Delete("2", context.TODO())

		assert.NotNil(t, err)
		assert.Equal(t, "Superheroe with ID 2 does not exist", err.Error())
	})

	t.Run("should delete a superheroe", func(t *testing.T) {
		mockRepo := new(mock.Repository)
		ctrl := controller.NewController(mockRepo, nil)
		mockRepo.On("DeleteSuperheroe", "1", context.TODO()).Return("Character deleted 1", nil)
		result, _ := ctrl.Delete("1", context.TODO())

		assert.Equal(t, "Character deleted 1", result)
	})
}

func BenchmarkByID(b *testing.B) {
	mockRepo := new(mock.Repository)
	ctrl := controller.NewController(mockRepo, nil)
	mockRepo.On("GetSuperheroes", context.TODO()).Return([]*entity.Superhero{&batman})

	for i := 0; i < b.N; i++ {
		ctrl.GetAll(context.TODO())
	}
}

func BenchmarkAdd(b *testing.B) {
	mockRepo := new(mock.Repository)
	ctrl := controller.NewController(mockRepo, nil)
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
	ctrl := controller.NewController(mockRepo, nil)
	oid, _ := primitive.ObjectIDFromHex("1")
	nh := entity.Superhero{
		ID:    oid,
		Name:  "Superman",
		Alias: "The Last Son Of Krypton",
	}
	mockRepo.On("EditSuperheroe", &nh).Return(&nh, nil)
	mockRepo.On("GetSuperheroes", context.TODO()).Return(sh)

	for i := 0; i < b.N; i++ {
		ctrl.Edit("1", &nh, context.TODO())
	}
}

func BenchmarkDelete(b *testing.B) {
	mockRepo := new(mock.Repository)
	ctrl := controller.NewController(mockRepo, nil)
	mockRepo.On("GetSuperheroes", context.TODO()).Return(sh)
	mockRepo.On("DeleteSuperheroe", "1").Return("Character deleted 1", nil)

	for i := 0; i < b.N; i++ {
		ctrl.Delete("1", context.TODO())
	}
}
