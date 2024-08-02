package zip

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Builder struct {
	f *os.File
	w *zip.Writer
}

func NewBuilder(name string) (*Builder, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	builder := Builder{
		f: f,
		w: zip.NewWriter(f),
	}
	return &builder, nil
}

func NewBuilderWriter(w io.Writer) *Builder { return &Builder{w: zip.NewWriter(w)} }

func (b *Builder) Close() error {
	err := b.w.Close()
	if err != nil {
		return err
	}
	if b.f == nil {
		return nil
	}
	return b.f.Close()
}

func (b *Builder) Append(base, name string) error { return b.AppendAs(base, name, name) }

func (b *Builder) AppendAs(base, name, zipName string) error {
	w, err := b.w.Create(zipName)
	if err != nil {
		return err
	}
	fileName := filepath.Join(base, name)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}
