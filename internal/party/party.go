package party

import "github.com/loneJogger/go-dungeon-crawler/internal/combat"

type Party struct {
	Members []*combat.Character
	Gold    int
}

func New() *Party {
	knight := combat.NewKnight()
	witch := combat.NewWitch()
	priest := combat.NewPriest()
	return &Party{
		Members: []*combat.Character{
			&knight.Character,
			&witch.Character,
			&priest.Character,
		},
		Gold: 0,
	}
}
