package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type FightFestPaper struct {
	CollectionBgPath   string                 `json:"CollectionBgPath"`
	CollectionFgPath   string                 `json:"CollectionFgPath"`
	GameAdFigurePath   string                 `json:"GameAdFigurePath"`
	InterviewBgPath    string                 `json:"InterviewBgPath"`
	InterviewFgPath    string                 `json:"InterviewFgPath"`
	IssueNumber        map[string]json.Number `json:"IssueNumber"`
	IssueNumberText    map[string]json.Number `json:"IssueNumberText"`
	MainBgPathList     []string               `json:"MainBgPathList"`
	MainFgPathList     []string               `json:"MainFgPathList"`
	MainPageDesc       map[string]json.Number `json:"MainPageDesc"`
	MainPageTitle      map[string]json.Number `json:"MainPageTitle"`
	PaperID            json.Number            `json:"PaperID"`
	UnlockSubMissionID json.Number            `json:"UnlockSubMissionID"`
}
type FightFestPaperAccessor struct {
	_data                 []FightFestPaper
	_dataCollectionBgPath map[string]FightFestPaper
	_dataCollectionFgPath map[string]FightFestPaper
	_dataInterviewFgPath  map[string]FightFestPaper
}

// LoadData retrieves the data. Must be called before FightFestPaper.GroupData
func (a *FightFestPaperAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FightFestPaper.json")
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
func (a *FightFestPaperAccessor) Raw() ([]FightFestPaper, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FightFestPaper{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with FightFestPaperAccessor.LoadData to preload everything
func (a *FightFestPaperAccessor) GroupData() {
	a._dataCollectionBgPath = map[string]FightFestPaper{}
	a._dataCollectionFgPath = map[string]FightFestPaper{}
	a._dataInterviewFgPath = map[string]FightFestPaper{}
	for _, d := range a._data {
		a._dataCollectionBgPath[d.CollectionBgPath] = d
		a._dataCollectionFgPath[d.CollectionFgPath] = d
		a._dataInterviewFgPath[d.InterviewFgPath] = d
	}
}

// ByCollectionBgPath returns the FightFestPaper uniquely identified by CollectionBgPath
//
// Error is only non-nil if the source errors out
func (a *FightFestPaperAccessor) ByCollectionBgPath(identifier string) (FightFestPaper, error) {
	if a._dataCollectionBgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestPaper{}, err
			}
		}
		a.GroupData()
	}
	return a._dataCollectionBgPath[identifier], nil
}

// ByCollectionFgPath returns the FightFestPaper uniquely identified by CollectionFgPath
//
// Error is only non-nil if the source errors out
func (a *FightFestPaperAccessor) ByCollectionFgPath(identifier string) (FightFestPaper, error) {
	if a._dataCollectionFgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestPaper{}, err
			}
		}
		a.GroupData()
	}
	return a._dataCollectionFgPath[identifier], nil
}

// ByInterviewFgPath returns the FightFestPaper uniquely identified by InterviewFgPath
//
// Error is only non-nil if the source errors out
func (a *FightFestPaperAccessor) ByInterviewFgPath(identifier string) (FightFestPaper, error) {
	if a._dataInterviewFgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestPaper{}, err
			}
		}
		a.GroupData()
	}
	return a._dataInterviewFgPath[identifier], nil
}
