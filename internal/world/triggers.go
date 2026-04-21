package world

import (
	"image"

	"github.com/lafriks/go-tiled"
)

type Trigger struct {
	Name   string
	Bounds image.Rectangle
}

func LoadTriggers(m *tiled.Map) []Trigger {
	var triggers []Trigger
	for _, group := range m.ObjectGroups {
		if group.Name != "triggers" {
			continue
		}
		for _, obj := range group.Objects {
			triggers = append(triggers, Trigger{
				Name: obj.Name,
				Bounds: image.Rect(
					int(obj.X),
					int(obj.Y),
					int(obj.X+obj.Width),
					int(obj.Y+obj.Height),
				),
			})
		}
	}
	return triggers
}

func CheckTrigger(triggers []Trigger, x, y float64) *Trigger {
	pt := image.Pt(int(x), int(y))
	for i := range triggers {
		if pt.In(triggers[i].Bounds) {
			return &triggers[i]
		}
	}
	return nil
}
