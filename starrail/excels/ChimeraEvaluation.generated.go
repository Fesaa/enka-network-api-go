package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type ChimeraEvaluation struct {
	ConditionJson  string                 `json:"ConditionJson"`
	EvaluationDesc map[string]json.Number `json:"EvaluationDesc"`
	EvaluationID   json.Number            `json:"EvaluationID"`
	EvaluationName map[string]json.Number `json:"EvaluationName"`
	GroupID        json.Number            `json:"GroupID"`
}
type ChimeraEvaluationAccessor struct {
	_data              []ChimeraEvaluation
	_dataConditionJson map[string]ChimeraEvaluation
	_dataEvaluationID  map[json.Number]ChimeraEvaluation
}

// LoadData retrieves the data. Must be called before ChimeraEvaluation.GroupData
func (a *ChimeraEvaluationAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ChimeraEvaluation.json")
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
func (a *ChimeraEvaluationAccessor) Raw() ([]ChimeraEvaluation, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ChimeraEvaluation{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ChimeraEvaluationAccessor.LoadData to preload everything
func (a *ChimeraEvaluationAccessor) GroupData() {
	a._dataConditionJson = map[string]ChimeraEvaluation{}
	a._dataEvaluationID = map[json.Number]ChimeraEvaluation{}
	for _, d := range a._data {
		a._dataConditionJson[d.ConditionJson] = d
		a._dataEvaluationID[d.EvaluationID] = d
	}
}

// ByConditionJson returns the ChimeraEvaluation uniquely identified by ConditionJson
//
// Error is only non-nil if the source errors out
func (a *ChimeraEvaluationAccessor) ByConditionJson(identifier string) (ChimeraEvaluation, error) {
	if a._dataConditionJson == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ChimeraEvaluation{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConditionJson[identifier], nil
}

// ByEvaluationID returns the ChimeraEvaluation uniquely identified by EvaluationID
//
// Error is only non-nil if the source errors out
func (a *ChimeraEvaluationAccessor) ByEvaluationID(identifier json.Number) (ChimeraEvaluation, error) {
	if a._dataEvaluationID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ChimeraEvaluation{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEvaluationID[identifier], nil
}
