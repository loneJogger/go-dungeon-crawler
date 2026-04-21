package ui

import (
	"image"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	charDelay        = 5
	charSize         = 16
	charsPerRow      = 16
	DialogPadding    = 8
	DialogLineHeight = 20
	DialogMaxChars   = 19
	DialogMaxLines   = 4
)

type TextScroll struct {
	pages       []string
	pageIndex   int
	charIndex   int
	timer       int
	waitingNext bool
	done        bool
	OnComplete  func()
	Color       [4]float32
}

func NewTextScroll(text string, onComplete func()) *TextScroll {
	return &TextScroll{
		pages:      paginate(text),
		OnComplete: onComplete,
		Color:      [4]float32{1, 1, 1, 1},
	}
}

func (t *TextScroll) Update() {
	if t.done {
		return
	}

	if t.waitingNext {
		if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
			if t.pageIndex < len(t.pages)-1 {
				t.pageIndex++
				t.charIndex = 0
				t.waitingNext = false
			} else {
				t.done = true
				if t.OnComplete != nil {
					t.OnComplete()
				}
			}
		}
		return
	}

	delay := charDelay
	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		delay = charDelay / 2
		if delay < 1 {
			delay = 1
		}
	}

	t.timer++
	if t.timer >= delay {
		t.timer = 0
		currentPage := t.pages[t.pageIndex]
		if t.charIndex < len(currentPage) {
			t.charIndex++
		} else {
			t.waitingNext = true
		}
	}
}

func (t *TextScroll) Draw(screen *ebiten.Image, font *ebiten.Image, boxX, boxY int) {
	currentPage := t.pages[t.pageIndex]
	col := 0
	row := 0

	for _, ch := range currentPage[:t.charIndex] {
		if ch == '\n' {
			col = 0
			row++
			continue
		}

		charIndex := int(ch) - 32
		sx := (charIndex % charsPerRow) * charSize
		sy := (charIndex / charsPerRow) * charSize

		src := font.SubImage(image.Rect(sx, sy, sx+charSize, sy+charSize)).(*ebiten.Image)

		op := &ebiten.DrawImageOptions{}
		op.ColorScale.Scale(t.Color[0], t.Color[1], t.Color[2], t.Color[3])
		op.GeoM.Translate(
			float64(boxX+DialogPadding+col*charSize),
			float64(boxY+DialogPadding+row*DialogLineHeight),
		)
		screen.DrawImage(src, op)

		col++
		if col >= DialogMaxChars {
			col = 0
			row++
		}
	}

	if t.waitingNext {
		// TODO: draw "press Z" prompt indicator
	}
}

func (t *TextScroll) IsDone() bool { return t.done }

func (t *TextScroll) IsWaiting() bool { return t.waitingNext }

// paginate splits text into pages that fit the dialog box.
// \n\n forces a manual page break, \n forces a line break.
// Lines longer than DialogMaxChars are word-wrapped automatically.
func paginate(text string) []string {
	manualPages := strings.Split(text, "\n\n")
	var pages []string

	for _, block := range manualPages {
		lines := wordWrap(block)
		for i := 0; i < len(lines); i += DialogMaxLines {
			end := i + DialogMaxLines
			if end > len(lines) {
				end = len(lines)
			}
			pages = append(pages, strings.Join(lines[i:end], "\n"))
		}
	}

	return pages
}

// wordWrap breaks a block of text into lines no longer than DialogMaxChars,
// splitting at word boundaries. Existing \n are preserved.
func wordWrap(text string) []string {
	var lines []string
	for _, line := range strings.Split(text, "\n") {
		words := strings.Fields(line)
		current := ""
		for _, word := range words {
			if len(current) == 0 {
				current = word
			} else if len(current)+1+len(word) <= DialogMaxChars {
				current += " " + word
			} else {
				lines = append(lines, current)
				current = word
			}
		}
		if len(current) > 0 {
			lines = append(lines, current)
		}
	}
	return lines
}
