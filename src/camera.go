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
	cam.transform = n.Transform{}
	cam.transform.Scale = Vector3{0.75, 0.25, 1}.Scale(1)
	return cam
}

//translate moves the camera
func (cam *camera) translate(v Vector3) {
	cam.transform.Position = cam.transform.Position.Add(v)
}

//rotate moves the camera
func (cam *camera) rotate(v Vector3) {
	rotation := n.NewQuaternionVector3ToVector3(Vector3{0, 1, 0}, v)
	cam.transform.Rotation = n.NewQuaternionFromMatrix(rotation.ToMatrix())

	//euler := cam.transform.Rotation.ToEuler()
	//euler = euler.Add(v)
	//cam.transform.Rotation = n.NewQuaternionVector3ToVector3(Vector3{0, 1, 0}, euler)
	//cam.transform.Rotation = n.NewQuaternionLookRotation(v, Vector3{0, 1, 0})
}

//view returns the view matrix
func (cam *camera) view() Matrix {
	return cam.transform.ToMatrix()
}

//projection returns the projection matrix
func (cam *camera) projection() Matrix {
	return cam.projMatrix
}
