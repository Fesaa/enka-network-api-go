package data

import (
	"github.com/Fesaa/enka-network-api-go/starrail"
	"log/slog"
	"testing"
)

var c EnkaData

func TestCorrectLoading(t *testing.T) {
	if c == nil {
		m, e := New(slog.Default())
		if e != nil {
			t.Error(e)
			t.FailNow()
		}
		c = m
	}

	//Jingliu
	d := c.StarRailData().CharacterData("1212")
	if d == nil {
		t.Error("Expected some data got nil")
		t.FailNow()
	}

	if d.Name() != "940795027444042330" {
		t.Errorf("Expected Jingliu(940795027444042330) got %s", d.Name())
		t.FailNow()
	}

	if d.Path != "Destruction" {
		t.Errorf("Expected Destruction got %s", d.Path)
		t.FailNow()
	}

}

func TestRelicExcels(t *testing.T) {
	if c == nil {
		m, e := New(slog.Default())
		if e != nil {
			t.Error(e)
			t.FailNow()
		}
		c = m
	}

	// Castorice
	skillTree := c.StarRailData().Excels().SkillTree("1407")

	if skillTree[starrail.AnchorBasic].PointID != 1407001 {
		t.Errorf("Expected Basic(1407001) got %d", skillTree[starrail.AnchorBasic].PointID)
		t.FailNow()
	}

	if skillTree[starrail.AnchorTalent].PointID != 1407004 {
		t.Errorf("Expected Talent(1407004) got %d", skillTree[starrail.AnchorTalent].PointID)
		t.FailNow()
	}

	if skillTree[starrail.AnchorMemoTalent].PointID != 1407302 {
		t.Errorf("Expected MemoTalent(1407302) got %d", skillTree[starrail.AnchorMemoTalent].PointID)
		t.FailNow()
	}

	if skillTree[starrail.AnchorMajor3].PointID != 1407103 {
		t.Errorf("Expected Major3(1407103) got %d", skillTree[starrail.AnchorMajor3].PointID)
		t.FailNow()
	}

	set, ok := c.StarRailData().Excels().RelicSetConfig("101")
	if !ok {
		t.Error("Expected some data got nil")
		t.FailNow()
	}

	if set.SetIconPath != "SpriteOutput/ItemIcon/71000.png" {
		t.Errorf("Expected SpriteOutput/ItemIcon/71000.png got %s", set.SetIconPath)
		t.FailNow()
	}

	setItem, ok := c.StarRailData().Excels().RelicConfig("31011")
	if !ok {
		t.Error("Expected some data got nil")
		t.FailNow()
	}

	if setItem.Rarity != "CombatPowerRelicRarity2" {
		t.Errorf("Expected CombatPowerRelicRarity2 got %s", setItem.Rarity)
		t.FailNow()
	}

}
