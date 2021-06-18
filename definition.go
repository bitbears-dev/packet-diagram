package packetdiagram

import (
	"io"

	"gopkg.in/yaml.v2"
)

type Definition struct {
	Theme         *ThemeSpec    `yaml:"theme,omitempty"`
	OctetsPerLine OctetsPerLine `yaml:"octets-per-line"`
	XAxis         XAxisSpec     `yaml:"x-axis"`
	YAxis         YAxisSpec     `yaml:"y-axis"`
	Placements    []Placement   `yaml:"placements"`
}

type ThemeSpec struct {
	Predefined string         `yaml:"predefined"`
	Background BackgroundSpec `yaml:"background"`
	Text       TextSpec       `yaml:"text"`
}

type BackgroundSpec struct {
	Color string `yaml:"color"`
}

const defaultBackgroundColor = "white"

type TextSpec struct {
	Color      string `yaml:"color"`
	Size       string `yaml:"size"`
	FontFamily string `yaml:"font-family"`
}

const defaultTextColor = "black"
const defaultTextSize = "16pt"
const defaultTextFontFamily = "Sans Serif"

type OctetsPerLine uint

const (
	FourOctetsPerLine OctetsPerLine = 4
)

type XAxisSpec struct {
	Bits   *XAxisBitsSpec   `yaml:"bits,omitempty"`
	Octets *XAxisOctetsSpec `yaml:"octets,omitempty"`
}

type YAxisSpec struct {
	Bits   *YAxisBitsSpec   `yaml:"bits,omitempty"`
	Octets *YAxisOctetsSpec `yaml:"octets,omitempty"`
}

type XAxisBitsSpec struct {
	Show      *bool              `yaml:"show,omitempty"`
	Direction XAxisBitsDirection `yaml:"direction"`
	Origin    uint               `yaml:"origin"`
	Unit      XAxisBitsUnit      `yaml:"unit"`
}

type XAxisBitsDirection string

const (
	XAxisBitsDirectionLeftToRight XAxisBitsDirection = "left-to-right"
)

type XAxisBitsUnit uint

type XAxisOctetsSpec struct {
	Show *bool `yaml:"show,omitempty"`
}

type YAxisBitsSpec struct {
	Show *bool `yaml:"show,omitempty"`
}

type YAxisOctetsSpec struct {
	Show *bool `yaml:"show,omitempty"`
}

type Placement struct {
	Label          string `yaml:"label"`
	Bits           uint   `yaml:"bits"`
	VariableLength bool   `yaml:"variable-length"`
}

func LoadDefinition(r io.Reader) (*Definition, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var def Definition
	err = yaml.Unmarshal(b, &def)
	if err != nil {
		return nil, err
	}

	return &def, nil
}

func (d *Definition) GetBitsPerLine() int {
	return int(d.OctetsPerLine) * 8
}

func (d *Definition) GetXAxisBitsOrigin() int {
	if d.XAxis.Bits == nil {
		return 0
	}

	return int(d.XAxis.Bits.Origin)
}

func (d *Definition) ShouldShowXAxisBits() bool {
	if d.XAxis.Bits == nil {
		return false
	}

	if d.XAxis.Bits.Show == nil {
		return true
	}

	return *d.XAxis.Bits.Show
}

func (d *Definition) ShouldShowXAxisOctets() bool {
	if d.XAxis.Octets == nil {
		return false
	}

	if d.XAxis.Octets.Show == nil {
		return true
	}

	return *d.XAxis.Octets.Show
}

func (d *Definition) GetBackgroundColor() string {
	if d.Theme == nil {
		return defaultBackgroundColor
	}

	if d.Theme.Predefined != "" {
		// TODO: To be implemented
		return defaultBackgroundColor
	}

	if d.Theme.Background.Color == "" {
		return defaultBackgroundColor
	}

	return d.Theme.Background.Color
}

func (d *Definition) GetTextColor() string {
	if d.Theme == nil {
		return defaultTextColor
	}

	if d.Theme.Predefined != "" {
		// TODO: To be implemented
		return defaultTextColor
	}

	if d.Theme.Text.Color == "" {
		return defaultTextColor
	}

	return d.Theme.Text.Color
}

func (d *Definition) GetTextSize() string {
	if d.Theme == nil {
		return defaultTextSize
	}

	if d.Theme.Predefined != "" {
		// TODO: To be implemented
		return defaultTextSize
	}

	if d.Theme.Text.Color == "" {
		return defaultTextSize
	}

	return d.Theme.Text.Size
}

func (d *Definition) GetTextFontFamily() string {
	if d.Theme == nil {
		return defaultTextFontFamily
	}

	if d.Theme.Predefined != "" {
		// TODO: To be implemented
		return defaultTextFontFamily
	}

	if d.Theme.Text.FontFamily == "" {
		return defaultTextFontFamily
	}

	return d.Theme.Text.FontFamily
}
