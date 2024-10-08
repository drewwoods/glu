// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package glu

// #cgo darwin LDFLAGS: -framework Carbon -framework OpenGL
// #cgo linux LDFLAGS: -lGLU
// #cgo windows LDFLAGS: -lglu32
//
// #ifdef __APPLE__
// #define GL_SILENCE_DEPRECATION
//   #include <OpenGL/glu.h>
// #else
//   #include <GL/glu.h>
// #endif
import "C"
import (
	"errors"
	"reflect"
	"unsafe"
)

type Float interface {
	~float64 | ~float32
}

func ptr(v interface{}) unsafe.Pointer {

	if v == nil {
		return unsafe.Pointer(nil)
	}

	rv := reflect.ValueOf(v)
	var et reflect.Value
	switch rv.Type().Kind() {
	case reflect.Uintptr:
		offset, _ := v.(uintptr)
		return unsafe.Pointer(offset)
	case reflect.Ptr:
		et = rv.Elem()
	case reflect.Slice:
		et = rv.Index(0)
	default:
		panic("type must be a pointer, a slice, uintptr or nil")
	}

	return unsafe.Pointer(et.UnsafeAddr())
}

func ErrorString(error uint32) (string, error) {
	e := unsafe.Pointer(C.gluErrorString(C.GLenum(error)))
	if e == nil {
		return "", errors.New("Invalid GL error code")
	}
	return C.GoString((*C.char)(e)), nil
}

func Build2DMipmaps(target uint32, internalFormat int, width, height int, format, typ uint32, data interface{}) int {
	return int(C.gluBuild2DMipmaps(
		C.GLenum(target),
		C.GLint(internalFormat),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLenum(format),
		C.GLenum(typ),
		ptr(data),
	))
}

func Perspective(fovy, aspect, zNear, zFar float64) {
	C.gluPerspective(
		C.GLdouble(fovy),
		C.GLdouble(aspect),
		C.GLdouble(zNear),
		C.GLdouble(zFar),
	)
}

func LookAt[num Float](eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ num) {
	C.gluLookAt(
		C.GLdouble(eyeX),
		C.GLdouble(eyeY),
		C.GLdouble(eyeZ),
		C.GLdouble(centerX),
		C.GLdouble(centerY),
		C.GLdouble(centerZ),
		C.GLdouble(upX),
		C.GLdouble(upY),
		C.GLdouble(upZ),
	)
}

func PickMatrix(x, y, delX, delY float64, view *[4]int32) {
	v := (*C.GLint)(unsafe.Pointer(view))

	C.gluPickMatrix(
		C.GLdouble(x),
		C.GLdouble(y),
		C.GLdouble(delX),
		C.GLdouble(delY),
		v,
	)
}

func UnProject(winX, winY, winZ float64, model, proj *[16]float64, view *[4]int32) (float64, float64, float64) {
	var ox, oy, oz C.GLdouble

	m := (*C.GLdouble)(unsafe.Pointer(model))
	p := (*C.GLdouble)(unsafe.Pointer(proj))
	v := (*C.GLint)(unsafe.Pointer(view))

	C.gluUnProject(
		C.GLdouble(winX),
		C.GLdouble(winY),
		C.GLdouble(winZ),
		m,
		p,
		v,
		&ox,
		&oy,
		&oz,
	)

	return float64(ox), float64(oy), float64(oz)
}

func Project(projX, projY, projZ float64, model, proj *[16]float64, view *[4]int32) (float64, float64, float64) {
	var ox, oy, oz C.GLdouble

	m := (*C.GLdouble)(unsafe.Pointer(model))
	p := (*C.GLdouble)(unsafe.Pointer(proj))
	v := (*C.GLint)(unsafe.Pointer(view))

	C.gluProject(
		C.GLdouble(projX),
		C.GLdouble(projY),
		C.GLdouble(projZ),
		m,
		p,
		v,
		&ox,
		&oy,
		&oz,
	)

	return float64(ox), float64(oy), float64(oz)
}

func NewQuadric() unsafe.Pointer {
	return unsafe.Pointer(C.gluNewQuadric())
}

func Sphere(q unsafe.Pointer, radius float32, slices, stacks int) {
	C.gluSphere((*C.GLUquadric)(q), C.GLdouble(radius), C.GLint(slices), C.GLint(stacks))
}

func Cylinder(q unsafe.Pointer, base, top, height float32, slices, stacks int) {
	C.gluCylinder((*C.GLUquadric)(q), C.GLdouble(base), C.GLdouble(top), C.GLdouble(height), C.GLint(slices), C.GLint(stacks))
}

func Disk(q unsafe.Pointer, inner, outer float32, slices, loops int) {
	C.gluDisk((*C.GLUquadric)(q), C.GLdouble(inner), C.GLdouble(outer), C.GLint(slices), C.GLint(loops))
}

func PartialDisk(q unsafe.Pointer, inner, outer float32, slices, loops int, startAngle, sweepAngle float32) {
	C.gluPartialDisk((*C.GLUquadric)(q), C.GLdouble(inner), C.GLdouble(outer), C.GLint(slices), C.GLint(loops), C.GLdouble(startAngle), C.GLdouble(sweepAngle))
}
