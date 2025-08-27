package command

import (
	"github.com/dmitriy-zverev/rss-cli/internal/config"
	"github.com/dmitriy-zverev/rss-cli/internal/database"
)

type State struct {
	Cfg *config.Config
	Db  *database.Queries
}
