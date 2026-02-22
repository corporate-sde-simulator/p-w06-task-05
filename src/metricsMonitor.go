package pool

// MetricsMonitor — tracks connection pool utilization metrics.
// This module is COMPLETE and WORKING. Your task is in autoScaler.go.

import (
	"sync"
	"time"
)

type PoolMetrics struct {
	TotalConnections    int
	ActiveConnections   int
	IdleConnections     int
	WaitingRequests     int
	AvgResponseTimeMs   float64
	ErrorRate           float64
	Timestamp           time.Time
}

type MetricsMonitor struct {
	mu       sync.RWMutex
	history  []PoolMetrics
	maxHistory int
}

func NewMetricsMonitor(maxHistory int) *MetricsMonitor {
	return &MetricsMonitor{
		history:    make([]PoolMetrics, 0),
		maxHistory: maxHistory,
	}
}

func (m *MetricsMonitor) RecordMetrics(metrics PoolMetrics) {
	m.mu.Lock()
	defer m.mu.Unlock()
	metrics.Timestamp = time.Now()
	m.history = append(m.history, metrics)
	if len(m.history) > m.maxHistory {
		m.history = m.history[1:]
	}
}

func (m *MetricsMonitor) GetLatest() *PoolMetrics {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if len(m.history) == 0 {
		return nil
	}
	latest := m.history[len(m.history)-1]
	return &latest
}

func (m *MetricsMonitor) GetUtilization() float64 {
	latest := m.GetLatest()
	if latest == nil || latest.TotalConnections == 0 {
		return 0
	}
	return float64(latest.ActiveConnections) / float64(latest.TotalConnections) * 100
}

func (m *MetricsMonitor) GetAvgUtilization(window time.Duration) float64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	cutoff := time.Now().Add(-window)
	sum := 0.0
	count := 0
	for _, pm := range m.history {
		if pm.Timestamp.After(cutoff) && pm.TotalConnections > 0 {
			util := float64(pm.ActiveConnections) / float64(pm.TotalConnections) * 100
			sum += util
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return sum / float64(count)
}

func (m *MetricsMonitor) GetHistory(window time.Duration) []PoolMetrics {
	m.mu.RLock()
	defer m.mu.RUnlock()
	cutoff := time.Now().Add(-window)
	var result []PoolMetrics
	for _, pm := range m.history {
		if pm.Timestamp.After(cutoff) {
			result = append(result, pm)
		}
	}
	return result
}
