package uidoc

import (
	"fmt"

	"github.com/andlabs/ui"
)

type layout struct {
	e          Element
	x, y, w, h float64
	path       *ui.Path
}

type Group struct {
	layouts []layout
	ElementBase
}

func NewGroup(elements []Element) *Group {
	g := &Group{layouts: make([]layout, len(elements))}
	for i, e := range elements {
		g.layouts[i] = layout{e: e}
	}
	return g
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func (g *Group) Layout(width float64) (w float64, h float64) {
	groupLeft := g.MarginLeft + g.PaddingLeft
	groupRight := width - g.MarginRight - g.PaddingRight
	x := groupLeft
	y := g.MarginTop + g.PaddingTop
	lineHeight := 0.0
	for i, l := range g.layouts {
		mTop, mRight, mBottom, mLeft := l.e.Margins()
		pTop, pRight, pBottom, pLeft := l.e.Padding()
		l.y = y + mTop + pTop
		l.x = x + mLeft + pLeft
		l.w, l.h = l.e.Layout(groupRight - x - mRight - pRight)
		lineHeight = max(lineHeight, l.h+mBottom+pBottom+mTop+pTop)
		w = max(w, x+l.w+mRight+pRight)
		switch l.e.Mode() {
		case LayoutBlock:
			x = groupLeft
			y += lineHeight
			lineHeight = 0
		case LayoutInline:
			eleWidth := l.w + mRight + pRight + mLeft + mRight
			if x+eleWidth > groupRight {
				x = groupLeft
				y += lineHeight
				lineHeight = 0
			} else {
				x += eleWidth
			}
		default:
			panic(fmt.Errorf("Invlalid LayoutMode %v", l.e.Mode()))
		}
		if b := l.e.Fill(); b != nil {
			l.path = ui.NewPath(ui.Winding)
			l.path.AddRectangle(l.x-pLeft, l.y-pTop, l.w+pLeft+pRight, l.h+pTop+pBottom)
			l.path.End()
		}
		g.layouts[i] = l
	}
	h = y + lineHeight
	return
}
func (g *Group) Render(dp *ui.AreaDrawParams, x, y float64) {
	for _, l := range g.layouts {
		if l.path != nil {
			dp.Context.Fill(l.path, l.e.Fill())
		}
		l.e.Render(dp, l.x, l.y)
	}
}
func (g *Group) Free() {
	for _, l := range g.layouts {
		if l.path != nil {
			l.path.Free()
		}
		l.e.Free()
	}
	g.layouts = nil
}

func (g *Group) Append(e Element) {
	g.layouts = append(g.layouts, layout{e: e})
}
