package db

import (
	"fmt"
	"github.com/dbl90/airfreight"
	"github.com/dbl90/airfreight/internal/models"
	"github.com/dbl90/airfreight/internal/models/config"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const username = "username"
const password = "password"
const dbname = "wms"

func TestDBClient_GetMawbByNumber(t *testing.T) {
	container, err := gnomock.Start(postgres.Preset(postgres.WithUser(username, password), postgres.WithDatabase(dbname)), gnomock.WithTimeout(time.Minute*2))

	if err != nil {
		t.Fatal(err.Error())
	}

	t.Cleanup(func() { _ = gnomock.Stop(container) })

	dbClient, err := NewDBClient(&config.DbConfig{
		Host:     container.Host,
		Port:     container.DefaultPort(),
		User:     username,
		Password: password,
		DBName:   dbname,
	})
	if err != nil {
		assert.NoError(t, err)
		t.Fatal(err.Error())
	}

	connectionStr := fmt.Sprintf("postgres://%s:%d/%s?sslmode=enable", container.Host, container.DefaultPort(), dbname)

	err = airfreight.MigrateDb(connectionStr)
	if err != nil {
		return
	}

	actualMAWBNum := "1234567890"
	insert := dbClient.Insert(&models.Mawb{Number: actualMAWBNum})
	assert.NoError(t, insert.Error)

	retreivedMawbNum := dbClient.GetMawbByNumber(actualMAWBNum)
	assert.Equal(t, actualMAWBNum, retreivedMawbNum.Number)
}
