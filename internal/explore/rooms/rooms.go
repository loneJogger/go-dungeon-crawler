package rooms

import (
	"github.com/loneJogger/go-dungeon-crawler/internal/explore/room"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
)

// SwitchFn is the callback rooms use to request a room transition.
// Explore.switchRoom satisfies this type.
type SwitchFn func(r *room.Room, tt transition.TransitionType, startX, startY float64)
