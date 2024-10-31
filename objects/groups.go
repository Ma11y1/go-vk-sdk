package objects

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-vk-sdk/constants"
)

type GroupFull struct {
	ID                   int                 `json:"id"`
	Name                 string              `json:"name"`
	ScreenName           string              `json:"screen_name"`
	Type                 string              `json:"type"`
	IsClosed             int                 `json:"is_closed"`
	AdminLevel           int                 `json:"admin_level,omitempty"`
	Deactivated          string              `json:"deactivated,omitempty"`
	FinishDate           int                 `json:"finish_date,omitempty"`
	Photo100             string              `json:"photo_100,omitempty"`
	Photo200             string              `json:"photo_200,omitempty"`
	Photo50              string              `json:"photo_50,omitempty"`
	StartDate            int                 `json:"start_date,omitempty"`
	Market               GroupMarketInfo     `json:"market,omitempty"`
	MemberStatus         int                 `json:"member_status,omitempty"`
	City                 City                `json:"city,omitempty"`
	Country              Country             `json:"country,omitempty"`
	IsAdmin              BoolInt             `json:"is_admin"`
	IsAdvertiser         BoolInt             `json:"is_advertiser,omitempty"`
	IsMember             BoolInt             `json:"is_member,omitempty"`
	IsFavorite           BoolInt             `json:"is_favorite,omitempty"`
	IsAdult              BoolInt             `json:"is_adult,omitempty"`
	IsSubscribed         BoolInt             `json:"is_subscribed,omitempty"`
	CanPost              BoolInt             `json:"can_post,omitempty"`
	CanSeeAllPosts       BoolInt             `json:"can_see_all_posts,omitempty"`
	CanCreateTopic       BoolInt             `json:"can_create_topic,omitempty"`
	CanUploadVideo       BoolInt             `json:"can_upload_video,omitempty"`
	CanUploadDoc         BoolInt             `json:"can_upload_doc,omitempty"`
	HasPhoto             BoolInt             `json:"has_photo,omitempty"`
	CanMessage           BoolInt             `json:"can_message,omitempty"`
	IsMessagesBlocked    BoolInt             `json:"is_messages_blocked,omitempty"`
	CanSendNotify        BoolInt             `json:"can_send_notify,omitempty"`
	IsSubscribedPodcasts BoolInt             `json:"is_subscribed_podcasts,omitempty"`
	CanSubscribePodcasts BoolInt             `json:"can_subscribe_podcasts,omitempty"`
	CanSubscribePosts    BoolInt             `json:"can_subscribe_posts,omitempty"`
	HasMarketApp         BoolInt             `json:"has_market_app,omitempty"`
	IsHiddenFromFeed     BoolInt             `json:"is_hidden_from_feed,omitempty"`
	IsMarketCartEnabled  BoolInt             `json:"is_market_cart_enabled,omitempty"`
	Verified             BoolInt             `json:"verified,omitempty"`
	Trending             BoolInt             `json:"trending,omitempty"`
	Description          string              `json:"description,omitempty"`
	WikiPage             string              `json:"wiki_page,omitempty"`
	MembersCount         int                 `json:"members_count,omitempty"`
	Counters             GroupCountersGroup  `json:"counters,omitempty"`
	Cover                GroupCover          `json:"cover,omitempty"`
	Activity             string              `json:"activity,omitempty"`
	FixedPost            int                 `json:"fixed_post,omitempty"`
	Status               string              `json:"status,omitempty"`
	MainAlbumID          int                 `json:"main_album_id,omitempty"`
	Links                []GroupLinksItem    `json:"links,omitempty"`
	Contacts             []GroupContactsItem `json:"contacts,omitempty"`
	Site                 string              `json:"site,omitempty"`
	MainSection          int                 `json:"main_section,omitempty"`
	OnlineStatus         GroupOnlineStatus   `json:"online_status,omitempty"`
	AgeLimits            int                 `json:"age_limits,omitempty"`
	BanInfo              *GroupBanInfo       `json:"ban_info,omitempty"`
	Addresses            GroupAddressesInfo  `json:"addresses,omitempty"`
	LiveCovers           GroupLiveCovers     `json:"live_covers,omitempty"`
	CropPhoto            UsersCropPhoto      `json:"crop_photo,omitempty"`
	Wall                 int                 `json:"wall,omitempty"`
	ActionButton         GroupActionButton   `json:"action_button,omitempty"`
	TrackCode            string              `json:"track_code,omitempty"`
	PublicDateLabel      string              `json:"public_date_label,omitempty"`
	AuthorID             int                 `json:"author_id,omitempty"`
	Phone                string              `json:"phone,omitempty"`
	Like                 GroupLike           `json:"like"`
}

type GroupAddress struct {
	AdditionalAddress string                `json:"additional_address"`
	Address           string                `json:"address"`
	CityID            int                   `json:"city_id"`
	CountryID         int                   `json:"country_id"`
	Distance          int                   `json:"distance"`
	ID                int                   `json:"id"`
	Latitude          float64               `json:"latitude"`
	Longitude         float64               `json:"longitude"`
	MetroStationID    int                   `json:"metro_station_id"`
	Phone             string                `json:"phone"`
	TimeOffset        int                   `json:"time_offset"`
	Timetable         GroupAddressTimetable `json:"timetable"`
	Title             string                `json:"title"`
	WorkInfoStatus    string                `json:"work_info_status"`
	PlaceID           int                   `json:"place_id"`
}

type GroupAddressTimetableDay struct {
	BreakCloseTime int `json:"break_close_time"`
	BreakOpenTime  int `json:"break_open_time"`
	CloseTime      int `json:"close_time"`
	OpenTime       int `json:"open_time"`
}

type GroupAddressTimetable struct {
	Fri GroupAddressTimetableDay `json:"fri"`
	Mon GroupAddressTimetableDay `json:"mon"`
	Sat GroupAddressTimetableDay `json:"sat"`
	Sun GroupAddressTimetableDay `json:"sun"`
	Thu GroupAddressTimetableDay `json:"thu"`
	Tue GroupAddressTimetableDay `json:"tue"`
	Wed GroupAddressTimetableDay `json:"wed"`
}

type GroupAddressesInfo struct {
	IsEnabled     BoolInt `json:"is_enabled"`
	MainAddressID int     `json:"main_address_id"`
}

func (group GroupFull) ToMention() string {
	return fmt.Sprintf("[club%d|%s]", group.ID, group.Name)
}

type GroupLike struct {
	IsLiked BoolInt          `json:"is_liked"`
	Friends GroupLikeFriends `json:"friends"`
}

type GroupLikeFriends struct {
	Count   int   `json:"count"`
	Preview []int `json:"preview"`
}

type GroupLiveCovers struct {
	IsEnabled  BoolInt  `json:"is_enabled"`
	IsScalable BoolInt  `json:"is_scalable"`
	StoryIDs   []string `json:"story_ids"`
}

type GroupsBanInfo struct {
	AdminID        int     `json:"admin_id"`
	Comment        string  `json:"comment"`
	Date           int     `json:"date"`
	EndDate        int     `json:"end_date"`
	Reason         int     `json:"reason"`
	CommentVisible BoolInt `json:"comment_visible"`
}

type GroupCallbackServer struct {
	CreatorID int    `json:"creator_id"`
	ID        int    `json:"id"`
	SecretKey string `json:"secret_key"`
	Status    string `json:"status"`
	Title     string `json:"title"`
	URL       string `json:"url"`
}

type GroupCallbackSettings struct {
	APIVersion string              `json:"api_version"`
	Events     GroupLongPollEvents `json:"events"`
}

type GroupContactsItem struct {
	Desc   string `json:"desc"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	UserID int    `json:"user_id"`
}

type GroupCountersGroup struct {
	Addresses      int `json:"addresses"`
	Albums         int `json:"albums"`
	Articles       int `json:"articles"`
	Audios         int `json:"audios"`
	Docs           int `json:"docs"`
	Market         int `json:"market"`
	Photos         int `json:"photos"`
	Topics         int `json:"topics"`
	Videos         int `json:"videos"`
	Narratives     int `json:"narratives"`
	Clips          int `json:"clips"`
	ClipsFollowers int `json:"clips_followers"`
}

func (personal *GroupCountersGroup) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("[]")) {
		return nil
	}

	type renamedGroupCountersGroup GroupCountersGroup

	var r renamedGroupCountersGroup

	err := json.Unmarshal(data, &r)
	if err != nil {
		return fmt.Errorf("objects.GroupCountersGroup: %w", err)
	}

	*personal = GroupCountersGroup(r)

	return nil
}

type GroupCover struct {
	Enabled BoolInt `json:"enabled"`
	Images  []Image `json:"images"`
}

type GroupBanInfo struct {
	Comment string `json:"comment"`
	EndDate int    `json:"end_date"`
}

type GroupCategory struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Subcategories []IDName `json:"subcategories"`
}

type GroupCategoryFull struct {
	ID            int                 `json:"id"`
	Name          string              `json:"name"`
	PageCount     int                 `json:"page_count"`
	PagePreviews  []GroupFull         `json:"page_previews"`
	Subcategories []GroupCategoryFull `json:"subcategories"`
}

type GroupCategoryType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GroupLink struct {
	Desc            string  `json:"desc"`
	EditTitle       BoolInt `json:"edit_title"`
	ImageProcessing BoolInt `json:"image_processing"`
	Name            string  `json:"name"`
	ID              int     `json:"id"`
	URL             string  `json:"url"`
}

type GroupPublicCategoryList struct {
	ID            int                 `json:"id"`
	Name          string              `json:"name"`
	Subcategories []GroupCategoryType `json:"subcategories"`
}

type GroupSettings struct {
	Access             int                       `json:"access"`
	Address            string                    `json:"address"`
	Audio              int                       `json:"audio"`
	Description        string                    `json:"description"`
	Docs               int                       `json:"docs"`
	ObsceneWords       []string                  `json:"obscene_words"`
	Photos             int                       `json:"photos"`
	PublicCategory     int                       `json:"public_category"`
	PublicCategoryList []GroupPublicCategoryList `json:"public_category_list"`
	PublicSubcategory  int                       `json:"public_subcategory"`
	Rss                string                    `json:"rss"`
	Subject            int                       `json:"subject"`
	SubjectList        []GroupSubjectItem        `json:"subject_list"`
	Title              string                    `json:"title"`
	Topics             int                       `json:"topics"`
	Video              int                       `json:"video"`
	Wall               int                       `json:"wall"`
	Website            string                    `json:"website"`
	Wiki               int                       `json:"wiki"`
	CountryID          int                       `json:"country_id"`
	CityID             int                       `json:"city_id"`
	Messages           int                       `json:"messages"`
	Articles           int                       `json:"articles"`
	Events             int                       `json:"events"`
	AgeLimits          int                       `json:"age_limits"`
	ObsceneFilter      BoolInt                   `json:"obscene_filter"`
	ObsceneStopwords   BoolInt                   `json:"obscene_stopwords"`
	LiveCovers         struct {
		IsEnabled BoolInt `json:"is_enabled"`
	} `json:"live_covers"`
	Market           GroupMarketInfo     `json:"market"`
	SectionsList     []IDTitle           `json:"sections_list"`
	MainSection      int                 `json:"main_section"`
	SecondarySection int                 `json:"secondary_section"`
	ActionButton     GroupActionButton   `json:"action_button"`
	Phone            string              `json:"phone"`
	RecognizePhoto   int                 `json:"recognize_photo"`
	MarketServices   GroupMarketServices `json:"market_services"`
	Narratives       int                 `json:"narratives"`
	Clips            int                 `json:"clips"`
	Textlives        int                 `json:"textlives"`
	Youla            GroupYoula          `json:"youla"`
}

type GroupMarketServices struct {
	Enabled         BoolInt            `json:"enabled"`
	CanMessage      BoolInt            `json:"can_message"`
	CommentsEnabled BoolInt            `json:"comments_enabled"`
	ContactID       int                `json:"contact_id"`
	Currency        MarketCurrency     `json:"currency"`
	ViewType        GroupSelectedItems `json:"view_type"`
	BlockName       GroupSelectedItems `json:"block_name"`
	ButtonLabel     GroupSelectedItems `json:"button_label"`
}

type GroupSelectedItems struct {
	SelectedItemID int64    `json:"selected_item_id"`
	Items          []IDName `json:"items"`
}

type GroupYoula struct {
	CategoryTree GroupYoulaCategory `json:"category_tree"`
	Groupettings GroupYoulaSettings `json:"group_settings"`
}

type GroupYoulaCategory struct {
	ID            int                     `json:"id"`
	Title         string                  `json:"title"`
	Subcategories []GroupYoulaSubcategory `json:"subcategories"`
}

type GroupYoulaSubcategory struct {
	ID            int                     `json:"id"`
	Title         string                  `json:"title"`
	ParentID      int                     `json:"parent_id"`
	Subcategories []GroupYoulaSubcategory `json:"subcategories"`
}

type GroupYoulaSettings struct {
	IsActive              BoolInt   `json:"is_active"`
	IsModerated           BoolInt   `json:"is_moderated"`
	ShowModerationSetting BoolInt   `json:"show_moderation_setting"`
	ModerationStatus      int       `json:"moderation_status"`
	DeclineReason         string    `json:"decline_reason"`
	GroupMode             int       `json:"group_mode"`
	SelectedCategoryIDs   []int     `json:"selected_category_ids"`
	Lat                   float64   `json:"lat"`
	Long                  float64   `json:"long"`
	Radius                float64   `json:"radius"`
	RadiusArea            string    `json:"radius_area"`
	Address               string    `json:"address"`
	Radiuses              []float64 `json:"radiuses"`
}

type GroupActionButton struct {
	ActionType constants.GroupActionType `json:"action_type"`
	Target     GroupActionButtonTarget   `json:"target"`
	Title      string                    `json:"title"`
	IsEnabled  BoolInt                   `json:"is_enabled,omitempty"`
}

type GroupActionButtonTarget struct {
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	UserID         int     `json:"user_id"`
	URL            string  `json:"url"`
	GoogleStoreURL string  `json:"google_store_url"`
	ItunesURL      string  `json:"itunes_url"`
	AppID          int     `json:"app_id"`
	IsInternal     BoolInt `json:"is_internal"`
}

type GroupXtrInvitedBy struct {
	AdminLevel   int     `json:"admin_level"`
	ID           int     `json:"id"`
	InvitedBy    int     `json:"invited_by"`
	Name         string  `json:"name"`
	Photo100     string  `json:"photo_100"`
	Photo200     string  `json:"photo_200"`
	Photo50      string  `json:"photo_50"`
	ScreenName   string  `json:"screen_name"`
	Type         string  `json:"type"`
	IsClosed     int     `json:"is_closed"`
	IsAdmin      BoolInt `json:"is_admin"`
	IsMember     BoolInt `json:"is_member"`
	IsAdvertiser BoolInt `json:"is_advertiser"`
}

func (group GroupXtrInvitedBy) ToMention() string {
	return fmt.Sprintf("[club%d|%s]", group.ID, group.Name)
}

type GroupLinksItem struct {
	Desc      string  `json:"desc"`
	EditTitle BoolInt `json:"edit_title"`
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Photo100  string  `json:"photo_100"`
	Photo50   string  `json:"photo_50"`
	URL       string  `json:"url"`
}

type GroupLongPollEvents struct {
	MessageNew                    BoolInt `json:"message_new"`
	MessageReply                  BoolInt `json:"message_reply"`
	PhotoNew                      BoolInt `json:"photo_new"`
	AudioNew                      BoolInt `json:"audio_new"`
	VideoNew                      BoolInt `json:"video_new"`
	WallReplyNew                  BoolInt `json:"wall_reply_new"`
	WallReplyEdit                 BoolInt `json:"wall_reply_edit"`
	WallReplyDelete               BoolInt `json:"wall_reply_delete"`
	WallReplyRestore              BoolInt `json:"wall_reply_restore"`
	WallPostNew                   BoolInt `json:"wall_post_new"`
	BoardPostNew                  BoolInt `json:"board_post_new"`
	BoardPostEdit                 BoolInt `json:"board_post_edit"`
	BoardPostRestore              BoolInt `json:"board_post_restore"`
	BoardPostDelete               BoolInt `json:"board_post_delete"`
	PhotoCommentNew               BoolInt `json:"photo_comment_new"`
	PhotoCommentEdit              BoolInt `json:"photo_comment_edit"`
	PhotoCommentDelete            BoolInt `json:"photo_comment_delete"`
	PhotoCommentRestore           BoolInt `json:"photo_comment_restore"`
	VideoCommentNew               BoolInt `json:"video_comment_new"`
	VideoCommentEdit              BoolInt `json:"video_comment_edit"`
	VideoCommentDelete            BoolInt `json:"video_comment_delete"`
	VideoCommentRestore           BoolInt `json:"video_comment_restore"`
	MarketCommentNew              BoolInt `json:"market_comment_new"`
	MarketCommentEdit             BoolInt `json:"market_comment_edit"`
	MarketCommentDelete           BoolInt `json:"market_comment_delete"`
	MarketCommentRestore          BoolInt `json:"market_comment_restore"`
	MarketOrderNew                BoolInt `json:"market_order_new"`
	MarketOrderEdit               BoolInt `json:"market_order_edit"`
	PollVoteNew                   BoolInt `json:"poll_vote_new"`
	GroupJoin                     BoolInt `json:"group_join"`
	GroupLeave                    BoolInt `json:"group_leave"`
	GroupChangeSettings           BoolInt `json:"group_change_settings"`
	GroupChangePhoto              BoolInt `json:"group_change_photo"`
	GroupOfficersEdit             BoolInt `json:"group_officers_edit"`
	MessageAllow                  BoolInt `json:"message_allow"`
	MessageDeny                   BoolInt `json:"message_deny"`
	WallRepost                    BoolInt `json:"wall_repost"`
	UserBlock                     BoolInt `json:"user_block"`
	UserUnblock                   BoolInt `json:"user_unblock"`
	MessageEdit                   BoolInt `json:"message_edit"`
	MessageTypingState            BoolInt `json:"message_typing_state"`
	LeadFormsNew                  BoolInt `json:"lead_forms_new"`
	LikeAdd                       BoolInt `json:"like_add"`
	LikeRemove                    BoolInt `json:"like_remove"`
	VkpayTransaction              BoolInt `json:"vkpay_transaction"`
	AppPayload                    BoolInt `json:"app_payload"`
	MessageRead                   BoolInt `json:"message_read"`
	MessageEvent                  BoolInt `json:"message_event"`
	DonutSubscriptionCreate       BoolInt `json:"donut_subscription_create"`
	DonutSubscriptionProlonged    BoolInt `json:"donut_subscription_prolonged"`
	DonutSubscriptionExpired      BoolInt `json:"donut_subscription_expired"`
	DonutSubscriptionCancelled    BoolInt `json:"donut_subscription_cancelled"`
	DonutSubscriptionPriceChanged BoolInt `json:"donut_subscription_price_changed"`
	DonutMoneyWithdraw            BoolInt `json:"donut_money_withdraw"`
	DonutMoneyWithdrawError       BoolInt `json:"donut_money_withdraw_error"`
}

type GroupLongPollServer struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	Ts     string `json:"ts"`
}

func (lp GroupLongPollServer) GetURL(wait int) string {
	return fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=%d", lp.Server, lp.Key, lp.Ts, wait)
}

type GroupLongPollSettings struct {
	APIVersion string              `json:"api_version"`
	Events     GroupLongPollEvents `json:"events"`
	IsEnabled  BoolInt             `json:"is_enabled"`
}

type GroupMarketInfo struct {
	Type            constants.GroupMarketType `json:"type,omitempty"`
	ContactID       int                       `json:"contact_id,omitempty"`
	Currency        MarketCurrency            `json:"currency,omitempty"`
	CurrencyText    string                    `json:"currency_text,omitempty"`
	Enabled         BoolInt                   `json:"enabled"`
	CommentsEnabled BoolInt                   `json:"comments_enabled,omitempty"`
	CanMessage      BoolInt                   `json:"can_message,omitempty"`
	IsHsEnabled     BoolInt                   `json:"is_hs_enabled,omitempty"`
	MainAlbumID     int                       `json:"main_album_id,omitempty"`
	PriceMax        string                    `json:"price_max,omitempty"`
	PriceMin        string                    `json:"price_min,omitempty"`
	Wiki            PagesWikipageFull         `json:"wiki,omitempty"`
	CityIDs         []int                     `json:"city_ids"`
	CountryIDs      []int                     `json:"country_ids,omitempty"`
	MinOrderPrice   MarketPrice               `json:"min_order_price,omitempty"`
}

type GroupMemberRole struct {
	ID          int      `json:"id"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

type GroupMemberRoleXtrUser struct {
	UserFull
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

type GroupMemberStatus struct {
	Member      BoolInt  `json:"member"`
	UserID      int      `json:"user_id"`
	Permissions []string `json:"permissions"`
}

type GroupMemberStatusFull struct {
	Invitation BoolInt `json:"invitation"`
	Member     BoolInt `json:"member"`
	Request    BoolInt `json:"request"`
	CanInvite  BoolInt `json:"can_invite"`
	CanRecall  BoolInt `json:"can_recall"`
	UserID     int     `json:"user_id"`
}

type GroupOnlineStatus struct {
	Minutes int    `json:"minutes"`
	Status  string `json:"status"`
}

type GroupOwnerXtrBanInfo struct {
	BanInfo *GroupsBanInfo `json:"ban_info"`
	Group   GroupFull      `json:"group"`
	Profile UserFull       `json:"profile"`
	Type    string         `json:"type"`
}

type GroupSubjectItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GroupTokenPermissionSetting struct {
	Name    string `json:"name"`
	Setting int    `json:"setting"`
}

type GroupTokenPermissions struct {
	Mask        int                           `json:"mask"`
	Permissions []GroupTokenPermissionSetting `json:"permissions"`
}

type GroupTag struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type GroupProfileItem struct {
	ID        int    `json:"id"`
	Photo50   string `json:"photo_50"`
	Photo100  string `json:"photo_100"`
	FirstName string `json:"first_name"`
}

type GroupAttach struct {
	ID         int    `json:"id"`
	IsFavorite bool   `json:"is_favorite"`
	Size       int    `json:"size"`
	Status     string `json:"status"`
	Text       string `json:"text"`
}
