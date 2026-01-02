# Configuration Manager
The Configuration Manager is the core of the configuration management of this project.

The Configuration Manager is the component responsible for loading and providing configuration values across the application.

# Options
At the very heart tof the configuration manager is present the `ConfigOptions` struct.

``` go
type ConfigOptions struct {
	EnvironmentFileLocaltion string
	ApplicationConfigs any
}
```
Here is a brief explanation of the possible parameters:

- EnvironmentFileLocation: this is the location of the .env file (including the name of the file itself).
If it is not provided or not set, it is assumed that a .env file is present in the root of the project.
Please note: if the file is not present, or if the application is running in a Docker environment where environment variables are set without a file, this property is ignored.

- ApplicationConfigs: this property allows you to assign any value to it.
It is meant to allow the creation of custom settings from a user project without changing the core code.


# Usage
To use the Configuration Manager, it is necessary to obtain an instance of it.
Moreover, you must call the following function at the entry point of your application:
```go
config.NewConfigManager(ConfigOptions or nil)
``` 
After this first initialization, it is possible to call:
```go
config.GetConfigManager().GetConfig()
```
to retrieve the configuration object.


# Example (Default settings)

Here is an example of usage:

```go
package main

import (
	"log"

	"github.com/JDarwind/go-skeleton-starter/pkg/config"
)

func main() {
	config.NewConfigManager(nil)

	cfg := config.GetConfigManager().GetConfig()
	addr := ":" + cfg.ServerConfig.Port

	log.Printf("starting server on %s", addr)
}
```

# Example (Custom settings)

```go
package main

import (
	"log"

	"github.com/JDarwind/go-skeleton-starter/pkg/config"
)

type MyCustomConfig struct {
	Message string
}

func main() {
	options := &config.ConfigOptions{
		EnvironmentFileLocation: ".env.example",
		ApplicationConfigs: MyCustomConfig{
			Message: "hello",
		},
	}

	config.NewConfigManager(options)

	cfg := config.GetConfigManager().GetConfig()

	applicationConfig, ok := cfg.ApplicationConfig.(MyCustomConfig)
	if !ok {
		log.Fatal("config not valid")
	}

	log.Printf("Message is %s", applicationConfig.Message)

	addr := ":" + cfg.ServerConfig.Port
	log.Printf("starting server on %s", addr)
}
```