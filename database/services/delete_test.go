package services

import (
	"testing"

	"github.com/mesg-foundation/core/service"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	hash, _ := Save(&service.Service{
		Name: "TestDelete",
	})
	err := Delete(hash)
	require.Nil(t, err)
}
