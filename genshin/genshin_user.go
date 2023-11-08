package genshin

type User struct {
	NickName              string
	Signature             string
	Level                 int
	WorldLevel            int
	NameCardId            int
	CompletedAchievements int
	TowerFloorIndex       int
	TowerLevelIndex       int
	ShowCasedCharacters   []ShowCaseCharacter
	// Ids of the users NameCards, use EnkaNetworkApi#GetNameCardName to get the name
	NameCardsId []int
	// Always present
	// As of version 4.1 not reliable.
	// Use EnkaNetworkApi#GetGenshinProfileIdentifier instead.
	ProfilePictureId Pair[int]
	Characters       []UserCharacter
}

type Pair[T any] struct {
	Left  T
	Right T
}

func filterID(pfp *RawProfilePicture) Pair[int] {
	if pfp.AvatarId != 0 {
		return Pair[int]{pfp.AvatarId, 0}
	} else {
		return Pair[int]{pfp.AvatarId, pfp.CostumeId}
	}
}

func UserFromRaw(rawUser *RawGenshinUser) *User {
	playerInfoData := rawUser.PlayerInfo
	profilePicture := playerInfoData.ProfilePicture
	signature := playerInfoData.Signature
	if signature == nil {
		temp := ""
		signature = &temp
	}

	var user User = User{
		NickName:              playerInfoData.Nickname,
		Level:                 playerInfoData.Level,
		Signature:             *signature,
		WorldLevel:            playerInfoData.WorldLevel,
		NameCardId:            playerInfoData.NameCardId,
		CompletedAchievements: playerInfoData.FinishedAchievementsCount,
		TowerFloorIndex:       playerInfoData.TowerFloorIndex,
		TowerLevelIndex:       playerInfoData.TowerLevelIndex,
		ProfilePictureId:      filterID(&profilePicture),
	}

	showAvatars := playerInfoData.ShowAvatarInfoList
	if showAvatars == nil {
		user.ShowCasedCharacters = make([]ShowCaseCharacter, 0)
	} else {
		user.ShowCasedCharacters = make([]ShowCaseCharacter, len(showAvatars))
		for i, showAvatar := range showAvatars {
			user.ShowCasedCharacters[i] = ShowCaseCharacter{
				AvatarId:  showAvatar.AvatarId,
				Level:     showAvatar.Level,
				CostumeId: showAvatar.CostumeId,
			}
		}
	}

	user.NameCardsId = playerInfoData.ShowNameCardIdList

	avatars := rawUser.AvatarInfoList
	if avatars == nil {
		user.Characters = make([]UserCharacter, 0)
	} else {
		user.Characters = make([]UserCharacter, len(avatars))
		for i, avatar := range avatars {
			user.Characters[i] = userCharacterFromRaw(avatar)
		}
	}

	return &user
}
