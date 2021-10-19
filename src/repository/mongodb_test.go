package repository_test

import (
	"context"
	"log"
	"os"
	"superheroe-api/superheroe-golang-api/src/entity"
	"superheroe-api/superheroe-golang-api/src/repository"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	oid, _ = primitive.ObjectIDFromHex("1")
	batman = entity.Superhero{
		ID:    oid,
		Name:  "Batman",
		Alias: "The Dark Knight",
	}

	user = "mongodb"
	pwd  = "mongodb"
	host = "localhost"
	port = "27017"

	repo repository.Repository
	ctx  = context.TODO()
)

func TestMain(m *testing.M) {
	os.Setenv("MONGO_USER", user)
	os.Setenv("MONGO_PWD", pwd)

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository:   "mongo",
		Tag:          "3.0",
		ExposedPorts: []string{"27017"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"27017": {
				{HostIP: "127.0.0.1", HostPort: port},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err.Error())
	}

	if err = pool.Retry(func() error {
		repo, _ = repository.NewMongoConnection(ctx)
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err.Error())
	}

	defer func() {
		repo.Close(ctx)
		os.Unsetenv("MONGO_USER")
		os.Unsetenv("MONGO_PWD")
	}()

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)

}

func TestAddSuperheroe(t *testing.T) {
	res, err := repo.AddSuperheroe(&batman, ctx)
	assert.NoError(t, err)
	assert.Equal(t, batman.Name, res.Name)
}
