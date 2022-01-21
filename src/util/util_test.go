package util_test

import (
	"fmt"
	"superheroe-api/superheroe-golang-api/src/entity"
	"superheroe-api/superheroe-golang-api/src/util"

	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	batman   entity.Character
	superman entity.Character
)

func init() {
	batmanOID, _ := primitive.ObjectIDFromHex("1")
	supermanOID, _ := primitive.ObjectIDFromHex("2")
	batman = entity.Character{
		ID:    batmanOID,
		Name:  "Batman",
		Alias: "The World's Greatest Detective",
	}

	superman = entity.Character{
		ID:    supermanOID,
		Name:  "Superman",
		Alias: "The Last Son Of Krypton",
	}
}

func TestVerifySuperheroe(t *testing.T) {
	t.Run("should return error when name is already taken", func(t *testing.T) {
		oid, _ := primitive.ObjectIDFromHex("1")
		thor := entity.Character{
			ID:    oid,
			Name:  "Thor",
			Alias: "God of Thunder",
		}
		sh := []entity.Character{thor}
		err := util.VerifySuperheroe(sh, thor)

		assert.NotNil(t, err)
		assert.Equal(t, "Name is already taken", err.Error())
	})

	t.Run("should not return error when a new heroe is verified", func(t *testing.T) {
		sh := []entity.Character{batman}
		err := util.VerifySuperheroe(sh, superman)

		assert.Nil(t, err)
	})
}

func TestSuperheroeExists(t *testing.T) {
	t.Run("should return true if a heroe exists", func(t *testing.T) {
		oid, _ := primitive.ObjectIDFromHex("1")
		hulk := entity.Character{
			ID:    oid,
			Name:  "The Hulk",
			Alias: "The Strongest There Is",
		}
		sh := []*entity.Character{&hulk}
		resp := util.SuperheroeExists(sh, "The Hulk")

		assert.True(t, resp)
	})

	t.Run("should return false if a heroe does not exists", func(t *testing.T) {
		oid, _ := primitive.ObjectIDFromHex("1")
		wonderWoman := entity.Character{
			ID:    oid,
			Name:  "Wonder Woman",
			Alias: "Princess of Themyscira",
		}
		sh := []*entity.Character{&wonderWoman}
		resp := util.SuperheroeExists(sh, "Thor")

		assert.False(t, resp)
	})
}

func BenchmarkVerifySuperheroe(b *testing.B) {
	sh := []entity.Character{batman}
	for i := 0; i < b.N; i++ {
		util.VerifySuperheroe(sh, superman)
	}
}

func BenchmarkSuperheroeExists(b *testing.B) {
	sh := []*entity.Character{&batman}
	for i := 0; i < b.N; i++ {
		util.SuperheroeExists(sh, "Batman")
	}
}

func ExampleSuperheroeExists() {
	sh := []*entity.Character{&batman}
	fmt.Println(util.SuperheroeExists(sh, "Batman"))
	fmt.Println(util.SuperheroeExists(sh, "Thor"))
	//Output:
	//true
	//false
}
