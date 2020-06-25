package main

import (
	n "github.com/lachee/noodle"
)

type camera struct {
	transform Transform

	projMatrix Matrix
}

//newCamera creates a new camera
func newCamera() *camera {
	cam := &camera{}
	cam.projMatrix = n.NewMatrixPerspective(45.0, float64(n.Width())/float64(n.Height()), 1, 100.0)
	cam.transform = n.NewTransformIdentity()
	cam.transform.Scale = Vector3{0.75, 0.25, 1}.Scale(1)
	return cam
}

//translate moves the camera
func (cam *camera) translate(v Vector3) {
	cam.transform.Position = cam.transform.Position.Add(v)
}

//rotate moves the camera
func (cam *camera) rotate(v Vector3) {

	euler := n.NewQuaternionEuler(v)
	rotation := euler.Multiply(cam.transform.Rotation)
	cam.transform.Rotation = rotation
}

//lookAt forces the camera to look at a point
func (cam *camera) lookAt(v Vector3) {
	cam.transform.Rotation = n.NewQuaternionLookAt(cam.transform.Position, v, Vector3{0, 1, 0})
}

//view returns the view matrix
func (cam *camera) view() Matrix {
	return cam.transform.ToMatrix()
}

//projection returns the projection matrix
func (cam *camera) projection() Matrix {
	return cam.projMatrix
}
