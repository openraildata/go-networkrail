package corpus_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/openraildata/go-networkrail/corpus"
)

func TestImport(t *testing.T) {
	file, err := os.Open("E:\\RailData\\NetworkRail\\CORPUS\\2024-06-11\\CORPUSExtract.json")
	assert.NoError(t, err, "unable to open JSON file")

	data := corpus.Data{}

	err = json.NewDecoder(file).Decode(&data)
	assert.NoError(t, err, "unable to decode JSON")
}
