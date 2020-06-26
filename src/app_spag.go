package main

import (
	"log"

	n "github.com/lachee/noodle"
)

//input handler from noodle
var input *n.InputHandler

//SpagApplication handles the game. Put your variables in here
type SpagApplication struct {
	renderer *meshRenderer
	camera   *camera
}

//Start allows for setup
func (app *SpagApplication) Start() bool {
	app.renderer = newMeshRenderer()
	app.camera = newCamera()
	input = n.Input()
	return true
}

//Update runs once a frame
func (app *SpagApplication) Update(dt float32) {

	axis := input.GetAxis2D(n.KeyA, n.KeyD, n.KeyS, n.KeyW)
	vertical := input.GetAxis(n.KeyC, n.KeySpace)
	app.camera.translateLocal(Vector3{axis.X, vertical, axis.Y}.Scale(dt * 10))
	//app.camera.translateLocal(Vector3{0, 0, 1}.Scale(dt * 10))

	//Just testing rotations
	//rotation := input.GetAxis2D(n.KeyArrowLeft, n.KeyArrowRight, n.KeyArrowDown, n.KeyArrowUp).Scale(90 * n.Deg2Rad * dt)
	//app.camera.rotate(rotation.ToVector3())

	//Always look at the center
	app.camera.lookAt(Vector3{0, 0, 0})
}

//Render draws the frame
func (app *SpagApplication) Render() {
	app.renderer.begin()

	transform := n.NewTransform(Vector3{0, 0, 0}, n.NewQuaternionIdentity(), Vector3{1, 1, 1})
	app.renderer.draw(&mesh{verts: rotCubeVerts, tris: rotCubeTris}, transform)

	app.renderer.draw(&mesh{verts: rotCubeVertsSquish, tris: rotCubeTris}, app.camera.transform)

	var rt Transform

	//X Axis
	for i := 0; i < 10; i++ {
		rt = n.NewTransform(Vector3{float32(i) * 0.1, 0, 0}, n.NewQuaternionIdentity(), Vector3{0.01, 0.01, 0.01})
		app.renderer.draw(&mesh{verts: rotCubeVerts, tris: rotCubeTris}, rt)
	}
	//Z Axis
	for i := 0; i < 10; i++ {
		rt = n.NewTransform(Vector3{0, 0, float32(i) * 0.1}, n.NewQuaternionIdentity(), Vector3{0.05, 0.05, 0.05})
		app.renderer.draw(&mesh{verts: rotCubeVerts, tris: rotCubeTris}, rt)
	}

	rt = n.NewTransformIdentity()
	rt.Rotation = n.NewQuaternionIdentity()
	rt.Scale = Vector3{0.025, 0.025, 0.025}
	for i := 0; i < 15; i++ {

		//Forward
		rt.Position = app.camera.transform.Position.Add(app.camera.forward().Scale(0.06 * float32(i)))
		app.renderer.draw(&mesh{verts: rotCubeVerts, tris: rotCubeTris}, rt)

		//Rightd
		rt.Position = app.camera.transform.Position.Add(app.camera.right().Scale(0.04 * float32(i)))
		app.renderer.draw(&mesh{verts: rotCubeVerts, tris: rotCubeTris}, rt)

		//Up
		rt.Position = app.camera.transform.Position.Add(app.camera.up().Scale(0.025 * float32(i)))
		app.renderer.draw(&mesh{verts: rotCubeVerts, tris: rotCubeTris}, rt)
	}

	log.Println(app.camera.transform.Position)
	app.renderer.end()
}
