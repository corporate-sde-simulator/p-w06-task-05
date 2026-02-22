# PR Review - Connection pool auto-scaler (by Suresh)

## Reviewer: Kavitha Rajan
---

**Overall:** Good foundation but critical bugs need fixing before merge.

### `autoScaler.go`

> **Bug #1:** Scale-up triggers when utilization is LOW instead of HIGH and adds connections when idle
> This is the higher priority fix. Check the logic carefully and compare against the design doc.

### `metricsMonitor.go`

> **Bug #2:** Cooldown timer resets on every metrics check instead of on last scaling action
> This is more subtle but will cause issues in production. Make sure to add a test case for this.

---

**Suresh**
> Acknowledged. I have documented the issues for whoever picks this up.
