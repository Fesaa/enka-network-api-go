package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type TrainPartyPassengerDiary struct {
	DiaryID   json.Number            `json:"DiaryID"`
	DiaryText map[string]json.Number `json:"DiaryText"`
}
type TrainPartyPassengerDiaryAccessor struct {
	_data []TrainPartyPassengerDiary
}

// LoadData retrieves the data. Must be called before TrainPartyPassengerDiary.GroupData
func (a *TrainPartyPassengerDiaryAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrainPartyPassengerDiary.json")
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
func (a *TrainPartyPassengerDiaryAccessor) Raw() ([]TrainPartyPassengerDiary, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrainPartyPassengerDiary{}, err
		}
	}
	return a._data, nil
}
