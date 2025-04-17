package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type PlanetFesAvatar struct {
	AnimConfig        string    `json:"AnimConfig"`
	Body              string    `json:"Body"`
	CD                float64   `json:"CD"`
	CargoIcon         string    `json:"CargoIcon"`
	Description       string    `json:"Description"`
	GachaUnlockIDList []float64 `json:"GachaUnlockIDList"`
	HeadIcon          string    `json:"HeadIcon"`
	ID                float64   `json:"ID"`
	IncomeParam       float64   `json:"IncomeParam"`
	ItemID            float64   `json:"ItemID"`
	LandType          string    `json:"LandType"`
	MidIcon           string    `json:"MidIcon"`
	MiniIcon          string    `json:"MiniIcon"`
	Name              hash.Hash `json:"Name"`
	PlanetType        string    `json:"PlanetType"`
	Rarity            float64   `json:"Rarity"`
	Skill1List        []float64 `json:"Skill1List"`
	Skill2List        []float64 `json:"Skill2List"`
}
type PlanetFesAvatarAccessor struct {
	_data            []PlanetFesAvatar
	_dataAnimConfig  map[string]PlanetFesAvatar
	_dataBody        map[string]PlanetFesAvatar
	_dataCargoIcon   map[string]PlanetFesAvatar
	_dataDescription map[string]PlanetFesAvatar
	_dataHeadIcon    map[string]PlanetFesAvatar
	_dataID          map[float64]PlanetFesAvatar
	_dataItemID      map[float64]PlanetFesAvatar
	_dataMidIcon     map[string]PlanetFesAvatar
	_dataMiniIcon    map[string]PlanetFesAvatar
}

// LoadData retrieves the data. Must be called before PlanetFesAvatar.GroupData
func (a *PlanetFesAvatarAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesAvatar.json")
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
func (a *PlanetFesAvatarAccessor) Raw() ([]PlanetFesAvatar, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesAvatar{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlanetFesAvatarAccessor.LoadData to preload everything
func (a *PlanetFesAvatarAccessor) GroupData() {
	for _, d := range a._data {
		a._dataAnimConfig[d.AnimConfig] = d
		a._dataBody[d.Body] = d
		a._dataCargoIcon[d.CargoIcon] = d
		a._dataDescription[d.Description] = d
		a._dataHeadIcon[d.HeadIcon] = d
		a._dataID[d.ID] = d
		a._dataItemID[d.ItemID] = d
		a._dataMidIcon[d.MidIcon] = d
		a._dataMiniIcon[d.MiniIcon] = d
	}
}

// ByAnimConfig returns the PlanetFesAvatar uniquely identified by AnimConfig
//
// Error is only non-nil if the source errors out
func (a *PlanetFesAvatarAccessor) ByAnimConfig(identifier string) (PlanetFesAvatar, error) {
	if a._dataAnimConfig == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesAvatar{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAnimConfig[identifier], nil
}

// ByBody returns the PlanetFesAvatar uniquely identified by Body
//
// Error is only non-nil if the source errors out
func (a *PlanetFesAvatarAccessor) ByBody(identifier string) (PlanetFesAvatar, error) {
	if a._dataBody == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesAvatar{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBody[identifier], nil
}

// ByCargoIcon returns the PlanetFesAvatar uniquely identified by CargoIcon
//
// Error is only non-nil if the source errors out
func (a *PlanetFesAvatarAccessor) ByCargoIcon(identifier string) (PlanetFesAvatar, error) {
	if a._dataCargoIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesAvatar{}, err
			}
		}
		a.GroupData()
	}
	return a._dataCargoIcon[identifier], nil
}

// ByDescription returns the PlanetFesAvatar uniquely identified by Description
//
// Error is only non-nil if the source errors out
func (a *PlanetFesAvatarAccessor) ByDescription(identifier string) (PlanetFesAvatar, error) {
	if a._dataDescription == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesAvatar{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDescription[identifier], nil
}

// ByHeadIcon returns the PlanetFesAvatar uniquely identified by HeadIcon
//
// Error is only non-nil if the source errors out
func (a *PlanetFesAvatarAccessor) ByHeadIcon(identifier string) (PlanetFesAvatar, error) {
	if a._dataHeadIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesAvatar{}, err
			}
		}
		a.GroupData()
	}
	return a._dataHeadIcon[identifier], nil
}

// ByID returns the PlanetFesAvatar uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesAvatarAccessor) ByID(identifier float64) (PlanetFesAvatar, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesAvatar{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByItemID returns the PlanetFesAvatar uniquely identified by ItemID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesAvatarAccessor) ByItemID(identifier float64) (PlanetFesAvatar, error) {
	if a._dataItemID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesAvatar{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemID[identifier], nil
}

// ByMidIcon returns the PlanetFesAvatar uniquely identified by MidIcon
//
// Error is only non-nil if the source errors out
func (a *PlanetFesAvatarAccessor) ByMidIcon(identifier string) (PlanetFesAvatar, error) {
	if a._dataMidIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesAvatar{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMidIcon[identifier], nil
}

// ByMiniIcon returns the PlanetFesAvatar uniquely identified by MiniIcon
//
// Error is only non-nil if the source errors out
func (a *PlanetFesAvatarAccessor) ByMiniIcon(identifier string) (PlanetFesAvatar, error) {
	if a._dataMiniIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesAvatar{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMiniIcon[identifier], nil
}
