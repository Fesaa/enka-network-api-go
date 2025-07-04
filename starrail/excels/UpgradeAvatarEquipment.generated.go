package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type UpgradeAvatarEquipment struct {
	ELADHEIOONA float64 `json:"ELADHEIOONA"`
	FJIFJPGEJID string  `json:"FJIFJPGEJID"`
}
type UpgradeAvatarEquipmentAccessor struct {
	_data            []UpgradeAvatarEquipment
	_dataELADHEIOONA map[float64]UpgradeAvatarEquipment
	_dataFJIFJPGEJID map[string]UpgradeAvatarEquipment
}

// LoadData retrieves the data. Must be called before UpgradeAvatarEquipment.GroupData
func (a *UpgradeAvatarEquipmentAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/UpgradeAvatarEquipment.json")
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
func (a *UpgradeAvatarEquipmentAccessor) Raw() ([]UpgradeAvatarEquipment, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []UpgradeAvatarEquipment{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with UpgradeAvatarEquipmentAccessor.LoadData to preload everything
func (a *UpgradeAvatarEquipmentAccessor) GroupData() {
	a._dataELADHEIOONA = map[float64]UpgradeAvatarEquipment{}
	a._dataFJIFJPGEJID = map[string]UpgradeAvatarEquipment{}
	for _, d := range a._data {
		a._dataELADHEIOONA[d.ELADHEIOONA] = d
		a._dataFJIFJPGEJID[d.FJIFJPGEJID] = d
	}
}

// ByELADHEIOONA returns the UpgradeAvatarEquipment uniquely identified by ELADHEIOONA
//
// Error is only non-nil if the source errors out
func (a *UpgradeAvatarEquipmentAccessor) ByELADHEIOONA(identifier float64) (UpgradeAvatarEquipment, error) {
	if a._dataELADHEIOONA == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return UpgradeAvatarEquipment{}, err
			}
		}
		a.GroupData()
	}
	return a._dataELADHEIOONA[identifier], nil
}

// ByFJIFJPGEJID returns the UpgradeAvatarEquipment uniquely identified by FJIFJPGEJID
//
// Error is only non-nil if the source errors out
func (a *UpgradeAvatarEquipmentAccessor) ByFJIFJPGEJID(identifier string) (UpgradeAvatarEquipment, error) {
	if a._dataFJIFJPGEJID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return UpgradeAvatarEquipment{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFJIFJPGEJID[identifier], nil
}
