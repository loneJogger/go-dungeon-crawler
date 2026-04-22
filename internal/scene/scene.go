package scene

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type SceneSwitcher interface {
	SetScene(Scene)
}

type SceneExiter interface {
	OnExit()
}

type SceneEnter interface {
	OnEnter()
}

type PlayerPositioner interface {
	SetPlayerPos(x, y float64)
}

type ExitConfig struct {
	TriggerName      string
	ReturnX, ReturnY float64
}
