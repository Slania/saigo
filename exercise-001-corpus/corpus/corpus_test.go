package corpus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorpus(t *testing.T) {
	result := Analyze("Are you serious? I knew you were.")
	assert.Equal(t, len(result), 6)
	assert.Equal(t, result[0].Word, "you")
	assert.Equal(t, result[0].Count, 2)
}

func BenchmarkCorpus(b *testing.B) {
	analysis_string := `How much wood could a wood chuck chuck if a wood chuck could chuck wood?
	She sells sea shells by the sea shore. Betty bought some butter, the butter was bitter, so
	she bought some better butter to make the bitter butter better.`

	for n := 0; n < b.N; n++ {
		Analyze(analysis_string)
	}
}
