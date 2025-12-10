package glu

import (
	"testing"
)

func TestNurbs(t *testing.T) {
	// The following code is adapted from the OpenGL SuperBible
	// and is used to test the NURBS functions.
	// It does not render anything, but it calls the functions
	// to ensure they don't crash.
	ctlpoints := [4][4][3]float32{}
	knots := []float32{0.0, 0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0}

	// Initialize the control points
	for u := 0; u < 4; u++ {
		for v := 0; v < 4; v++ {
			ctlpoints[u][v][0] = 2.0 * (float32(u) - 1.5)
			ctlpoints[u][v][1] = 2.0 * (float32(v) - 1.5)
			if (u == 1 || u == 2) && (v == 1 || v == 2) {
				ctlpoints[u][v][2] = 3.0
			} else {
				ctlpoints[u][v][2] = -3.0
			}
		}
	}

	// Create a new NURBS renderer
	nurbs := NewNurbsRenderer()

	// Set the NURBS properties
	nurbs.NurbsProperty(SAMPLING_TOLERANCE, 25.0)
	nurbs.NurbsProperty(DISPLAY_MODE, FILL)

	// Begin the surface
	nurbs.BeginSurface()

	// Define the NURBS surface
	// Flatten the ctlpoints array
	flatCtlPoints := make([]float32, 4*4*3)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 3; k++ {
				flatCtlPoints[(i*4*3)+(j*3)+k] = ctlpoints[i][j][k]
			}
		}
	}

	nurbs.NurbsSurface(8, knots, 8, knots, 4*3, 3, flatCtlPoints, 4, 4, MAP1_VERTEX_3)

	// End the surface
	nurbs.EndSurface()
}
