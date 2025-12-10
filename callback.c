// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "callback.h"

void setGluTessCallback(GLUtesselator *tess, GLenum which) {
	switch(which) {
	case GLU_TESS_BEGIN_DATA:
		gluTessCallback(tess, which, (void (*)())goTessBeginData);
		break;
	case GLU_TESS_VERTEX_DATA:
		gluTessCallback(tess, which, (void (*)())goTessVertexData);
		break;
	case GLU_TESS_END_DATA:
		gluTessCallback(tess, which, (void (*)())goTessEndData);
		break;
	case GLU_TESS_ERROR_DATA:
		gluTessCallback(tess, which, (void (*)())goTessErrorData);
		break;
	case GLU_TESS_EDGE_FLAG_DATA:
		gluTessCallback(tess, which, (void (*)())goTessEdgeFlagData);
		break;
	case GLU_TESS_COMBINE_DATA:
		gluTessCallback(tess, which, (void (*)())goTessCombineData);
		break;
	}
}

void setGluNurbsCallback(GLUnurbs *nurbs, GLenum which) {
	switch(which) {
	case GLU_NURBS_BEGIN_DATA:
		gluNurbsCallback(nurbs, which, (void (*)())goNurbsBeginData);
		break;
	case GLU_NURBS_VERTEX_DATA:
		gluNurbsCallback(nurbs, which, (void (*)())goNurbsVertexData);
		break;
	case GLU_NURBS_NORMAL_DATA:
		gluNurbsCallback(nurbs, which, (void (*)())goNurbsNormalData);
		break;
	case GLU_NURBS_COLOR_DATA:
		gluNurbsCallback(nurbs, which, (void (*)())goNurbsColorData);
		break;
	case GLU_NURBS_TEXTURE_COORD_DATA:
		gluNurbsCallback(nurbs, which, (void (*)())goNurbsTextureCoordData);
		break;
	case GLU_NURBS_END_DATA:
		gluNurbsCallback(nurbs, which, (void (*)())goNurbsEndData);
		break;
	case GLU_NURBS_ERROR:
		gluNurbsCallback(nurbs, which, (void (*)())goNurbsErrorData);
		break;
	}
}

