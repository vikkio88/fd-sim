package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewHyperlinkSegment(text string, onTapped func()) *HyperlinkSegment {
	return &HyperlinkSegment{
		Alignment: fyne.TextAlignLeading,
		Text:      text,
		OnTapped:  onTapped,
	}
}

type HyperlinkSegment struct {
	Alignment fyne.TextAlign
	Text      string
	OnTapped  func()
}

// Inline returns true as hyperlinks are inside other elements.
func (h *HyperlinkSegment) Inline() bool {
	return true
}

// Textual returns the content of this segment rendered to plain text.
func (h *HyperlinkSegment) Textual() string {
	return h.Text
}

// Visual returns the hyperlink widget required to render this segment.
func (h *HyperlinkSegment) Visual() fyne.CanvasObject {
	link := widget.NewHyperlink(h.Text, nil)
	link.OnTapped = h.OnTapped
	link.Alignment = h.Alignment
	return &fyne.Container{Layout: &unpadTextWidgetLayout{}, Objects: []fyne.CanvasObject{link}}
}

// Update applies the current state of this hyperlink segment to an existing visual.
func (h *HyperlinkSegment) Update(o fyne.CanvasObject) {
	link := o.(*fyne.Container).Objects[0].(*widget.Hyperlink)
	link.Text = h.Text
	link.URL = nil
	link.OnTapped = h.OnTapped
	link.Alignment = h.Alignment
	link.Refresh()
}

// Select tells the segment that the user is selecting the content between the two positions.
func (h *HyperlinkSegment) Select(begin, end fyne.Position) {
	// no-op: this will be added when we progress to editor
}

// SelectedText should return the text representation of any content currently selected through the Select call.
func (h *HyperlinkSegment) SelectedText() string {
	// no-op: this will be added when we progress to editor
	return ""
}

// Unselect tells the segment that the user is has cancelled the previous selection.
func (h *HyperlinkSegment) Unselect() {
	// no-op: this will be added when we progress to editor
}

type unpadTextWidgetLayout struct {
}

func (u *unpadTextWidgetLayout) Layout(o []fyne.CanvasObject, s fyne.Size) {
	pad := theme.InnerPadding() * -1
	pad2 := pad * -2

	o[0].Move(fyne.NewPos(pad, pad))
	o[0].Resize(s.Add(fyne.NewSize(pad2, pad2)))
}

func (u *unpadTextWidgetLayout) MinSize(o []fyne.CanvasObject) fyne.Size {
	pad := theme.InnerPadding() * 2
	return o[0].MinSize().Subtract(fyne.NewSize(pad, pad))
}
