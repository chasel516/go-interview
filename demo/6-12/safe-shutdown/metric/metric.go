package metric

import (
	"gitee.com/phper95/pkg/prome"
	"github.com/prometheus/client_golang/prometheus"
)

const AppName = "safe-shutdown"

var TestCostTime = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:        "test_cost_time",
		Help:        "histogram for apui cost time",
		Buckets:     prome.DefaultBuckets,
		ConstLabels: prometheus.Labels{"machine": prome.GetHostName(), "app": AppName},
	},
	[]string{"tag"},
)
