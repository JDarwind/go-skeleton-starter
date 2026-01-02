package integration

import "github.com/JDarwind/go-skeleton-starter/pkg/database_drivers/postgres"

func (s *DatabaseIntegrationSuite) TestPostgresPGXNativeConnection() {
	pg := s.Config.Postgres

	cfg := postgres.PostgresConfig{
		Username: pg.User,
		Password: pg.Password,
		Host:     pg.Host,
		Port:     pg.Port,
		Database: pg.Database,
		Schema:   pg.Schema,

		ConnectionName: "postgres-integration-test",
		DriverType:     postgres.PostgresPGXNative,
	}

	driver, err := postgres.NewPostgresDriver(postgres.PostgresPGXNative, cfg)
	s.Require().NoError(err, "failed to create Postgres driver")

	s.T().Logf(
		"Connecting to Postgres (%s:%s/%s)",
		cfg.Host, cfg.Port, cfg.Database,
	)

	s.Require().NoError(driver.Connect(), "failed to connect to Postgres")
	defer driver.Close()
}
