package database

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testConn = "test"

var testConnURI string

func TestMain(m *testing.M) {
	//Set the connection string to your mongodb instance
	testConnURI = "test@localhost"
	retCode := m.Run()
	os.Exit(retCode)
}

func TestNew(t *testing.T) {
	mongo := New()
	assert.NotNil(t, mongo)
	assert.Nil(t, mongo.GetClient())
}

func TestGetConnectinoURI(t *testing.T) {
	os.Setenv(connectionURI, testConn)
	mongo := New()
	assert.Equal(t, mongo.getConnectionURI(), testConn)
	os.Unsetenv(testConn)
}

func TestConnect(t *testing.T) {
	os.Setenv(connectionURI, testConnURI)
	mongo := New()
	ctx := context.TODO()
	err := mongo.Connect(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, mongo.GetClient())
	coll := mongo.GetCollection(Database, testConn)
	assert.NotNil(t, coll)
	err = mongo.Disconnect(ctx)
	assert.Nil(t, err)
	os.Unsetenv(connectionURI)
	err = mongo.Connect(ctx)
	assert.NotNil(t, err)
	err = mongo.Disconnect(ctx)
	assert.Equal(t, err == ErrClientNotInitialized, true)
}
