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
	cam.projMatrix = n.NewMatrixPerspective(45.0, float32(n.Width())/float32(n.Height()), 1, 100.0)
	cam.transform = n.NewTransformIdentity()
	cam.transform.Scale = Vector3{0.75, 0.25, 1}.Scale(1)
	return cam
}

//forward direction
func (cam *camera) forward() Vector3 {
	return cam.transform.Rotation.Forward()
}

//right is the forward direction, rotated 90d
func (cam *camera) right() Vector3 {
	return cam.transform.Rotation.Right()
}

//up is the forward direction, rotated 90d on the pitch (z)
func (cam *camera) up() Vector3 {
	return cam.transform.Rotation.Up()
}

//translate moves the camera
func (cam *camera) translate(v Vector3) {
	cam.transform.Position = cam.transform.Position.Add(v)
}

//translateLocal rotates the vector and then moves by it. As a result, it is relative.
func (cam *camera) translateLocal(v Vector3) {
	right := cam.right().Scale(v.X)
	up := cam.up().Scale(v.Y)
	forward := cam.forward().Scale(v.Z)
	cam.transform.Position = cam.transform.Position.Add(right)
	cam.transform.Position = cam.transform.Position.Add(up)
	cam.transform.Position = cam.transform.Position.Add(forward)
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
