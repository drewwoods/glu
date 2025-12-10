package glu

// #ifdef __APPLE__
// #define GL_SILENCE_DEPRECATION
//   #include <OpenGL/glu.h>
// #else
//   #include <GL/glu.h>
// #endif
// #include <stdlib.h>
import "C"
import "unsafe"

// Nurbs holds the GLUnurbs object.
type Nurbs struct {
	nurbs *C.GLUnurbs
}

// NewNurbsRenderer creates a new NURBS object.
func NewNurbsRenderer() *Nurbs {
	n := &Nurbs{
		nurbs: C.gluNewNurbsRenderer(),
	}

	if n.nurbs == nil {
		panic("Out of memory or GLU not initialized.")
	}

	return n
}

// NurbsProperty sets a NURBS property.
func (n *Nurbs) NurbsProperty(property uint32, value float32) {
	C.gluNurbsProperty(n.nurbs, C.GLenum(property), C.GLfloat(value))
}

// BeginSurface begins a NURBS surface definition.
func (n *Nurbs) BeginSurface() {
	C.gluBeginSurface(n.nurbs)
}

// EndSurface ends a NURBS surface definition.
func (n *Nurbs) EndSurface() {
	C.gluEndSurface(n.nurbs)
}

// NurbsSurface defines a NURBS surface.
func (n *Nurbs) NurbsSurface(sKnotCount int, sKnots []float32, tKnotCount int, tKnots []float32, sStride int, tStride int, ctlarray []float32, sOrder int, tOrder int, type0 uint32) {
	C.gluNurbsSurface(
		n.nurbs,
		C.GLint(sKnotCount),
		(*C.GLfloat)(unsafe.Pointer(&sKnots[0])),
		C.GLint(tKnotCount),
		(*C.GLfloat)(unsafe.Pointer(&tKnots[0])),
		C.GLint(sStride),
		C.GLint(tStride),
		(*C.GLfloat)(unsafe.Pointer(&ctlarray[0])),
		C.GLint(sOrder),
		C.GLint(tOrder),
		C.GLenum(type0),
	)
}
