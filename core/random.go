package core

import (
	"math/rand"
	"sync"
)

var once sync.Once

type randomizer struct {
	on bool
}

var instance *randomizer

func Random() *randomizer {

	once.Do(func() {

		instance = &randomizer{on: true}
	})

	return instance
}

func (r *randomizer) VecInUnitSphere() Vec3 {
	if !r.on {
		return Vec3{0.0, 0.0, 0.0}
	}

	vec := Vec3{1.0, 0.0, 0.0}
	for vec.LenSqr() >= 1.0 {
		vec = Vec3{rand.Float32(), rand.Float32(), rand.Float32()}.Mul(2.0).Sub(Vec3{1.0, 1.0, 1.0})
	}
	return vec
}

func (r *randomizer) Enable() {
	r.on = true
}

func (r *randomizer) Disable() {
	r.on = false
}