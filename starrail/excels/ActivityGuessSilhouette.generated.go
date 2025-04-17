package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ActivityGuessSilhouette struct {
	ActivityID         float64 `json:"ActivityID"`
	ActivityModuleID   float64 `json:"ActivityModuleID"`
	Aim01              string  `json:"Aim01"`
	Aim02              string  `json:"Aim02"`
	BranchQuestID      float64 `json:"BranchQuestID"`
	Daily              string  `json:"Daily"`
	Day                float64 `json:"Day"`
	FinishSubMissionID float64 `json:"FinishSubMissionID"`
	KeyIconPath        string  `json:"KeyIconPath"`
	KeyIconPath2       string  `json:"KeyIconPath2"`
	MissionID          float64 `json:"MissionID"`
	Order              float64 `json:"Order"`
	QuestID            float64 `json:"QuestID"`
	SilhouetteID       float64 `json:"SilhouetteID"`
	SilhouetteIconPath string  `json:"SilhouetteIconPath"`
	Tab                string  `json:"Tab"`
	Title              string  `json:"Title"`
	Unlock             string  `json:"Unlock"`
}
type ActivityGuessSilhouetteAccessor struct {
	_data                   []ActivityGuessSilhouette
	_dataAim02              map[string]ActivityGuessSilhouette
	_dataKeyIconPath        map[string]ActivityGuessSilhouette
	_dataSilhouetteID       map[float64]ActivityGuessSilhouette
	_dataSilhouetteIconPath map[string]ActivityGuessSilhouette
	_dataTitle              map[string]ActivityGuessSilhouette
}

// LoadData retrieves the data. Must be called before ActivityGuessSilhouette.GroupData
func (a *ActivityGuessSilhouetteAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityGuessSilhouette.json")
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
func (a *ActivityGuessSilhouetteAccessor) Raw() ([]ActivityGuessSilhouette, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityGuessSilhouette{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityGuessSilhouetteAccessor.LoadData to preload everything
func (a *ActivityGuessSilhouetteAccessor) GroupData() {
	a._dataAim02 = map[string]ActivityGuessSilhouette{}
	a._dataKeyIconPath = map[string]ActivityGuessSilhouette{}
	a._dataSilhouetteID = map[float64]ActivityGuessSilhouette{}
	a._dataSilhouetteIconPath = map[string]ActivityGuessSilhouette{}
	a._dataTitle = map[string]ActivityGuessSilhouette{}
	for _, d := range a._data {
		a._dataAim02[d.Aim02] = d
		a._dataKeyIconPath[d.KeyIconPath] = d
		a._dataSilhouetteID[d.SilhouetteID] = d
		a._dataSilhouetteIconPath[d.SilhouetteIconPath] = d
		a._dataTitle[d.Title] = d
	}
}

// ByAim02 returns the ActivityGuessSilhouette uniquely identified by Aim02
//
// Error is only non-nil if the source errors out
func (a *ActivityGuessSilhouetteAccessor) ByAim02(identifier string) (ActivityGuessSilhouette, error) {
	if a._dataAim02 == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityGuessSilhouette{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAim02[identifier], nil
}

// ByKeyIconPath returns the ActivityGuessSilhouette uniquely identified by KeyIconPath
//
// Error is only non-nil if the source errors out
func (a *ActivityGuessSilhouetteAccessor) ByKeyIconPath(identifier string) (ActivityGuessSilhouette, error) {
	if a._dataKeyIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityGuessSilhouette{}, err
			}
		}
		a.GroupData()
	}
	return a._dataKeyIconPath[identifier], nil
}

// BySilhouetteID returns the ActivityGuessSilhouette uniquely identified by SilhouetteID
//
// Error is only non-nil if the source errors out
func (a *ActivityGuessSilhouetteAccessor) BySilhouetteID(identifier float64) (ActivityGuessSilhouette, error) {
	if a._dataSilhouetteID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityGuessSilhouette{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSilhouetteID[identifier], nil
}

// BySilhouetteIconPath returns the ActivityGuessSilhouette uniquely identified by SilhouetteIconPath
//
// Error is only non-nil if the source errors out
func (a *ActivityGuessSilhouetteAccessor) BySilhouetteIconPath(identifier string) (ActivityGuessSilhouette, error) {
	if a._dataSilhouetteIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityGuessSilhouette{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSilhouetteIconPath[identifier], nil
}

// ByTitle returns the ActivityGuessSilhouette uniquely identified by Title
//
// Error is only non-nil if the source errors out
func (a *ActivityGuessSilhouetteAccessor) ByTitle(identifier string) (ActivityGuessSilhouette, error) {
	if a._dataTitle == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityGuessSilhouette{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTitle[identifier], nil
}
