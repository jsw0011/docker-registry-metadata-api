package metadata_api

import (
	"context"
	"database/sql"
	"docker_image_metadata_api/models"
	"docker_image_metadata_api/restapi/operations/metadata"
	opmetadata "docker_image_metadata_api/restapi/operations/metadata"
	"docker_image_metadata_api/server/configuration"
	"docker_image_metadata_api/server/database_adapter"
	"docker_image_metadata_api/server/log"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/lib/pq"
)

type MetadataAPI struct {
	dbConnector *pq.Connector
	ctx         context.Context
}

func GetAPI(cfg *configuration.AppConfiguration) *MetadataAPI {
	dbConStr := database_adapter.GetDBConnStr(&cfg.Database)
	conn, err := pq.NewConnector(dbConStr)
	if err != nil {
		log.Logger.Criticalf("Unable to create DB connector")
		os.Exit(6)
	}
	return &MetadataAPI{
		dbConnector: conn,
		ctx:         context.Background(),
	}
}

func (handler *MetadataAPI) PostMetadataImageTagKeys(ctx context.Context, params metadata.PostMetadataImageTagKeysParams) middleware.Responder {
	log.Logger.Noticef("[API] POST PostMetadataImageTagKeys()")

	db := sql.OpenDB(handler.dbConnector)

	_, qerr := db.Query("INSERT INTO  metadata (image_path, tag, met_key, met_default_value) VALUES ($1,$2,$3,$4)",
		params.Image.ImageName, params.Image.TagName, params.Image.Key, params.Image.Value)

	defer db.Close()

	if qerr != nil {
		er, ok := qerr.(*pq.Error)
		if ok {
			// https://www.postgresql.org/docs/9.3/errcodes-appendix.html
			log.Logger.Debugf(" PG ercode: %s", er.Code)
		}
		if ok && er.Code == "23503" { // foreign_key_violation
			r := models.NotFoundResponse{Message: "Image or tag does not exists in database"}
			return opmetadata.NewPostMetadataImageTagKeysNotFound().WithPayload(&r)
		} else if ok && er.Code == "23505" { // unique_violation
			r := models.ConflictResponse{Message: "Key already exsísts"}
			return opmetadata.NewPostMetadataImageTagKeysConflict().WithPayload(&r)
		} else {
			e := database_adapter.DBconnectErr(qerr)
			return opmetadata.NewPostMetadataImageTagKeysInternalServerError().WithPayload(&e) // ERR
		}
	}

	return opmetadata.NewPostMetadataImageTagKeysCreated()
}

func (handler *MetadataAPI) PostMetadataImageTag(ctx context.Context, params metadata.PostMetadataImageTagParams) middleware.Responder {
	log.Logger.Noticef("[API] POST PostMetadataImageTag()")

	db := sql.OpenDB(handler.dbConnector)

	rows, qerr := db.Query("SELECT met_key as Key, met_default_value as Value FROM metadata WHERE image_path=$1 AND tag=$2",
		params.Image.ImageName, params.Image.TagName)

	defer db.Close()

	if qerr != nil {
		e := database_adapter.DBconnectErr(qerr)
		return opmetadata.NewPostMetadataImageTagInternalServerError().WithPayload(&e) // ERR
	}

	keys := models.Metadata{}
	for rows.Next() {
		var keyPair models.MetadataObj
		e := rows.Scan(&keyPair.Key, &keyPair.Value)
		if e != nil {
			log.Logger.Errorf(" Error occured while scanning selected data from DB: %s", e.Error())
			scanErr := models.ErrorResponse{Message: "Something went wrong while parsing result of database"}
			return opmetadata.NewPostMetadataImageTagInternalServerError().WithPayload(&scanErr) // ERR
		}
		keys = append(keys, &keyPair)
	}

	return opmetadata.NewPostMetadataImageTagOK().WithPayload(keys)
}

func (handler *MetadataAPI) PutMetadataImageTagKeys(ctx context.Context, params metadata.PutMetadataImageTagKeysParams) middleware.Responder {
	log.Logger.Noticef("[API] PUT PutMetadataImageTagKeys()")

	db := sql.OpenDB(handler.dbConnector)

	rows, qerr := db.Query("UPDATE  metadata SET met_default_value=$1 WHERE image_path=$2 AND tag=$3 AND met_key=$4 RETURNING met_key",
		params.Image.Value, params.Image.ImageName, params.Image.TagName, params.Image.Key)

	defer db.Close()

	if qerr != nil {
		e := database_adapter.DBconnectErr(qerr)
		return opmetadata.NewPostMetadataImageTagKeysInternalServerError().WithPayload(&e) // ERR
	}
	// check if anything was returned, when yes, then it was successfull
	if rows.Next() {
		return opmetadata.NewPutMetadataImageTagKeysOK() // updated succesfully
	}

	// Not Found
	e := models.NotFoundResponse{Message: "Image, tag, or key doesn't exists"}
	return opmetadata.NewPutMetadataImageTagKeysNotFound().WithPayload(&e)
}

func (handler *MetadataAPI) DeleteMetadataImageTagKeys(ctx context.Context, params metadata.DeleteMetadataImageTagKeysParams) middleware.Responder {
	log.Logger.Noticef("[API] DELETE DeleteMetadataImageTagKeys()")

	db := sql.OpenDB(handler.dbConnector)

	rows, qerr := db.Query("DELETE FROM  metadata WHERE image_path=$1 AND tag=$2 AND met_key=$3 RETURNING met_key",
		params.Image.ImageName, params.Image.TagName, params.Image.Key)

	defer db.Close()

	if qerr != nil {
		e := database_adapter.DBconnectErr(qerr)
		return opmetadata.NewDeleteMetadataImageTagKeysInternalServerError().WithPayload(&e) // ERR
	}
	// check if anything was returned, when yes, then it was successfull
	if rows.Next() {
		return opmetadata.NewDeleteMetadataImageTagKeysOK() // updated succesfully
	}

	// Not Found
	e := models.NotFoundResponse{Message: "Image, tag, or key doesn't exists"}
	return opmetadata.NewDeleteMetadataImageTagKeysNotFound().WithPayload(&e)
}
