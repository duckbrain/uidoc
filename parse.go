package uidoc

import (
	"encoding/json"
	"fmt"

	"github.com/andlabs/ui"
)

type parseEle struct {
	Type                                                 string
	Margins                                              float64
	MarginTop, MarginRight, MarginBottom, MarginLeft     *float64
	Paddings                                             float64
	PaddingTop, PaddingRight, PaddingBottom, PaddingLeft *float64
	Mode                                                 string

	// Group Properties
	Children []parseEle

	// Text Properties
	Text       string
	FontFamily string
	FontSize   float64
	FontWeight string
}

func Parse(data []byte) (ele Element, err error) {
	defer func() {
		if v := recover(); v != nil {
			err = v.(error)
		}
	}()

	var root parseEle

	if err := json.Unmarshal(data, &root); err != nil {
		return nil, err
	}

	ele = root.Element()
	return
}

func (p parseEle) Element() Element {
	switch p.Type {
	case "Text":
		f := &ui.FontDescriptor{
			Family: p.FontFamily,
			Size:   p.FontSize,
		}
		switch p.FontWeight {
		case "Bold":
			f.Weight = ui.TextWeightBold
		case "Normal", "":
			f.Weight = ui.TextWeightNormal
		}
		base := NewText(p.Text, ui.LoadClosestFont(f))
		base.ElementBase = parseElementBase(p)
		return base
	case "Group":
		children := make([]Element, len(p.Children))
		for i, child := range p.Children {
			children[i] = child.Element()
		}
		base := NewGroup(children)
		base.ElementBase = parseElementBase(p)
		return base
	default:
		panic(fmt.Errorf("Unknown Type %v", p.Type))
	}
}

func parseElementBase(p parseEle) ElementBase {
	base := ElementBase{}
	base.MarginTop = marginVal(p.Margins, p.MarginTop)
	base.MarginRight = marginVal(p.Margins, p.MarginRight)
	base.MarginBottom = marginVal(p.Margins, p.MarginBottom)
	base.MarginLeft = marginVal(p.Margins, p.MarginLeft)
	base.PaddingTop = marginVal(p.Paddings, p.PaddingTop)
	base.PaddingRight = marginVal(p.Paddings, p.PaddingRight)
	base.PaddingBottom = marginVal(p.Paddings, p.PaddingBottom)
	base.PaddingLeft = marginVal(p.Paddings, p.PaddingLeft)
	switch p.Mode {
	case "Block", "":
		base.LayoutMode = LayoutBlock
	case "Inline":
		base.LayoutMode = LayoutInline
	default:
		panic(fmt.Errorf("Unknown Layout Mode %v", p.Mode))
	}
	return base
}

func marginVal(overall float64, specific *float64) float64 {
	if specific != nil {
		return *specific
	}
	return overall
}
