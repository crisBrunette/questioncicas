package question_generator

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorCorrectAnswer(t *testing.T) {

	assert.NotNil(t, Generator(&http.Request{}))
}
