package data

import (
	"github.com/Fesaa/enka-network-api-go/starrail/excels"
	"github.com/rs/zerolog"
	"testing"
)

var c EnkaData

func TestCorrectLoading(t *testing.T) {
	if c == nil {
		m, e := New(zerolog.Nop())
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
		m, e := New(zerolog.Nop())
		if e != nil {
			t.Error(e)
			t.FailNow()
		}
		c = m
	}

	// Castorice
	accessor := excels.AvatarSkillTreeConfigAccessor{}
	skillTrees, err := accessor.Raw()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	castoriceData := map[string]excels.AvatarSkillTreeConfig{}
	for _, skillTree := range skillTrees {
		if skillTree.AvatarID != 1407 {
			continue
		}

		castoriceData[skillTree.Anchor] = skillTree
	}

	if castoriceData["Point01"].PointID != 1407001 {
		t.Errorf("Expected Basic(1407001) got %f", castoriceData["Point01"].PointID)
		t.FailNow()
	}

	if castoriceData["Point04"].PointID != 1407004 {
		t.Errorf("Expected Talent(1407004) got %f", castoriceData["Point04"].PointID)
		t.FailNow()
	}

	if castoriceData["Point20"].PointID != 1407302 {
		t.Errorf("Expected MemoTalent(1407302) got %f", castoriceData["Point20"].PointID)
		t.FailNow()
	}

	if castoriceData["Point08"].PointID != 1407103 {
		t.Errorf("Expected Major3(1407103) got %f", castoriceData["Point08"].PointID)
		t.FailNow()
	}

	relicAccessor := excels.RelicSetConfigAccessor{}
	set, err := relicAccessor.BySetID(101)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if set.SetIconPath != "SpriteOutput/ItemIcon/71000.png" {
		t.Errorf("Expected SpriteOutput/ItemIcon/71000.png got %s", set.SetIconPath)
		t.FailNow()
	}

	relicSetItemAccessor := excels.RelicConfigAccessor{}
	setItem, err := relicSetItemAccessor.ByID(31011)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if setItem.Rarity != "CombatPowerRelicRarity2" {
		t.Errorf("Expected CombatPowerRelicRarity2 got %s", setItem.Rarity)
		t.FailNow()
	}

}
