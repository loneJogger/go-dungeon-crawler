package characters

import (
	"github.com/loneJogger/go-dungeon-crawler/internal/combat/items"
)

type Knight struct {
	Character
}

func NewKnight() *Knight {

	return &Knight{
		Character: Character{
			Name:          "Bonnie",
			Level:         15,
			Experience:    0,
			TotalHP:       135,
			CurrentHP:     135,
			TotalMP:       0,
			CurrentMP:     0,
			Strength:      45,
			Intelligence:  10,
			Defense:       40,
			Spirit:        25,
			Dexterity:     15,
			Accuracy:      25,
			Luck:          15,
			Actions:       []*Action{},
			Equipment:     []*items.Equipement{},
			StatusEffects: []*StatusEffect{},
			Spells:        []*Spell{},
			Resistences: Resistences{
				Fire:  0.875,
				Water: 1.0,
				Earth: 1.125,
				Air:   1.0,
				Light: 1.25,
				Dark:  0.875,
			},
		},
	}
}
