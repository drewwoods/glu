// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package glu

// #ifdef __APPLE__
// #define GL_SILENCE_DEPRECATION
// #endif
//#include "callback.h"
import "C"
import (
	"unsafe"
)

// =============================================================================
//
// Section: Tesselator callbacks
//
// =============================================================================

type TessBeginHandler func(tessType uint32, polygonData interface{})

//export goTessBeginData
func goTessBeginData(tessType C.GLenum, tessPtr unsafe.Pointer) {
	var tess *Tesselator = (*Tesselator)(tessPtr)
	if tess == nil || tess.beginData == nil {
		return
	}
	tess.beginData((uint32)(tessType), tess.polyData)
}

// ===========================================================================

type TessVertexHandler func(vertexData interface{}, polygonData interface{})

//export goTessVertexData
func goTessVertexData(vertexDataPtr, tessPtr unsafe.Pointer) {
	var tess *Tesselator = (*Tesselator)(tessPtr)
	if tess == nil || tess.vertexData == nil {
		return
	}
	var wrapper *vertexDataWrapper = (*vertexDataWrapper)(vertexDataPtr)
	tess.vertexData(wrapper.data, tess.polyData)
}

// ===========================================================================

type TessEndHandler func(polygonData interface{})

//export goTessEndData
func goTessEndData(tessPtr unsafe.Pointer) {
	var tess *Tesselator = (*Tesselator)(tessPtr)
	if tess == nil || tess.endData == nil {
		return
	}
	tess.endData(tess.polyData)
}

// ===========================================================================

type TessErrorHandler func(errorNumber uint32, polygonData interface{})

//export goTessErrorData
func goTessErrorData(errorNumber C.GLenum, tessPtr unsafe.Pointer) {
	var tess *Tesselator = (*Tesselator)(tessPtr)
	if tess == nil || tess.errorData == nil {
		return
	}
	tess.errorData(uint32(errorNumber), tess.polyData)
}

// ===========================================================================

type TessEdgeFlagHandler func(flag bool, polygonData interface{})

//export goTessEdgeFlagData
func goTessEdgeFlagData(flag C.GLboolean, tessPtr unsafe.Pointer) {
	var tess *Tesselator = (*Tesselator)(tessPtr)
	if tess == nil || tess.edgeFlagData == nil {
		return
	}
	var goFlag bool
	if C.GLboolean(0) == flag {
		goFlag = false
	} else {
		goFlag = true
	}

	tess.edgeFlagData(goFlag, tess.polyData)
}

// ===========================================================================

type TessCombineHandler func(coords [3]float64,
	vertexData [4]interface{},
	weight [4]float32,
	polygonData interface{}) (outData interface{})

//export goTessCombineData
func goTessCombineData(coords, vertexData, weight, outData, tessPtr unsafe.Pointer) {
	var tess *Tesselator = (*Tesselator)(tessPtr)
	if tess == nil || tess.combineData == nil {
		return
	}

	var _coords *[3]float64 = (*[3]float64)(coords)
	var _weight *[4]float32 = (*[4]float32)(weight)

	var wrappers *[4]*vertexDataWrapper = (*[4]*vertexDataWrapper)(vertexData)
	var _vertexData [4]interface{}

	for i, wrapper := range *wrappers {
		// Work around for https://bugs.freedesktop.org/show_bug.cgi?id=51641
		// According to documentation, all vertex pointers should be valid.
		if wrapper == nil {
			_vertexData[i] = _vertexData[0]
		} else {
			_vertexData[i] = wrapper.data
		}
	}

	out := tess.combineData(*_coords, _vertexData, *_weight, tess.polyData)
	outWrapper := &vertexDataWrapper{out}

	tess.vertData = append(tess.vertData, outWrapper)
	_outData := (**vertexDataWrapper)(outData)
	*_outData = outWrapper
}

// =============================================================================
//
// Section: NURBS callbacks
//
// =============================================================================

type NurbsBeginDataHandler func(tessType uint32, polygonData interface{})
type NurbsVertexDataHandler func(vertexData []float32, polygonData interface{})
type NurbsNormalDataHandler func(normalData []float32, polygonData interface{})
type NurbsColorDataHandler func(colorData []float32, polygonData interface{})
type NurbsTextureCoordDataHandler func(texCoordData []float32, polygonData interface{})
type NurbsEndDataHandler func(polygonData interface{})
type NurbsErrorHandler func(errorNumber uint32, polygonData interface{})

//export goNurbsBeginData
func goNurbsBeginData(tessType C.GLenum, nurbsPtr unsafe.Pointer) {
	var nurbs *Nurbs = (*Nurbs)(nurbsPtr)
	if nurbs == nil || nurbs.beginData == nil {
		return
	}
	nurbs.beginData((uint32)(tessType), nurbs.polyData)
}

//export goNurbsVertexData
func goNurbsVertexData(vertexDataPtr, nurbsPtr unsafe.Pointer) {
	var nurbs *Nurbs = (*Nurbs)(nurbsPtr)
	if nurbs == nil || nurbs.vertexData == nil {
		return
	}
	var vertex []float32 = (*[3]float32)(vertexDataPtr)[:]
	nurbs.vertexData(vertex, nurbs.polyData)
}

//export goNurbsNormalData
func goNurbsNormalData(normalDataPtr, nurbsPtr unsafe.Pointer) {
	var nurbs *Nurbs = (*Nurbs)(nurbsPtr)
	if nurbs == nil || nurbs.normalData == nil {
		return
	}
	var normal []float32 = (*[3]float32)(normalDataPtr)[:]
	nurbs.normalData(normal, nurbs.polyData)
}

//export goNurbsColorData
func goNurbsColorData(colorDataPtr, nurbsPtr unsafe.Pointer) {
	var nurbs *Nurbs = (*Nurbs)(nurbsPtr)
	if nurbs == nil || nurbs.colorData == nil {
		return
	}
	var color []float32 = (*[4]float32)(colorDataPtr)[:]
	nurbs.colorData(color, nurbs.polyData)
}

//export goNurbsTextureCoordData
func goNurbsTextureCoordData(texCoordDataPtr, nurbsPtr unsafe.Pointer) {
	var nurbs *Nurbs = (*Nurbs)(nurbsPtr)
	if nurbs == nil || nurbs.textureCoordData == nil {
		return
	}
	var texCoord []float32 = (*[4]float32)(texCoordDataPtr)[:]
	nurbs.textureCoordData(texCoord, nurbs.polyData)
}

//export goNurbsEndData
func goNurbsEndData(nurbsPtr unsafe.Pointer) {
	var nurbs *Nurbs = (*Nurbs)(nurbsPtr)
	if nurbs == nil || nurbs.endData == nil {
		return
	}
	nurbs.endData(nurbs.polyData)
}

//export goNurbsErrorData
func goNurbsErrorData(errorNumber C.GLenum, nurbsPtr unsafe.Pointer) {
	var nurbs *Nurbs = (*Nurbs)(nurbsPtr)
	if nurbs == nil || nurbs.errorData == nil {
		return
	}
	nurbs.errorData(uint32(errorNumber), nurbs.polyData)
}

// =============================================================================

// Sets the callback for TESS_BEGIN_DATA.
func (tess *Tesselator) SetBeginCallback(f TessBeginHandler) {
	if tess.tess == nil {
		panic("Uninitialised Tesselator. @see glu.NewTess.")
	}
	tess.beginData = f
	C.setGluTessCallback(tess.tess, C.GLenum(TESS_BEGIN_DATA))
}

// Sets the callback for TESS_VERTEX_DATA.
func (tess *Tesselator) SetVertexCallback(f TessVertexHandler) {
	if tess.tess == nil {
		panic("Uninitialised Tesselator. @see glu.NewTess.")
	}
	tess.vertexData = f
	C.setGluTessCallback(tess.tess, C.GLenum(TESS_VERTEX_DATA))
}

// Sets the callback for TESS_END_DATA.
func (tess *Tesselator) SetEndCallback(f TessEndHandler) {
	if tess.tess == nil {
		panic("Uninitialised Tesselator. @see glu.NewTess.")
	}
	tess.endData = f
	C.setGluTessCallback(tess.tess, C.GLenum(TESS_END_DATA))
}

// Sets the callback for TESS_ERROR_DATA.
func (tess *Tesselator) SetErrorCallback(f TessErrorHandler) {
	if tess.tess == nil {
		panic("Uninitialised Tesselator. @see glu.NewTess.")
	}
	tess.errorData = f
	C.setGluTessCallback(tess.tess, C.GLenum(TESS_ERROR_DATA))
}

// Sets the callback for TESS_EDGE_FLAG_DATA.
func (tess *Tesselator) SetEdgeFlagCallback(f TessEdgeFlagHandler) {
	if tess.tess == nil {
		panic("Uninitialised Tesselator. @see glu.NewTess.")
	}
	tess.edgeFlagData = f
	C.setGluTessCallback(tess.tess, C.GLenum(TESS_EDGE_FLAG_DATA))
}

// Sets the callback for TESS_COMBINE_DATA.
func (tess *Tesselator) SetCombineCallback(f TessCombineHandler) {
	if tess.tess == nil {
		panic("Uninitialised Tesselator. @see glu.NewTess.")
	}
	tess.combineData = f
	C.setGluTessCallback(tess.tess, C.GLenum(TESS_COMBINE_DATA))
}
