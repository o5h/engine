package zip

import (
	"archive/zip"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Package struct {
	info   map[int64]string
	reader *zip.ReadCloser
}

func BuildZipPackage(name string, zipName string) error {
	infoFile, err := os.Open(name)
	if err != nil {
		return err
	}
	defer infoFile.Close()
	info := map[int64]string{}
	yaml.NewDecoder(infoFile).Decode(&info)

	zipFile, err := os.Create(zipName)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(zipFile)

	dir := filepath.Dir(name)
	addFileToZip(zw, dir, filepath.Base(name))
	for _, name := range info {
		addFileToZip(zw, dir, name)
	}
	defer zw.Close()
	return nil
}

func addFileToZip(zw *zip.Writer, dir, name string) error {
	w, err := zw.Create(name)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(filepath.Join(dir, name))
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

func OpenZipPackage(name string) (*Package, error) {
	reader, err := zip.OpenReader(name)
	if err != nil {
		return nil, err
	}
	pkg := Package{
		info:   map[int64]string{},
		reader: reader}
	return &pkg, pkg.init()
}

func (pkg *Package) init() error {
	f, err := pkg.reader.Open("pkg-info.yaml")
	if err != nil {
		return err
	}
	defer f.Close()
	return yaml.NewDecoder(f).Decode(&pkg.info)
}

func (pkg *Package) Open(name string) (fs.File, error) {
	return pkg.reader.Open(name)
}

func (pkg *Package) Close() error {
	return pkg.reader.Close()
}
