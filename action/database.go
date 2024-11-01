package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Database Doc: https://dev.vk.com/ru/method/database
type Database struct {
	BaseAction
}

// GetChairs Doc: https://dev.vk.com/ru/method/database.getChairs
func (d *Database) GetChairs(actor actor.Actor) *request.DatabaseGetChairsRequest {
	return request.NewDatabaseGetChairsRequest(d.api, actor)
}

// GetCities Doc: https://dev.vk.com/ru/method/database.getCities
func (d *Database) GetCities(actor actor.Actor) *request.DatabaseGetCitiesRequest {
	return request.NewDatabaseGetCitiesRequest(d.api, actor)
}

// GetCitiesById Doc: https://dev.vk.com/ru/method/database.getCitiesById
func (d *Database) GetCitiesById(actor actor.Actor) *request.DatabaseGetCitiesByIdRequest {
	return request.NewDatabaseGetCitiesByIdRequest(d.api, actor)
}

// GetCountries Doc: https://dev.vk.com/ru/method/database.getCountries
func (d *Database) GetCountries(actor actor.Actor) *request.DatabaseGetCountriesRequest {
	return request.NewDatabaseGetCountriesRequest(d.api, actor)
}

// GetCountriesById Doc: https://dev.vk.com/ru/method/database.getCountriesById
func (d *Database) GetCountriesById(actor actor.Actor) *request.DatabaseGetCountriesByIdRequest {
	return request.NewDatabaseGetCountriesByIdRequest(d.api, actor)
}

// GetFaculties Doc: https://dev.vk.com/ru/method/database.getFaculties
func (d *Database) GetFaculties(actor actor.Actor) *request.DatabaseGetFacultiesRequest {
	return request.NewDatabaseGetFacultiesRequest(d.api, actor)
}

// GetMetroStations Doc: https://dev.vk.com/ru/method/database.getMetroStations
func (d *Database) GetMetroStations(actor actor.Actor) *request.DatabaseGetMetroStationsRequest {
	return request.NewDatabaseGetMetroStationsRequest(d.api, actor)
}

// GetMetroStationsById Doc: https://dev.vk.com/ru/method/database.getMetroStationsById
func (d *Database) GetMetroStationsById(actor actor.Actor) *request.DatabaseGetMetroStationsByIdRequest {
	return request.NewDatabaseGetMetroStationsByIdRequest(d.api, actor)
}

// GetRegions Doc: https://dev.vk.com/ru/method/database.getRegions
func (d *Database) GetRegions(actor actor.Actor) *request.DatabaseGetRegionsRequest {
	return request.NewDatabaseGetRegionsRequest(d.api, actor)
}

// GetSchoolClasses Doc: https://dev.vk.com/ru/method/database.getSchoolClasses
func (d *Database) GetSchoolClasses(actor actor.Actor) *request.DatabaseGetSchoolClassesRequest {
	return request.NewDatabaseGetSchoolClassesRequest(d.api, actor)
}

// GetSchools Doc: https://dev.vk.com/ru/method/database.getSchools
func (d *Database) GetSchools(actor actor.Actor) *request.DatabaseGetSchoolsRequest {
	return request.NewDatabaseGetSchoolsRequest(d.api, actor)
}

// GetUniversities Doc: https://dev.vk.com/ru/method/database.getUniversities
func (d *Database) GetUniversities(actor actor.Actor) *request.DatabaseGetUniversitiesRequest {
	return request.NewDatabaseGetUniversitiesRequest(d.api, actor)
}
