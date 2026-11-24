package api

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var (
	storeMu  sync.RWMutex
	releases = map[string]Release{}
)

func init() {
	seed := []Release{{ID: "rel-payments-240", Service: "payments-api", Version: "v2.4.0", Environment: "production", ErrorBudgetRemainingPct: 8.4, P95LatencyMs: 410, P95BudgetMs: 250, TestPassRatePct: 96.2, HasRollbackPlan: false, EnvValidated: true, TLSDaysRemaining: 19}, {ID: "rel-checkout-1182", Service: "checkout-web", Version: "v1.18.2", Environment: "production", ErrorBudgetRemainingPct: 78.6, P95LatencyMs: 182, P95BudgetMs: 300, TestPassRatePct: 99.4, HasRollbackPlan: true, EnvValidated: true, TLSDaysRemaining: 142, CanaryPct: 5}, {ID: "rel-ledger-310", Service: "ledger-worker", Version: "v3.1.0", Environment: "staging", ErrorBudgetRemainingPct: 42, P95LatencyMs: 285, P95BudgetMs: 250, TestPassRatePct: 97.1, HasRollbackPlan: true, EnvValidated: false, TLSDaysRemaining: 31}}
	for _, v := range seed {
		v.UpdatedAt = time.Now()
		releases[v.ID] = v
	}
}
func listReleases() []Release {
	storeMu.RLock()
	defer storeMu.RUnlock()
	out := make([]Release, 0, len(releases))
	for _, v := range releases {
		out = append(out, v)
	}
	return out
}
func saveRelease(v Release) (Release, error) {
	if strings.TrimSpace(v.Service) == "" || strings.TrimSpace(v.Version) == "" {
		return v, fmt.Errorf("service and version are required")
	}
	if v.ErrorBudgetRemainingPct < 0 || v.ErrorBudgetRemainingPct > 100 || v.TestPassRatePct < 0 || v.TestPassRatePct > 100 || v.P95LatencyMs < 0 || v.P95BudgetMs <= 0 || v.TLSDaysRemaining < 0 {
		return v, fmt.Errorf("metrics are outside accepted ranges")
	}
	if v.ID == "" {
		v.ID = fmt.Sprintf("rel-%d", time.Now().UnixNano())
	}
	v.UpdatedAt = time.Now()
	storeMu.Lock()
	releases[v.ID] = v
	storeMu.Unlock()
	return v, nil
}
