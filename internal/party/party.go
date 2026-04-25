package party

import "github.com/loneJogger/go-dungeon-crawler/internal/combat/characters"

type Party struct {
	Members []*characters.Character
	Gold    int
}

func New() *Party {
	knight := characters.NewKnight()
	witch := characters.NewWitch()
	priest := characters.NewPriest()
	return &Party{
		Members: []*characters.Character{
			&knight.Character,
			&witch.Character,
			&priest.Character,
		},
		Gold: 0,
	}
}
