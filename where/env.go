package where

import (
	"strings"

	"github.com/priuatus/submind/app"
)

// EnvConfigPath is the environment variable name for the config path
var EnvConfigPath = strings.ToUpper(app.Name) + "_CONFIG_PATH"
