package localization

import (
	"fmt"
	"github.com/Fesaa/enka-network-api-go/hash"
)

func (l *Localization) loadHonkaiLocalization() {
	if _, ok := l.honkaiLocalizationCache[l.key]; ok {
		return
	}

	l.log.Info("(Honkai) Loading new localization", "key", l.key)
	localizationMap, err := l.fetchJson("honkai_", fmt.Sprintf(HONKAI_BASE_URL, string(l.key)))
	if err != nil {
		if l.key != l.defaultKey {
			l.log.Error("(Honkai) Couldn't load localization, failling back",
				"key", l.key, "fallback_key", l.defaultKey, "error", err)
			l.key = l.defaultKey
			l.loadHonkaiLocalization()
			return
		}
		l.key = l.defaultKey
		l.honkaiLocalizationCache = make(map[LocalizationKey]LocalizationMap)
		l.log.Error("(Honkai) Localization is not available", "key", l.key)
		return
	}

	l.honkaiLocalizationCache[l.key] = *localizationMap
	l.log.Info("(Honkai) Loaded localization!", "key", l.key)

}

// GetHonkaiLocale tries to retrieve the string for a hash
// in Honkai: Star Rail
//
// # Takes a nameable and returns a nilable string
//
// Methods on structs will use GetHonkaiLocaleOrHash, call this directly
// for full control over it's behavior
func GetHonkaiLocale(nameable hash.Nameable) *string {
	if m, ok := localization.honkaiLocalizationCache[localization.key]; ok {
		if str, ok := m[nameable.GetHash()]; ok {
			return &str
		}
	}
	return nil
}

// GetHonkaiLocaleOrHash tries to retrieve the string for a hash
// in Honkai: Star Rail
//
// Returns the hash if string is not found
func GetHonkaiLocaleOrHash(nameable hash.Nameable) string {
	name := GetHonkaiLocale(nameable)
	if name != nil {
		return *name
	}

	return nameable.GetHash()
}
