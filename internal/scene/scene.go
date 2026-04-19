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
