package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type SwordTrainingStatus struct {
	InitialValue      float64   `json:"InitialValue"`
	MaximumValue      float64   `json:"MaximumValue"`
	StatusID          float64   `json:"StatusID"`
	StatusIcon        string    `json:"StatusIcon"`
	StatusName        hash.Hash `json:"StatusName"`
	StatusOutLineIcon string    `json:"StatusOutLineIcon"`
}
type SwordTrainingStatusAccessor struct {
	_data                  []SwordTrainingStatus
	_dataStatusOutLineIcon map[string]SwordTrainingStatus
	_dataStatusID          map[float64]SwordTrainingStatus
	_dataStatusIcon        map[string]SwordTrainingStatus
}

// LoadData retrieves the data. Must be called before SwordTrainingStatus.GroupData
func (a *SwordTrainingStatusAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SwordTrainingStatus.json")
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
func (a *SwordTrainingStatusAccessor) Raw() ([]SwordTrainingStatus, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SwordTrainingStatus{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SwordTrainingStatusAccessor.LoadData to preload everything
func (a *SwordTrainingStatusAccessor) GroupData() {
	for _, d := range a._data {
		a._dataStatusOutLineIcon[d.StatusOutLineIcon] = d
		a._dataStatusID[d.StatusID] = d
		a._dataStatusIcon[d.StatusIcon] = d
	}
}

// ByStatusOutLineIcon returns the SwordTrainingStatus uniquely identified by StatusOutLineIcon
//
// Error is only non-nil if the source errors out
func (a *SwordTrainingStatusAccessor) ByStatusOutLineIcon(identifier string) (SwordTrainingStatus, error) {
	if a._dataStatusOutLineIcon == nil {
		err := a.LoadData()
		if err != nil {
			return SwordTrainingStatus{}, err
		}
		a.GroupData()
	}
	return a._dataStatusOutLineIcon[identifier], nil
}

// ByStatusID returns the SwordTrainingStatus uniquely identified by StatusID
//
// Error is only non-nil if the source errors out
func (a *SwordTrainingStatusAccessor) ByStatusID(identifier float64) (SwordTrainingStatus, error) {
	if a._dataStatusID == nil {
		err := a.LoadData()
		if err != nil {
			return SwordTrainingStatus{}, err
		}
		a.GroupData()
	}
	return a._dataStatusID[identifier], nil
}

// ByStatusIcon returns the SwordTrainingStatus uniquely identified by StatusIcon
//
// Error is only non-nil if the source errors out
func (a *SwordTrainingStatusAccessor) ByStatusIcon(identifier string) (SwordTrainingStatus, error) {
	if a._dataStatusIcon == nil {
		err := a.LoadData()
		if err != nil {
			return SwordTrainingStatus{}, err
		}
		a.GroupData()
	}
	return a._dataStatusIcon[identifier], nil
}
