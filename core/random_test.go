package core

import (
	"testing"
)

func TestRandom_VecInUnitSphere(t *testing.T) {
	t.Log("We can generate a random vector in a unit sphere:")
	for i := 0; i < 10; i++ {
		randomVec := Random().VecInUnitSphere()
		if randomVec.LenSqr() < 1 {
			t.Logf("\tPASSED: generated %v, length is %v.\n", randomVec, randomVec.Len())
		} else {
			t.Fatalf("\tFAILED: generated %v, length is %v.\n", randomVec, randomVec.Len())
		}
	}
}

func TestRandom_VecInUnitSphere_Disable(t *testing.T) {
	t.Log("When we generate a random vector in a unit sphere, we can disable randomness:")
	Random().Disable()
	defer Random().Enable()

	randomVec := Random().VecInUnitSphere()
	expectedVec := Vec3{0.0, 0.0, 0.0}
	if randomVec == expectedVec {
		t.Logf("\tPASSED: result is %v, expected %v.\n", randomVec, expectedVec)
	} else {
		t.Fatalf("\tFAILED: result is %v, expected %v.\n", randomVec, expectedVec)
	}
}
