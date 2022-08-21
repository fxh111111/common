package common

import (
	"github.com/fxh111111/common/maps"
	"testing"
)

func TestMaps(t *testing.T) {
	a := make(map[string]string)
	a["1"] = "1"
	maps.Keys(a)
	maps.Values(a)
}
