package enum

import "testing"

func TestBehaviorPolymorphism(t *testing.T) {
	objs := []Renderable{
		Sphere{},
		Sphere{},
		Box{},
	}
	toRender := []Renderable{
		PolyPlane{},
		Box{},
		Collection{xs: objs},
	}
	RenderAll(toRender)

	// won't compile: int does not implement SceneObject
	// ScreenSpaceOcclusion(1)

	// testing the typeset
	ScreenSpaceOcclusion(Sphere{})
	ScreenSpaceOcclusion(Box{})
	ScreenSpaceOcclusion(PolyPlane{})
	ScreenSpaceOcclusion(Torus{})

	// Collection is not in the typeset! Hence it won't compile.
	// Collection does not implement SceneObject
	// ScreenSpaceOcclusion(Collection{xs: []Renderable{}})
}
