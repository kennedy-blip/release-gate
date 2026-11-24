package api
import "time"
type Release struct { ID string `json:"id"`; Service string `json:"service"`; Version string `json:"version"`; Environment string `json:"environment"`; ErrorBudgetRemainingPct float64 `json:"errorBudgetRemainingPct"`; P95LatencyMs float64 `json:"p95LatencyMs"`; P95BudgetMs float64 `json:"p95BudgetMs"`; TestPassRatePct float64 `json:"testPassRatePct"`; HasRollbackPlan bool `json:"hasRollbackPlan"`; EnvValidated bool `json:"envValidated"`; TLSDaysRemaining int `json:"tlsDaysRemaining"`; CanaryPct float64 `json:"canaryPct"`; UpdatedAt time.Time `json:"updatedAt"` }
type EvaluationRequest struct { Release; MinErrorBudgetPct float64 `json:"minErrorBudgetPct"`; MinTestPassRatePct float64 `json:"minTestPassRatePct"` }
type GateCheck struct { ID string `json:"id"`; Name string `json:"name"`; Status string `json:"status"`; Impact int `json:"impact"`; Detail string `json:"detail"` }
type ActionItem struct { Severity string `json:"severity"`; Message string `json:"message"` }
type RolloutStep struct { Percent int `json:"percent"`; Soak string `json:"soak"`; Requirement string `json:"requirement"` }
type EvaluationResponse struct { Service string `json:"service"`; Version string `json:"version"`; Score int `json:"score"`; Verdict string `json:"verdict"`; RecommendedCanaryPct int `json:"recommendedCanaryPct"`; Checks []GateCheck `json:"checks"`; Actions []ActionItem `json:"actions"`; Rollout []RolloutStep `json:"rollout"`; EvaluatedAt time.Time `json:"evaluatedAt"` }
