package integration

import (
	"testing"
	"time"

	testconfig "github.com/JDarwind/go-skeleton-starter/tests/config"
	"github.com/stretchr/testify/suite"
)

type DatabaseIntegrationSuite struct {
	suite.Suite
	Config testconfig.TestConfig
}

func (s *DatabaseIntegrationSuite) SetupSuite() {
	s.T().Log("Loading shared test configuration")

	cfg, err := testconfig.Load()
	s.Require().NoError(err, "failed to load test configuration")

	s.Config = cfg

	s.T().Log("Waiting for databases to be ready")
	time.Sleep(5 * time.Second)
}

func TestDatabaseIntegrationSuite(t *testing.T) {
	suite.Run(t, new(DatabaseIntegrationSuite))
}
