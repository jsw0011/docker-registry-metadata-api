package image_api

import (
	"context"
	"database/sql"
	"docker_image_metadata_api/models"
	opimage "docker_image_metadata_api/restapi/operations/image"
	"docker_image_metadata_api/server/docker_reg_models"
	"docker_image_metadata_api/server/log"
	"os"

	"github.com/lib/pq"

	"docker_image_metadata_api/server/configuration"
	"docker_image_metadata_api/server/database_adapter"

	dockReg "github.com/bloodorangeio/reggie"
	"github.com/go-openapi/runtime/middleware"
)

type ImageAPI struct {
	dbConnector *pq.Connector
	ctx         context.Context
	regClient   *dockReg.Client
}

func GetAPI(cfg *configuration.AppConfiguration) *ImageAPI {
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

	return &ImageAPI{
		dbConnector: conn,
		regClient:   rClientPre,
		ctx:         context.Background(),
	}
}

func (handler *ImageAPI) GetMetadataImages(ctx context.Context, params opimage.GetMetadataImagesParams) middleware.Responder {
	log.Logger.Noticef("[API] GET GetMetadataImage()")

	db := sql.OpenDB(handler.dbConnector)

	rows, qerr := db.Query("select distinct image_path from images")

	if qerr != nil {
		e := database_adapter.DBconnectErr(qerr)
		return opimage.NewGetMetadataImagesInternalServerError().WithPayload(&e) // ERR
	}

	defer db.Close()

	res := models.ListOfImages{}

	// Read result row by row
	var image string
	for rows.Next() {
		err := rows.Scan(&image)

		if err != nil {
			log.Logger.Errorf("Something went wrong while fetching data from DB. Error: %v", err)
			e := models.ErrorResponse{Message: "Fetching data went wrong"}
			return opimage.NewGetMetadataImagesInternalServerError().WithPayload(&e) // ERR
		}

		res = append(res, image)
	}

	return opimage.NewGetMetadataImagesOK().WithPayload(res)
}

func (handler *ImageAPI) DeleteMetadataImage(ctx context.Context, params opimage.DeleteMetadataImageParams) middleware.Responder {
	log.Logger.Noticef("[API] DELETE DeleteMetadataImage()")

	db := sql.OpenDB(handler.dbConnector)

	rows, qerr := db.Query("DELETE FROM images WHERE image_path=$1 RETURNING tag",
		params.Image.ImageName)

	defer db.Close()

	if qerr != nil {
		e := database_adapter.DBconnectErr(qerr)
		return opimage.NewDeleteMetadataImageInternalServerError().WithPayload(&e) // ERR
	}

	log.Logger.Debugf(" DB request without errors")

	// check if anything was returned, when yes, then it was successfull
	var deletedTags []string
	for rows.Next() {
		var tag string
		e := rows.Scan(&tag)
		if e != nil {
			log.Logger.Errorf(" Error occured while scanning selected data from DB: %s", e.Error())
			scanErr := models.ErrorResponse{Message: "Something went wrong while parsing result of database"}
			return opimage.NewDeleteMetadataImageInternalServerError().WithPayload(&scanErr) // ERR
		}
		deletedTags = append(deletedTags, tag)
	}

	if len(deletedTags) > 0 {
		return opimage.NewDeleteMetadataImageOK().WithPayload(deletedTags) // FOUND AND DELETED OK!
	}

	// Not Found
	log.Logger.Debugf(" Image \"%s\" with tags %#v not found in DB", params.Image.ImageName, deletedTags)
	e := models.NotFoundResponse{Message: "Image doesn't exists"}
	return opimage.NewDeleteMetadataImageNotFound().WithPayload(&e)

}
