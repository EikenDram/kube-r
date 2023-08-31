package database

import (
	"github.com/EikenDram/kube-r/manager/config"

	"database/sql"
	"fmt"

	_ "github.com/ibmdb/go_ibm_db"
)

func Init(cfg *config.ConfigStruct) (*sql.DB, error) {
	con := fmt.Sprintf("HOSTNAME=%s;DATABASE=%s;PORT=%s;UID=%s;PWD=%s", cfg.Database.Host, cfg.Database.Name, cfg.Database.Port, cfg.Database.User, cfg.Database.Pass)
	return sql.Open("go_ibm_db", con)
}
