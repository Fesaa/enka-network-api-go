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
	HONKAI_BASE_URL  = "https://raw.githubusercontent.com/Dimbreath/StarRailData/master/TextMap/TextMap%s.json"
	GENSHIN_BASE_URL = "https://gitlab.com/Dimbreath/AnimeGameData/-/raw/main/TextMap/TextMap%s.json"
)

var localization *Localization

func Get() *Localization {
	return localization
}

type LocalizationMap map[string]string

type Localization struct {
	defaultKey LocalizationKey
	key        LocalizationKey

	honkaiLocalizationCache  map[LocalizationKey]LocalizationMap
	genshinLocalizationCache map[LocalizationKey]LocalizationMap

	log *slog.Logger
}

func Init(logger *slog.Logger) {
	l := Localization{
		defaultKey: ENGLISH,

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

func (l *Localization) fetchJson(url string) (*LocalizationMap, error) {

	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	if req.StatusCode != 200 {
		l.log.Debug("Returned a non 200 status code. Got ", "status_code", req.StatusCode)
		return nil, errors.New("enka-network-api-go: Non 200 status code returned: " + req.Status)
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var localizationJson LocalizationMap
	err = json.Unmarshal(data, &localizationJson)
	if err != nil {
		return nil, err
	}

	return &localizationJson, nil
}
