package pool

// AutoScaler — dynamically adjusts connection pool size based on load.
//
// YOU MUST IMPLEMENT the methods marked with TODO.
// MetricsMonitor is working — use it to read utilization data.

import (
	"time"
)

type ScaleAction int

const (
	Hold ScaleAction = iota
	ScaleUp
	ScaleDown
)

type ScaleDecision struct {
	Action           ScaleAction
	CurrentSize      int
	RecommendedSize  int
	Reason           string
	Utilization      float64
}

type AutoScaler struct {
	monitor          *MetricsMonitor
	minPoolSize      int
	maxPoolSize      int
	currentPoolSize  int
	scaleUpThreshold   float64   // Utilization % above which to scale up
	scaleDownThreshold float64   // Utilization % below which to scale down
	cooldownPeriod     time.Duration
	lastScaleTime      time.Time
	scaleHistory       []ScaleDecision
}

func NewAutoScaler(monitor *MetricsMonitor, minSize, maxSize, currentSize int) *AutoScaler {
	return &AutoScaler{
		monitor:            monitor,
		minPoolSize:        minSize,
		maxPoolSize:        maxSize,
		currentPoolSize:    currentSize,
		scaleUpThreshold:   80.0,
		scaleDownThreshold: 30.0,
		cooldownPeriod:     60 * time.Second,
		scaleHistory:       make([]ScaleDecision, 0),
	}
}

// EvaluateScaling decides whether to scale the pool up, down, or hold.
//
// 1. Get average utilization from monitor over the last 5 minutes
// 2. Check cooldown — if last scale action was within cooldownPeriod, return Hold
// 3. If utilization > scaleUpThreshold:
//    - Calculate new size: currentPoolSize * 1.5 (rounded up)
//    - Cap at maxPoolSize
//    - Return ScaleUp decision
// 4. If utilization < scaleDownThreshold:
//    - Calculate new size: currentPoolSize * 0.75 (rounded down)
//    - Floor at minPoolSize
//    - Return ScaleDown decision
// 5. Otherwise return Hold
func (as *AutoScaler) EvaluateScaling() ScaleDecision {
	return ScaleDecision{Action: Hold, CurrentSize: as.currentPoolSize}
}

// ApplyDecision applies a scaling decision and updates internal state.
//
// 1. If decision.Action is Hold, do nothing
// 2. Update currentPoolSize to decision.RecommendedSize
// 3. Update lastScaleTime to now
// 4. Append to scaleHistory
func (as *AutoScaler) ApplyDecision(decision ScaleDecision) {
}

// GetScaleHistory returns the history of scaling decisions.
func (as *AutoScaler) GetScaleHistory() []ScaleDecision {
	return as.scaleHistory
}

// GetCurrentPoolSize returns the current pool size.
func (as *AutoScaler) GetCurrentPoolSize() int {
	return as.currentPoolSize
}
