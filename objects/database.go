package objects

type DatabaseCity struct {
	ID        int            `json:"id"`
	Title     string         `json:"title"`
	Area      string         `json:"area,omitempty"`
	Region    string         `json:"region,omitempty"`
	Important NumberFlagBool `json:"important,omitempty"`
}

type DatabaseMetroStation struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	CityID int    `json:"city_id"`
}

type DatabaseFaculty struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type DatabaseRegion struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type DatabaseSchool struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type DatabaseStation struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	CityID int    `json:"city_id"`
	Color  string `json:"color"`
}

type DatabaseUniversity struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
