package keyboard

type Code uint32

// # Key codes
// ## Android
// https://source.android.com/devices/input/keyboard-devices.html
// ## Windows
// http://nehe.gamedev.net/article/msdn_virtualkey_codes/15009/
// http://www.kbdedit.com/manual/low_level_vk_list.html

const (
	Unknown Code = 0

	CodeA Code = 4
	CodeB Code = 5
	CodeC Code = 6
	CodeD Code = 7
	CodeE Code = 8
	CodeF Code = 9
	CodeG Code = 10
	CodeH Code = 11
	CodeI Code = 12
	CodeJ Code = 13
	CodeK Code = 14
	CodeL Code = 15
	CodeM Code = 16
	CodeN Code = 17
	CodeO Code = 18
	CodeP Code = 19
	CodeQ Code = 20
	CodeR Code = 21
	CodeS Code = 22
	CodeT Code = 23
	CodeU Code = 24
	CodeV Code = 25
	CodeW Code = 26
	CodeX Code = 27
	CodeY Code = 28
	CodeZ Code = 29

	Code1 Code = 30
	Code2 Code = 31
	Code3 Code = 32
	Code4 Code = 33
	Code5 Code = 34
	Code6 Code = 35
	Code7 Code = 36
	Code8 Code = 37
	Code9 Code = 38
	Code0 Code = 39

	ReturnEnter        Code = 40
	Escape             Code = 41
	DeleteBackspace    Code = 42
	Tab                Code = 43
	Spacebar           Code = 44
	HyphenMinus        Code = 45 // -
	EqualSign          Code = 46 // =
	LeftSquareBracket  Code = 47 // [
	RightSquareBracket Code = 48 // ]
	Backslash          Code = 49 // \
	Semicolon          Code = 51 // ;
	Apostrophe         Code = 52 // '
	GraveAccent        Code = 53 // `
	Comma              Code = 54 // ,
	FullStop           Code = 55 // .
	Slash              Code = 56 // /
	CapsLock           Code = 57

	F1          Code = 58
	F2          Code = 59
	F3          Code = 60
	F4          Code = 61
	F5          Code = 62
	F6          Code = 63
	F7          Code = 64
	F8          Code = 65
	F9          Code = 66
	F10         Code = 67
	F11         Code = 68
	F12         Code = 69
	PrintScreen Code = 70

	Pause         Code = 72
	Insert        Code = 73
	Home          Code = 74
	PageUp        Code = 75
	DeleteForward Code = 76
	End           Code = 77
	PageDown      Code = 78

	RightArrow Code = 79
	LeftArrow  Code = 80
	DownArrow  Code = 81
	UpArrow    Code = 82

	KeypadNumLock     Code = 83
	KeypadSlash       Code = 84 // /
	KeypadAsterisk    Code = 85 // *
	KeypadHyphenMinus Code = 86 // -
	KeypadPlusSign    Code = 87 // +
	KeypadEnter       Code = 88
	Keypad1           Code = 89
	Keypad2           Code = 90
	Keypad3           Code = 91
	Keypad4           Code = 92
	Keypad5           Code = 93
	Keypad6           Code = 94
	Keypad7           Code = 95
	Keypad8           Code = 96
	Keypad9           Code = 97
	Keypad0           Code = 98
	KeypadFullStop    Code = 99  // .
	KeypadEqualSign   Code = 103 // =

	F13 Code = 104
	F14 Code = 105
	F15 Code = 106
	F16 Code = 107
	F17 Code = 108
	F18 Code = 109
	F19 Code = 110
	F20 Code = 111
	F21 Code = 112
	F22 Code = 113
	F23 Code = 114
	F24 Code = 115

	Help Code = 117

	Mute       Code = 127
	VolumeUp   Code = 128
	VolumeDown Code = 129

	LeftControl  Code = 224
	LeftShift    Code = 225
	LeftAlt      Code = 226
	LeftGUI      Code = 227
	RightControl Code = 228
	RightShift   Code = 229
	RightAlt     Code = 230
	RightGUI     Code = 231
)
