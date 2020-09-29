package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMongo(t *testing.T) {
	url := "mongo://localhost:27017/"
	actual, err := GetMongo(url)
	if err != nil {
		t.Error(err)
	}
	expected := MongoDB{ConnectionString: "mongodb://localhost:27017"}
	assert.Equal(t, expected, *actual)
}
