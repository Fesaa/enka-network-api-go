package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type TalkSentenceMultiVoice struct {
	TalkSentenceID float64   `json:"TalkSentenceID"`
	VoiceIDList    []float64 `json:"VoiceIDList"`
}
type TalkSentenceMultiVoiceAccessor struct {
	_data               []TalkSentenceMultiVoice
	_dataTalkSentenceID map[float64]TalkSentenceMultiVoice
}

// LoadData retrieves the data. Must be called before TalkSentenceMultiVoice.GroupData
func (a *TalkSentenceMultiVoiceAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TalkSentenceMultiVoice.json")
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
func (a *TalkSentenceMultiVoiceAccessor) Raw() ([]TalkSentenceMultiVoice, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TalkSentenceMultiVoice{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TalkSentenceMultiVoiceAccessor.LoadData to preload everything
func (a *TalkSentenceMultiVoiceAccessor) GroupData() {
	a._dataTalkSentenceID = map[float64]TalkSentenceMultiVoice{}
	for _, d := range a._data {
		a._dataTalkSentenceID[d.TalkSentenceID] = d
	}
}

// ByTalkSentenceID returns the TalkSentenceMultiVoice uniquely identified by TalkSentenceID
//
// Error is only non-nil if the source errors out
func (a *TalkSentenceMultiVoiceAccessor) ByTalkSentenceID(identifier float64) (TalkSentenceMultiVoice, error) {
	if a._dataTalkSentenceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TalkSentenceMultiVoice{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTalkSentenceID[identifier], nil
}
