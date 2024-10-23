package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/ru/method/database
// The methods in this section provide access to the VK database of educational institutions.
//	Access to data is free and does not require authorization, however, the number of requests from one IP address may be limited;
//	if you need to make a large number of requests, it is recommended to perform requests from the client side using JSONP.

// DatabaseGetChairsRequest defines the request for database.getChairs
//
// The method returns a list of university departments by the specified faculty.
// Doc: https://dev.vk.com/method/database.getChairs
type DatabaseGetChairsRequest struct {
	BaseRequest
}

// NewDatabaseGetChairsRequest creates a new request for database.getChairs
func NewDatabaseGetChairsRequest(a *api.API, actor actor.Actor) *DatabaseGetChairsRequest {
	return &DatabaseGetChairsRequest{*NewMethodBaseRequest(a, actor, "database.getChairs")}
}

// Exec executes the request and unmarshals the response into DatabaseGetChairsResponse
func (r *DatabaseGetChairsRequest) Exec(ctx context.Context) (response response.DatabaseGetChairsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DatabaseGetCitiesRequest defines the request for database.getCities
//
// The method returns a list of cities.
// Doc: https://dev.vk.com/method/database.getCities
type DatabaseGetCitiesRequest struct {
	BaseRequest
}

// NewDatabaseGetCitiesRequest creates a new request for database.getCities
func NewDatabaseGetCitiesRequest(a *api.API, actor actor.Actor) *DatabaseGetCitiesRequest {
	return &DatabaseGetCitiesRequest{*NewMethodBaseRequest(a, actor, "database.getCities")}
}

// Exec executes the request and unmarshals the response into DatabaseGetCitiesResponse
func (r *DatabaseGetCitiesRequest) Exec(ctx context.Context) (response response.DatabaseGetCitiesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DatabaseGetCitiesByIdRequest defines the request for database.getCitiesById
//
// The method returns information about cities and regions by their IDs.
// Doc: https://dev.vk.com/method/database.getCitiesById
type DatabaseGetCitiesByIdRequest struct {
	BaseRequest
}

// NewDatabaseGetCitiesByIdRequest creates a new request for database.getCitiesById
func NewDatabaseGetCitiesByIdRequest(a *api.API, actor actor.Actor) *DatabaseGetCitiesByIdRequest {
	return &DatabaseGetCitiesByIdRequest{*NewMethodBaseRequest(a, actor, "database.getCitiesById")}
}

// Exec executes the request and unmarshals the response into DatabaseGetCitiesByIdResponse
func (r *DatabaseGetCitiesByIdRequest) Exec(ctx context.Context) (response response.DatabaseGetCitiesByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DatabaseGetCountriesRequest defines the request for database.getCountries
//
// The method returns a list of countries.
// Doc: https://dev.vk.com/method/database.getCountries
type DatabaseGetCountriesRequest struct {
	BaseRequest
}

// NewDatabaseGetCountriesRequest creates a new request for database.getCountries
func NewDatabaseGetCountriesRequest(a *api.API, actor actor.Actor) *DatabaseGetCountriesRequest {
	return &DatabaseGetCountriesRequest{*NewMethodBaseRequest(a, actor, "database.getCountries")}
}

// Exec executes the request and unmarshals the response into DatabaseGetCountriesResponse
func (r *DatabaseGetCountriesRequest) Exec(ctx context.Context) (response response.DatabaseGetCountriesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DatabaseGetCountriesByIdRequest defines the request for database.getCountriesById
//
// The method returns information about countries by their IDs.
// Doc: https://dev.vk.com/method/database.getCountriesById
type DatabaseGetCountriesByIdRequest struct {
	BaseRequest
}

// NewDatabaseGetCountriesByIdRequest creates a new request for database.getCountriesById
func NewDatabaseGetCountriesByIdRequest(a *api.API, actor actor.Actor) *DatabaseGetCountriesByIdRequest {
	return &DatabaseGetCountriesByIdRequest{*NewMethodBaseRequest(a, actor, "database.getCountriesById")}
}

// Exec executes the request and unmarshals the response into DatabaseGetCountriesByIdResponse
func (r *DatabaseGetCountriesByIdRequest) Exec(ctx context.Context) (response response.DatabaseGetCountriesByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DatabaseGetFacultiesRequest defines the request for database.getFaculties
//
// The method returns a list of faculties.
// Doc: https://dev.vk.com/method/database.getFaculties
type DatabaseGetFacultiesRequest struct {
	BaseRequest
}

// NewDatabaseGetFacultiesRequest creates a new request for database.getFaculties
func NewDatabaseGetFacultiesRequest(a *api.API, actor actor.Actor) *DatabaseGetFacultiesRequest {
	return &DatabaseGetFacultiesRequest{*NewMethodBaseRequest(a, actor, "database.getFaculties")}
}

// Exec executes the request and unmarshals the response into DatabaseGetFacultiesResponse
func (r *DatabaseGetFacultiesRequest) Exec(ctx context.Context) (response response.DatabaseGetFacultiesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DatabaseGetMetroStationsRequest defines the request for database.getMetroStations
//
// The method returns a list of metro stations.
// Doc: https://dev.vk.com/method/database.getMetroStations
type DatabaseGetMetroStationsRequest struct {
	BaseRequest
}

// NewDatabaseGetMetroStationsRequest creates a new request for database.getMetroStations
func NewDatabaseGetMetroStationsRequest(a *api.API, actor actor.Actor) *DatabaseGetMetroStationsRequest {
	return &DatabaseGetMetroStationsRequest{*NewMethodBaseRequest(a, actor, "database.getMetroStations")}
}

// Exec executes the request and unmarshals the response into DatabaseGetMetroStationsResponse
func (r *DatabaseGetMetroStationsRequest) Exec(ctx context.Context) (response response.DatabaseGetMetroStationsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DatabaseGetMetroStationsByIdRequest defines the request for database.getMetroStationsById
//
// The method returns information about one or more metro stations by their IDs.
// Doc: https://dev.vk.com/method/database.getMetroStationsById
type DatabaseGetMetroStationsByIdRequest struct {
	BaseRequest
}

// NewDatabaseGetMetroStationsByIdRequest creates a new request for database.getMetroStationsById
func NewDatabaseGetMetroStationsByIdRequest(a *api.API, actor actor.Actor) *DatabaseGetMetroStationsByIdRequest {
	return &DatabaseGetMetroStationsByIdRequest{*NewMethodBaseRequest(a, actor, "database.getMetroStationsById")}
}

// Exec executes the request and unmarshals the response into DatabaseGetMetroStationsByIdResponse
func (r *DatabaseGetMetroStationsByIdRequest) Exec(ctx context.Context) (response response.DatabaseGetMetroStationsByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DatabaseGetRegionsRequest defines the request for database.getRegions
//
// The method returns a list of regions.
// Doc: https://dev.vk.com/method/database.getRegions
type DatabaseGetRegionsRequest struct {
	BaseRequest
}

// NewDatabaseGetRegionsRequest creates a new request for database.getRegions
func NewDatabaseGetRegionsRequest(a *api.API, actor actor.Actor) *DatabaseGetRegionsRequest {
	return &DatabaseGetRegionsRequest{*NewMethodBaseRequest(a, actor, "database.getRegions")}
}

// Exec executes the request and unmarshals the response into DatabaseGetRegionsResponse
func (r *DatabaseGetRegionsRequest) Exec(ctx context.Context) (response response.DatabaseGetRegionsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DatabaseGetSchoolClassesRequest defines the request for database.getSchoolClasses
//
// The method returns a list of classes typical for schools in a specified country.
// Doc: https://dev.vk.com/method/database.getSchoolClasses
type DatabaseGetSchoolClassesRequest struct {
	BaseRequest
}

// NewDatabaseGetSchoolClassesRequest creates a new request for database.getSchoolClasses
func NewDatabaseGetSchoolClassesRequest(a *api.API, actor actor.Actor) *DatabaseGetSchoolClassesRequest {
	return &DatabaseGetSchoolClassesRequest{*NewMethodBaseRequest(a, actor, "database.getSchoolClasses")}
}

// Exec executes the request and unmarshals the response into DatabaseGetSchoolClassesResponse
func (r *DatabaseGetSchoolClassesRequest) Exec(ctx context.Context) (response response.DatabaseGetSchoolClassesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DatabaseGetSchoolsRequest defines the request for database.getSchools
//
// The method returns a list of schools.
// Doc: https://dev.vk.com/method/database.getSchools
type DatabaseGetSchoolsRequest struct {
	BaseRequest
}

// NewDatabaseGetSchoolsRequest creates a new request for database.getSchools
func NewDatabaseGetSchoolsRequest(a *api.API, actor actor.Actor) *DatabaseGetSchoolsRequest {
	return &DatabaseGetSchoolsRequest{*NewMethodBaseRequest(a, actor, "database.getSchools")}
}

// Exec executes the request and unmarshals the response into DatabaseGetSchoolsResponse
func (r *DatabaseGetSchoolsRequest) Exec(ctx context.Context) (response response.DatabaseGetSchoolsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DatabaseGetUniversitiesRequest defines the request for database.getUniversities
//
// The method returns a list of universities.
// Doc: https://dev.vk.com/method/database.getUniversities
type DatabaseGetUniversitiesRequest struct {
	BaseRequest
}

// NewDatabaseGetUniversitiesRequest creates a new request for database.getUniversities
func NewDatabaseGetUniversitiesRequest(a *api.API, actor actor.Actor) *DatabaseGetUniversitiesRequest {
	return &DatabaseGetUniversitiesRequest{*NewMethodBaseRequest(a, actor, "database.getUniversities")}
}

// Exec executes the request and unmarshals the response into DatabaseGetUniversitiesResponse
func (r *DatabaseGetUniversitiesRequest) Exec(ctx context.Context) (response response.DatabaseGetUniversitiesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
