package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type DrinkMakerGuest struct {
	BartenderGuestName hash.Hash   `json:"BartenderGuestName"`
	BigIconPath        string      `json:"BigIconPath"`
	EmotionProblemList []hash.Hash `json:"EmotionProblemList"`
	FavorTagList       []float64   `json:"FavorTagList"`
	FinishQuestID      float64     `json:"FinishQuestID"`
	FinishSubMissionID float64     `json:"FinishSubMissionID"`
	GuestID            float64     `json:"GuestID"`
	IconPath           string      `json:"IconPath"`
	LinePath           string      `json:"LinePath"`
	MaxFaith           float64     `json:"MaxFaith"`
	MaxFaithReward     float64     `json:"MaxFaithReward"`
}
type DrinkMakerGuestAccessor struct {
	_data                   []DrinkMakerGuest
	_dataBigIconPath        map[string]DrinkMakerGuest
	_dataFinishQuestID      map[float64]DrinkMakerGuest
	_dataFinishSubMissionID map[float64]DrinkMakerGuest
	_dataGuestID            map[float64]DrinkMakerGuest
	_dataIconPath           map[string]DrinkMakerGuest
	_dataLinePath           map[string]DrinkMakerGuest
}

// LoadData retrieves the data. Must be called before DrinkMakerGuest.GroupData
func (a *DrinkMakerGuestAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/DrinkMakerGuest.json")
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
func (a *DrinkMakerGuestAccessor) Raw() ([]DrinkMakerGuest, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []DrinkMakerGuest{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with DrinkMakerGuestAccessor.LoadData to preload everything
func (a *DrinkMakerGuestAccessor) GroupData() {
	a._dataBigIconPath = map[string]DrinkMakerGuest{}
	a._dataFinishQuestID = map[float64]DrinkMakerGuest{}
	a._dataFinishSubMissionID = map[float64]DrinkMakerGuest{}
	a._dataGuestID = map[float64]DrinkMakerGuest{}
	a._dataIconPath = map[string]DrinkMakerGuest{}
	a._dataLinePath = map[string]DrinkMakerGuest{}
	for _, d := range a._data {
		a._dataBigIconPath[d.BigIconPath] = d
		a._dataFinishQuestID[d.FinishQuestID] = d
		a._dataFinishSubMissionID[d.FinishSubMissionID] = d
		a._dataGuestID[d.GuestID] = d
		a._dataIconPath[d.IconPath] = d
		a._dataLinePath[d.LinePath] = d
	}
}

// ByBigIconPath returns the DrinkMakerGuest uniquely identified by BigIconPath
//
// Error is only non-nil if the source errors out
func (a *DrinkMakerGuestAccessor) ByBigIconPath(identifier string) (DrinkMakerGuest, error) {
	if a._dataBigIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DrinkMakerGuest{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBigIconPath[identifier], nil
}

// ByFinishQuestID returns the DrinkMakerGuest uniquely identified by FinishQuestID
//
// Error is only non-nil if the source errors out
func (a *DrinkMakerGuestAccessor) ByFinishQuestID(identifier float64) (DrinkMakerGuest, error) {
	if a._dataFinishQuestID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DrinkMakerGuest{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFinishQuestID[identifier], nil
}

// ByFinishSubMissionID returns the DrinkMakerGuest uniquely identified by FinishSubMissionID
//
// Error is only non-nil if the source errors out
func (a *DrinkMakerGuestAccessor) ByFinishSubMissionID(identifier float64) (DrinkMakerGuest, error) {
	if a._dataFinishSubMissionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DrinkMakerGuest{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFinishSubMissionID[identifier], nil
}

// ByGuestID returns the DrinkMakerGuest uniquely identified by GuestID
//
// Error is only non-nil if the source errors out
func (a *DrinkMakerGuestAccessor) ByGuestID(identifier float64) (DrinkMakerGuest, error) {
	if a._dataGuestID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DrinkMakerGuest{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGuestID[identifier], nil
}

// ByIconPath returns the DrinkMakerGuest uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *DrinkMakerGuestAccessor) ByIconPath(identifier string) (DrinkMakerGuest, error) {
	if a._dataIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DrinkMakerGuest{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}

// ByLinePath returns the DrinkMakerGuest uniquely identified by LinePath
//
// Error is only non-nil if the source errors out
func (a *DrinkMakerGuestAccessor) ByLinePath(identifier string) (DrinkMakerGuest, error) {
	if a._dataLinePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DrinkMakerGuest{}, err
			}
		}
		a.GroupData()
	}
	return a._dataLinePath[identifier], nil
}
