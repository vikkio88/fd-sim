package widgets

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type SortType int

const (
	SortNone SortType = iota
	SortAscending
	SortDescending
)

type ListHeaderSort struct {
	ColNumber int
	Type      SortType
}

type ListColumn struct {
	Text             string
	Alignment        fyne.TextAlign
	CanToggleVisible bool
}

func NewListCol(text string, align fyne.TextAlign) ListColumn {
	return ListColumn{
		Text:      text,
		Alignment: align,
	}
}

type ListHeader struct {
	widget.BaseWidget

	DisableSorting bool

	OnColumnSortChanged         func(ListHeaderSort)
	OnColumnVisibilityChanged   func(int, bool)
	OnColumnVisibilityMenuShown func(*widget.PopUp)

	sort          ListHeaderSort
	columns       []ListColumn
	columnVisible []bool
	columnsLayout *ColumnsLayout

	columnsContainer *fyne.Container
	container        *fyne.Container
	popUpMenu        *fyne.Container
}

func NewListHeader(cols []ListColumn, layout *ColumnsLayout) *ListHeader {
	l := &ListHeader{
		columns:          cols,
		columnsLayout:    layout,
		columnsContainer: container.New(layout),
		DisableSorting:   true,
	}
	l.columnVisible = make([]bool, len(cols))
	for i := range l.columnVisible {
		l.columnVisible[i] = true
	}
	l.container = container.NewMax(canvas.NewRectangle(theme.BackgroundColor()), l.columnsContainer)
	l.ExtendBaseWidget(l)
	l.buildColumns()
	return l
}

func (l *ListHeader) SetColumnVisible(colNum int, visible bool) {
	if colNum >= len(l.columns) {
		log.Println("error: ListHeader.SetColumnVisible: column index out of range")
		return
	}
	if visible {
		l.columnsContainer.Objects[colNum].Show()
	} else {
		l.columnsContainer.Objects[colNum].Hide()
	}
	l.columnVisible[colNum] = visible
	l.columnsContainer.Refresh()
}

func (l *ListHeader) buildColumns() {
	for i, c := range l.columns {
		hdr := newColHeader(c, &l.DisableSorting)
		hdr.OnSortChanged = func(i int) func(SortType) {
			return func(sort SortType) { l.SetSorting(ListHeaderSort{ColNumber: i, Type: sort}) }
		}(i)
		hdr.OnTappedSecondary = l.TappedSecondary
		l.columnsContainer.Add(hdr)
	}
}

// Sets the sorting for the ListHeader. Will invoke
// OnColumnSortChanged if set.
func (l *ListHeader) SetSorting(sort ListHeaderSort) {
	if l.sort == sort {
		return
	}
	l.sort = sort
	for i, c := range l.columnsContainer.Objects {
		if i == sort.ColNumber {
			c.(*colHeader).Sort = sort.Type
		} else {
			c.(*colHeader).Sort = SortNone
		}
	}
	l.Refresh()
	if l.OnColumnSortChanged != nil {
		l.OnColumnSortChanged(sort)
	}
}

func (l *ListHeader) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(l.container)
}

func (l *ListHeader) TappedSecondary(e *fyne.PointEvent) {
	l.setupPopUpMenu()
	if len(l.popUpMenu.Objects) == 0 {
		return
	}
	pop := widget.NewPopUp(l.popUpMenu, fyne.CurrentApp().Driver().CanvasForObject(l))
	pop.ShowAtPosition(e.AbsolutePosition)
	if l.OnColumnVisibilityMenuShown != nil {
		l.OnColumnVisibilityMenuShown(pop)
	}
}

func (l *ListHeader) setupPopUpMenu() {
	if l.popUpMenu == nil {
		l.popUpMenu = container.New(&VboxCustomPadding{ExtraPad: -10})
		for i, c := range l.columns {
			if c.CanToggleVisible {
				l.popUpMenu.Add(widget.NewCheck(c.Text, l.createOnChangedCallbk(i)))
			}
		}
	}
	objIdx := 0
	for i, col := range l.columns {
		if col.CanToggleVisible {
			l.popUpMenu.Objects[objIdx].(*widget.Check).Checked = l.columnVisible[i]
			objIdx++
		}
	}
}

func (l *ListHeader) createOnChangedCallbk(colNum int) func(bool) {
	return func(val bool) {
		l.columnVisible[colNum] = val
		l.SetColumnVisible(colNum, val)
		if l.OnColumnVisibilityChanged != nil {
			l.OnColumnVisibilityChanged(colNum, val)
		}
	}
}

type colHeader struct {
	widget.BaseWidget

	Sort              SortType
	OnSortChanged     func(SortType)
	OnTappedSecondary func(*fyne.PointEvent)

	sortDisabled *bool
	columnCfg    ListColumn

	label             *widget.RichText
	sortIcon          *widget.Icon
	sortIconNegSpacer fyne.CanvasObject
	container         *fyne.Container
}

func newColHeader(columnCfg ListColumn, sortDisabled *bool) *colHeader {
	c := &colHeader{columnCfg: columnCfg, sortDisabled: sortDisabled}
	c.ExtendBaseWidget(c)

	c.label = widget.NewRichTextWithText(columnCfg.Text)
	c.label.Segments[0].(*widget.TextSegment).Style = widget.RichTextStyle{
		TextStyle: fyne.TextStyle{Bold: true},
		Alignment: columnCfg.Alignment,
	}
	c.sortIcon = widget.NewIcon(theme.MenuDropDownIcon())
	// hack to remove extra icon space
	// should be hidden whenever sortIcon is hidden
	c.sortIconNegSpacer = NewHSpace(0)

	return c
}

func (c *colHeader) Tapped(*fyne.PointEvent) {
	if *c.sortDisabled {
		return
	}
	switch c.Sort {
	case SortNone:
		c.Sort = SortAscending
	case SortAscending:
		c.Sort = SortDescending
	case SortDescending:
		c.Sort = SortNone
	default:
		log.Println("notReached colHeader.Tapped")
	}
	c.Refresh()
	if c.OnSortChanged != nil {
		c.OnSortChanged(c.Sort)
	}
}

func (c *colHeader) TappedSecondary(e *fyne.PointEvent) {
	if c.OnTappedSecondary != nil {
		c.OnTappedSecondary(e)
	}
}

func (c *colHeader) Refresh() {
	if c.Sort == SortDescending {
		c.sortIcon.Resource = theme.MenuDropDownIcon()
	} else {
		c.sortIcon.Resource = theme.MenuDropUpIcon()
	}

	if c.Sort > 0 && c.sortIcon.Hidden {
		c.sortIcon.Show()
		c.container.Add(c.sortIconNegSpacer)
	} else if (c.Sort == SortNone || *c.sortDisabled) && !c.sortIcon.Hidden {
		c.sortIcon.Hide()
		c.container.Remove(c.sortIconNegSpacer)
	}

	c.BaseWidget.Refresh()
}

func (c *colHeader) CreateRenderer() fyne.WidgetRenderer {
	if c.container == nil {
		c.container = container.New(&HboxCustomPadding{DisableThemePad: true, ExtraPad: -8})
		if c.columnCfg.Alignment != fyne.TextAlignLeading {
			c.container.Add(layout.NewSpacer())
		}
		c.container.Add(c.label)
		// c.container.Add(c.sortIcon)
		c.container.Add(c.sortIconNegSpacer)
		if c.columnCfg.Alignment == fyne.TextAlignCenter {
			c.container.Add(layout.NewSpacer())
		}
	}
	return widget.NewSimpleRenderer(c.container)
}

type HSpace struct {
	widget.BaseWidget

	Width float32
}

func NewHSpace(w float32) *HSpace {
	h := &HSpace{Width: w}
	h.ExtendBaseWidget(h)
	return h
}

func (h *HSpace) MinSize() fyne.Size {
	return fyne.NewSize(h.Width, 0)
}

func (h *HSpace) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(layout.NewSpacer())
}

// ColumnsLayout lays out a number of items into columns.
// There are two types of columns: fixed-width and variable width.
// A fixed width column is any with a non-negative width and will
// be laid out with that width. A variable width column is created
// by using any negative number as its width. Variable width columns
// are all laid out with the same width, splitting the "leftover" space
// equally between themselves after accounting for the fixed-width columns.
// Hidden items are not shown and take up 0 space.
type ColumnsLayout struct {
	ColumnWidths []float32
}

func NewColumnsLayout(widths []float32) *ColumnsLayout {
	return &ColumnsLayout{ColumnWidths: widths}
}

func (c *ColumnsLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	var width float32
	var height float32
	for i := 0; i < len(objects); i++ {
		if !objects[i].Visible() {
			continue
		}
		s := objects[i].MinSize()
		height = fyne.Max(height, s.Height)
		w := s.Width
		if i < len(c.ColumnWidths) && c.ColumnWidths[i] > w {
			w = c.ColumnWidths[i]
		}
		width += w
	}
	return fyne.NewSize(width, height)
}

func (c *ColumnsLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	var fixedW float32
	var expandObjCount int
	for i := 0; i < min(len(objects), len(c.ColumnWidths)); i++ {
		if !objects[i].Visible() {
			continue
		}
		if c.ColumnWidths[i] < 0 {
			expandObjCount++
		} else {
			itemW := objects[i].MinSize().Width
			fixedW += fyne.Max(itemW, c.ColumnWidths[i])
		}
	}
	extraW := size.Width - fixedW
	expandObjW := extraW / float32(expandObjCount)

	var x float32
	for i := 0; i < len(objects); i++ {
		if !objects[i].Visible() {
			continue
		}

		w := objects[i].MinSize().Width
		if i >= len(c.ColumnWidths) || c.ColumnWidths[i] < 0 {
			// expanding width column
			w = fyne.Max(expandObjW, w)
		} else {
			w = fyne.Max(c.ColumnWidths[i], w)
		}
		objects[i].Resize(fyne.NewSize(w, size.Height))
		objects[i].Move(fyne.NewPos(x, 0))
		x += w
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
