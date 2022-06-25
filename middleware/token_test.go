package middleware

import (
	"testing"
	"time"

	"github.com/Philip-21/proj1/config"

	"github.com/Philip-21/proj1/models"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store *models.ContentUser) { //*handlers.Repository {
	config := config.Envconfig{
		TokenSymmetricKey:   "avsvbfhtyruterwfhrytwiquytruyhtit",
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}
