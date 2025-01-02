package panel

import "math"

func (p *Panel) scaleColorToRaw(color Color) Color {
	return Color{
		R: color.R * float32(p.ColorScale),
		G: color.G * float32(p.ColorScale),
		B: color.B * float32(p.ColorScale),
	}
}

func colorDiff(c1 Color, c2 Color) float32 {
	return float32(math.Sqrt(
		math.Pow((float64(c1.R)-float64(c2.R)), 2) +
			math.Pow((float64(c1.G)-float64(c2.G)), 2) +
			math.Pow((float64(c1.B)-float64(c2.B)), 2),
	))
}

func (p *Panel) Set(color Color) {

	diff := colorDiff(color, p.CurrentColor)

	if diff < 0.0000001 {
		return
	}

	p.CurrentColor = color
	p.RawColor = p.scaleColorToRaw(color)

	p.Dirty = true
}

func (p *Panel) SetRGB(r float32, g float32, b float32) {
	p.Set(Color{
		R: r, G: g, B: b,
	})
}
