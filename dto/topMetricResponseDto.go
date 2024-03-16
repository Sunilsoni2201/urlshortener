package dto

type TopMetricResponse struct {
	TopMetrics map[string]int64 `json:"top_metrics"`
}

func NewTopMetricRespons(metrics map[string]int64) *TopMetricResponse {
	return &TopMetricResponse{
		TopMetrics: metrics,
	}
}
