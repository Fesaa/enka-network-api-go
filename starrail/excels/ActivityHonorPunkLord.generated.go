package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type ActivityHonorPunkLord struct {
	DisplayPriority json.Number            `json:"DisplayPriority"`
	HonorID         json.Number            `json:"HonorID"`
	HonorName       map[string]json.Number `json:"HonorName"`
}
type ActivityHonorPunkLordAccessor struct {
	_data []ActivityHonorPunkLord
}

// LoadData retrieves the data. Must be called before ActivityHonorPunkLord.GroupData
func (a *ActivityHonorPunkLordAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityHonorPunkLord.json")
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
func (a *ActivityHonorPunkLordAccessor) Raw() ([]ActivityHonorPunkLord, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityHonorPunkLord{}, err
		}
	}
	return a._data, nil
}
