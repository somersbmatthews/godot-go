package godot

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewPhysics2DDirectBodyStateSWFromPointer(ptr gdnative.Pointer) Physics2DDirectBodyStateSW {
func newPhysics2DDirectBodyStateSWFromPointer(ptr gdnative.Pointer) Physics2DDirectBodyStateSW {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Physics2DDirectBodyStateSW{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Software implementation of [Physics2DDirectBodyState]. This object exposes no new methods or properties and should not be used, as [Physics2DDirectBodyState] selects the best implementation available.
*/
type Physics2DDirectBodyStateSW struct {
	Physics2DDirectBodyState
	owner gdnative.Object
}

func (o *Physics2DDirectBodyStateSW) BaseClass() string {
	return "Physics2DDirectBodyStateSW"
}

// Physics2DDirectBodyStateSWImplementer is an interface that implements the methods
// of the Physics2DDirectBodyStateSW class.
type Physics2DDirectBodyStateSWImplementer interface {
	Physics2DDirectBodyStateImplementer
}
