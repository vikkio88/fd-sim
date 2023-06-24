package widgets

import "fyne.io/fyne/v2/widget"

func NewTSegment(text string) *widget.TextSegment {
	return &widget.TextSegment{Text: text}
}

func NewSepSegment() *widget.SeparatorSegment {
	return &widget.SeparatorSegment{}
}
