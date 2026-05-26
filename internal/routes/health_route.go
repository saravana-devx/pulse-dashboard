package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"pulseDashboard/internal/rabbitmq"
	"pulseDashboard/internal/redis"
)

const (
	statusUp      = "up"
	statusDown    = "down"
	statusSkipped = "skipped"
)

type serviceStatus struct {
	Status  string `json:"status"`
	Latency string `json:"latency,omitempty"`
	Error   string `json:"error,omitempty"`
}

func RegisterHealthRoute(r *gin.Engine, db *gorm.DB, rdb *redis.Redis, mq *rabbitmq.RabbitMQ) {
	r.GET("/health", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()

		results := map[string]serviceStatus{
			"postgres":   checkPostgres(ctx, db),
			"redis":      checkRedis(ctx, rdb),
			"rabbitmq":   checkRabbitMQ(mq),
			"clickhouse": checkClickHouse(ctx),
		}

		overall := statusUp
		code := http.StatusOK
		for _, s := range results {
			if s.Status == statusDown {
				overall = statusDown
				code = http.StatusServiceUnavailable
				break
			}
		}

		c.JSON(code, gin.H{
			"status":   overall,
			"services": results,
		})
	})
}

func checkPostgres(ctx context.Context, db *gorm.DB) serviceStatus {
	start := time.Now()
	sqlDB, err := db.DB()
	if err != nil {
		return serviceStatus{Status: statusDown, Error: err.Error()}
	}
	if err := sqlDB.PingContext(ctx); err != nil {
		return serviceStatus{Status: statusDown, Error: err.Error()}
	}
	return serviceStatus{Status: statusUp, Latency: time.Since(start).String()}
}

func checkRedis(ctx context.Context, rdb *redis.Redis) serviceStatus {
	start := time.Now()
	if err := rdb.Client.Ping(ctx).Err(); err != nil {
		return serviceStatus{Status: statusDown, Error: err.Error()}
	}
	return serviceStatus{Status: statusUp, Latency: time.Since(start).String()}
}

// checkRabbitMQ probes the broker. amqp091-go's connection/channel calls aren't
// context-aware, so the route's timeout doesn't bound this directly; we keep the
// probe cheap instead. Opening a channel is a round-trip to the broker, so it
// confirms the broker actually answers — not just that our local socket is open.
func checkRabbitMQ(mq *rabbitmq.RabbitMQ) serviceStatus {
	start := time.Now()
	if mq == nil || mq.Conn == nil || mq.Conn.IsClosed() {
		return serviceStatus{Status: statusDown, Error: "connection is closed"}
	}
	ch, err := mq.Conn.Channel()
	if err != nil {
		return serviceStatus{Status: statusDown, Error: err.Error()}
	}
	_ = ch.Close()
	return serviceStatus{Status: statusUp, Latency: time.Since(start).String()}
}

// checkClickHouse is a stub. Wire it up once a ClickHouse client is added to
// the project — until then it reports "skipped" so the route still works.
func checkClickHouse(_ context.Context) serviceStatus {
	return serviceStatus{Status: statusSkipped, Error: "clickhouse client not configured"}
}
