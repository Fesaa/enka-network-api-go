package genshin

type ProfilePicture struct {
	Id       int
	IconPath string
	// ROFILE_PICTURE_UNLOCK_BY_ITEM
	// PROFILE_PICTURE_UNLOCK_BY_COSTUME
	UnlockType string
	InternalId int
}
