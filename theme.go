package packetdiagram

const (
	defaultBackgroundColor                = "white"
	defaultTextColor                      = "black"
	defaultTextSize                       = "16pt"
	defaultTextSizeInPixels          uint = 12
	defaultTextFontFamily                 = "Sans Serif"
	defaultAxisTitleTextSize              = "8pt"
	defaultAxisTitleTextSizeInPixels uint = 6
)

type ThemeSpec struct {
	Predefined *string         `yaml:"predefined,omitempty"`
	Background *BackgroundSpec `yaml:"background,omitempty"`
	Text       *TextSpec       `yaml:"text,omitempty"`
}

type BackgroundSpec struct {
	Color *string `yaml:"color,omitempty"`
}

type TextSpec struct {
	Color         *string `yaml:"color,omitempty"`
	Size          *string `yaml:"size,omitempty"`
	FontFamily    *string `yaml:"font-family,omitempty"`
	AxisTitleSize *string `yaml:"axis-title-size,omitempty"`
}

func (t ThemeSpec) GetBackgroundColor() string {
	if t.Background == nil || t.Background.Color == nil {
		return defaultBackgroundColor
	}
	return *t.Background.Color
}

func (t ThemeSpec) GetTextColor() string {
	if t.Text == nil || t.Text.Color == nil {
		return defaultTextColor
	}
	return *t.Text.Color
}

func (t ThemeSpec) GetTextSize() string {
	if t.Text == nil || t.Text.Size == nil {
		return defaultTextSize
	}
	return *t.Text.Size
}

func (t ThemeSpec) GetTextFontFamily() string {
	if t.Text == nil || t.Text.FontFamily == nil {
		return defaultTextFontFamily
	}
	return *t.Text.FontFamily
}

func (t ThemeSpec) GetAxisTitleTextSize() string {
	if t.Text == nil || t.Text.AxisTitleSize == nil {
		return defaultAxisTitleTextSize
	}
	return *t.Text.Size
}

var defaultTheme = &ThemeSpec{
	Background: &BackgroundSpec{
		Color: stringp(defaultBackgroundColor),
	},
	Text: &TextSpec{
		Color:         stringp(defaultTextColor),
		Size:          stringp(defaultTextSize),
		FontFamily:    stringp(defaultTextFontFamily),
		AxisTitleSize: stringp(defaultAxisTitleTextSize),
	},
}

func stringp(s string) *string {
	return &s
}
