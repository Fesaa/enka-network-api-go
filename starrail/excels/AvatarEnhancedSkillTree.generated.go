package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AvatarEnhancedSkillTree struct {
	AvatarID     float64   `json:"AvatarID"`
	Comment01    hash.Hash `json:"Comment01"`
	CommentIndex float64   `json:"CommentIndex"`
	DescAfter    hash.Hash `json:"DescAfter"`
	DescBefore   hash.Hash `json:"DescBefore"`
	SkillTreeID  float64   `json:"SkillTreeID"`
}
type AvatarEnhancedSkillTreeAccessor struct {
	_data            []AvatarEnhancedSkillTree
	_dataSkillTreeID map[float64]AvatarEnhancedSkillTree
}

// LoadData retrieves the data. Must be called before AvatarEnhancedSkillTree.GroupData
func (a *AvatarEnhancedSkillTreeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AvatarEnhancedSkillTree.json")
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
func (a *AvatarEnhancedSkillTreeAccessor) Raw() ([]AvatarEnhancedSkillTree, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AvatarEnhancedSkillTree{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AvatarEnhancedSkillTreeAccessor.LoadData to preload everything
func (a *AvatarEnhancedSkillTreeAccessor) GroupData() {
	a._dataSkillTreeID = map[float64]AvatarEnhancedSkillTree{}
	for _, d := range a._data {
		a._dataSkillTreeID[d.SkillTreeID] = d
	}
}

// BySkillTreeID returns the AvatarEnhancedSkillTree uniquely identified by SkillTreeID
//
// Error is only non-nil if the source errors out
func (a *AvatarEnhancedSkillTreeAccessor) BySkillTreeID(identifier float64) (AvatarEnhancedSkillTree, error) {
	if a._dataSkillTreeID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarEnhancedSkillTree{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSkillTreeID[identifier], nil
}
