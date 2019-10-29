package hotpepper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuuis/RecommendSystem/infrastructures"
)

func TestHotpepperUnit(t *testing.T) {
	infrastructures.InitEnvWithPath("../../")

	payload := &Payload{
		Keywords: "肉まん",
	}

	res, err := Request(payload)
	if err != nil {
		t.Errorf("cannot access api: %+v", err)
	}
	if len(res.Results.Error) != 0 {
		t.Errorf("api errors: %+v", res.Results.Error)
	}

	assert.Equal(t, 1, res.Results.ResultsStart)
}
