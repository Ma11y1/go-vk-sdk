package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/database

type DatabaseGetChairsResponse struct {
	BaseResponse
	Response struct {
		Count int               `json:"count"`
		Items []objects.IDTitle `json:"items"`
	} `json:"response"`
}

type DatabaseGetCitiesResponse struct {
	BaseResponse
	Response struct {
		Count int                    `json:"count"`
		Items []objects.DatabaseCity `json:"items"`
	} `json:"response"`
}

type DatabaseGetCitiesByIDResponse struct {
	BaseResponse
	Response []objects.DatabaseCity `json:"response"`
}

type DatabaseGetCountriesResponse struct {
	BaseResponse
	Response struct {
		Count int               `json:"count"`
		Items []objects.IDTitle `json:"items"`
	} `json:"response"`
}

type DatabaseGetCountriesByIDResponse struct {
	BaseResponse
	Response []objects.IDTitle `json:"response"`
}

type DatabaseGetFacultiesResponse struct {
	BaseResponse
	Response struct {
		Count int                       `json:"count"`
		Items []objects.DatabaseFaculty `json:"items"`
	} `json:"response"`
}

type DatabaseGetMetroStationsResponse struct {
	BaseResponse
	Response struct {
		Count int                            `json:"count"`
		Items []objects.DatabaseMetroStation `json:"items"`
	} `json:"response"`
}

type DatabaseGetMetroStationsByIDResponse struct {
	BaseResponse
	Response []objects.DatabaseMetroStation `json:"response"`
}

type DatabaseGetRegionsResponse struct {
	BaseResponse
	Response struct {
		Count int                      `json:"count"`
		Items []objects.DatabaseRegion `json:"items"`
	} `json:"response"`
}

type DatabaseGetSchoolClassesResponse struct {
	BaseResponse
	Response []objects.IDTitle `json:"response"`
}

type DatabaseGetSchoolsResponse struct {
	BaseResponse
	Response struct {
		Count int                      `json:"count"`
		Items []objects.DatabaseSchool `json:"items"`
	} `json:"response"`
}

type DatabaseGetUniversitiesResponse struct {
	BaseResponse
	Response struct {
		Count int                          `json:"count"`
		Items []objects.DatabaseUniversity `json:"items"`
	} `json:"response"`
}
