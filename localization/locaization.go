package localization

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"sync"
)

const (
	HONKAI_BASE_URL  = "https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/TextMap/TextMap%s.json"
	GENSHIN_BASE_URL = "https://gitlab.com/Dimbreath/AnimeGameData/-/raw/master/TextMap/TextMap%s.json"
)

var localization *Localization

func Get() *Localization {
	return localization
}

type LocalizationMap map[string]string

type Localization struct {
	defaultKey LocalizationKey
	key        LocalizationKey

	cache LocalizationCache

	honkaiLocalizationCache  map[LocalizationKey]LocalizationMap
	genshinLocalizationCache map[LocalizationKey]LocalizationMap

	log *slog.Logger
}

// LocalizationCache is an interface for a cache
// that can be used to store localization data
func Init(logger *slog.Logger, caches ...LocalizationCache) {
	var cache LocalizationCache
	if len(caches) > 0 {
		cache = caches[0]
	} else {
		cache = &diskCache{}
	}

	l := Localization{
		defaultKey: ENGLISH,
		cache:      cache,

		honkaiLocalizationCache:  map[LocalizationKey]LocalizationMap{},
		genshinLocalizationCache: map[LocalizationKey]LocalizationMap{},

		log: logger,
	}

	l.key = l.defaultKey
	l.loadLocalizations()

	localization = &l
}

func (l *Localization) loadLocalizations() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		l.loadHonkaiLocalization()
		wg.Add(-1)
	}()

	go func() {
		l.loadGenshinLocalization()
		wg.Add(-1)
	}()

	wg.Wait()
}

func SetLocalization(locale LocalizationKey) {
	localization.key = locale
	localization.loadLocalizations()
}

func (l *Localization) fetchJson(s string, url string) (*LocalizationMap, error) {
	cacheKey := s + string(l.key)

	data, err := l.cache.Load(cacheKey)
	if data != nil {
		l.log.Debug("Loaded from disk", "url", url)
		return unmarshal(data)
	}

	if err != nil {
		l.log.Debug("Failed to load from disk", "url", url, "error", err)
	} else {
		l.log.Debug("No data found on disk", "url", url)
	}

	l.log.Info("Fetching localization from network", "url", url)
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()
	if req.StatusCode != 200 {
		l.log.Debug("Returned a non 200 status code. Got ", "status_code", req.StatusCode)
		return nil, errors.New("enka-network-api-go: Non 200 status code returned: " + req.Status)
	}

	data, err = io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	defer func() {
		l.log.Debug("Saving localization to disk", "url", url)
		err, _ := l.cache.Save(cacheKey, data)
		if err != nil {
			l.log.Warn("Failed to save to disk", "url", url, "error", err)
		}
	}()

	return unmarshal(data)
}

func unmarshal(data []byte) (*LocalizationMap, error) {
	var localizationJson LocalizationMap
	err := json.Unmarshal(data, &localizationJson)
	if err != nil {
		return nil, err
	}

	return &localizationJson, nil
}
