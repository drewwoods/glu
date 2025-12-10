package glu

// #ifdef __APPLE__
// #define GL_SILENCE_DEPRECATION
//   #include <OpenGL/glu.h>
// #else
//   #include <GL/glu.h>
// #endif
// #include <stdlib.h>
// #include "callback.h"
import "C"
import "unsafe"

// Nurbs holds the GLUnurbs object.
type Nurbs struct {
	nurbs *C.GLUnurbs
	polyData interface{}

	beginData    NurbsBeginDataHandler
	vertexData   NurbsVertexDataHandler
	normalData   NurbsNormalDataHandler
	colorData    NurbsColorDataHandler
	textureCoordData NurbsTextureCoordDataHandler
	endData      NurbsEndDataHandler
	errorData    NurbsErrorHandler
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

// BeginCurve begins a NURBS curve definition.
func (n *Nurbs) BeginCurve() {
	C.gluBeginCurve(n.nurbs)
}

// BeginTrim begins a NURBS trim definition.
func (n *Nurbs) BeginTrim() {
	C.gluBeginTrim(n.nurbs)
}

// Delete deletes the NURBS object.
func (n *Nurbs) Delete() {
	C.gluDeleteNurbsRenderer(n.nurbs)
	n.nurbs = nil
}

// EndCurve ends a NURBS curve definition.
func (n *Nurbs) EndCurve() {
	C.gluEndCurve(n.nurbs)
}

// EndTrim ends a NURBS trim definition.
func (n *Nurbs) EndTrim() {
	C.gluEndTrim(n.nurbs)
}

// GetNurbsProperty returns a NURBS property value.
func (n *Nurbs) GetNurbsProperty(property uint32) float32 {
	var value C.GLfloat
	C.gluGetNurbsProperty(n.nurbs, C.GLenum(property), &value)
	return float32(value)
}

// LoadSamplingMatrices loads the sampling matrices.
func (n *Nurbs) LoadSamplingMatrices(model, perspective *[16]float32, view *[4]int32) {
	C.gluLoadSamplingMatrices(
		n.nurbs,
		(*C.GLfloat)(unsafe.Pointer(model)),
		(*C.GLfloat)(unsafe.Pointer(perspective)),
		(*C.GLint)(unsafe.Pointer(view)),
	)
}

// NurbsCurve defines a NURBS curve.
func (n *Nurbs) NurbsCurve(knotCount int, knots []float32, stride int, control []float32, order int, type0 uint32) {
	C.gluNurbsCurve(
		n.nurbs,
		C.GLint(knotCount),
		(*C.GLfloat)(unsafe.Pointer(&knots[0])),
		C.GLint(stride),
		(*C.GLfloat)(unsafe.Pointer(&control[0])),
		C.GLint(order),
		C.GLenum(type0),
	)
}

// PwlCurve defines a piecewise-linear curve.
func (n *Nurbs) PwlCurve(count int, data []float32, stride int, type0 uint32) {
	C.gluPwlCurve(
		n.nurbs,
		C.GLint(count),
		(*C.GLfloat)(unsafe.Pointer(&data[0])),
		C.GLint(stride),
		C.GLenum(type0),
	)
}

// NurbsCallbackData sets the user data for the callbacks.
func (n *Nurbs) NurbsCallbackData(userData interface{}) {
	n.polyData = userData
	C.gluNurbsCallbackData(n.nurbs, unsafe.Pointer(n))
}

// SetBeginCallback sets the callback for NURBS_BEGIN_DATA.
func (n *Nurbs) SetBeginCallback(f NurbsBeginDataHandler) {
	if n.nurbs == nil {
		panic("Uninitialised Nurbs. @see glu.NewNurbsRenderer.")
	}
	n.beginData = f
	C.setGluNurbsCallback(n.nurbs, C.GLenum(NURBS_BEGIN_DATA))
}

// SetVertexCallback sets the callback for NURBS_VERTEX_DATA.
func (n *Nurbs) SetVertexCallback(f NurbsVertexDataHandler) {
	if n.nurbs == nil {
		panic("Uninitialised Nurbs. @see glu.NewNurbsRenderer.")
	}
	n.vertexData = f
	C.setGluNurbsCallback(n.nurbs, C.GLenum(NURBS_VERTEX_DATA))
}

// SetNormalCallback sets the callback for NURBS_NORMAL_DATA.
func (n *Nurbs) SetNormalCallback(f NurbsNormalDataHandler) {
	if n.nurbs == nil {
		panic("Uninitialised Nurbs. @see glu.NewNurbsRenderer.")
	}
	n.normalData = f
	C.setGluNurbsCallback(n.nurbs, C.GLenum(NURBS_NORMAL_DATA))
}

// SetColorCallback sets the callback for NURBS_COLOR_DATA.
func (n *Nurbs) SetColorCallback(f NurbsColorDataHandler) {
	if n.nurbs == nil {
		panic("Uninitialised Nurbs. @see glu.NewNurbsRenderer.")
	}
	n.colorData = f
	C.setGluNurbsCallback(n.nurbs, C.GLenum(NURBS_COLOR_DATA))
}

// SetTextureCoordCallback sets the callback for NURBS_TEXTURE_COORD_DATA.
func (n *Nurbs) SetTextureCoordCallback(f NurbsTextureCoordDataHandler) {
	if n.nurbs == nil {
		panic("Uninitialised Nurbs. @see glu.NewNurbsRenderer.")
	}
	n.textureCoordData = f
	C.setGluNurbsCallback(n.nurbs, C.GLenum(NURBS_TEXTURE_COORD_DATA))
}

// SetEndCallback sets the callback for NURBS_END_DATA.
func (n *Nurbs) SetEndCallback(f NurbsEndDataHandler) {
	if n.nurbs == nil {
		panic("Uninitialised Nurbs. @see glu.NewNurbsRenderer.")
	}
	n.endData = f
	C.setGluNurbsCallback(n.nurbs, C.GLenum(NURBS_END_DATA))
}

// SetErrorCallback sets the callback for NURBS_ERROR.
func (n *Nurbs) SetErrorCallback(f NurbsErrorHandler) {
	if n.nurbs == nil {
		panic("Uninitialised Nurbs. @see glu.NewNurbsRenderer.")
	}
	n.errorData = f
	C.setGluNurbsCallback(n.nurbs, C.GLenum(NURBS_ERROR))
}
