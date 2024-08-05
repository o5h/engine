package zip

import (
	"io"
	"os"
	"testing"

	"github.com/o5h/testing/assert"
)

func TestBuildZipPackage(t *testing.T) {
	zipFile := "testdata/test.zip"

	assert.Nil(t, BuildZipPackage("testdata/pkg-info.yaml", zipFile))
	pkg, err := OpenZipPackage(zipFile)
	assert.Nil(t, err)
	rs, err := pkg.Open("hello.txt")
	assert.Nil(t, err)
	s, err := io.ReadAll(rs)
	assert.Nil(t, err)
	assert.Eq(t, string(s), "Hello from Zip Package!")
	assert.Nil(t, rs.Close())
	assert.Nil(t, pkg.Close())
	assert.Nil(t, os.Remove(zipFile))
}
