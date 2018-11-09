package main

import (
	"math/rand"
	"time"
)

func stickfern() *IFS {
	return &IFS{
		Rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
		Depth: 13,
		Scale: 300,
		Type:  TypeFork,
		Groups: []Group{
			Group{
				Args: []float32{
					0.0100,
					-0.4100,
					0.3900,
					0.0000,
					-0.2800,
					-0.1850,
				},
			},
			Group{
				Args: []float32{
					0.7000,
					0.3300,
					-0.3500,
					0.7000,
					0.1850,
					0.0150,
				},
			},
			Group{
				Args: []float32{
					0.0000,
					0.1750,
					0.0130,
					0.4600,
					-0.0950,
					-0.2850,
				},
			},
		},
	}
}
