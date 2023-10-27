package enkanetworkapigo

// In memory cache with maps
// This is not thread save. Be careful
type MemoryCache struct {
	HonkaiUsers map[string]CachedData[*RawHonkaiUser]
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		HonkaiUsers: map[string]CachedData[*RawHonkaiUser]{},
	}
}

func (m *MemoryCache) GetHonkaiUser(uid string) *RawHonkaiUser {
	if cache, ok := m.HonkaiUsers[uid]; ok {
		if !cache.IsExpired() {
			return cache.GetData()
		}
		delete(m.HonkaiUsers, uid)
		return nil
	}
	return nil
}

func (m *MemoryCache) AddHonkaiUser(user *RawHonkaiUser) {
	m.HonkaiUsers[user.Uid] = NewCachedData[*RawHonkaiUser](user)
}
