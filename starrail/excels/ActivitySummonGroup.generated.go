package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ActivitySummonGroup struct {
	ActivityModuleID         float64     `json:"ActivityModuleID"`
	AvatarSkillList          []float64   `json:"AvatarSkillList"`
	BackgroundTrashImageList []string    `json:"BackgroundTrashImageList"`
	Desc                     hash.Hash   `json:"Desc"`
	EntranceID               float64     `json:"EntranceID"`
	GotoID                   float64     `json:"GotoID"`
	GroupID                  float64     `json:"GroupID"`
	ImagePath                string      `json:"ImagePath"`
	MappingInfo              float64     `json:"MappingInfo"`
	MasterImagePath          []string    `json:"MasterImagePath"`
	MazeBuffID               float64     `json:"MazeBuffID"`
	MiniImagePath            string      `json:"MiniImagePath"`
	MonsterDesc              hash.Hash   `json:"MonsterDesc"`
	MonsterEventID           float64     `json:"MonsterEventID"`
	MonsterImagePath         string      `json:"MonsterImagePath"`
	MonsterMiddleIcon        string      `json:"MonsterMiddleIcon"`
	MonsterName              hash.Hash   `json:"MonsterName"`
	MonsterSkillDescList     []hash.Hash `json:"MonsterSkillDescList"`
	MonsterSkillList         []float64   `json:"MonsterSkillList"`
	OriginalDesc             hash.Hash   `json:"OriginalDesc"`
	OriginalImagePath        string      `json:"OriginalImagePath"`
	OriginalMiniImagePath    string      `json:"OriginalMiniImagePath"`
	OriginalStageName        hash.Hash   `json:"OriginalStageName"`
	StageName                hash.Hash   `json:"StageName"`
	SubMissionID             float64     `json:"SubMissionID"`
}
type ActivitySummonGroupAccessor struct {
	_data                      []ActivitySummonGroup
	_dataActivityModuleID      map[float64]ActivitySummonGroup
	_dataEntranceID            map[float64]ActivitySummonGroup
	_dataGotoID                map[float64]ActivitySummonGroup
	_dataGroupID               map[float64]ActivitySummonGroup
	_dataImagePath             map[string]ActivitySummonGroup
	_dataMappingInfo           map[float64]ActivitySummonGroup
	_dataMazeBuffID            map[float64]ActivitySummonGroup
	_dataMiniImagePath         map[string]ActivitySummonGroup
	_dataMonsterEventID        map[float64]ActivitySummonGroup
	_dataMonsterImagePath      map[string]ActivitySummonGroup
	_dataMonsterMiddleIcon     map[string]ActivitySummonGroup
	_dataOriginalImagePath     map[string]ActivitySummonGroup
	_dataOriginalMiniImagePath map[string]ActivitySummonGroup
	_dataSubMissionID          map[float64]ActivitySummonGroup
}

// LoadData retrieves the data. Must be called before ActivitySummonGroup.GroupData
func (a *ActivitySummonGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivitySummonGroup.json")
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
func (a *ActivitySummonGroupAccessor) Raw() ([]ActivitySummonGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivitySummonGroup{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivitySummonGroupAccessor.LoadData to preload everything
func (a *ActivitySummonGroupAccessor) GroupData() {
	for _, d := range a._data {
		a._dataActivityModuleID[d.ActivityModuleID] = d
		a._dataEntranceID[d.EntranceID] = d
		a._dataGotoID[d.GotoID] = d
		a._dataGroupID[d.GroupID] = d
		a._dataImagePath[d.ImagePath] = d
		a._dataMappingInfo[d.MappingInfo] = d
		a._dataMazeBuffID[d.MazeBuffID] = d
		a._dataMiniImagePath[d.MiniImagePath] = d
		a._dataMonsterEventID[d.MonsterEventID] = d
		a._dataMonsterImagePath[d.MonsterImagePath] = d
		a._dataMonsterMiddleIcon[d.MonsterMiddleIcon] = d
		a._dataOriginalImagePath[d.OriginalImagePath] = d
		a._dataOriginalMiniImagePath[d.OriginalMiniImagePath] = d
		a._dataSubMissionID[d.SubMissionID] = d
	}
}

// ByActivityModuleID returns the ActivitySummonGroup uniquely identified by ActivityModuleID
//
// Error is only non-nil if the source errors out
func (a *ActivitySummonGroupAccessor) ByActivityModuleID(identifier float64) (ActivitySummonGroup, error) {
	if a._dataActivityModuleID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivitySummonGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataActivityModuleID[identifier], nil
}

// ByEntranceID returns the ActivitySummonGroup uniquely identified by EntranceID
//
// Error is only non-nil if the source errors out
func (a *ActivitySummonGroupAccessor) ByEntranceID(identifier float64) (ActivitySummonGroup, error) {
	if a._dataEntranceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivitySummonGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEntranceID[identifier], nil
}

// ByGotoID returns the ActivitySummonGroup uniquely identified by GotoID
//
// Error is only non-nil if the source errors out
func (a *ActivitySummonGroupAccessor) ByGotoID(identifier float64) (ActivitySummonGroup, error) {
	if a._dataGotoID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivitySummonGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGotoID[identifier], nil
}

// ByGroupID returns the ActivitySummonGroup uniquely identified by GroupID
//
// Error is only non-nil if the source errors out
func (a *ActivitySummonGroupAccessor) ByGroupID(identifier float64) (ActivitySummonGroup, error) {
	if a._dataGroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivitySummonGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGroupID[identifier], nil
}

// ByImagePath returns the ActivitySummonGroup uniquely identified by ImagePath
//
// Error is only non-nil if the source errors out
func (a *ActivitySummonGroupAccessor) ByImagePath(identifier string) (ActivitySummonGroup, error) {
	if a._dataImagePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivitySummonGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataImagePath[identifier], nil
}

// ByMappingInfo returns the ActivitySummonGroup uniquely identified by MappingInfo
//
// Error is only non-nil if the source errors out
func (a *ActivitySummonGroupAccessor) ByMappingInfo(identifier float64) (ActivitySummonGroup, error) {
	if a._dataMappingInfo == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivitySummonGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMappingInfo[identifier], nil
}

// ByMazeBuffID returns the ActivitySummonGroup uniquely identified by MazeBuffID
//
// Error is only non-nil if the source errors out
func (a *ActivitySummonGroupAccessor) ByMazeBuffID(identifier float64) (ActivitySummonGroup, error) {
	if a._dataMazeBuffID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivitySummonGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMazeBuffID[identifier], nil
}

// ByMiniImagePath returns the ActivitySummonGroup uniquely identified by MiniImagePath
//
// Error is only non-nil if the source errors out
func (a *ActivitySummonGroupAccessor) ByMiniImagePath(identifier string) (ActivitySummonGroup, error) {
	if a._dataMiniImagePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivitySummonGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMiniImagePath[identifier], nil
}

// ByMonsterEventID returns the ActivitySummonGroup uniquely identified by MonsterEventID
//
// Error is only non-nil if the source errors out
func (a *ActivitySummonGroupAccessor) ByMonsterEventID(identifier float64) (ActivitySummonGroup, error) {
	if a._dataMonsterEventID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivitySummonGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMonsterEventID[identifier], nil
}

// ByMonsterImagePath returns the ActivitySummonGroup uniquely identified by MonsterImagePath
//
// Error is only non-nil if the source errors out
func (a *ActivitySummonGroupAccessor) ByMonsterImagePath(identifier string) (ActivitySummonGroup, error) {
	if a._dataMonsterImagePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivitySummonGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMonsterImagePath[identifier], nil
}

// ByMonsterMiddleIcon returns the ActivitySummonGroup uniquely identified by MonsterMiddleIcon
//
// Error is only non-nil if the source errors out
func (a *ActivitySummonGroupAccessor) ByMonsterMiddleIcon(identifier string) (ActivitySummonGroup, error) {
	if a._dataMonsterMiddleIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivitySummonGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMonsterMiddleIcon[identifier], nil
}

// ByOriginalImagePath returns the ActivitySummonGroup uniquely identified by OriginalImagePath
//
// Error is only non-nil if the source errors out
func (a *ActivitySummonGroupAccessor) ByOriginalImagePath(identifier string) (ActivitySummonGroup, error) {
	if a._dataOriginalImagePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivitySummonGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataOriginalImagePath[identifier], nil
}

// ByOriginalMiniImagePath returns the ActivitySummonGroup uniquely identified by OriginalMiniImagePath
//
// Error is only non-nil if the source errors out
func (a *ActivitySummonGroupAccessor) ByOriginalMiniImagePath(identifier string) (ActivitySummonGroup, error) {
	if a._dataOriginalMiniImagePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivitySummonGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataOriginalMiniImagePath[identifier], nil
}

// BySubMissionID returns the ActivitySummonGroup uniquely identified by SubMissionID
//
// Error is only non-nil if the source errors out
func (a *ActivitySummonGroupAccessor) BySubMissionID(identifier float64) (ActivitySummonGroup, error) {
	if a._dataSubMissionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivitySummonGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSubMissionID[identifier], nil
}
