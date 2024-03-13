package localization

import "fmt"

func (l *Localization) loadGenshinLocalization() {
	if _, ok := l.genshinLocalizationCache[l.key]; ok {
		return
	}

	l.log.Info("(Genshin) Loading localization", "key", l.key)
	localizationMap, err := l.fetchJson(fmt.Sprintf(GENSHIN_BASE_URL, string(l.key)))
	if err != nil {
		if l.key != l.defaultKey {
			l.log.Error("(Genshin) Couldn't load localization, failling back",
				"key", l.key, "fallback_key", l.defaultKey, "error", err)
			l.key = l.defaultKey
			l.loadGenshinLocalization()
			return
		}
		l.key = l.defaultKey
		l.honkaiLocalizationCache = map[LocalizationKey]LocalizationMap{}
		l.log.Error("(Genshin) Localization is not available", "key", l.key)
		return
	}

	l.genshinLocalizationCache[l.key] = *localizationMap
	l.log.Info("(Genshin) Loaded localization!", "key", l.key)
}

// GetGenshinLocale tries to retrieve the string for a hash
// in Genish impact
//
// # Takes a nameable and returns a nilable string
//
// Methods on structs will use GetGensinLocaleOrHash, call this directly
// for full control over it's behavior
func GetGenshinLocale(nameable Nameable) *string {
	if m, ok := localization.genshinLocalizationCache[localization.key]; ok {
		if name, ok := m[nameable.GetHash()]; ok {
			return &name
		}
	}
	return nil
}

// GetGenshinLocaleOrHash tries to retrieve the string for a hash
// in Genish impact
//
// Returns the hash if the string is not found
func GetGenshinLocaleOrHash(nameable Nameable) string {
	name := GetGenshinLocale(nameable)
	if name != nil {
		return *name
	}
	return nameable.GetHash()
}
