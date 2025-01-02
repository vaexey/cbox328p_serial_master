package panel

// Color object for both raw and decimal colors
type Color struct {
	R float32 `json:"r"`
	G float32 `json:"g"`
	B float32 `json:"b"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

// Single array panel
type Panel struct {
	Id       string   `json:"id"`       // Main panel ID (3 char)
	GroupIds []string `json:"groupIds"` // Secondary panel IDs, eg. groups (3 char)

	Position Position `json:"position"` // Panel position relative to first element of array

	ColorScale int `json:"colorScale"` // Maximum value for color component (default 255)

	RawColor     Color `json:"-"`            // Raw color currently set on panel (after color correction and normalization to 0-255)
	CurrentColor Color `json:"currentColor"` // Color currently set on panel (range 0-1)

	Dirty bool `json:"-"` // Whether panel color has been changed and not yet sent
}
