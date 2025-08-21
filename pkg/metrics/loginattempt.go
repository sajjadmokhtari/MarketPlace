package metrics

import "github.com/prometheus/client_golang/prometheus"


var LoginAttempts = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "login_attempts_total",
		Help: "Number of login attempts, success or fail",
	},
	[]string{"status"}, // status میتونه "success" یا "fail" باشه
)
