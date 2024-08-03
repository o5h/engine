module github.com/o5h/engine

go 1.22

toolchain go1.22.5

require (
	github.com/o5h/glx v0.0.0
	github.com/o5h/opengles v0.0.0
	golang.org/x/sys v0.22.0
	gopkg.in/yaml.v3 v3.0.1
)

require github.com/o5h/glm v0.0.0-20200112235157-ed7a80d112f3 // indirect

replace (
	github.com/o5h/glm => ../glm
	github.com/o5h/glx => ../glx
	github.com/o5h/opengles => ../opengles

)
