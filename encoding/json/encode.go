package json

import (
	"encoding/json"

	"github.com/JudicaelT/betterstandards/assert"
)

func MustMarshal(value any) []byte {
	var valueJson []byte
	valueJson, err := json.Marshal(value)
	assert.NoError(err)
	return valueJson
}
