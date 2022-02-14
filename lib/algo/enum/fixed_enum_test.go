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
}
