attribute vec3 position;

uniform mat4 uProj;
uniform mat4 uView;
uniform mat4 uModel;

void main(void) {
	gl_Position = uProj * uView * uModel * vec4(position, 1.0);
}
