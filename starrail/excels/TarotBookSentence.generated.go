package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type TarotBookSentence struct {
	ID       json.Number            `json:"ID"`
	Sentence map[string]json.Number `json:"Sentence"`
	VoiceID  json.Number            `json:"VoiceID"`
}
type TarotBookSentenceAccessor struct {
	_data   []TarotBookSentence
	_dataID map[json.Number]TarotBookSentence
}

// LoadData retrieves the data. Must be called before TarotBookSentence.GroupData
func (a *TarotBookSentenceAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TarotBookSentence.json")
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
func (a *TarotBookSentenceAccessor) Raw() ([]TarotBookSentence, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TarotBookSentence{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TarotBookSentenceAccessor.LoadData to preload everything
func (a *TarotBookSentenceAccessor) GroupData() {
	a._dataID = map[json.Number]TarotBookSentence{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the TarotBookSentence uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *TarotBookSentenceAccessor) ByID(identifier json.Number) (TarotBookSentence, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TarotBookSentence{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
