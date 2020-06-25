package main

import (
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
	app.camera.translate(Vector3{axis.X, vertical, axis.Y})

	rotation := input.GetAxis2D(n.KeyArrowLeft, n.KeyArrowRight, n.KeyArrowDown, n.KeyArrowUp)
	app.camera.rotate(rotation.ToVector3())
}

//Render draws the frame
func (app *SpagApplication) Render() {
	app.renderer.begin()

	transform := n.NewTransform(Vector3{0, 0, 0}, n.NewQuaternionFromEuler(Vector3{0, 0, 0}), Vector3{1, 1, 1})
	app.renderer.draw(&mesh{verts: rotCubeVerts, tris: rotCubeTris}, transform)

	app.renderer.draw(&mesh{verts: rotCubeVerts, tris: rotCubeTris}, app.camera.transform)

	app.renderer.end()
}