package characters

import (
	"github.com/loneJogger/go-dungeon-crawler/internal/combat/items"
)

type Witch struct {
	Character
}

func NewWitch() *Witch {
	return &Witch{
		Character: Character{
			Name:          "Isolde",
			Level:         15,
			Experience:    0,
			TotalHP:       72,
			CurrentHP:     72,
			TotalMP:       98,
			CurrentMP:     98,
			Strength:      12,
			Intelligence:  52,
			Defense:       14,
			Spirit:        38,
			Dexterity:     22,
			Accuracy:      20,
			Luck:          18,
			Actions:       []*Action{},
			Equipment:     []*items.Equipement{},
			StatusEffects: []*StatusEffect{},
			Spells:        []*Spell{},
			Resistences: Resistences{
				Fire:  1.125,
				Water: 0.875,
				Earth: 1.0,
				Air:   1.0,
				Light: 0.75,
				Dark:  1.5,
			},
		},
	}
}
