package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MatchThreeAvatarSkillDialog struct {
	AvatarPic     string  `json:"AvatarPic"`
	EnvironmentID float64 `json:"EnvironmentID"`
	ID            float64 `json:"ID"`
}
type MatchThreeAvatarSkillDialogAccessor struct {
	_data   []MatchThreeAvatarSkillDialog
	_dataID map[float64]MatchThreeAvatarSkillDialog
}

// LoadData retrieves the data. Must be called before MatchThreeAvatarSkillDialog.GroupData
func (a *MatchThreeAvatarSkillDialogAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MatchThreeAvatarSkillDialog.json")
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
func (a *MatchThreeAvatarSkillDialogAccessor) Raw() ([]MatchThreeAvatarSkillDialog, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MatchThreeAvatarSkillDialog{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MatchThreeAvatarSkillDialogAccessor.LoadData to preload everything
func (a *MatchThreeAvatarSkillDialogAccessor) GroupData() {
	a._dataID = map[float64]MatchThreeAvatarSkillDialog{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the MatchThreeAvatarSkillDialog uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MatchThreeAvatarSkillDialogAccessor) ByID(identifier float64) (MatchThreeAvatarSkillDialog, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MatchThreeAvatarSkillDialog{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
