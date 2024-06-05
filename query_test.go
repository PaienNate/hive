package hive

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind(t *testing.T) {
	var users []User
	err := DB.Find(&users).Error
	assert.NoError(t, err)
	assert.NotEmpty(t, users)
}
