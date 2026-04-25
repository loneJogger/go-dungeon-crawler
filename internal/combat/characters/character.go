package characters

import (
	"github.com/loneJogger/go-dungeon-crawler/internal/combat/items"
)

const StatCap = 100

type Character struct {
	Name                                                                      string
	Level                                                                     int
	Experience                                                                int
	TotalHP, CurrentHP                                                        int
	TotalMP, CurrentMP                                                        int
	Strength, Defense, Intelligence, Spirit, Dexterity, Accuracy, Speed, Luck int
	Actions                                                                   []*Action
	Equipment                                                                 []*items.Equipement
	StatusEffects                                                             []*StatusEffect
	Spells                                                                    []*Spell
	Resistences                                                               Resistences
}

type Action struct {
	Name string
}

type Spell struct {
	Name string
}

type Element int

const (
	Fire Element = iota
	Water
	Earth
	Air
	Light
	Dark
)

type Resistences map[Element]float32

type StatusEffect struct {
	Name string
}
