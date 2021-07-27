package util_test

import (
	"fmt"
	"superheroe-api/superheroe-golang-api/entity"
	"superheroe-api/superheroe-golang-api/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	batman   entity.Superhero
	superman entity.Superhero
)

func init() {
	batman = entity.Superhero{
		ID:    "1",
		Name:  "Batman",
		Alias: "The World's Greatest Detective",
	}

	superman = entity.Superhero{
		ID:    "2",
		Name:  "Superman",
		Alias: "The Last Son Of Krypton",
	}
}

func TestVerifySuperheroe(t *testing.T) {
	t.Run("should return error when name is already taken", func(t *testing.T) {
		thor := entity.Superhero{
			ID:    "1",
			Name:  "Thor",
			Alias: "God of Thunder",
		}
		sh := []*entity.Superhero{&thor}
		err := util.VerifySuperheroe(sh, thor)

		assert.NotNil(t, err)
		assert.Equal(t, "Name is already taken", err.Error())
	})

	t.Run("should not return error when a new heroe is verified", func(t *testing.T) {
		sh := []*entity.Superhero{&batman}
		err := util.VerifySuperheroe(sh, superman)

		assert.Nil(t, err)
	})
}

func TestSuperheroeExists(t *testing.T) {
	t.Run("should return true if a heroe exists", func(t *testing.T) {
		hulk := entity.Superhero{
			ID:    "1",
			Name:  "The Hulk",
			Alias: "The Strongest There Is",
		}
		sh := []*entity.Superhero{&hulk}
		resp := util.SuperheroeExists(sh, "1")

		assert.True(t, resp)
	})

	t.Run("should return false if a heroe does not exists", func(t *testing.T) {
		wonderWoman := entity.Superhero{
			ID:    "1",
			Name:  "Wonder Woman",
			Alias: "Princess of Themyscira",
		}
		sh := []*entity.Superhero{&wonderWoman}
		resp := util.SuperheroeExists(sh, "3")

		assert.False(t, resp)
	})
}

func BenchmarkVerifySuperheroe(b *testing.B) {
	sh := []*entity.Superhero{&batman}
	for i := 0; i < b.N; i++ {
		util.VerifySuperheroe(sh, superman)
	}
}

func BenchmarkSuperheroeExists(b *testing.B) {
	sh := []*entity.Superhero{&batman}
	for i := 0; i < b.N; i++ {
		util.SuperheroeExists(sh, "3")
	}
}

func ExampleSuperheroeExists() {
	sh := []*entity.Superhero{&batman}
	fmt.Println(util.SuperheroeExists(sh, "1"))
	fmt.Println(util.SuperheroeExists(sh, "2"))
	//Output:
	//true
	//false
}
