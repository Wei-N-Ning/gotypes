package enum

// Renderable defines the behavior
type Renderable interface {
	Render()
}

type Sphere struct{}

func (sphere Sphere) Render() {}

type Box struct{}

func (box Box) Render() {}

type PolyPlane struct{}

func (pPlane PolyPlane) Render() {}

type Torus struct {
	Radius float32
}

func (torus Torus) Render() {}

type Collection struct {
	xs []Renderable
}

func (coll Collection) Render() {}

// SceneObject defines the typeset
type SceneObject interface {
	Sphere | Box | PolyPlane | Torus
}

// RenderAll uses the behavior-polymorphism
func RenderAll(objs []Renderable) {
	for _, obj := range objs {
		obj.Render()
	}
}

// ScreenSpaceOcclusion uses the typeset (SceneObject) to limit the type of
// Renderable supported by this function
func ScreenSpaceOcclusion[T SceneObject](x T) float32 {
	return 0
}
