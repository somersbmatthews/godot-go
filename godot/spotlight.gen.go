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

//func NewSpotLightFromPointer(ptr gdnative.Pointer) SpotLight {
func newSpotLightFromPointer(ptr gdnative.Pointer) SpotLight {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := SpotLight{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A SpotLight light is a type of [Light] node that emits lights in a specific direction, in the shape of a cone. The light is attenuated through the distance and this attenuation can be configured by changing the energy, radius and attenuation parameters of [Light]. TODO: Image of a spotlight.
*/
type SpotLight struct {
	Light
	owner gdnative.Object
}

func (o *SpotLight) BaseClass() string {
	return "SpotLight"
}

// SpotLightImplementer is an interface that implements the methods
// of the SpotLight class.
type SpotLightImplementer interface {
	LightImplementer
}