# go-dungeon-crawler

A retro JRPG built in Go with Ebitengine. Learning project.

## Stack

- [Ebitengine](https://ebitengine.org/) — rendering
- [go-tiled](https://github.com/lafriks/go-tiled) — Tiled map loading

## Run

```
go run ./cmd/main.go
```

## Controls

| Key | Action |
|-----|--------|
| Arrow keys | Move |
| Z | Interact / Confirm |
| X | Cancel |
| A | Open menu |

## Next Steps

### UI

- update string parsing to handle color syntax
- move text colors, improve palette
- main menu layout
- implement items menu

### Management

- Inventory, Money, State?

### Explore
- Expand interactions (give items, ask questions)

### Battle
- ATB system
- Tagerting UI
- Add Actions to characters
- Create Enemies
- Setup battle scene phases and interactions
- Enemy behavior

### Game
- Saving
- refactor Audio handling