package url

import (
	"io/fs"
	"os"
	"path"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetTags(t *testing.T) {
	data := `
tags:
  hello: world.com`
	tempFileFullPath := getTempFilePath()
	t.Cleanup(func() {
		os.Remove(tempFileFullPath)
	})

	os.WriteFile(tempFileFullPath, []byte(data), fs.FileMode(os.O_CREATE))
	tags := getTags(tempFileFullPath)

	assert.NotNil(t, tags)
	val, ok := tags.Tags["hello"]
	assert.True(t, ok)
	assert.Equal(t, "world.com", val)
}

func getTempFilePath() string {
	tempDir := os.TempDir()
	tempFileName := uuid.New().String()
	tempFileFullPath := path.Join(tempDir, tempFileName)
	return tempFileFullPath
}

func TestSaveTags(t *testing.T) {

	tags := &Tags{
		Tags: make(map[string]string),
	}
	tempFileFullPath := getTempFilePath()

	t.Cleanup(func() {
		os.Remove(tempFileFullPath)
	})

	tags.Tags["hello"] = "world"
	saveTags(tags, tempFileFullPath)
	data, err := os.ReadFile(tempFileFullPath)
	assert.Nil(t, err)
	contents := string(data)
	assert.Contains(t, contents, "hello:")
	assert.Contains(t, contents, "world")
}