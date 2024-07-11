package egl

import "errors"

type Display uintptr
type Surface uintptr
type Context uintptr
type Config uintptr

const (
	DONT_CARE = -1
	FALSE     = 0
	TRUE      = 1

	NO_SURFACE      = Surface(0)
	NO_CONTEXT      = Context(0)
	NO_DISPLAY      = Display(0)
	NO_CONFIG       = Config(0)
	DEFAULT_DISPLAY = NativeDisplay(0)

	// BindAPI/QueryAPI targets
	OPENGL_ES_API = 0x30A0

	CONTEXT_CLIENT_VERSION = 0x3098
	VG_ALPHA_FORMAT        = 0x3088
	VG_ALPHA_FORMAT_NONPRE = 0x308B
	VG_ALPHA_FORMAT_PRE    = 0x308C
	VG_COLORSPACE          = 0x3087
	VG_COLORSPACE_sRGB     = 0x3089
	VG_COLORSPACE_LINEAR   = 0x308A

	/* Config attributes */
	BUFFER_SIZE             = 0x3020
	ALPHA_SIZE              = 0x3021
	BLUE_SIZE               = 0x3022
	GREEN_SIZE              = 0x3023
	RED_SIZE                = 0x3024
	DEPTH_SIZE              = 0x3025
	STENCIL_SIZE            = 0x3026
	CONFIG_CAVEAT           = 0x3027
	CONFIG_ID               = 0x3028
	LEVEL                   = 0x3029
	MAX_PBUFFER_HEIGHT      = 0x302A
	MAX_PBUFFER_PIXELS      = 0x302B
	MAX_PBUFFER_WIDTH       = 0x302C
	NATIVE_RENDERABLE       = 0x302D
	NATIVE_VISUAL_ID        = 0x302E
	NATIVE_VISUAL_TYPE      = 0x302F
	SAMPLES                 = 0x3031
	SAMPLE_BUFFERS          = 0x3032
	SURFACE_TYPE            = 0x3033
	TRANSPARENT_TYPE        = 0x3034
	TRANSPARENT_BLUE_VALUE  = 0x3035
	TRANSPARENT_GREEN_VALUE = 0x3036
	TRANSPARENT_RED_VALUE   = 0x3037
	NONE                    = 0x3038 /* Attrib list terminator */
	BIND_TO_TEXTURE_RGB     = 0x3039
	BIND_TO_TEXTURE_RGBA    = 0x303A
	MIN_SWAP_INTERVAL       = 0x303B
	MAX_SWAP_INTERVAL       = 0x303C
	LUMINANCE_SIZE          = 0x303D
	ALPHA_MASK_SIZE         = 0x303E
	COLOR_BUFFER_TYPE       = 0x303F
	RENDERABLE_TYPE         = 0x3040
	MATCH_NATIVE_PIXMAP     = 0x3041 /* Pseudo-attribute (not queryable) */
	CONFORMANT              = 0x3042

	/* Config attribute mask bits */
	PBUFFER_BIT                 = 0x0001
	PIXMAP_BIT                  = 0x0002
	WINDOW_BIT                  = 0x0004
	VG_COLORSPACE_LINEAR_BIT    = 0x0020
	VG_ALPHA_FORMAT_PRE_BIT     = 0x0040
	MULTISAMPLE_RESOLVE_BOX_BIT = 0x0200
	SWAP_BEHAVIOR_PRESERVED_BIT = 0x0400

	OPENGL_ES_BIT  = 0x0001
	OPENVG_BIT     = 0x0002
	OPENGL_ES2_BIT = 0x0004
	OPENGL_BIT     = 0x0008

	SUCCESS             = 0x3000
	NOT_INITIALIZED     = 0x3001
	BAD_ACCESS          = 0x3002
	BAD_ALLOC           = 0x3003
	BAD_ATTRIBUTE       = 0x3004
	BAD_CONFIG          = 0x3005
	BAD_CONTEXT         = 0x3006
	BAD_CURRENT_SURFACE = 0x3007
	BAD_DISPLAY         = 0x3008
	BAD_MATCH           = 0x3009
	BAD_NATIVE_PIXMAP   = 0x300A
	BAD_NATIVE_WINDOW   = 0x300B
	BAD_PARAMETER       = 0x300C
	BAD_SURFACE         = 0x300D
	CONTEXT_LOST        = 0x300E
)

var (
	ErrNotInitialized    = errors.New("NOT_INITIALIZED")
	ErrBadAccess         = errors.New("BAD_ACCESS")
	ErrBadAlloc          = errors.New("BAD_ALLOC")
	ErrBadAttribute      = errors.New("BAD_ATTRIBUTE")
	ErrBadConfig         = errors.New("BAD_CONFIG")
	ErrBadContext        = errors.New("BAD_CONTEXT")
	ErrBadCurrentSurface = errors.New("BAD_CURRENT_SURFACE")
	ErrBadDisplay        = errors.New("BAD_DISPLAY")
	ErrBadMatch          = errors.New("BAD_MATCH")
	ErrBadNativePixmap   = errors.New("BAD_NATIVE_PIXMAP")
	ErrBadNativeWindow   = errors.New("BAD_NATIVE_WINDOW")
	ErrBadParameter      = errors.New("BAD_PARAMETER")
	ErrBadSurface        = errors.New("BAD_SURFACE")
	ErrContextLost       = errors.New("CONTEXT_LOST")
	ErrUnknown           = errors.New("Unknown")
)
