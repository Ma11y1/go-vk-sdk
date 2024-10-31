package objects

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type UserFull struct {
	ID                     int     `json:"id"`
	FirstName              string  `json:"first_name"`
	LastName               string  `json:"last_name"`
	FirstNameNom           string  `json:"first_name_nom"`
	FirstNameGen           string  `json:"first_name_gen"`
	FirstNameDat           string  `json:"first_name_dat"`
	FirstNameAcc           string  `json:"first_name_acc"`
	FirstNameIns           string  `json:"first_name_ins"`
	FirstNameAbl           string  `json:"first_name_abl"`
	LastNameNom            string  `json:"last_name_nom"`
	LastNameGen            string  `json:"last_name_gen"`
	LastNameDat            string  `json:"last_name_dat"`
	LastNameAcc            string  `json:"last_name_acc"`
	LastNameIns            string  `json:"last_name_ins"`
	LastNameAbl            string  `json:"last_name_abl"`
	MaidenName             string  `json:"maiden_name"`
	Sex                    int     `json:"sex"`
	Nickname               string  `json:"nickname"`
	Domain                 string  `json:"domain"`
	ScreenName             string  `json:"screen_name"`
	Bdate                  string  `json:"bdate"`
	City                   City    `json:"city"`
	Country                Country `json:"country"`
	Photo50                string  `json:"photo_50"`
	Photo100               string  `json:"photo_100"`
	Photo200               string  `json:"photo_200"`
	PhotoMax               string  `json:"photo_max"`
	Photo200Orig           string  `json:"photo_200_orig"`
	Photo400Orig           string  `json:"photo_400_orig"`
	PhotoMaxOrig           string  `json:"photo_max_orig"`
	PhotoID                string  `json:"photo_id"`
	FriendStatus           int     `json:"friend_status"`
	OnlineApp              int     `json:"online_app"`
	Online                 BoolInt `json:"online"`
	OnlineMobile           BoolInt `json:"online_mobile"`
	HasPhoto               BoolInt `json:"has_photo"`
	HasMobile              BoolInt `json:"has_mobile"`
	IsClosed               BoolInt `json:"is_closed"`
	IsFriend               BoolInt `json:"is_friend"`
	IsFavorite             BoolInt `json:"is_favorite"`
	IsHiddenFromFeed       BoolInt `json:"is_hidden_from_feed"`
	CanAccessClosed        BoolInt `json:"can_access_closed"`
	CanBeInvitedGroup      BoolInt `json:"can_be_invited_group"`
	CanPost                BoolInt `json:"can_post"`
	CanSeeAllPosts         BoolInt `json:"can_see_all_posts"`
	CanSeeAudio            BoolInt `json:"can_see_audio"`
	CanWritePrivateMessage BoolInt `json:"can_write_private_message"`
	CanSendFriendRequest   BoolInt `json:"can_send_friend_request"`
	CanCallFromGroup       BoolInt `json:"can_call_from_group"`
	Verified               BoolInt `json:"verified"`
	Trending               BoolInt `json:"trending"`
	Blacklisted            BoolInt `json:"blacklisted"`
	BlacklistedByMe        BoolInt `json:"blacklisted_by_me"`

	Facebook string `json:"facebook"`

	FacebookName string `json:"facebook_name"`

	Instagram       string            `json:"instagram"`
	Twitter         string            `json:"twitter"`
	Site            string            `json:"site"`
	Status          string            `json:"status"`
	StatusAudio     Audio             `json:"status_audio"`
	LastSeen        UsersLastSeen     `json:"last_seen"`
	CropPhoto       UsersCropPhoto    `json:"crop_photo"`
	FollowersCount  int               `json:"followers_count"`
	CommonCount     int               `json:"common_count"`
	Occupation      UsersOccupation   `json:"occupation"`
	Career          []UsersCareer     `json:"career"`
	Military        []UsersMilitary   `json:"military"`
	University      int               `json:"university"`
	UniversityName  string            `json:"university_name"`
	Faculty         int               `json:"faculty"`
	FacultyName     string            `json:"faculty_name"`
	Graduation      int               `json:"graduation"`
	EducationForm   string            `json:"education_form"`
	EducationStatus string            `json:"education_status"`
	HomeTown        string            `json:"home_town"`
	Relation        int               `json:"relation"`
	Personal        UsersPersonal     `json:"personal"`
	Interests       string            `json:"interests"`
	Music           string            `json:"music"`
	Activities      string            `json:"activities"`
	Movies          string            `json:"movies"`
	Tv              string            `json:"tv"`
	Books           string            `json:"books"`
	Games           string            `json:"games"`
	Universities    []UsersUniversity `json:"universities"`
	Schools         []UsersSchool     `json:"schools"`
	About           string            `json:"about"`
	Relatives       []UsersRelative   `json:"relatives"`
	Quotes          string            `json:"quotes"`
	Lists           []int             `json:"lists"`
	Deactivated     string            `json:"deactivated"`
	WallDefault     string            `json:"wall_default"`
	Timezone        int               `json:"timezone"`
	Exports         UsersExports      `json:"exports"`
	Counters        UsersUserCounters `json:"counters"`
	MobilePhone     string            `json:"mobile_phone"`
	HomePhone       string            `json:"home_phone"`
	FoundWith       int               `json:"found_with"`
	ImageStatus     ImageStatusInfo   `json:"image_status"`
	OnlineInfo      UsersOnlineInfo   `json:"online_info"`
	Mutual          RequestMutual     `json:"mutual"`
	TrackCode       string            `json:"track_code"`
	RelationPartner UsersUserMin      `json:"relation_partner"`
	Type            string            `json:"type"`
	Skype           string            `json:"skype"`
}

func (user *UserFull) ToMention() string {
	return fmt.Sprintf("[id%d|%s %s]", user.ID, user.FirstName, user.LastName)
}

// User min version
type User struct {
	ID           int           `json:"id"`
	FirstName    string        `json:"first_name"`
	LastName     string        `json:"last_name"`
	FriendStatus int           `json:"friend_status"`
	Mutual       RequestMutual `json:"mutual"`
	OnlineApp    int           `json:"online_app"`
	Online       BoolInt       `json:"online"`
	OnlineMobile BoolInt       `json:"online_mobile"`
	Sex          int           `json:"sex"`
	ScreenName   string        `json:"screen_name"`
	Photo50      string        `json:"photo_50"`
	Photo100     string        `json:"photo_100"`
	Verified     BoolInt       `json:"verified"`
	Trending     BoolInt       `json:"trending"`
}

func (user *User) ToMention() string {
	return fmt.Sprintf("[id%d|%s %s]", user.ID, user.FirstName, user.LastName)
}

type ImageStatusInfo struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Images []Image `json:"images"`
}

type UsersOnlineInfo struct {
	AppID    int     `json:"app_id"`
	LastSeen int     `json:"last_seen"`
	Status   string  `json:"status"`
	Visible  BoolInt `json:"visible"`
	IsOnline BoolInt `json:"is_online"`
	IsMobile BoolInt `json:"is_mobile"`
}

type UsersUserMin struct {
	Deactivated string `json:"deactivated"`
	FirstName   string `json:"first_name"`
	Hidden      int    `json:"hidden"`
	ID          int    `json:"id"`
	LastName    string `json:"last_name"`
}

func (user UsersUserMin) ToMention() string {
	return fmt.Sprintf("[id%d|%s %s]", user.ID, user.FirstName, user.LastName)
}

type UsersCareer struct {
	CityID    int    `json:"city_id"`
	CityName  string `json:"city_name"`
	Company   string `json:"company"`
	CountryID int    `json:"country_id"`
	From      int    `json:"from"`
	GroupID   int    `json:"group_id"`
	ID        int    `json:"id"`
	Position  string `json:"position"`
	Until     int    `json:"until"`
}

type UsersCropPhoto struct {
	Crop  UsersCropPhotoCrop `json:"crop"`
	Photo Photo              `json:"photo"`
	Rect  UsersCropPhotoRect `json:"rect"`
}

type UsersCropPhotoCrop struct {
	X  float64 `json:"x"`
	X2 float64 `json:"x2"`
	Y  float64 `json:"y"`
	Y2 float64 `json:"y2"`
}

type UsersCropPhotoRect struct {
	X  float64 `json:"x"`
	X2 float64 `json:"x2"`
	Y  float64 `json:"y"`
	Y2 float64 `json:"y2"`
}

type UsersExports struct {
	Facebook    int `json:"facebook"`
	Livejournal int `json:"livejournal"`
	Twitter     int `json:"twitter"`
}

type UsersLastSeen struct {
	Platform int `json:"platform"`
	Time     int `json:"time"`
}

type UsersMilitary struct {
	CountryID int    `json:"country_id"`
	From      int    `json:"from"`
	ID        int    `json:"id"`
	Unit      string `json:"unit"`
	UnitID    int    `json:"unit_id"`
	Until     int    `json:"until"`
}

type UsersOccupation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type UsersPersonal struct {
	Alcohol    int      `json:"alcohol"`
	InspiredBy string   `json:"inspired_by"`
	Langs      []string `json:"langs"`
	LifeMain   int      `json:"life_main"`
	PeopleMain int      `json:"people_main"`
	Political  int      `json:"political"`
	Religion   string   `json:"religion"`
	Smoking    int      `json:"smoking"`
	ReligionID int      `json:"religion_id"`
}

func (personal *UsersPersonal) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("[]")) {
		return nil
	}

	type renamedUsersPersonal UsersPersonal

	var r renamedUsersPersonal

	err := json.Unmarshal(data, &r)
	if err != nil {
		return fmt.Errorf("objects.UsersPersonal: %w", err)
	}

	*personal = UsersPersonal(r)

	return nil
}

type UsersRelative struct {
	BirthDate string `json:"birth_date"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
}

type UsersSchool struct {
	City          int    `json:"city"`
	Class         string `json:"class"`
	Country       int    `json:"country"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	Type          int    `json:"type"`
	TypeStr       string `json:"type_str"`
	YearFrom      int    `json:"year_from"`
	YearGraduated int    `json:"year_graduated"`
	YearTo        int    `json:"year_to"`
	Speciality    string `json:"speciality,omitempty"`
}

type UsersUniversity struct {
	Chair           int    `json:"chair"`
	ChairName       string `json:"chair_name"`
	City            int    `json:"city"`
	Country         int    `json:"country"`
	EducationForm   string `json:"education_form"`
	EducationStatus string `json:"education_status"`
	Faculty         int    `json:"faculty"`
	FacultyName     string `json:"faculty_name"`
	Graduation      int    `json:"graduation"`
	ID              int    `json:"id"`
	Name            string `json:"name"`
}

type UsersUserCounters struct {
	Albums        int `json:"albums"`
	Audios        int `json:"audios"`
	Followers     int `json:"followers"`
	Friends       int `json:"friends"`
	Gifts         int `json:"gifts"`
	Groups        int `json:"groups"`
	Notes         int `json:"notes"`
	OnlineFriends int `json:"online_friends"`
	Pages         int `json:"pages"`
	Photos        int `json:"photos"`
	Subscriptions int `json:"subscriptions"`
	UserPhotos    int `json:"user_photos"`
	UserVideos    int `json:"user_videos"`
	Videos        int `json:"videos"`
	MutualFriends int `json:"mutual_friends"`
}

type UsersUserLim struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	NameGen string `json:"name_gen"`
	Photo   string `json:"photo"`
}
