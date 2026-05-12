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
- implement items menu
- implment equipment menu
- implement save menu

### Management

- Inventory, Money, State?
- Save and Continue functionality

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