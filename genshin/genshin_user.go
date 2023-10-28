package genshin

import "github.com/Fesaa/enka-network-api-go/cache"

const BASE_GENSHIN_UI_URL = "https://enka.network/ui/"

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
	NameCards             []NameCard
	// Always present
	// As of version 4.1 not reliable.
	// Use EnkaNetworkApi#GetGenshinProfileIdentifier instead.
	ProfilePictureId int
	Characters       []UserCharacter
}

func UserFromRaw(rawUser *RawGenshinUser) *User {
	playerInfoData := rawUser.PlayerInfo
	profilePicture := playerInfoData.ProfilePicture
	signature := playerInfoData.Signature
	if signature == nil {
		signature = ""
	}

	var profileId int
	if profilePicture.AvatarId != 0 {
		profileId = profilePicture.AvatarId
	} else {
		profileId = profilePicture.Id
	}

	var user User = User{
		NickName:              playerInfoData.Nickname,
		Level:                 playerInfoData.Level,
		Signature:             signature,
		WorldLevel:            playerInfoData.WorldLevel,
		NameCardId:            playerInfoData.NameCardId,
		CompletedAchievements: playerInfoData.FinishedAchievementsCount,
		TowerFloorIndex:       playerInfoData.TowerFloorIndex,
		TowerLevelIndex:       playerInfoData.TowerLevelIndex,
		ProfilePictureId:      profileId,
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

	nameCards := playerInfoData.ShowNameCardIdList
	if nameCards == nil {
		user.NameCards = make([]NameCard, 0)
	} else {
		user.NameCards = make([]NameCard, len(nameCards))
		for i, nameCard := range nameCards {
			name := cache.Get().GetNameCardName(nameCard)
			url := BASE_GENSHIN_UI_URL + name + ".png"
			user.NameCards[i] = NameCard{
				RawId: nameCard,
				Url:   url,
			}
		}
	}

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
