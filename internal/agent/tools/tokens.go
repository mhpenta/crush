package tools

import (
	"cmp"
	"math"
)

// TokenEstimator approximates token usage for warnings.
type TokenEstimator struct {
	SafetyMargin       float64
	BytesPerToken      float64
	APIStructureTokens int
}

// NewTokenEstimator returns an estimator with sensible defaults.
func NewTokenEstimator() *TokenEstimator {
	return &TokenEstimator{
		SafetyMargin:       1.2,
		BytesPerToken:      4,
		APIStructureTokens: 3,
	}
}

// Estimate returns an approximate token count for the given text.
// Uses byte length which scales reasonably across languages.
func (e *TokenEstimator) Estimate(text string) int {
	if text == "" {
		return 0
	}
	tokens := float64(len(text)) / cmp.Or(e.BytesPerToken, 3) * cmp.Or(e.SafetyMargin, 1.2)
	return int(math.Ceil(tokens)) + e.APIStructureTokens
}
