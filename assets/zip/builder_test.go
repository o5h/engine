package zip_test

import (
	"testing"

	"github.com/o5h/engine/assets/zip"
	"github.com/o5h/testing/assert"
)

func TestBuilder_Append(t *testing.T) {
	builder, err := zip.NewBuilder("testdata/test.zip")
	assert.Nil(t, err)
	defer builder.Close()
	assert.Nil(t, builder.Append("./testdata/", "hello.txt"))
	assert.Nil(t, builder.AppendAs("./testdata/subfolder", "text.txt", "hello2.txt"))
	assert.Nil(t, builder.Close())
}
