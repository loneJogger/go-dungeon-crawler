package characters

import (
	"github.com/loneJogger/go-dungeon-crawler/internal/combat/items"
)

type Priest struct {
	Character
}

func NewPriest() *Priest {
	return &Priest{
		Character: Character{
			Name:          "Anthony",
			Level:         15,
			Experience:    0,
			TotalHP:       88,
			CurrentHP:     88,
			TotalMP:       72,
			CurrentMP:     72,
			Strength:      16,
			Intelligence:  38,
			Defense:       20,
			Spirit:        52,
			Dexterity:     18,
			Accuracy:      18,
			Luck:          22,
			Actions:       []*Action{},
			Equipment:     []*items.Equipement{},
			StatusEffects: []*StatusEffect{},
			Spells:        []*Spell{},
			Resistences: Resistences{
				Fire:  1.125,
				Water: 1.125,
				Earth: 1.125,
				Air:   1.125,
				Light: 1.25,
				Dark:  0.75,
			},
		},
	}
}
