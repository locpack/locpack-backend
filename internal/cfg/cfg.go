package cfg

import (
	"placelists-back/pkg/cfg"
)

type Cfg struct {
	Database cfg.Database `env-prefix:"PS_DATABASE_"`
	API      cfg.API      `env-prefix:"PS_API_"`
}
