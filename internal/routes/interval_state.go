// setup:feature:demo

package routes

import (
	"encoding/json"
	"os"
	"sync"
)

type intervalState struct {
	Master masterState    `json:"master"`
	Charts map[string]int `json:"charts"`
	Tiles  map[string]int `json:"tiles"`
	Admin  map[string]int `json:"admin"`
}

type masterState struct {
	Enabled    bool `json:"enabled"`
	IntervalMs int  `json:"intervalMs"`
}

var (
	statePath = "db/intervals.json"
	stateMu   sync.Mutex
)

func loadIntervalState() *intervalState {
	data, err := os.ReadFile(statePath)
	if err != nil {
		return nil
	}
	var s intervalState
	if err := json.Unmarshal(data, &s); err != nil {
		return nil
	}
	return &s
}

func saveIntervalState() {
	stateMu.Lock()
	defer stateMu.Unlock()

	s := intervalState{
		Charts: make(map[string]int),
		Tiles:  make(map[string]int),
		Admin:  make(map[string]int),
	}

	// Read master state
	rtMaster.mu.RLock()
	s.Master.Enabled = rtMaster.enabled
	s.Master.IntervalMs = rtMaster.intervalMs
	rtMaster.mu.RUnlock()

	// Read chart intervals
	rtIntervals.mu.RLock()
	for k, v := range rtIntervals.intervals {
		s.Charts[k] = v
	}
	rtIntervals.mu.RUnlock()

	// Read tile intervals
	numTileIntervals.mu.RLock()
	for k, v := range numTileIntervals.intervals {
		s.Tiles[k] = v
	}
	numTileIntervals.mu.RUnlock()

	// Read admin intervals
	adminIntervals.mu.RLock()
	for k, v := range adminIntervals.intervals {
		s.Admin[k] = v
	}
	adminIntervals.mu.RUnlock()

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return
	}
	os.WriteFile(statePath, data, 0644) //nolint:errcheck
}
