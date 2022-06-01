/*
 * Data Catalog Service - Asset Details
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"context"
	"net/http"
)



// DefaultApiRouter defines the required methods for binding the api requests to a responses for the DefaultApi
// The DefaultApiRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultApiServicer to perform the required actions, then write the service results to the http response.
type DefaultApiRouter interface { 
	CreateAsset(http.ResponseWriter, *http.Request)
	DeleteAsset(http.ResponseWriter, *http.Request)
	GetAssetInfo(http.ResponseWriter, *http.Request)
	UpdateAsset(http.ResponseWriter, *http.Request)
}


// DefaultApiServicer defines the api actions for the DefaultApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultApiServicer interface { 
	CreateAsset(context.Context, string, CreateAssetRequest) (ImplResponse, error)
	DeleteAsset(context.Context, string, DeleteAssetRequest) (ImplResponse, error)
	GetAssetInfo(context.Context, string, GetAssetRequest) (ImplResponse, error)
	UpdateAsset(context.Context, string, UpdateAssetRequest) (ImplResponse, error)
}
