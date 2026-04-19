package ui

type Menu struct {
	Items   []MenuItem
	focused int
}

type MenuItem struct {
	Label    string
	OnSelect func()
}

func (m *Menu) MoveUp() {}

func (m *Menu) MoveDown() {}

func (m *Menu) Select() { m.Items[m.focused].OnSelect() }

func (m *Menu) Focused() int { return m.focused }
