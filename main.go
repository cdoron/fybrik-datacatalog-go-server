/*
 * Data Catalog Service - Asset Details
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	api "github.com/cdoron/datacatalog-go/api"
)

func main() {
	log.Printf("Server started")

	DefaultApiService := api.NewDefaultApiService()
	DefaultApiController := api.NewDefaultApiController(DefaultApiService)

	router := api.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
