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

- improve and expand text color palette
- main menu layout
- implement items menu

### Management

- Inventory, Money, State?

### Explore

- Expand interactions (NPCs give items, NPCs ask questions)

### Battle

- ATB system
- Tagerting UI
- Add Actions to characters
- Create Enemies
- Setup battle scene phases and interactions
- Enemy behavior

### Game

- Saving