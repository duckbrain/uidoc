package uidoc

import (
	"fmt"

	"github.com/andlabs/ui"
)

var _ fmt.Formatter

type Doc struct {
	area, measureArea *ui.Area
	box               *ui.Box
	doc               Element
	width, height     float64
}

func NewDoc() *Doc {
	r := &Doc{}
	// Scrolling Area
	r.area = ui.NewScrollingArea(r, 400, 400)
	r.measureArea = ui.NewArea(&measureHandler{r})
	r.box = ui.NewVerticalBox()
	toolbarContainer := ui.NewHorizontalBox()
	toolbarHeightRetainer := ui.NewHorizontalSeparator()
	toolbarContainer.Append(toolbarHeightRetainer, false)
	toolbarContainer.Append(r.measureArea, true)
	r.box.Append(r.area, true)
	r.box.Append(toolbarContainer, false)
	return r
}

func (r *Doc) SetDocument(doc Element) {
	r.doc = doc
	r.Layout()
}

func (r *Doc) Layout() {
	r.layout(r.width)
}

func (r *Doc) layout(width float64) {
	if r.doc == nil {
		return
	}
	_, h := r.doc.Layout(width)
	r.width = width
	r.height = h
	r.area.QueueRedrawAll()
	r.area.SetSize(int(r.width), int(r.height))
	r.area.QueueRedrawAll()
}

func (r *Doc) Draw(a *ui.Area, dp *ui.AreaDrawParams) {
	if r.doc == nil {
		return
	}
	r.doc.Render(dp, 0, 0)
}

func (r *Doc) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {
}
func (r *Doc) MouseCrossed(a *ui.Area, left bool) {
}
func (r *Doc) DragBroken(a *ui.Area) {
}
func (r *Doc) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) {
	return false
}

// Wrap the box to make this element behave as a control

func (r *Doc) Destroy() {
	r.box.Destroy()
}
func (r *Doc) LibuiControl() uintptr {
	return r.box.LibuiControl()
}
func (r *Doc) Handle() uintptr {
	return r.box.Handle()
}
func (r *Doc) Show() {
	r.box.Show()
}
func (r *Doc) Hide() {
	r.box.Hide()
}
func (r *Doc) Enable() {
	r.box.Enable()
}
func (r *Doc) Disable() {
	r.box.Disable()
}

type measureHandler struct {
	parent *Doc
}

func (r *measureHandler) Draw(a *ui.Area, dp *ui.AreaDrawParams) {
	if r.parent.width != dp.AreaWidth {
		fmt.Sprintf("Size changed: %v\n", dp.AreaWidth)
		r.parent.layout(dp.AreaWidth)
	}

	//Fill background
	p := ui.NewPath(ui.Winding)
	p.AddRectangle(0, 0, dp.AreaWidth, dp.AreaHeight)
	p.End()
	dp.Context.Fill(p, &ui.Brush{
		A: 1,
	})
}
func (r *measureHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {
}
func (r *measureHandler) MouseCrossed(a *ui.Area, left bool) {
}
func (r *measureHandler) DragBroken(a *ui.Area) {
}
func (r *measureHandler) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) bool {
	return false
}
