package explore

import (
	"fmt"

	"github.com/loneJogger/go-dungeon-crawler/internal/combat/characters"
	"github.com/loneJogger/go-dungeon-crawler/internal/ctx"
	"github.com/loneJogger/go-dungeon-crawler/internal/ui"
)

func BuildCharacterListMenu(c *ctx.GameContext, stack *ui.MenuStack) *ui.Menu {
	items := make([]ui.MenuItem, len(c.Party.Members))
	for i, member := range c.Party.Members {
		m := member
		items[i] = ui.MenuItem{
			Label: m.Name,
			OnSelect: func() {
				stack.Push(buildCharacterStatsMenu(m))
			},
		}
	}
	menu := ui.NewMenu(items)
	menu.NavSound = c.Assets.MenuNav
	menu.SelectSound = c.Assets.MenuSelect
	return menu
}

func buildCharacterStatsMenu(char *characters.Character) *ui.Menu {
	noop := func() {}
	items := []ui.MenuItem{
		{Label: char.Name, OnSelect: noop},
		{Label: fmt.Sprintf("HP  %3d/%3d", char.CurrentHP, char.TotalHP), OnSelect: noop},
		{Label: fmt.Sprintf("MP  %3d/%3d", char.CurrentMP, char.TotalMP), OnSelect: noop},
		{Label: fmt.Sprintf("STR %3d", char.Strength), OnSelect: noop},
		{Label: fmt.Sprintf("INT %3d", char.Intelligence), OnSelect: noop},
		{Label: fmt.Sprintf("DEF %3d", char.Defense), OnSelect: noop},
		{Label: fmt.Sprintf("SPR %3d", char.Spirit), OnSelect: noop},
		{Label: fmt.Sprintf("DEX %3d", char.Dexterity), OnSelect: noop},
		{Label: fmt.Sprintf("LCK %3d", char.Luck), OnSelect: noop},
	}
	return ui.NewMenu(items)
}
