package promclient

import "time"
import "github.com/gin-gonic/gin"
import "github.com/prometheus/client_golang/prometheus"
import ginprometheus "github.com/zsais/go-gin-prometheus"

var numOfErrors = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "numOfErrors",
		Help: "count error by type",
	},
	[]string{"type"},
)

var summaryByType = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "summaryByType",
		Help: "Summury by type",
	},
	[]string{"type", "measure_unit"},
)

var dbRequests = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "dbResponses",
		Help: "Database requests time in seconds",
	},
	[]string{"dbReqName"},
)

func IncError(errorType string) {
	numOfErrors.WithLabelValues(errorType).Inc()
}

func SummaryByType(event string, num float64, mu string) {
	summaryByType.WithLabelValues(event, mu).Observe(num)
}

func EndDbRequestTimer(dbReqName string, startTime time.Time) {
	dbRequests.WithLabelValues(dbReqName).Observe(float64(time.Since(startTime).Seconds()))
}

func UseGin(gin *gin.Engine) {
	p := ginprometheus.NewPrometheus("gin")
	p.Use(gin)
	prometheus.Register(numOfErrors)
	prometheus.Register(summaryByType)
	prometheus.Register(dbRequests)
}
