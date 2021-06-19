package packetdiagram

import (
	"io"

	"gopkg.in/yaml.v2"
)

const (
	defaultBackgroundColor    = "white"
	defaultTextColor          = "black"
	defaultTextSize           = "16pt"
	defaultTextFontFamily     = "Sans Serif"
	defaultOctetsPerLine      = 4
	defaultXAxisBitsHeight    = 20
	defaultXAxisBitsDirection = XAxisBitsDirectionLeftToRight
	defaultXAxisBitsUnit      = 32
	defaultXAxisBitsOrigin    = 0
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

type TextSpec struct {
	Color      string `yaml:"color"`
	Size       string `yaml:"size"`
	FontFamily string `yaml:"font-family"`
}

type OctetsPerLine uint

type XAxisSpec struct {
	Bits   *XAxisBitsSpec   `yaml:"bits,omitempty"`
	Octets *XAxisOctetsSpec `yaml:"octets,omitempty"`
}

type YAxisSpec struct {
	Bits   *YAxisBitsSpec   `yaml:"bits,omitempty"`
	Octets *YAxisOctetsSpec `yaml:"octets,omitempty"`
}

type XAxisBitsSpec struct {
	Show      *bool               `yaml:"show,omitempty"`
	Height    *uint               `yaml:"height,omitempty"`
	Direction *XAxisBitsDirection `yaml:"direction,omitempty"`
	Origin    *uint               `yaml:"origin,omitempty"`
	Unit      *XAxisBitsUnit      `yaml:"unit,omitempty"`
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
	if d.OctetsPerLine == 0 {
		return defaultOctetsPerLine
	}
	return int(d.OctetsPerLine) * 8
}

func (d *Definition) GetXAxisBitsHeight() uint {
	if d.XAxis.Bits == nil || d.XAxis.Bits.Height == nil {
		return defaultXAxisBitsHeight
	}

	return *d.XAxis.Bits.Height
}

func (d *Definition) GetXAxisBitsDirection() XAxisBitsDirection {
	if d.XAxis.Bits == nil || d.XAxis.Bits.Direction == nil {
		return defaultXAxisBitsDirection
	}

	return *d.XAxis.Bits.Direction
}

func (d *Definition) GetXAxisBitsUnit() XAxisBitsUnit {
	if d.XAxis.Bits == nil || d.XAxis.Bits.Unit == nil {
		return defaultXAxisBitsUnit
	}

	return *d.XAxis.Bits.Unit
}

func (d *Definition) GetXAxisBitsOrigin() int {
	if d.XAxis.Bits == nil || d.XAxis.Bits.Origin == nil {
		return defaultXAxisBitsOrigin
	}

	return int(*d.XAxis.Bits.Origin)
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
