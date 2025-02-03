package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ProcessAndValidateIncomingTasks(t *testing.T) {
	// test a request with a invalid uuid
	_, err := ProcessAndValidateIncomingTasks([]byte(`[
  {
    "id": "invalid-236f-46da-8600-25b60f92c091",
    "text": "I have to do the dishes",
    "done": false,
    "updatedAt": "2025-02-03T17:40:50.961Z",
    "createdAt": "2025-02-03T17:40:50.961Z"
  }
]`))
	assert.NotNil(t, err)

	// test a valid input
	_, err = ProcessAndValidateIncomingTasks([]byte(`[
  {
    "id": "8631fdef-236f-46da-8600-25b60f92c091",
    "text": "I have to do the dishes",
    "done": false,
    "updatedAt": "2025-02-03T17:40:50.961Z",
    "createdAt": "..."
  }
]`))
	assert.Nil(t, err)

}
