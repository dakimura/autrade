package plot

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"os"
)

func PlotGraph(title, xLabel, yLabel, legend string, points []int, saveFilePath string) error {
	p, err := plot.New()
	if err != nil {
		os.Exit(1)
	}

	p.Title.Text = title
	p.X.Label.Text = xLabel
	p.Y.Label.Text = yLabel

	//plotutil.DefaultGlyphShapes GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

	l, err := plotter.NewLine(convertToXYs(points))
	if err != nil {
		return err
	}
	p.Add(l)

	//err = plotutil.AddLinePoints(p,
	//	legend, convertToXYs(points))
	//if err != nil {
	//	return err
	//}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, saveFilePath); err != nil {
		return err
	}

	return nil
}

func convertToXYs(points []int) plotter.XYs {
	pts := make(plotter.XYs, len(points))
	for i := range pts {
		pts[i].X = float64(i)
		pts[i].Y = float64(points[i])
	}
	return pts
}
