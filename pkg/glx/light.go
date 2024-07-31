package glx

type Light interface {
	Apply(*Renderer)
}
