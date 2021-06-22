package packetdiagram

import (
	"fmt"
	"strings"

	svg "github.com/ajstarks/svgo"
)

func defineStyles(def *Definition, dim Dimensions, canvas *svg.SVG) {
	style := getStyleForXAxisBits(def, dim) + "\n"
	style += getStyleForXAxisOctets(def, dim) + "\n"
	style += getStyleForYAxisBits(def, dim) + "\n"
	style += getStyleForYAxisOctets(def, dim) + "\n"
	style += getStyleForPlacements(def, dim) + "\n"
	style += getStyleForBreakMark(def, dim) + "\n"
	canvas.Style("text/css", style)
}

func shrinkStyle(style string) string {
	s := strings.ReplaceAll(style, "\r", "")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\t", "")
	return s
}

func getStyleForXAxisBits(def *Definition, dim Dimensions) string {
	return shrinkStyle(fmt.Sprintf(`
text.x-bit{
	fill:%s;
	font-size:%s;
	text-anchor: middle;
}
text.x-bit-title{
	fill:%s;
	font-size:%s;
	text-anchor: start;
}
line.x-bit{
	stroke:black;
}`,
		def.GetTextColor(),
		def.GetTextSize(),
		def.GetTextColor(),
		def.GetAxisTitleTextSize(),
	))
}

func getStyleForXAxisOctets(def *Definition, dim Dimensions) string {
	return shrinkStyle(fmt.Sprintf(`
text.x-octet{
	fill:%s;
	font-size:%s;
	text-anchor: middle;
}
text.x-octet-title{
	fill:%s;
	font-size:%s;
	text-anchor: start;
}
line.x-octet{
	stroke:black;
}`,
		def.GetTextColor(),
		def.GetTextSize(),
		def.GetTextColor(),
		def.GetAxisTitleTextSize(),
	))
}

func getStyleForYAxisBits(def *Definition, dim Dimensions) string {
	return shrinkStyle(fmt.Sprintf(`
text.y-bit{
	fill:%s;
	font-size:%s;
	text-anchor: end;
}
text.y-bit-title{
	fill:%s;
	font-size:%s;
	text-anchor: end;
}
line.y-bit{
	stroke:black;
}`,
		def.GetTextColor(),
		def.GetTextSize(),
		def.GetTextColor(),
		def.GetAxisTitleTextSize(),
	))
}

func getStyleForYAxisOctets(def *Definition, dim Dimensions) string {
	return shrinkStyle(fmt.Sprintf(`
text.y-octet{
	fill:%s;
	font-size:%s;
	text-anchor: end;
}
text.y-octet-title{
	fill:%s;
	font-size:%s;
	text-anchor: end;
}
line.y-octet{
	stroke:black;
}`,
		def.GetTextColor(),
		def.GetTextSize(),
		def.GetTextColor(),
		def.GetAxisTitleTextSize(),
	))
}

func getStyleForPlacements(def *Definition, dim Dimensions) string {
	return shrinkStyle(fmt.Sprintf(`
polygon.placement{
	fill:white;
	stroke:black;
}

text.placement{
	fill:%s;
	font-family:%s;
	font-size:%s;
	text-anchor:middle;
}`,
		def.GetTextColor(),
		def.GetTextFontFamily(),
		def.GetTextSize(),
	))
}

func getStyleForBreakMark(def *Definition, dim Dimensions) string {
	return shrinkStyle(`
path.breakmark{
	fill:none;
	stroke:black;
}`,
	)
}
