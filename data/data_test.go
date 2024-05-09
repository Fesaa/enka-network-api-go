package data

import (
	"log/slog"
	"testing"

	"github.com/Fesaa/enka-network-api-go/localization"
)

var c EnkaData
var l *localization.Localization

func TestCorrectLoading(t *testing.T) {
	if c == nil {
		m, e := New(slog.Default())
		if e != nil {
			t.Error(e)
			t.FailNow()
		}
		c = m
	}

	if l == nil {
		localization.Init(slog.Default())
		l = localization.Get()
	}

	//Jingliu
	d := c.GetStarRailCharacterData("1212")
	if d == nil {
		t.Error("Expected some data got nil")
		t.FailNow()
	}

	if d.Name() != "Jingliu" {
		t.Errorf("Expected Jingliu got %s", d.Name())
		t.FailNow()
	}

	if d.Path != "Destruction" {
		t.Errorf("Expected Destruction got %s", d.Path)
		t.FailNow()
	}

}
