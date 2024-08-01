package glx

import (
	"log"

	"github.com/o5h/opengles/gl"
)

// type Uniform struct {
// 	Name     string
// 	Location int
// 	Type     int
// }

type UniformBinding func(*Uniforms, int)

type UniformApplier struct {
	Name     uniform
	Location int
	Bind     UniformBinding
}

type Program struct {
	Attribs          []Attribute
	MaterialUniforms []*UniformApplier
	ModelUniforms    []*UniformApplier
	LightUniforms    []*UniformApplier

	VertexShaderId   uint32
	FragmentShaderId uint32
	ID               uint32
}

func CreateProgram(vert, frag string) *Program {
	program := &Program{}
	program.ID = gl.CreateProgram()
	var err error
	program.VertexShaderId, err = compileShader(vert, gl.VERTEX_SHADER)
	must(err)
	program.FragmentShaderId, err = compileShader(frag, gl.FRAGMENT_SHADER)
	must(err)
	for _, a := range Attributes {
		gl.BindAttribLocation(program.ID, uint32(a.Location()), string(a))
	}

	gl.AttachShader(program.ID, program.FragmentShaderId)
	gl.AttachShader(program.ID, program.VertexShaderId)
	gl.LinkProgram(program.ID)

	if gl.GetProgramiv(program.ID, gl.LINK_STATUS) == gl.FALSE {
		msg := gl.GetProgramInfoLog(program.ID)
		log.Println("ERROR:", msg)
		panic(msg)
	}
	program.initializeAttributes()
	program.initializeUniforms()
	return program
}

func (p *Program) addUniform(u uniform, loc int) {
	applier := &UniformApplier{
		Name:     u,
		Location: loc,
		Bind:     uniformBindings[u]}

	switch u.Type() {
	case ut_Mesh:
		p.ModelUniforms = append(p.ModelUniforms, applier)
	case ut_Material:
		p.MaterialUniforms = append(p.MaterialUniforms, applier)
	case ut_Light:
		p.LightUniforms = append(p.LightUniforms, applier)
	}
}

func (p *Program) initializeAttributes() {
	num := gl.GetProgramiv(p.ID, gl.ACTIVE_ATTRIBUTES)
	for i := 0; i < int(num); i++ {
		name, size, ty := gl.GetActiveAttrib(p.ID, uint32(i))
		location := gl.GetAttribLocation(p.ID, name)
		log.Printf("ATTRIBUTE %v, LOCATION:=%d, size=%d, type=0x%x", name, location, size, ty)
		attrib := Attribute(name)
		p.Attribs = append(p.Attribs, attrib)
	}
}

func (p *Program) initializeUniforms() {
	num := gl.GetProgramiv(p.ID, gl.ACTIVE_UNIFORMS)
	for i := 0; i < int(num); i++ {
		name, size, ty := gl.GetActiveUniform(p.ID, uint32(i))
		location := gl.GetUniformLocation(p.ID, name)
		p.addUniform(uniform(name), location)
		log.Printf("UNIFORM %v, LOCATION:=%d, size=%d, type=0x%x", name, location, size, ty)
	}
}
