package packetdiagram

import (
	"testing"

	"github.com/tj/assert"
)

func TestGetBitDistributions(t *testing.T) {
	testData := []struct {
		Name        string
		BitsPerLine int
		Cursor      *Cursor
		Placement   Placement
		Expected    []int
	}{
		{
			Name:        "fit in a single line",
			BitsPerLine: 32,
			Cursor: &Cursor{
				x: 0,
				y: 0,
			},
			Placement: Placement{
				Bits: 16,
			},
			Expected: []int{16},
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			t.Parallel()

			d := getBitDistributions(data.BitsPerLine, data.Cursor, data.Placement)
			assert.Equal(t, data.Expected, d)
		})
	}
}

func TestCalculateXAxisBitLabelDimensions(t *testing.T) {
	testData := []struct {
		Name           string
		Definition     *Definition
		Dimensions     Dimensions
		ExpectedXs     []int
		ExpectedYs     []int
		ExpectedLabels []string
	}{
		{
			Name: "OctetsPerLine=4,BitOrigin=0,YAxisWidth=50,CellDimension=30,30",
			Definition: &Definition{
				OctetsPerLine: 4,
			},
			Dimensions: Dimensions{
				YAxis: Dimension{
					Width: 50,
				},
				Cell: Dimension{
					Width:  30,
					Height: 30,
				},
			},
			ExpectedXs:     []int{65, 95, 125, 155, 185, 215, 245, 275, 305, 335, 365, 395, 425, 455, 485, 515, 545, 575, 605, 635, 665, 695, 725, 755, 785, 815, 845, 875, 905, 935, 965, 995},
			ExpectedYs:     []int{26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26},
			ExpectedLabels: []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"},
		},
		{
			Name: "OctetsPerLine=4,BitOrigin=1,YAxisWidth=50,CellDimension=30,30",
			Definition: &Definition{
				OctetsPerLine: 4,
				XAxis: XAxisSpec{
					Bits: &XAxisBitsSpec{
						Origin: 1,
					},
				},
			},
			Dimensions: Dimensions{
				YAxis: Dimension{
					Width: 50,
				},
				Cell: Dimension{
					Width:  30,
					Height: 30,
				},
			},
			ExpectedXs:     []int{65, 95, 125, 155, 185, 215, 245, 275, 305, 335, 365, 395, 425, 455, 485, 515, 545, 575, 605, 635, 665, 695, 725, 755, 785, 815, 845, 875, 905, 935, 965, 995},
			ExpectedYs:     []int{26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26},
			ExpectedLabels: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32"},
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			t.Parallel()

			xs, ys, labels := calculateXAxisBitLabelDimensions(data.Definition, data.Dimensions)
			assert.Equal(t, data.ExpectedXs, xs)
			assert.Equal(t, data.ExpectedYs, ys)
			assert.Equal(t, data.ExpectedLabels, labels)
		})
	}
}
