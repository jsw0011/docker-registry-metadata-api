package database_adapter

import (
	"docker_image_metadata_api/models"
	"docker_image_metadata_api/server/configuration"
	"docker_image_metadata_api/server/log"
	"fmt"
)

func DBconnectErr(err error) models.ErrorResponse {
	log.Logger.Errorf("Unable to connect to the DB: %v", err)
	return models.ErrorResponse{Message: "Communication error between API and internal DB"}
}

func GetDBConnStr(conf *configuration.ConfDB) string {
	sslmode := "disable"
	if conf.SSL {
		sslmode = "require"
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", conf.User, conf.Pass, conf.Host, conf.Port, conf.Database, sslmode)
}
