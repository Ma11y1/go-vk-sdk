package objects

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type UtilsDomainResolved struct {
	GroupID  int    `json:"group_id"`
	ObjectID int    `json:"object_id"`
	Type     string `json:"type"`
}

func (link *UtilsDomainResolved) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("[]")) {
		return nil
	}

	type renamedUtilsDomainResolved UtilsDomainResolved

	var r renamedUtilsDomainResolved

	err := json.Unmarshal(data, &r)
	if err != nil {
		return fmt.Errorf("objects.UtilsDomainResolved: %w", err)
	}

	*link = UtilsDomainResolved(r)

	return nil
}

type UtilsLastShortenedLink struct {
	AccessKey string `json:"access_key"`
	Key       string `json:"key"`
	ShortURL  string `json:"short_url"`
	Timestamp int    `json:"timestamp"`
	URL       string `json:"url"`
	Views     int    `json:"views"`
}

type UtilsLinkChecked struct {
	Link   string `json:"link"`
	Status string `json:"status"`
}

type UtilsLinkStats struct {
	Key   string       `json:"key"`
	Stats []UtilsStats `json:"stats"`
}

type UtilsLinkStatsExtended struct {
	Key   string               `json:"key"`
	Stats []UtilsStatsExtended `json:"stats"`
}

type UtilsShortLink struct {
	AccessKey string `json:"access_key"`
	Key       string `json:"key"`
	ShortURL  string `json:"short_url"`
	URL       string `json:"url"`
}

type UtilsStats struct {
	Timestamp int `json:"timestamp"`
	Views     int `json:"views"`
}

type UtilsStatsCity struct {
	CityID int `json:"city_id"`
	Views  int `json:"views"`
}

type UtilsStatsCountry struct {
	CountryID int `json:"country_id"`
	Views     int `json:"views"`
}

type UtilsStatsExtended struct {
	Cities    []UtilsStatsCity    `json:"cities"`
	Countries []UtilsStatsCountry `json:"countries"`
	SexAge    []UtilsStatsSexAge  `json:"sex_age"`
	Timestamp int                 `json:"timestamp"`
	Views     int                 `json:"views"`
}

type UtilsStatsSexAge struct {
	AgeRange string `json:"age_range"`
	Female   int    `json:"female"`
	Male     int    `json:"male"`
}
