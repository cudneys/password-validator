package monitoring

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"strconv"
)

type Metric struct {
	MetricCollector prometheus.Collector
	ID              string
	Name            string
	Description     string
	Type            string
	Args            []string
}

var reqCnt = &Metric{
	ID:          "reqCnt",
	Name:        "requests_total",
	Description: "the number of HTTP requests processed",
	Type:        "counter_vec",
	Args:        []string{"status"}}

type Prometheus struct {
	reqCnt        *prometheus.CounterVec
	router        *gin.Engine
	listenAddress string

	Metric      *Metric
	MetricsPath string
}

func NewPrometheus(subsystem string) *Prometheus {
	p := &Prometheus{
		Metric:        reqCnt,
		MetricsPath:   "/metrics",
		listenAddress: ":9901",
	}

	p.registerMetrics(subsystem)
	p.router = gin.Default()

	return p
}

func (p *Prometheus) registerMetrics(subsystem string) {
	metric := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      reqCnt.Name,
			Help:      reqCnt.Description,
		},
		reqCnt.Args,
	)
	if err := prometheus.Register(metric); err != nil {
		log.Print(fmt.Sprintf("%s could not be registered: ", reqCnt, err))
	} else {
		log.Print(fmt.Sprintf("%s registered.", reqCnt))
	}
	p.reqCnt = metric

	reqCnt.MetricCollector = metric
}

func (p *Prometheus) Use(e *gin.Engine) {
	e.Use(p.handlerFunc())
	p.setMetricsPath(e)
}

func (p *Prometheus) handlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.String() == p.MetricsPath {
			c.Next()
			return
		}
		c.Next()

		status := strconv.Itoa(c.Writer.Status())
		p.reqCnt.WithLabelValues(status).Inc()
	}
}

func (p *Prometheus) setMetricsPath(e *gin.Engine) {
	p.router.GET(p.MetricsPath, prometheusHandler())
	go p.router.Run(p.listenAddress)
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
