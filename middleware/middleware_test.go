package middleware

import (
	"testing"
	"time"

	"github.com/Philip-21/proj1/config"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func newTestServer(t *testing.T, store *gorm.DB) *Server {
	config := config.TokenConfig{
		TokenSymmetricKey:   "avsvbfhtyruterwfhrytwiquytruyhtit",
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}
