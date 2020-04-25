package main

import (
	"GoSlice/data"
)

type option func(s *GoSlice)

func (s *GoSlice) With(o ...option) {
	for _, option := range o {
		option(s)
	}
}

// MeldDistance is the distance which two points have to be within to
// count them as one point.
func ḾeldDistance(m data.Micrometer) option {
	return func(s *GoSlice) {
		s.o.MeldDistance = m
	}
}

// Center is the point where the model is finally placed
func Center(p data.MicroVec3) option {
	return func(s *GoSlice) {
		s.o.Printer.Center = p
	}
}

// InitialLayerThickness is the layer thickness for the first layer
func InitialLayerThickness(m data.Micrometer) option {
	return func(s *GoSlice) {
		s.o.Print.InitialLayerThickness = m
	}
}

// LayerThickness is the thickness for all but the first layer
func LayerThickness(m data.Micrometer) option {
	return func(s *GoSlice) {
		s.o.Print.LayerThickness = m
	}
}

// JoinPolygonSnapDistance is the distance used to check if two open
// polygons can be snapped together to one bigger polygon.
// Checked by the start and endpoints of the polygons.
func JoinPolygonSnapDistance(m data.Micrometer) option {
	return func(s *GoSlice) {
		s.o.JoinPolygonSnapDistance = m
	}
}

// FinishPolygonSnapDistance is the max distance between start end endpoint of
// a polygon used to check if a open polygon can be closed.
func FinishPolygonSnapDistance(m data.Micrometer) option {
	return func(s *GoSlice) {
		s.o.FinishPolygonSnapDistance = m
	}
}

// FilamentDiameter is the filament diameter used by the printer
func FilamentDiameter(m data.Millimeter) option {
	return func(s *GoSlice) {
		s.o.Filament.FilamentDiameter = m.ToMicrometer()
	}
}

// ExtrusionWidth is the diameter of your nozzle
func ExtrusionWidth(m data.Micrometer) option {
	return func(s *GoSlice) {
		s.o.Printer.ExtrusionWidth = m
	}
}

// InsetCount is the number of perimeters
func InsetCount(n int) option {
	return func(s *GoSlice) {
		s.o.Print.InsetCount = n
	}
}

// InitialLayerSpeed is the speed only for the first layer in mm per second
func InitialLayerSpeed(mmPerS data.Millimeter) option {
	return func(s *GoSlice) {
		s.o.Print.IntialLayerSpeed = mmPerS
	}
}

// LayerSpeed is the speed for all but the first layer in mm per second
func LayerSpeed(mmPerS data.Millimeter) option {
	return func(s *GoSlice) {
		s.o.Print.LayerSpeed = mmPerS
	}
}

// OuterPerimeterSpeed is the speed only for outer perimeters
func OuterPerimeterSpeed(mmPerS data.Millimeter) option {
	return func(s *GoSlice) {
		s.o.Print.OuterPerimeterSpeed = mmPerS
	}
}

// InfillOverlapPercent says how much the infill should overlap into the perimeters
func InfillOverlapPercent(percent int) option {
	return func(s *GoSlice) {
		s.o.Print.InfillOverlapPercent = percent
	}
}