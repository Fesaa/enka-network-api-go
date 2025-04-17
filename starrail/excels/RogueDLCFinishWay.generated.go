package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueDLCFinishWay struct {
	FinishType    string        `json:"FinishType"`
	ID            json.Number   `json:"ID"`
	IsBackTrack   bool          `json:"IsBackTrack"`
	ParamInt1     json.Number   `json:"ParamInt1"`
	ParamInt2     json.Number   `json:"ParamInt2"`
	ParamIntList  []json.Number `json:"ParamIntList"`
	ParamItemList []interface{} `json:"ParamItemList"`
	ParamStr1     string        `json:"ParamStr1"`
	ParamType     string        `json:"ParamType"`
	Progress      json.Number   `json:"Progress"`
}
type RogueDLCFinishWayAccessor struct {
	_data []RogueDLCFinishWay
}

// LoadData retrieves the data. Must be called before RogueDLCFinishWay.GroupData
func (a *RogueDLCFinishWayAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDLCFinishWay.json")
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
func (a *RogueDLCFinishWayAccessor) Raw() ([]RogueDLCFinishWay, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDLCFinishWay{}, err
		}
	}
	return a._data, nil
}
