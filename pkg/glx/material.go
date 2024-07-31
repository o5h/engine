package glx

type Material interface {
	Apply(*Renderer)
}
