package packetdiagram

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	defaultOctetsPerLine      = 4
	defaultXAxisBitsHeight    = 25
	defaultXAxisBitsDirection = XAxisBitsDirectionLeftToRight
	defaultXAxisBitsUnit      = 32
	defaultXAxisBitsOrigin    = 0
	defaultXAxisOctetsHeight  = 20
	defaultYAxisBitsOrigin    = 0
	defaultYAxisOctetsOrigin  = 0
	defaultCellWidth          = 30
	defaultCellHeight         = 30
	defaultBreakMarkWidth     = 10
	defaultBreakMarkHeight    = 10
)

type Definition struct {
	Theme         *ThemeSpec    `yaml:"theme,omitempty"`
	OctetsPerLine *uint         `yaml:"octets-per-line,omitempty"`
	XAxis         XAxisSpec     `yaml:"x-axis"`
	YAxis         YAxisSpec     `yaml:"y-axis"`
	Cell          CellSpec      `yaml:"cell"`
	BreakMark     BreakMarkSpec `yaml:"break-mark"`
	Placements    []Placement   `yaml:"placements"`
}

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
	Show   *bool `yaml:"show,omitempty"`
	Height *uint `yaml:"height,omitempty"`
}

type YAxisBitsSpec struct {
	Show   *bool `yaml:"show,omitempty"`
	Width  *uint `yaml:"width,omitempty"`
	Origin *uint `yaml:"origin,omitempty"`
}

type YAxisOctetsSpec struct {
	Show   *bool `yaml:"show,omitempty"`
	Width  *uint `yaml:"width,omitempty"`
	Origin *uint `yaml:"origin,omitempty"`
}

type CellSpec struct {
	Width  *uint `yaml:"width,omitempty"`
	Height *uint `yaml:"height,omitempty"`
}

type BreakMarkSpec struct {
	Width  *uint `yaml:"width,omitempty"`
	Height *uint `yaml:"height,omitempty"`
}

type Placement struct {
	Label          string                       `yaml:"label"`
	Bits           *uint                        `yaml:"bits,omitempty"`
	VariableLength *VariableLengthPlacementSpec `yaml:"variable-length,omitempty"`
}

type VariableLengthPlacementSpec struct {
	MaxBits uint `yaml:"max-bits"`
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

	err = def.validate()
	if err != nil {
		return nil, err
	}

	return &def, nil
}

func (d *Definition) validate() error {
	for _, p := range d.Placements {
		if p.Bits == nil && p.VariableLength == nil {
			return errors.New("either `bits` or `valiable-length` field is required for a placement")
		}
	}

	return nil
}

func (d *Definition) GetOctetsPerLine() uint {
	if d.OctetsPerLine == nil {
		return defaultOctetsPerLine
	}
	return uint(*d.OctetsPerLine)
}

func (d *Definition) GetBitsPerLine() uint {
	return d.GetOctetsPerLine() * 8
}

func (d *Definition) GetXAxisBitsHeight() uint {
	if d.XAxis.Bits == nil || d.XAxis.Bits.Height == nil {
		return defaultXAxisBitsHeight
	}

	return *d.XAxis.Bits.Height
}

func (d *Definition) GetXAxisOctetsHeight() uint {
	if d.XAxis.Octets == nil || d.XAxis.Octets.Height == nil {
		return defaultXAxisOctetsHeight
	}

	return *d.XAxis.Octets.Height
}

func (d *Definition) GetYAxisBitsWidth() uint {
	if d.YAxis.Bits == nil || d.YAxis.Bits.Width == nil {
		tb := d.GetTotalPlacementBits()
		cols := uint(len(fmt.Sprintf("%d", tb)))
		numWidth := cols * d.GetTextSizeInPixels()
		titleWidth := d.GetAxisTitleTextSizeInPixels() * (3 + 2)

		return maxUint(numWidth, titleWidth)
	}

	return *d.YAxis.Bits.Width
}

func (d *Definition) GetYAxisOctetsWidth() uint {
	if d.YAxis.Octets == nil || d.YAxis.Octets.Width == nil {
		to := d.GetTotalPlacementOctets()
		cols := uint(len(fmt.Sprintf("%d", to)))
		numWidth := cols * d.GetTextSizeInPixels()
		titleWidth := d.GetAxisTitleTextSizeInPixels() * (5 + 2) // len("octet") + somehow we need "2"

		return maxUint(numWidth, titleWidth)
	}

	return *d.YAxis.Bits.Width
}

func maxUint(nums ...uint) uint {
	max := uint(0)
	for _, x := range nums {
		if x > max {
			max = x
		}
	}
	return max
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

func (d *Definition) GetXAxisBitsOrigin() uint {
	if d.XAxis.Bits == nil || d.XAxis.Bits.Origin == nil {
		return defaultXAxisBitsOrigin
	}

	return *d.XAxis.Bits.Origin
}

func (d *Definition) GetYAxisBitsOrigin() uint {
	if d.YAxis.Bits == nil || d.YAxis.Bits.Origin == nil {
		return defaultYAxisBitsOrigin
	}

	return *d.YAxis.Bits.Origin
}

func (d *Definition) GetYAxisOctetsOrigin() uint {
	if d.YAxis.Octets == nil || d.YAxis.Octets.Origin == nil {
		return defaultYAxisOctetsOrigin
	}

	return *d.YAxis.Octets.Origin
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

func (d *Definition) ShouldShowYAxisBits() bool {
	if d.YAxis.Bits == nil {
		return false
	}

	if d.YAxis.Bits.Show == nil {
		return true
	}

	return *d.YAxis.Bits.Show
}

func (d *Definition) ShouldShowYAxisOctets() bool {
	if d.YAxis.Octets == nil {
		return false
	}

	if d.YAxis.Octets.Show == nil {
		return true
	}

	return *d.YAxis.Octets.Show
}

func (d *Definition) GetTheme() *ThemeSpec {
	if d.Theme != nil {
		return d.Theme
	}

	return defaultTheme
}

func (d *Definition) GetBackgroundColor() string {
	return d.GetTheme().GetBackgroundColor()
}

func (d *Definition) GetTextColor() string {
	return d.GetTheme().GetTextColor()
}

func (d *Definition) GetTextSize() string {
	return d.GetTheme().GetTextSize()
}

func (d *Definition) GetTextSizeInPixels() uint {
	size, err := cssTextSizeToPixels(d.GetTheme().GetTextSize())
	if err != nil {
		return defaultTextSizeInPixels
	}
	return size
}

func (d *Definition) GetTextFontFamily() string {
	return d.GetTheme().GetTextFontFamily()
}

func (d *Definition) GetAxisTitleTextSize() string {
	return d.GetTheme().GetAxisTitleTextSize()
}

func (d *Definition) GetAxisTitleTextSizeInPixels() uint {
	size, err := cssTextSizeToPixels(d.GetTheme().GetAxisTitleTextSize())
	if err != nil {
		return defaultAxisTitleTextSizeInPixels
	}
	return size
}

func cssTextSizeToPixels(size string) (uint, error) {
	switch {
	case strings.HasSuffix(size, "px"):
		px := strings.ReplaceAll(size, "px", "")
		pxf, err := strconv.ParseFloat(px, 32)
		if err != nil {
			return 0, err
		}
		return uint(pxf), nil
	case strings.HasSuffix(size, "pt"):
		px, err := pointsToPixels(size)
		if err != nil {
			return 0, err
		}
		return px, nil
	default:
		log.Printf("unsupported text size unit: %s", size)
		return 0, errors.Errorf("unsupported text size unit: %s", size)
	}
}

func pointsToPixels(points string) (uint, error) {
	pointsWithoutUnit := strings.ReplaceAll(points, "pt", "")
	p, err := strconv.ParseFloat(pointsWithoutUnit, 32)
	if err != nil {
		return 0, err
	}
	return uint(p * 0.75), nil
}

func (d *Definition) GetCellWidth() uint {
	if d.Cell.Width == nil {
		return defaultCellWidth
	}

	return *d.Cell.Width
}

func (d *Definition) GetCellHeight() uint {
	if d.Cell.Height == nil {
		return defaultCellHeight
	}

	return *d.Cell.Height
}

func (d *Definition) GetBreakMarkWidth() uint {
	if d.BreakMark.Width == nil {
		return defaultBreakMarkWidth
	}

	return *d.BreakMark.Width
}

func (d *Definition) GetBreakMarkHeight() uint {
	if d.BreakMark.Height == nil {
		return defaultBreakMarkHeight
	}

	return *d.BreakMark.Height
}

func (d *Definition) GetTotalPlacementOctets() uint {
	totalBits := d.GetTotalPlacementBits()
	return totalBits / 8
}

func (d *Definition) GetTotalPlacementBits() uint {
	sum := uint(0)
	for _, p := range d.Placements {
		if p.VariableLength == nil {
			if *p.Bits > 0 {
				sum += *p.Bits
			}
		} else {
			sum += p.VariableLength.MaxBits
		}
	}
	return sum
}

func (d *Definition) GetTotalRows() uint {
	rows := uint(1)
	bits := uint(0)
	for _, p := range d.Placements {
		if p.VariableLength == nil {
			bits += *p.Bits
		} else {
			bits += p.VariableLength.MaxBits
		}

		for bits >= d.GetBitsPerLine() {
			rows++
			bits -= d.GetBitsPerLine()
		}
	}

	if bits == 0 {
		rows--
	}
	return rows
}
