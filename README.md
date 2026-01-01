Go skeleton Repository

# Important
This project is maintained by a third party and is not affiliated in any way with the Go authors or maintainers.

This repository contains opinionated code intended to speed up the bootstrapping of your application. Nevertheless, this repository is not a complete or standalone application and must not be considered production-ready by design.

Before using any part of this code in production, make sure you fully understand which components are required and whether they fit your project’s requirements.

# Using this repository as template

To use this repository, you need to have the gonew tool installed.

Once installed, you can run the following command:

```bash
gonew github.com/JDarwind/go-skeleton-starter <your-namespace> <path-to-your-project>
```

## Example 
Assume your namespace is:
```bash
github.com/your_user/your_awesome_project
```
You can run:

```bash
gonew github.com/JDarwind/go-skeleton-starter github.com/your_user/your_awesome_project
```
This will create a folder named `your_awesome_project` in the directory where you ran the command.

## Example (Current Directory)
If you want to create the project in the current directory instead of a subfolder, you can run:

```bash
gonew github.com/JDarwind/go-skeleton-starter github.com/your_user/your_awesome_project .
```

# Installing GONEW
To install gonew, follow the installation instructions in the official documentation available at this link:

https://pkg.go.dev/golang.org/x/tools/cmd/gonew

# NOTES
- The official documentation link may change over time, as this project is not maintained or affiliated in any way with the Go authors or maintainers.

- Support for any tool such as `gonew` will not be provided nor guaranteed, since there is no affiliation between this project, or it's author and the gonew authors.

- `gonew` is used in this README only as an example. Other alternatives may exist or be introduced in the future.


# Please note
- the project itself is small and contains a minimal exmaple of usage. 
- Here i reported a quick explaination of the main components and classes in the reopository.


# Documentation: HTTP Server, Requests, Responses, and Databases

**HTTP Server**
- **Config:** use `pkg/config.NewConfigManager` to initialize configuration (optionally pass `ConfigOptions`).
- **Example router:** `internals/routes.NewRouter()` is a simple example using `pkg/network/httpkit` and middlewares in `internals/middlewares`.
- **Init mux:** pass your router to `pkg/server.InitMuxWithRoutes(router)` to get a `*http.ServeMux` that respects the configured prefix.
- **Start server:** read the port from `config.GetConfigManager().GetConfig().ServerConfig.Port` and call `http.ListenAndServe`.

Minimal example:

```go
package main

import (
	"net/http"
	"log"

	"github.com/JDarwind/go-skeleton-starter/pkg/config"
	"github.com/JDarwind/go-skeleton-starter/pkg/server"
	"github.com/JDarwind/go-skeleton-starter/internals/routes"
)

func main() {
	// Initialize configuration (optional ConfigOptions)
	config.NewConfigManager(nil)

	// Build example router
	router := routes.NewRouter()

	// Wrap router with mux that applies the configured prefix
	srv := server.InitMuxWithRoutes(router)

	cfg := config.GetConfigManager().GetConfig()
	addr := ":" + cfg.ServerConfig.Port

	log.Printf("starting server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, srv))
}
```

**Requests and Validation**
- Use `pkg/network/httpkit.ValidateRequest[T]` to bind and validate incoming requests. For non-GET methods it decodes the JSON body into the provided struct; it also binds query params when struct fields contain `query` tags.
- `ValidateRequest` returns the parsed value (as `any`) and a `map[string][]string` of validation errors (or `nil`).
- To provide custom validation messages implement the `httpkit.ValidationMessages` interface on your request struct (define `Messages() map[string]map[string]string`).

Example request struct (see `internals/requests/hello_request.go`):

```go
type HelloRequest struct {
	Name string `json:"name" query:"name" validate:"omitempty,min=3"`
}
```

Typical usage in a handler:

```go
data, errors := httpkit.ValidateRequest(r, &HelloRequest{})
if errors != nil {
	httpkit.NewResponse(w, r).
		Status(http.StatusUnprocessableEntity).
		Error(errors)
	return
}

req := data.(HelloRequest)
```

**Responses**
- Use `httpkit.NewResponse(w, r)` to build JSON (or other content-type) responses.
- Response chain methods: `.Status(code)`, `.ContentType(ct)`, `.Header(key, value)`, `.RemoveHeader(key)`, `.Success(data)`, `.Error(err, status...)`.
- `Success` wraps the data in a standard payload `{result: true, data: ...}`. `Error` returns `{result: false, errors: ...}` and accepts an optional status code.

Example success response:

```go
httpkit.NewResponse(w, r).Status(http.StatusOK).Success(map[string]string{"message": "ok"})
```

Example error response:

```go
httpkit.NewResponse(w, r).Status(http.StatusBadRequest).Error("invalid request")
```

**Middlewares and chaining**
- Use `httpkit.Chain(handler, middlewares...)` to apply middleware functions of type `func(http.Handler) http.Handler` (see `pkg/network/httpkit/middleware.go`).
- Example middleware: `internals/middlewares.Logging` logs requests and execution time.

**Database**
- **Manager:** `pkg/database.GetDatabaseManager()` holds registered connections by name.
- **Drivers:** Postgres and MySQL drivers are provided under `pkg/database_drivers/postgres` and `pkg/database_drivers/mysql`.
- **Registration:** after a successful `Connect()`, drivers register themselves with the `DatabaseManager` (via `AddDatabaseToList`).

Postgres (pgx native) example:

```go
import (
	"log"
	"github.com/JDarwind/go-skeleton-starter/pkg/database"
	"github.com/JDarwind/go-skeleton-starter/pkg/database_drivers/postgres"
)

cfg := postgres.PostgresConfig{
	Username:       "user",
	Password:       "pass",
	Host:           "localhost",
	Port:           "5432",
	Database:       "db",
	ConnectionName: "default",
}

pg := postgres.NewPostgresPGXNative(cfg)
if err := pg.Connect(); err != nil {
	log.Fatalf("postgres connect: %v", err)
}
defer pg.Close()

dbConn, err := database.GetDatabaseManager().GetDatabaseConnection("default")
if err != nil {
	log.Fatal(err)
}

pgDriver := dbConn.(*postgres.PostgresPGXNativeDriver)
pool := pgDriver.GetDriver() // *pgxpool.Pool
_ = pool
```

MySQL example:

```go
import (
	"log"
	"github.com/JDarwind/go-skeleton-starter/pkg/database"
	"github.com/JDarwind/go-skeleton-starter/pkg/database_drivers/mysql"
)

cfg := mysql.MysqlConfig{
	Username:       "user",
	Password:       "pass",
	Host:           "localhost",
	Port:           3306,
	Database:       "db",
	ConnectionName: "default",
}

m := mysql.NewMysqlDriver(cfg)
if err := m.Connect(); err != nil {
	log.Fatalf("mysql connect: %v", err)
}
defer m.Close()

dbConn, err := database.GetDatabaseManager().GetDatabaseConnection("default")
if err != nil {
	log.Fatal(err)
}

mysqlDriver := dbConn.(*mysql.MysqlDriver)
sqlDB := mysqlDriver.GetDriver() // *sql.DB
_ = sqlDB
```

Notes:
- Connection names default to `default` when empty.
- Drivers perform a `Ping` during `Connect()` and only register on successful connection.

---
For concrete examples, see `internals/routes/routes.go`, `internals/middlewares/logger.go` and `internals/requests/hello_request.go`.