package main

import (
	"math/rand"
	"time"
)

func fern() *IFS {
	return &IFS{
		Rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
		Depth: 3000,
		Scale: 30,
		Type:  TypeRandom,
		Groups: []Group{
			Group{
				Args: []float32{
					0.0,
					0.0,
					0.0,
					0.16,
					0.0,
					0.0,
				},
				Probability: 0.01,
			},
			Group{
				Args: []float32{
					0.2,
					-0.26,
					0.23,
					0.22,
					0.0,
					1.6,
				},
				Probability: 0.07,
			},
			Group{
				Args: []float32{
					-0.15,
					0.28,
					0.26,
					0.24,
					0.0,
					0.44,
				},
				Probability: 0.07,
			},
			Group{
				Args: []float32{
					0.85,
					0.04,
					-0.04,
					0.85,
					0.0,
					1.6,
				},
				Probability: 0.85,
			},
		},
	}
}
