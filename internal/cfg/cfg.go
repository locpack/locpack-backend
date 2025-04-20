package cfg

import (
	"locpack-backend/pkg/cfg"
)

type Cfg struct {
	Database cfg.Database `env-prefix:"LP_DATABASE_"`
	API      cfg.API      `env-prefix:"LP_API_"`
}
