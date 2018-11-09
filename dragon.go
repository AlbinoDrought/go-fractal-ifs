package main

import (
	"math/rand"
	"time"
)

func dragon() *IFS {
	return &IFS{
		Rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
		Depth: 2000,
		Scale: 30,
		Type:  TypeRandom,
		Groups: []Group{
			Group{
				Args: []float32{
					0.824074,
					0.281428,
					-0.212346,
					0.864198,
					-1.882290,
					-0.110607,
				},
				Probability: 0.8,
			},
			Group{
				Args: []float32{
					0.088272,
					0.520988,
					-0.463889,
					-0.377778,
					0.785360,
					8.095795,
				},
				Probability: 0.2,
			},
		},
	}
}
