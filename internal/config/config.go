package config

// global constants
const TileSize = 16
const ScreenWidth = 320
const ScreenHeight = 240

// text colors — add new entries here to make them available as %COLOR NAME% tags
var TextColors = map[string][4]float32{
	"DEFAULT":  {1, 1, 1, 1},
	"PINK":     {0.98, 0.653, 0.724, 1},
	"MINT":     {0.641, 0.949, 0.678, 1},
	"ICY_BLUE": {0.522, 0.682, 0.949, 1},
	"BLOODY":   {0.6, 0.063, 0.047, 1},
}
