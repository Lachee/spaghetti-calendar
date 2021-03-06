package main

import (
	"log"

	n "github.com/lachee/noodle"
)

type meshRenderer struct {
	shader       *n.Shader
	vertexBuffer n.WebGLBuffer
	indexBuffer  n.WebGLBuffer

	previousMesh *mesh

	uModelMatrixLoc      n.WebGLUniformLocation
	uViewMatrixLoc       n.WebGLUniformLocation
	uProjectionMatrixLoc n.WebGLUniformLocation

	projectionMatrix Matrix
	cameraMatrix     Matrix
}

//newMeshRenderer creates the renderer and sets up the buffers
func newMeshRenderer() *meshRenderer {
	mr := &meshRenderer{}

	//Can be combined, see the UI Renderer as an example
	mr.vertexBuffer = n.GL.CreateBuffer()
	mr.indexBuffer = n.GL.CreateBuffer()

	//Create the shader
	shader, err := n.LoadShaderFromURL("resources/shader/mesh.vert", "/resources/shader/mesh.frag")
	if err != nil {
		log.Fatalln("Failed to load shader. ", err)
	}
	mr.shader = shader

	mr.shader.BindVertexData("position", n.GlArrayBuffer, mr.vertexBuffer, 3, n.GlFloat, false, 0, 0)

	mr.uProjectionMatrixLoc = mr.shader.GetUniformLocation("uProj")
	mr.uViewMatrixLoc = mr.shader.GetUniformLocation("uView")
	mr.uModelMatrixLoc = mr.shader.GetUniformLocation("uModel")

	mr.shader.Use()

	//Setup initial camera
	mr.projectionMatrix = n.NewMatrixPerspective(45, float32(n.GL.Width())/float32(n.GL.Height()), 1, 100.0)
	//mr.cameraMatrix = n.NewMatrixTranslate(Vector3{10, -1, 0})

	//p := Vector3{-10, 10, 0}
	//q := n.NewQuaternionLookAt(p, Vector3{0, 0, 0}, Vector3{0, 1, 0})
	//mr.cameraMatrix = n.NewMatrixTranslate(p).Multiply(n.NewMatrixRotation(q))

	mr.cameraMatrix = n.NewMatrixTranslate(Vector3{0, 0, 0})
	log.Println(mr.cameraMatrix)

	//mr.cameraMatrix = mr.cameraMatrix.Multiply(n.NewMatrixLookAt(Vector3{20.0, 20.0, 20.0}, Vector3{0, 0, 0}, Vector3{0, 1, 0}))
	//log.Println(mr.cameraMatrix)
	//mr.cameraMatrix = n.NewMatrixRotation(n.NewQuaternionAxis(n.PI/3, Vector3{1, 0, 0})).Multiply(n.NewMatrixTranslate(Vector3{0, 20, -20}))

	return mr
}

//begin prepares the renderer
func (mr *meshRenderer) begin() {

	//Clear the scene
	n.GL.ClearDepth(1)
	n.GL.DepthFunc(n.GlLEqual)

	//Clear
	n.GL.Enable(n.GlDepthTest)
	n.GL.Clear(n.GlColorBufferBit | n.GlDepthBufferBit)

	//Use SHader
	mr.shader.Use()

	//Update the camera
	n.GL.UniformMatrix4fv(mr.uProjectionMatrixLoc, mr.projectionMatrix)
	n.GL.UniformMatrix4fv(mr.uViewMatrixLoc, mr.cameraMatrix)
}

//end flushes the renderer
func (mr *meshRenderer) end() {
	mr.previousMesh = nil
}

//draw a particular mesh to the screen
func (mr *meshRenderer) draw(mesh *mesh, modelMatrix Matrix) {

	//New Mesh, Who Dis?
	if mr.previousMesh != mesh {
		mr.previousMesh = mesh
		n.GL.BindBuffer(n.GlArrayBuffer, mr.vertexBuffer)
		n.GL.BufferData(n.GlArrayBuffer, mr.previousMesh.verts, n.GlStaticDraw)

		n.GL.BindBuffer(n.GlElementArrayBuffer, mr.indexBuffer)
		n.GL.BufferData(n.GlElementArrayBuffer, mr.previousMesh.tris, n.GlStaticDraw)
	}

	//Set the shader deets
	n.GL.UniformMatrix4fv(mr.uModelMatrixLoc, modelMatrix)

	//Render
	n.GL.DrawElements(n.GlLines, len(mr.previousMesh.tris), n.GlUnsignedShort, 0)
}
