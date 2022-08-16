package docker_reg_models

import (
	"docker_image_metadata_api/server/configuration"
	"docker_image_metadata_api/server/log"
	"strconv"

	dockReg "github.com/bloodorangeio/reggie"
)

func InitDockRegConnection(cfg *configuration.ConfRegistry) *dockReg.Client {
	// construct url for registry
	urlProt := "http"
	if cfg.SSL {
		urlProt = "https"
	}

	regUrl := urlProt + "://" + cfg.Host + ":" + strconv.FormatUint(uint64(cfg.Port), 10)

	// create client for registry
	rClientPre, regErr := dockReg.NewClient(
		regUrl,
		dockReg.WithUsernamePassword(cfg.User, cfg.Pass),
	)

	if regErr != nil {
		log.Logger.Criticalf("Error occured while creating docker registry client: %b", regErr)
		return nil
	}
	return rClientPre
}
