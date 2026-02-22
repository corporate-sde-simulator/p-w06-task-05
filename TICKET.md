# PLATFORM-2941: Build connection pool auto-scaler

**Status:** In Progress · **Priority:** High
**Sprint:** Sprint 28 · **Story Points:** 5
**Reporter:** Vikram Patel (Infra Lead) · **Assignee:** You (Intern)
**Due:** End of sprint (Friday)
**Labels:** `backend`, `golang`, `database`, `performance`
**Task Type:** Feature Ship

---

## Description

The `MetricsMonitor` tracks connection pool utilization. Build the `AutoScaler` that adjusts pool size dynamically based on load — scaling up when utilization is high and down when idle. Implement the TODOs in `autoScaler.go`.

## Acceptance Criteria

- [ ] `EvaluateScaling()` decides whether to scale up, down, or hold
- [ ] Scale-up triggers at >80% utilization, scale-down at <30%
- [ ] Cooldown period prevents thrashing (no changes within 60s of last change)
- [ ] Min/max pool size bounds are enforced
- [ ] All unit tests pass
