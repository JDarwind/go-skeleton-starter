package postgres

import (
	"fmt"

	"github.com/JDarwind/go-skeleton-starter/pkg/database"
)



func NewPostgresDriver(driverType PostgresDriverType, cfg PostgresConfig) (database.Database, error) {
	normalizeConfig(&cfg)
	switch driverType {
		case PostgresPGXNative:
			return NewPostgresPGXNative(cfg), nil
		case PostgresPGXStdlib:
			return NewPostgresPGXStdlib(cfg), nil
		default:
			return nil, fmt.Errorf("unsupported postgres driver: %s", driverType)
	}
}