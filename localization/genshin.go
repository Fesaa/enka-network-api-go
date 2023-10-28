package localization

import "fmt"

func (l *Localization) loadGenshinLocalization() {
	if _, ok := l.genshinLocalizationCache[l.key]; ok {
		return
	}

	l.log.Infof("(Genshin) Loading localization for %s", l.key)
	localizationMap, err := l.fetchJson(fmt.Sprintf(GENSHIN_BASE_URL, string(l.key)))
	if err != nil {
		l.log.Errorf("(Genshin) Couldn't load localization for %s, failling back to %s \n Error: %s", l.key, l.defaultKey, err)
		l.key = l.defaultKey
		l.honkaiLocalizationCache = map[LocalizationKey]LocalizationMap{}
		return
	}

	l.genshinLocalizationCache[l.key] = *localizationMap
	l.log.Infof("(Genshin) Loaded localization for %s", l.key)
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
