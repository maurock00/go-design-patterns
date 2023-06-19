package proxy

import (
	"strings"
	"unicode"
)

type TextRange struct {
	Start, End int
	Capitalize bool
}

func (t TextRange) Covers(position int) bool {
	return position >= t.Start && position <= t.End
}

type FormattedText struct {
	plainText  string
	formatting []*TextRange
}

func NewFormattedText(plainText string) *FormattedText {
	return &FormattedText{plainText: plainText}
}

func (f *FormattedText) Range(start, end int) *TextRange {
	r := &TextRange{Start: start, End: end, Capitalize: false}
	f.formatting = append(f.formatting, r)
	return r
}

func (f *FormattedText) String() string {
	sb := strings.Builder{}

	for i := 0; i < len(f.plainText); i++ {
		c := f.plainText[i]
		for _, r := range f.formatting {
			if r.Covers(i) && r.Capitalize {
				c = uint8(unicode.ToUpper(rune(c)))
			}
		}
		sb.WriteRune(rune(c))
	}

	return sb.String()
}

func main() {
	formatter := NewFormattedText("Esto es un textio de prueba")

	formatter.Range(1, 4).Capitalize = true
	formatter.Range(10, 12).Capitalize = true

	println(formatter.String())
}
