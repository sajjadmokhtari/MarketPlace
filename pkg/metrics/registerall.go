package metrics

import "github.com/prometheus/client_golang/prometheus"

func RegisterAll() {
    prometheus.MustRegister(LoginAttempts)
    prometheus.MustRegister(HttpDuration)
    prometheus.MustRegister(DbCall)
}