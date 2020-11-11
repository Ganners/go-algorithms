package lda

import (
	"math/rand"
	"testing"
	"time"
)

func TestLDA(t *testing.T) {
	inputDocuments := []string{
		"ball ball ball planet galaxy",
		"referendum planet planet referendum referendum",
		"planet planet galaxy planet ball",
		"planet galaxy referendum planet ball",
	}

	rand.Seed(time.Now().UnixNano())
	classes := LDA(inputDocuments, 3, 1)
	t.Error(classes)
}
