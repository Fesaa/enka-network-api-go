package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ChallengeActMark struct {
	MarkIconPath string  `json:"MarkIconPath"`
	MarkType     float64 `json:"MarkType"`
}
type ChallengeActMarkAccessor struct {
	_data []ChallengeActMark
}

// LoadData retrieves the data. Must be called before ChallengeActMark.GroupData
func (a *ChallengeActMarkAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ChallengeActMark.json")
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
func (a *ChallengeActMarkAccessor) Raw() ([]ChallengeActMark, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ChallengeActMark{}, err
		}
	}
	return a._data, nil
}
