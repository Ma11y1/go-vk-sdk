package longPollGroup

type FailedType int

const (
	FailedTypeOutdatedStory    FailedType = 1
	FailedTypeExpiredKey       FailedType = 2
	FailedTypeOutdatedUserInfo FailedType = 3
)
