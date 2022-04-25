package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	promHeadersReceived = promauto.NewCounter(prometheus.CounterOpts{
		Name: "app_headers_received",
		Help: "The total number of received heads",
	})
	promMissingBlockError = promauto.NewCounter(prometheus.CounterOpts{
		Name: "app_missing_block_error",
		Help: "The total number of missing blocks",
	})
	promFilterLogsError = promauto.NewCounter(prometheus.CounterOpts{
		Name: "app_filter_logs_error",
		Help: "The total number of filter logs errors",
	})
	promNewTokenFetched = promauto.NewCounter(prometheus.CounterOpts{
		Name: "app_new_token_fetched",
		Help: "The total number of new fetched tokens",
	})
)
