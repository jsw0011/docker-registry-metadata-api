package image_tag_api

import (
	"context"
	"database/sql"
	"docker_image_metadata_api/models"
	"docker_image_metadata_api/restapi/operations/image_tag"
	"docker_image_metadata_api/server/configuration"
	"docker_image_metadata_api/server/database_adapter"
	"docker_image_metadata_api/server/docker_reg_models"
	"docker_image_metadata_api/server/log"
	"os"

	dockReg "github.com/bloodorangeio/reggie"
	"github.com/go-openapi/runtime/middleware"
	"github.com/lib/pq"
)

type ImageTagAPI struct {
	dbConnector *pq.Connector
	regClient   *dockReg.Client
	ctx         context.Context
}

func GetAPI(cfg *configuration.AppConfiguration) *ImageTagAPI {
	dbConStr := database_adapter.GetDBConnStr(&cfg.Database)
	conn, err := pq.NewConnector(dbConStr)
	if err != nil {
		log.Logger.Criticalf("Unable to create DB connector")
		os.Exit(6)
	}

	// create client for registry
	rClientPre := docker_reg_models.InitDockRegConnection(&cfg.Registry)

	if rClientPre == nil {
		os.Exit(7)
	}

	return &ImageTagAPI{
		dbConnector: conn,
		regClient:   rClientPre,
		ctx:         context.Background(),
	}
}

func (handler *ImageTagAPI) PostMetadata(ctx context.Context, params image_tag.PostMetadataParams) middleware.Responder {
	log.Logger.Noticef("[API] POST PostMetadataImage()")

	req := handler.regClient.NewRequest(dockReg.GET, "/v2/<name>/manifests/<reference>",
		dockReg.WithName(params.Image.ImageName), dockReg.WithReference(params.Image.TagName))

	req.Header.Set("Accept", "application/vnd.docker.distribution.manifest.v2+json") //v2 schema of manifest

	res, dockErr := handler.regClient.Do(req)
	if dockErr != nil {
		errMsg := models.ErrorResponse{Message: "Unable to retrieve information from docker registry!"}
		log.Logger.Errorf(" %s: %v", errMsg.Message, dockErr)

		return image_tag.NewPostMetadataInternalServerError().WithPayload(&errMsg) // ERR
	}

	if res.StatusCode() == 404 {
		errMsg := models.FailedDependency{Message: "Image or tag not found in docker repository!"}
		log.Logger.Noticef("Image \"%s\" or tag \"%s\" not found in registry", params.Image.ImageName, params.Image.TagName)
		return image_tag.NewPostMetadataFailedDependency().WithPayload(&errMsg) // ERR
	} else if !res.IsSuccess() {
		errMsg := models.ErrorResponse{Message: "Unable to retrieve information from docker registry!"}
		log.Logger.Errorf(" %s: %v", errMsg.Message, dockErr)

		return image_tag.NewPostMetadataInternalServerError().WithPayload(&errMsg) // ERR
	}

	log.Logger.Debugf("Image \"%s\" with tag \"%s\" in docker registry exists", params.Image.ImageName, params.Image.TagName)

	// get SHA256 of the manifest
	imageDigest := res.Header().Get("Docker-Content-Digest")

	// Image and tag exists, lets try insert value into database
	log.Logger.Debugf("Registering metadata into DB")

	db := sql.OpenDB(handler.dbConnector)

	_, qerr := db.Query("INSERT INTO images (image_path, tag, image_sha) VALUES ($1, $2, $3)",
		params.Image.ImageName, params.Image.TagName, imageDigest)

	defer db.Close()

	if qerr != nil {
		er, ok := qerr.(*pq.Error)
		if ok {
			log.Logger.Debugf(" PG ercode: %s", er.Code)
		}
		if ok && er.Code == "23505" {
			errMsg := models.ConflictResponse{Message: "Image with tag already exists in database"}
			return image_tag.NewPostMetadataConflict().WithPayload(&errMsg) // ERR
		} else {
			e := database_adapter.DBconnectErr(qerr)
			return image_tag.NewPostMetadataInternalServerError().WithPayload(&e) // ERR
		}
	}

	log.Logger.Debugf("Metadata registered")

	return image_tag.NewPostMetadataCreated()

}

func (handler *ImageTagAPI) PostMetadataImage(ctx context.Context, params image_tag.PostMetadataImageParams) middleware.Responder {
	log.Logger.Noticef("[API] POST GetMetadataImageImagename()")

	db := sql.OpenDB(handler.dbConnector)

	log.Logger.Debugf(" Executing DB query")

	rows, qerr := db.Query("SELECT tag FROM images WHERE image_path=$1", params.Image.ImageName)

	defer db.Close()

	if qerr != nil {
		e := database_adapter.DBconnectErr(qerr)
		return image_tag.NewPostMetadataImageInternalServerError().WithPayload(&e) // ERR
	}

	var tagsarr []string
	var scanErr error

	for rows.Next() {
		var tag string
		scanErr = rows.Scan(&tag)
		if (scanErr) != nil {
			log.Logger.Errorf("Something went wrong while fetching data from DB. Error: %v", scanErr)
			e := models.ErrorResponse{Message: "Fetching data went wrong"}
			return image_tag.NewPostMetadataImageInternalServerError().WithPayload(&e) // ERR
		}
		tagsarr = append(tagsarr, tag)
	}

	if len(tagsarr) == 0 {
		resp := models.NotFoundResponse{Message: "Image does not exists in database"}
		return image_tag.NewPostMetadataImageNotFound().WithPayload(&resp)
	}

	return image_tag.NewPostMetadataImageOK().WithPayload(tagsarr)
}

func (handler *ImageTagAPI) DeleteMetadataImageTag(ctx context.Context, params image_tag.DeleteMetadataImageTagParams) middleware.Responder {
	log.Logger.Noticef("[API] DELETE DeleteMetadataImageTag()")

	db := sql.OpenDB(handler.dbConnector)

	log.Logger.Debugf(" Executing DB query")

	rows, qerr := db.Query("DELETE FROM images WHERE image_path=$1 AND tag=$2 RETURNING tag", params.Image.ImageName, params.Image.TagName)

	defer db.Close()

	if qerr != nil {
		e := database_adapter.DBconnectErr(qerr)
		return image_tag.NewDeleteMetadataImageTagInternalServerError().WithPayload(&e) // ERR
	}

	// check if anything was returned, when yes, then it was successfull
	if rows.Next() {
		return image_tag.NewDeleteMetadataImageTagOK() // FOUND AND DELETED OK!
	}

	// Not Found
	e := models.NotFoundResponse{Message: "Image or tag doesn't exists"}
	return image_tag.NewDeleteMetadataImageTagNotFound().WithPayload(&e)

}
