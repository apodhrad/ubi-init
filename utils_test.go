package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEsists(t *testing.T) {
	tmpDir := t.TempDir()

	file := filepath.Join(tmpDir, "foo.txt")
	info, ok := exists(file)
	assert.Nil(t, info)
	assert.False(t, ok)

	err := os.WriteFile(file, []byte(""), 0644)
	assert.Nil(t, err)
	info, ok = exists(file)
	assert.NotNil(t, info)
	assert.False(t, info.IsDir())
	assert.True(t, ok)

	dir := filepath.Join(tmpDir, "foo")
	info, ok = exists(dir)
	assert.Nil(t, info)
	assert.False(t, ok)

	err = os.Mkdir(dir, 0755)
	assert.Nil(t, err)
	info, ok = exists(dir)
	assert.NotNil(t, info)
	assert.True(t, info.IsDir())
	assert.True(t, ok)
}

func TestTemplatePath(t *testing.T) {
	d := &testDirEntry{isDir: false}
	s := templatePath("/root", "/root/foo.txt", d)
	assert.Equal(t, "foo.txt", s)

	d = &testDirEntry{isDir: true}
	s = templatePath("/root", "/root/foo.txt", d)
	assert.Equal(t, "foo.txt/", s)
}

type testDirEntry struct {
	isDir bool
}

func (d *testDirEntry) Name() string               { return "" }
func (d *testDirEntry) IsDir() bool                { return d.isDir }
func (d *testDirEntry) Type() fs.FileMode          { return 0 }
func (d *testDirEntry) Info() (fs.FileInfo, error) { return nil, nil }
