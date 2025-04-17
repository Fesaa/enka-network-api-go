package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type CutsceneActor struct {
	ActorID               string `json:"ActorID"`
	ActorModelPath        string `json:"ActorModelPath"`
	ResidentEffectKey     string `json:"ResidentEffectKey"`
	ResidentPossessionKey string `json:"ResidentPossessionKey"`
}
type CutsceneActorAccessor struct {
	_data               []CutsceneActor
	_dataActorID        map[string]CutsceneActor
	_dataActorModelPath map[string]CutsceneActor
}

// LoadData retrieves the data. Must be called before CutsceneActor.GroupData
func (a *CutsceneActorAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/CutsceneActor.json")
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
func (a *CutsceneActorAccessor) Raw() ([]CutsceneActor, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []CutsceneActor{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with CutsceneActorAccessor.LoadData to preload everything
func (a *CutsceneActorAccessor) GroupData() {
	for _, d := range a._data {
		a._dataActorID[d.ActorID] = d
		a._dataActorModelPath[d.ActorModelPath] = d
	}
}

// ByActorID returns the CutsceneActor uniquely identified by ActorID
//
// Error is only non-nil if the source errors out
func (a *CutsceneActorAccessor) ByActorID(identifier string) (CutsceneActor, error) {
	if a._dataActorID == nil {
		err := a.LoadData()
		if err != nil {
			return CutsceneActor{}, err
		}
		a.GroupData()
	}
	return a._dataActorID[identifier], nil
}

// ByActorModelPath returns the CutsceneActor uniquely identified by ActorModelPath
//
// Error is only non-nil if the source errors out
func (a *CutsceneActorAccessor) ByActorModelPath(identifier string) (CutsceneActor, error) {
	if a._dataActorModelPath == nil {
		err := a.LoadData()
		if err != nil {
			return CutsceneActor{}, err
		}
		a.GroupData()
	}
	return a._dataActorModelPath[identifier], nil
}
