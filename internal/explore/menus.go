package explore

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/combat"
	"github.com/loneJogger/go-dungeon-crawler/internal/config"
	"github.com/loneJogger/go-dungeon-crawler/internal/ctx"
	"github.com/loneJogger/go-dungeon-crawler/internal/ui"
)

func BuildCharacterOverviewMenu(c *ctx.GameContext, stack *ui.MenuStack) *ui.Menu {
	items := make([]ui.MenuItem, len(c.Party.Members))
	for i, member := range c.Party.Members {
		char := member
		items[i] = ui.MenuItem{
			HeightPx: 3*ui.DialogLineHeight + 8,
			OnSelect: func() { stack.Push(buildCharacterStatsMenu(char)) },
			DrawItem: func(screen *ebiten.Image, font *ebiten.Image, x, y int, focused, cursorOn bool) {
				cursor := " "
				if focused && cursorOn {
					cursor = ">"
				}
				ui.DrawText(screen, font, cursor, x, y)
				ui.DrawText(screen, font, fmt.Sprintf("%-7sL%2d", char.Name, char.Level), x+16, y)
				ui.DrawText(screen, font, fmt.Sprintf("HP %3d/%3d", char.CurrentHP, char.TotalHP), x+16, y+ui.DialogLineHeight)
				ui.DrawText(screen, font, fmt.Sprintf("MP %3d/%3d", char.CurrentMP, char.TotalMP), x+16, y+2*ui.DialogLineHeight)
			},
		}
	}
	menu := ui.NewMenu(items)
	menu.NavSound = c.Assets.MenuNav
	menu.SelectSound = c.Assets.MenuSelect
	return menu
}

func drawHoverContent(screen *ebiten.Image, c *ctx.GameContext, x, y, focused int) {
	switch focused {
	case 0: // Status
		drawStatusContent(screen, c.Assets.Font, c, x, y)
	case 1: // Items
		drawItemsContent(screen, c.Assets.Font, c.Party.Gold, x, y)
	}
}

func drawStatusContent(screen *ebiten.Image, font *ebiten.Image, c *ctx.GameContext, x, y int) {
	for i, char := range c.Party.Members {
		rowY := y + i*(3*ui.DialogLineHeight+8)
		ui.DrawText(screen, font, fmt.Sprintf("%-7sL%2d", char.Name, char.Level), x+16, rowY)
		ui.DrawText(screen, font, fmt.Sprintf("HP %3d/%3d", char.CurrentHP, char.TotalHP), x+16, rowY+ui.DialogLineHeight)
		ui.DrawText(screen, font, fmt.Sprintf("MP %3d/%3d", char.CurrentMP, char.TotalMP), x+16, rowY+2*ui.DialogLineHeight)
	}
}

func drawItemsContent(screen *ebiten.Image, font *ebiten.Image, gold int, x, y int) {
	ui.DrawText(screen, font, "No items", x, y)
	goldY := config.ScreenHeight - ui.DialogLineHeight - 16
	ui.DrawText(screen, font, fmt.Sprintf("Gold: %d", gold), x, goldY)
}

func buildCharacterStatsMenu(char *combat.Character) *ui.Menu {
	noop := func() {}
	menu := ui.NewMenu([]ui.MenuItem{
		{Label: char.Name, OnSelect: noop},
		{Label: fmt.Sprintf("Lv. %3d", char.Level), OnSelect: noop},
		{Label: fmt.Sprintf("HP %3d/%3d", char.CurrentHP, char.TotalHP), OnSelect: noop},
		{Label: fmt.Sprintf("MP %3d/%3d", char.CurrentMP, char.TotalMP), OnSelect: noop},
		{Label: fmt.Sprintf("STR %3d", char.Strength), OnSelect: noop},
		{Label: fmt.Sprintf("INT %3d", char.Intelligence), OnSelect: noop},
		{Label: fmt.Sprintf("DEF %3d", char.Defense), OnSelect: noop},
		{Label: fmt.Sprintf("SPR %3d", char.Spirit), OnSelect: noop},
		{Label: fmt.Sprintf("DEX %3d", char.Dexterity), OnSelect: noop},
		{Label: fmt.Sprintf("LCK %3d", char.Luck), OnSelect: noop},
	})
	menu.Static = true
	return menu
}
