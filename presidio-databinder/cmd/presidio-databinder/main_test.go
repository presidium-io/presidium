package main

import (
	"context"
	"testing"

	message_types "github.com/presid-io/presidio-genproto/golang"
	"github.com/stretchr/testify/assert"
)

func TestIsDatabase(t *testing.T) {
	assert.True(t, isDatabase("mssql"))
	assert.True(t, isDatabase("mysql"))
	assert.True(t, isDatabase("oracle"))
	assert.True(t, isDatabase("postgres"))
	assert.True(t, isDatabase("sqlite3"))
	assert.False(t, isDatabase("kafka"))
}

func TestIsCloudStorage(t *testing.T) {
	assert.True(t, isCloudStorage("s3"))
	assert.True(t, isCloudStorage("azureblob"))
	assert.True(t, isCloudStorage("googlestorage"))
	assert.True(t, isCloudStorage("postgres"))
}

func TestDataBinderInit(t *testing.T) {
	var s *server
	databinderTemplate := &message_types.DatabinderTemplate{}

	// validate databinder is initialized
	_, err := s.Init(context.Background(), databinderTemplate)
	assert.EqualError(t, err, "databinderTemplate must me set")

	databinder := [](*message_types.Databinder){
		&message_types.Databinder{
			BindType: "sqlite3",
		},
	}

	// validate connection string is set
	databinderTemplate.Databinder = databinder
	_, err = s.Init(context.Background(), databinderTemplate)
	assert.EqualError(t, err, "connectionString var must me set")

	// databinders array is empty
	assert.Empty(t, databinderArray)

	databinder[0].DbConfig = &message_types.DBConfig{
		ConnectionString: "./test.db?cache=shared&mode=rwc",
	}
	databinderTemplate.Databinder = databinder
	s.Init(context.Background(), databinderTemplate)

	// validate databinder was created successfully
	assert.Equal(t, len(databinderArray), 1)
}
