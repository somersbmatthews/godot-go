package editor

import (
	"log"
	"reflect"

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

/*
Object that holds the project-independent editor settings. These settings are generally visible in the Editor Settings menu. Accessing the settings is done by using the regular [Object] API, such as: [codeblock] settings.set(prop,value) settings.get(prop) list_of_settings = settings.get_property_list() [/codeblock]
*/
type EditorSettings struct {
	Resource
}

func (o *EditorSettings) BaseClass() string {
	return "EditorSettings"
}

/*
   Add a custom property info to a property. The dictionary must contain: name:[String](the name of the property) and type:[int](see TYPE_* in [@GlobalScope]), and optionally hint:[int](see PROPERTY_HINT_* in [@GlobalScope]), hint_string:[String]. Example: [codeblock] editor_settings.set("category/property_name", 0) var property_info = { "name": "category/property_name", "type": TYPE_INT, "hint": PROPERTY_HINT_ENUM, "hint_string": "one,two,three" } editor_settings.add_property_info(property_info) [/codeblock]
*/
func (o *EditorSettings) AddPropertyInfo(info *Dictionary) {
	log.Println("Calling EditorSettings.AddPropertyInfo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(info)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "add_property_info", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Erase a given setting (pass full property path).
*/
func (o *EditorSettings) Erase(property gdnative.String) {
	log.Println("Calling EditorSettings.Erase()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(property)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "erase", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Get the list of favorite directories for this project.
*/
func (o *EditorSettings) GetFavoriteDirs() *PoolStringArray {
	log.Println("Calling EditorSettings.GetFavoriteDirs()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_favorite_dirs", goArguments, "*PoolStringArray")

	returnValue := goRet.Interface().(*PoolStringArray)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the specific project settings path. Projects all have a unique sub-directory inside the settings path where project specific settings are saved.
*/
func (o *EditorSettings) GetProjectSettingsDir() gdnative.String {
	log.Println("Calling EditorSettings.GetProjectSettingsDir()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_project_settings_dir", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the list of recently visited folders in the file dialog for this project.
*/
func (o *EditorSettings) GetRecentDirs() *PoolStringArray {
	log.Println("Calling EditorSettings.GetRecentDirs()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_recent_dirs", goArguments, "*PoolStringArray")

	returnValue := goRet.Interface().(*PoolStringArray)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *EditorSettings) GetSetting(name gdnative.String) *Variant {
	log.Println("Calling EditorSettings.GetSetting()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_setting", goArguments, "*Variant")

	returnValue := goRet.Interface().(*Variant)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the global settings path for the engine. Inside this path you can find some standard paths such as: settings/tmp - used for temporary storage of files settings/templates - where export templates are located
*/
func (o *EditorSettings) GetSettingsDir() gdnative.String {
	log.Println("Calling EditorSettings.GetSettingsDir()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_settings_dir", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *EditorSettings) HasSetting(name gdnative.String) gdnative.Bool {
	log.Println("Calling EditorSettings.HasSetting()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "has_setting", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *EditorSettings) PropertyCanRevert(name gdnative.String) gdnative.Bool {
	log.Println("Calling EditorSettings.PropertyCanRevert()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "property_can_revert", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *EditorSettings) PropertyGetRevert(name gdnative.String) *Variant {
	log.Println("Calling EditorSettings.PropertyGetRevert()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "property_get_revert", goArguments, "*Variant")

	returnValue := goRet.Interface().(*Variant)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Set the list of favorite directories for this project.
*/
func (o *EditorSettings) SetFavoriteDirs(dirs *PoolStringArray) {
	log.Println("Calling EditorSettings.SetFavoriteDirs()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(dirs)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_favorite_dirs", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *EditorSettings) SetInitialValue(name gdnative.String, value *Variant, updateCurrent gdnative.Bool) {
	log.Println("Calling EditorSettings.SetInitialValue()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(name)
	goArguments[1] = reflect.ValueOf(value)
	goArguments[2] = reflect.ValueOf(updateCurrent)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_initial_value", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the list of recently visited folders in the file dialog for this project.
*/
func (o *EditorSettings) SetRecentDirs(dirs *PoolStringArray) {
	log.Println("Calling EditorSettings.SetRecentDirs()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(dirs)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_recent_dirs", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *EditorSettings) SetSetting(name gdnative.String, value *Variant) {
	log.Println("Calling EditorSettings.SetSetting()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(name)
	goArguments[1] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_setting", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   EditorSettingsImplementer is an interface for EditorSettings objects.
*/
type EditorSettingsImplementer interface {
	Class
}