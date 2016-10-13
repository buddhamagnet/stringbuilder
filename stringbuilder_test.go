package stringbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	assert.Equal(t, "node:economist:12345", buildConcat("node", "economist", "12345"))
}

func TestSprintf(t *testing.T) {
	assert.Equal(t, "node:economist:12345", buildSprintf("node", "economist", "12345"))
}

func TestJoin(t *testing.T) {
	assert.Equal(t, "node:economist:12345", buildJoin("node", "economist", "12345"))
}

func TestBuffer(t *testing.T) {
	assert.Equal(t, "node:economist:12345", buildBuffer("node", "economist", "12345"))
}

func BenchmarkBuildConcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buildConcat("node", "economist", "12345")
	}
}

func BenchmarkBuildSprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buildSprintf("node", "economist", "12345")
	}
}

func BenchmarkBuildJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buildJoin("node", "economist", "12345")
	}
}

func BenchmarkBuildBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buildBuffer("node", "economist", "12345")
	}
}
