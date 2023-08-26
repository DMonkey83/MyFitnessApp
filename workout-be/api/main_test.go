package api

import (
	"os"
	"testing"
	"time"

	"github.com/DMonkey83/MyFitnessApp/workout-be/config"
	db "github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := config.Config{
		TokenSymmetricKey:   util.GetRandomUsername(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
