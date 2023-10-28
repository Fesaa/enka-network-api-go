package cache

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/Fesaa/enka-network-api-go/genshin"
)

func (m *MemoryCache) loadGenshinResources() error {
	cards, err := loadCards()
	if err != nil {
		return err
	}
	m.GenshinNameCards = cards

	return nil
}

func loadCards() (map[int]string, error) {
	file, err := os.ReadFile("resources/genshin_namecards.json")
	if err != nil {
		return nil, err
	}

	type Icon struct {
		IconKey string `json:"icon"`
	}
	var genshinNameCards map[string]Icon
	err = json.Unmarshal(file, &genshinNameCards)
	if err != nil {
		return nil, err
	}

	var cards map[int]string = make(map[int]string, len(genshinNameCards))
	for id, card := range genshinNameCards {
		intId, _ := strconv.Atoi(id)
		cards[intId] = card.IconKey
	}
	return cards, nil
}

func (m *MemoryCache) GetNameCardName(id int) *string {
	if name, ok := m.GenshinNameCards[id]; ok {
		return &name
	}
	return nil
}

func (m *MemoryCache) HasNameCard(id int) bool {
	_, ok := m.GenshinNameCards[id]
	return ok
}

func (m *MemoryCache) AddGenshinUser(user *genshin.RawGenshinUser) {
	m.GenshinUsers[user.Uid] = NewCachedData[*genshin.RawGenshinUser](user)
}

func (m *MemoryCache) GetGenshinUser(uid string) *genshin.RawGenshinUser {
	if cache, ok := m.GenshinUsers[uid]; ok {
		if !cache.IsExpired() {
			return cache.GetData()
		}
		delete(m.GenshinUsers, uid)
		return nil
	}
	return nil
}
