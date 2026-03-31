package routes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnownTables(t *testing.T) {
	assert.True(t, knownTables["Tasks"])
	assert.True(t, knownTables["Items"])
	assert.True(t, knownTables["People"])
	assert.False(t, knownTables["Unknown"])
	assert.False(t, knownTables[""])
	assert.False(t, knownTables["Tasks; DROP TABLE"])
}
