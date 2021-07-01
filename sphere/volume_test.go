package sphere

import (
	"math"
	"testing"
)

func TestComputeVolume(t *testing.T) {
	r := 2.
	expectedVolume := 33.510322
	volume := ComputeVolume(r)
	if math.Abs(volume-expectedVolume) > 1e-5 {
		t.Errorf("volume %f does not match expected %f", volume, expectedVolume)
	}
}

func TestComputeVolumeCubeRatios(t *testing.T) {
	for i := 1.; i < 5000.; i++ {
		trueVolume := ComputeVolume(i)
		cubeVolume := ComputeVolumeCubeRatios(i)

		if math.Abs(trueVolume-cubeVolume) > 1e-2 {
			t.Errorf("error is true great for r=%f. %f != %f", i, trueVolume, cubeVolume)
		}
	}
}
