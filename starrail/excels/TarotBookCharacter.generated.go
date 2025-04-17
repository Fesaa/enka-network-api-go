package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type TarotBookCharacter struct {
	ID               json.Number            `json:"ID"`
	MainCatalogTitle map[string]json.Number `json:"MainCatalogTitle"`
	MaxLevel         json.Number            `json:"MaxLevel"`
	Name             map[string]json.Number `json:"Name"`
	Position         json.Number            `json:"Position"`
	PrefabPath       string                 `json:"PrefabPath"`
	RectIconPath     string                 `json:"RectIconPath"`
	RoundIconPath    string                 `json:"RoundIconPath"`
	StoryList        []json.Number          `json:"StoryList"`
	SubCatalogTitle  map[string]json.Number `json:"SubCatalogTitle"`
	TabIconPath      string                 `json:"TabIconPath"`
	Tag              json.Number            `json:"Tag"`
}
type TarotBookCharacterAccessor struct {
	_data              []TarotBookCharacter
	_dataPrefabPath    map[string]TarotBookCharacter
	_dataRectIconPath  map[string]TarotBookCharacter
	_dataRoundIconPath map[string]TarotBookCharacter
	_dataTabIconPath   map[string]TarotBookCharacter
}

// LoadData retrieves the data. Must be called before TarotBookCharacter.GroupData
func (a *TarotBookCharacterAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TarotBookCharacter.json")
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &a._data)
}

// Raw returns the raw data.
func (a *TarotBookCharacterAccessor) Raw() ([]TarotBookCharacter, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TarotBookCharacter{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TarotBookCharacterAccessor.LoadData to preload everything
func (a *TarotBookCharacterAccessor) GroupData() {
	a._dataPrefabPath = map[string]TarotBookCharacter{}
	a._dataRectIconPath = map[string]TarotBookCharacter{}
	a._dataRoundIconPath = map[string]TarotBookCharacter{}
	a._dataTabIconPath = map[string]TarotBookCharacter{}
	for _, d := range a._data {
		a._dataPrefabPath[d.PrefabPath] = d
		a._dataRectIconPath[d.RectIconPath] = d
		a._dataRoundIconPath[d.RoundIconPath] = d
		a._dataTabIconPath[d.TabIconPath] = d
	}
}

// ByPrefabPath returns the TarotBookCharacter uniquely identified by PrefabPath
//
// Error is only non-nil if the source errors out
func (a *TarotBookCharacterAccessor) ByPrefabPath(identifier string) (TarotBookCharacter, error) {
	if a._dataPrefabPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TarotBookCharacter{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPrefabPath[identifier], nil
}

// ByRectIconPath returns the TarotBookCharacter uniquely identified by RectIconPath
//
// Error is only non-nil if the source errors out
func (a *TarotBookCharacterAccessor) ByRectIconPath(identifier string) (TarotBookCharacter, error) {
	if a._dataRectIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TarotBookCharacter{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRectIconPath[identifier], nil
}

// ByRoundIconPath returns the TarotBookCharacter uniquely identified by RoundIconPath
//
// Error is only non-nil if the source errors out
func (a *TarotBookCharacterAccessor) ByRoundIconPath(identifier string) (TarotBookCharacter, error) {
	if a._dataRoundIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TarotBookCharacter{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRoundIconPath[identifier], nil
}

// ByTabIconPath returns the TarotBookCharacter uniquely identified by TabIconPath
//
// Error is only non-nil if the source errors out
func (a *TarotBookCharacterAccessor) ByTabIconPath(identifier string) (TarotBookCharacter, error) {
	if a._dataTabIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TarotBookCharacter{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTabIconPath[identifier], nil
}
