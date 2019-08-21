package progress

import (
	"fmt"

	"github.com/dairaga/js/dom"
	"github.com/dairaga/js/dom/bs"
)

// Bar represents a Bootstrap progress bar.
type Bar struct {
	*bs.Component
	max, min, value int
}

// NewBar returns a Bootstrap progress bar.
func NewBar(bgColor bs.Color, min, max, value int) *Bar {
	b := &Bar{bs.ComponentOf(dom.CreateElement("div")), 0, 0, 0}
	b.SetMax(max).
		SetMin(min).
		SetVal(value).
		Color(bgColor).
		AddClass("progress-bar").
		SetAttr("role", "progressbar")

	return b
}

// ----------------------------------------------------------------------------

// Stripped apply a stripe style.
func (b *Bar) Stripped() *Bar {
	b.AddClass("progress-bar-striped")
	return b
}

// Animate animates the stripes.
func (b *Bar) Animate() *Bar {
	b.AddClass("progress-bar-animated")
	return b
}

// Stop ...
func (b *Bar) Stop() *Bar {
	b.RemoveClass("progress-bar-animated")
	return b
}

// SetMax ...
func (b *Bar) SetMax(max int) *Bar {
	b.max = max
	b.SetAttr("aria-valuemax", fmt.Sprintf("%d", max))
	return b
}

// Max ...
func (b *Bar) Max() int {
	return b.max
}

// SetMin ...
func (b *Bar) SetMin(min int) *Bar {
	b.min = min
	b.SetAttr("aria-valuemin", fmt.Sprintf("%d", min))
	return b
}

// Min ...
func (b *Bar) Min() int {
	return b.min
}

// SetVal ...
func (b *Bar) SetVal(value int) *Bar {
	b.value = value
	b.SetAttr("aria-valuenow", fmt.Sprintf("%d", value))
	b.SetAttr("style", fmt.Sprintf("width: %d%%", value))
	return b
}

// Val ...
func (b *Bar) Val() int {
	return b.value
}

// ----------------------------------------------------------------------------

// Progress represents a Bootstrap progress conainer.
type Progress struct {
	*bs.Component
	bars []*Bar
}

// New returns a progress.
func New(bgColor bs.Color, min, max, value int) *Progress {
	b := NewBar(bgColor, min, max, value)
	p := &Progress{bs.ComponentOf(dom.CreateElement("div")), nil}
	p.AddClass("progress")
	p.Add(b)
	return p
}

// ----------------------------------------------------------------------------

// Add adds a progress bar.
func (p *Progress) Add(bar *Bar) *Progress {
	p.Append(bar)
	p.bars = append(p.bars, bar)
	return p
}

// Bar returns the n-index bar in progress container.
func (p *Progress) Bar(idx int) *Bar {
	if idx < 0 || idx >= len(p.bars) {
		return nil
	}
	return p.bars[idx]
}
