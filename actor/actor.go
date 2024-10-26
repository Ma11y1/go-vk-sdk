package actor

type Type int

const (
	UserType Type = iota + 1
	GroupType
	GroupsType
	UserVKIDType
	ServiceType
)

type Actor interface {
	GetType() Type
	GetAccessToken() string
	GetID() int
}
