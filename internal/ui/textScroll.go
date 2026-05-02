package ui

import (
	"image"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/loneJogger/go-dungeon-crawler/internal/config"
)

const (
	charDelay        = 8
	charSize         = 16
	charsPerRow      = 16
	DialogPadding    = 8
	DialogLineHeight = 20
	DialogMaxChars   = 19
	DialogMaxLines   = 4
)

type ColoredChar struct {
	ch    rune
	color [4]float32
}

type TextScroll struct {
	pages       [][]ColoredChar
	pageIndex   int
	charIndex   int
	timer       int
	waitingNext bool
	done        bool
	OnComplete  func()
	BeepSound   *audio.Player
}

func NewTextScroll(text string, onComplete func()) *TextScroll {
	return &TextScroll{
		pages:      paginate(text),
		OnComplete: onComplete,
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
			if t.BeepSound != nil {
				t.BeepSound.Rewind()
				t.BeepSound.Play()
			}
		} else {
			t.waitingNext = true
		}
	}
}

func (t *TextScroll) Draw(screen *ebiten.Image, font *ebiten.Image, boxX, boxY int) {
	currentPage := t.pages[t.pageIndex]
	col := 0
	row := 0

	for _, cc := range currentPage[:t.charIndex] {
		if cc.ch == '\n' {
			col = 0
			row++
			continue
		}

		charIndex := int(cc.ch) - 32
		sx := (charIndex % charsPerRow) * charSize
		sy := (charIndex / charsPerRow) * charSize

		src := font.SubImage(image.Rect(sx, sy, sx+charSize, sy+charSize)).(*ebiten.Image)

		op := &ebiten.DrawImageOptions{}
		op.ColorScale.Scale(cc.color[0], cc.color[1], cc.color[2], cc.color[3])
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

// parseColorTags converts a raw string with {COLOR NAME} tags into a flat
// sequence of ColoredChar values. Tags are consumed and not rendered.
// An unclosed { (no matching closing }) is treated as a literal { character.
func parseColorTags(text string) []ColoredChar {
	defaultColor := config.TextColors["DEFAULT"]
	currentColor := defaultColor
	runes := []rune(text)
	var result []ColoredChar
	i := 0
	for i < len(runes) {
		if runes[i] == '{' {
			j := i + 1
			for j < len(runes) && runes[j] != '}' {
				j++
			}
			if j < len(runes) {
				tag := strings.TrimSpace(string(runes[i+1 : j]))
				parts := strings.Fields(tag)
				if len(parts) == 2 && strings.ToUpper(parts[0]) == "COLOR" {
					name := strings.ToUpper(parts[1])
					if c, ok := config.TextColors[name]; ok {
						currentColor = c
					}
				}
				i = j + 1
			} else {
				result = append(result, ColoredChar{'{', currentColor})
				i++
			}
		} else {
			result = append(result, ColoredChar{runes[i], currentColor})
			i++
		}
	}
	return result
}

// paginate splits text into pages that fit the dialog box.
// \n\n forces a manual page break, \n forces a line break.
// Lines longer than DialogMaxChars are word-wrapped automatically.
func paginate(text string) [][]ColoredChar {
	chars := parseColorTags(text)
	defaultColor := config.TextColors["DEFAULT"]

	// split on consecutive \n\n pairs for manual page breaks
	var blocks [][]ColoredChar
	start := 0
	for i := 0; i < len(chars)-1; i++ {
		if chars[i].ch == '\n' && chars[i+1].ch == '\n' {
			blocks = append(blocks, chars[start:i])
			start = i + 2
		}
	}
	blocks = append(blocks, chars[start:])

	var pages [][]ColoredChar
	for _, block := range blocks {
		lines := wordWrapColored(block)
		for i := 0; i < len(lines); i += DialogMaxLines {
			end := i + DialogMaxLines
			if end > len(lines) {
				end = len(lines)
			}
			var page []ColoredChar
			for li, line := range lines[i:end] {
				if li > 0 {
					page = append(page, ColoredChar{'\n', defaultColor})
				}
				page = append(page, line...)
			}
			pages = append(pages, page)
		}
	}

	return pages
}

// wordWrapColored breaks a block of ColoredChars into lines no longer than
// DialogMaxChars, splitting at word boundaries. Existing \n are preserved.
func wordWrapColored(chars []ColoredChar) [][]ColoredChar {
	var lines [][]ColoredChar

	// split on explicit \n into paragraphs
	var paragraphs [][]ColoredChar
	start := 0
	for i, cc := range chars {
		if cc.ch == '\n' {
			paragraphs = append(paragraphs, chars[start:i])
			start = i + 1
		}
	}
	paragraphs = append(paragraphs, chars[start:])

	for _, para := range paragraphs {
		// extract words as slices of ColoredChar
		var words [][]ColoredChar
		wordStart := -1
		for i, cc := range para {
			if cc.ch == ' ' {
				if wordStart >= 0 {
					words = append(words, para[wordStart:i])
					wordStart = -1
				}
			} else {
				if wordStart < 0 {
					wordStart = i
				}
			}
		}
		if wordStart >= 0 {
			words = append(words, para[wordStart:])
		}

		var currentLine []ColoredChar
		currentLen := 0
		for _, word := range words {
			wLen := len(word)
			if currentLen == 0 {
				currentLine = append(currentLine, word...)
				currentLen = wLen
			} else if currentLen+1+wLen <= DialogMaxChars {
				currentLine = append(currentLine, ColoredChar{' ', word[0].color})
				currentLine = append(currentLine, word...)
				currentLen += 1 + wLen
			} else {
				lines = append(lines, currentLine)
				currentLine = append([]ColoredChar{}, word...)
				currentLen = wLen
			}
		}
		lines = append(lines, currentLine)
	}

	return lines
}
