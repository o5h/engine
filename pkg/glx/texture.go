package glx

import (
	"image"
	"unsafe"

	"github.com/o5h/engine/internal/opengl/gl"
)

type TextureFormat uint32
type TextureDataFormat uint32
type TextureWrap uint32
type TextureFilter uint32

const (
	TextureFormat_Alpha          TextureFormat = gl.ALPHA
	TextureFormat_RGB            TextureFormat = gl.RGB
	TextureFormat_RGBA           TextureFormat = gl.RGBA
	TextureFormat_Luminance      TextureFormat = gl.LUMINANCE
	TextureFormat_LuminanceAlpha TextureFormat = gl.LUMINANCE_ALPHA

	TextureDataFormat_UByte      TextureDataFormat = gl.UNSIGNED_BYTE
	TextureDataFormat_UShort565  TextureDataFormat = gl.UNSIGNED_SHORT_5_6_5
	TextureDataFormat_UShort4444 TextureDataFormat = gl.UNSIGNED_SHORT_4_4_4_4
	TextureDataFormat_UShort5551 TextureDataFormat = gl.UNSIGNED_SHORT_5_5_5_1

	TextureWrap_ClampToEdge    TextureWrap = gl.CLAMP_TO_EDGE
	TextureWrap_MirroredRepeat TextureWrap = gl.MIRRORED_REPEAT
	TextureWrap_Repeat         TextureWrap = gl.REPEAT

	TextureFilter_Nearest TextureFilter = gl.NEAREST
	TextureFilter_Linear  TextureFilter = gl.LINEAR
)

type Texture struct {
	Width          int
	Height         int
	InternalFormat TextureFormat
	Format         TextureFormat
	Min            TextureFilter
	Mag            TextureFilter
	WrapS          TextureWrap
	WrapT          TextureWrap
	DataFormat     TextureDataFormat
	Data           []byte
	ID             uint32
}

func CreateTexture(img image.Image) *Texture {
	texture := &Texture{
		Width:          img.Bounds().Max.X,
		Height:         img.Bounds().Max.Y,
		InternalFormat: TextureFormat_RGBA,
		Format:         TextureFormat_RGBA,
		Mag:            TextureFilter_Linear,
		Min:            TextureFilter_Linear,
		WrapS:          TextureWrap_ClampToEdge,
		WrapT:          TextureWrap_ClampToEdge,
		DataFormat:     TextureDataFormat_UByte}
	texture.Data = ConvertImage(texture.Format, img)
	BindTexture(texture)
	return texture
}

func BindTexture(t *Texture) {
	t.ID = gl.GenTexture()
	gl.BindTexture(gl.TEXTURE_2D, t.ID)
	gl.TexImage2D(gl.TEXTURE_2D,
		0,
		uint32(t.InternalFormat),
		t.Width,
		t.Height,
		0,
		uint32(t.Format),
		uint32(t.DataFormat),
		uintptr(unsafe.Pointer(&t.Data[0])))

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, int(t.Mag))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, int(t.Min))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, int(t.WrapS))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, int(t.WrapT))
	gl.GenerateMipmap(gl.TEXTURE_2D)
}

// func NewRGBATexture(img image.Image) *Texture {
// texture := Texture{
// 	Width:          img.Bounds().Max.X,
// 	Height:         img.Bounds().Max.Y,
// 	InternalFormat: TextureFormat_RGBA,
// 	Format:         TextureFormat_RGBA,
// 	Mag:            TextureFilter_Linear,
// 	Min:            TextureFilter_Linear,
// 	WrapS:          TextureWrap_ClampToEdge,
// 	WrapT:          TextureWrap_ClampToEdge,
// 	DataFormat:     TextureDataFormat_UByte}
// texture.Data = ConvertImage(texture.Format, img)
// return &texture
// }

// func ConvertImage(f TextureFormat, img image.Image) (b []byte) {
// 	switch f {
// 	case TextureFormat_RGBA:
// 		b = convertRGBA(img)
// 	default:
// 		panic("Unsupported")
// 	}
// 	return
// }

// func convertRGBA(img image.Image) []byte {
// 	bounds := img.Bounds()
// 	rgba := image.NewRGBA(bounds)
// 	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
// 	w := bounds.Dx()
// 	h := bounds.Dy()
// 	bb := make([]byte, 0, w*h*4)
// 	for dy := h - 1; dy >= 0; dy-- {
// 		s := dy * w * 4
// 		bb = append(bb, rgba.Pix[s:s+w*4]...)
// 	}
// 	return bb
// }

// func (t *Texture) bind() {
// 	t.ID = gl.GenTexture()
// 	gl.BindTexture(gl.TEXTURE_2D, t.ID)

// 	gl.TexImage2D(gl.TEXTURE_2D,
// 		0,
// 		uint32(t.InternalFormat),
// 		t.Width,
// 		t.Height,
// 		0,
// 		uint32(t.Format),
// 		uint32(t.DataFormat),
// 		uintptr(unsafe.Pointer(&t.Data[0])))

// 	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, int(t.Mag))
// 	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, int(t.Min))
// 	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, int(t.WrapS))
// 	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, int(t.WrapT))
// 	gl.GenerateMipmap(gl.TEXTURE_2D)
// }

// func (t *Texture) apply(i int) {
// 	gl.ActiveTexture(uint32(gl.TEXTURE0 + i))
// 	gl.BindTexture(gl.TEXTURE_2D, t.ID)
// }

func (format TextureFormat) String() string {
	switch format {
	case TextureFormat_Alpha:
		return "ALPHA"
	case TextureFormat_RGB:
		return "RGB"
	case TextureFormat_RGBA:
		return "RGBA"
	case TextureFormat_Luminance:
		return "LUMINANCE"
	case TextureFormat_LuminanceAlpha:
		return "LUMINANCE_ALPHA"
	default:
		panic("Unsupported texture format")
	}
}

func ParseTextureFormat(s string) TextureFormat {
	switch s {
	case "ALPHA":
		return TextureFormat_Alpha
	case "RGB":
		return TextureFormat_RGB
	case "RGBA":
		return TextureFormat_RGBA
	case "LUMINANCE":
		return TextureFormat_Luminance
	case "LUMINANCE_ALPHA":
		return TextureFormat_LuminanceAlpha
	default:
		panic("Unsupported texture format")
	}
}

func (filter TextureFilter) String() string {
	switch filter {
	case TextureFilter_Nearest:
		return "NEAREST"
	case TextureFilter_Linear:
		return "LINEAR"
	default:
		panic("Unsupported texture filter format")
	}
}

func ParseTextureFilter(s string) TextureFilter {
	switch s {
	case "NEAREST":
		return TextureFilter_Nearest
	case "LINEAR":
		return TextureFilter_Linear
	default:
		panic("Unsupported texture filter format")
	}
}
