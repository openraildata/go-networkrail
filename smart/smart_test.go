package smart_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/openraildata/go-networkrail/smart"
)

func TestImport(t *testing.T) {
	file, err := os.Open("E:\\RailData\\NetworkRail\\SMART\\2024-06-11\\SMARTExtract.json")
	assert.NoError(t, err, "unable to open JSON file")

	data := smart.Data{}

	err = json.NewDecoder(file).Decode(&data)
	assert.NoError(t, err, "unable to decode JSON")
}
