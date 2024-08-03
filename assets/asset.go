package assets

import (
	"io"
	"io/fs"
	"os"
)

type Decoder[T any] func(io.Reader) (T, error)
type Encoder[T any] func(io.Writer, T) error

var assetsFS fs.FS

func init() {
	assetsFS = os.DirFS("./")
}

func SetFS(f fs.FS) { assetsFS = f }

func MustDecode[T any](decoder Decoder[T], name string) T {
	t, err := Decode[T](decoder, name)
	if err != nil {
		panic(err)
	}
	return t
}

func Decode[T any](decoder Decoder[T], name string) (T, error) {
	var value T
	r, err := assetsFS.Open(name)
	if err != nil {
		return value, err
	}
	defer r.Close()
	value, err = decoder(r)
	if err != nil {
		return value, err
	}
	return value, nil
}
