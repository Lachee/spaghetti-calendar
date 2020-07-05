package main

import (
	n "github.com/lachee/noodle"
)

type camera struct {
	transform *Transform
	target    Vector3

	x, y, xSpeed, ySpeed, scrollSpeed, yMinLimit, yMaxLimit, distanceMin, distanceMax, distance float32

	projMatrix Matrix
}

//newCamera creates a new camera
func newCamera() *camera {
	cam := &camera{}
	cam.projMatrix = n.NewMatrixPerspective(45.0, float32(n.GL.Width())/float32(n.GL.Height()), 1, 100.0)
	cam.transform = n.NewTransform()
	cam.target = Vector3{0, 0, 0}

	cam.xSpeed = 0.5
	cam.ySpeed = 0.5
	cam.scrollSpeed = 0.5

	cam.yMinLimit = -(n.PI / 2)
	cam.yMaxLimit = n.PI / 2

	cam.distanceMin = 0.1
	cam.distanceMax = 50
	cam.distance = cam.distanceMin //cam.distanceMin + ((cam.distanceMax - cam.distanceMin) / 2)
	cam.x = 0
	cam.y = 0
	return cam
}

//update moves the camera
func (cam *camera) update() {
	if n.Input().GetButton(0) {
		delta := n.Input().GetMouseDelta()
		cam.x += float32(delta.X) * -cam.xSpeed * 0.02
		cam.y += float32(delta.Y) * cam.ySpeed * 0.02
		cam.y = n.Clamp32(cam.y, cam.yMinLimit, cam.yMaxLimit)
	}

	scroll := n.Input().GetMouseScroll()
	if scroll > 0 {
		cam.distance -= cam.scrollSpeed
	}
	if scroll < 0 {
		cam.distance += cam.scrollSpeed
	}

	rotation := n.NewQuaternionEuler(Vector3{cam.y, cam.x, 0})
	cam.transform.SetLocalRotation(rotation)

	negDistance := Vector3{0, 0, cam.distance}
	position := rotation.Rotate(negDistance).Add(cam.target)
	cam.transform.SetLocalPosition(position)
	//cam.transform.LookAt(cam.target, Vector3{0, 1, 0})
}

//view returns the view matrix
func (cam *camera) view() Matrix {
	return cam.transform.GetWorldMatrix()
}

//projection returns the projection matrix
func (cam *camera) projection() Matrix {
	return cam.projMatrix

}
