package ui

type Menu struct {
	Items   []MenuItem
	focused int
}

type MenuItem struct {
	Label    string
	OnSelect func()
}

func NewMenu(items []MenuItem) *Menu {
	return &Menu{Items: items}
}

func (m *Menu) MoveUp() {
	m.focused = (m.focused - 1 + len(m.Items)) % len(m.Items)
}

func (m *Menu) MoveDown() {
	m.focused = (m.focused + 1) % len(m.Items)
}

func (m *Menu) Select() { m.Items[m.focused].OnSelect() }

func (m *Menu) Focused() int { return m.focused }
