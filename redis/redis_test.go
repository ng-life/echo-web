package redis

import (
	"context"
	"github.com/go-redis/redis/v9"
	"reflect"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	tests := []struct {
		name string
		want *redis.Client
	}{
		{"client", Client()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Client()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RedisClient() = %v, want %v", got, tt.want)
			}
			got.Set(context.Background(), "lifeng", "ng-life", 1*time.Minute)
			result, _ := got.Get(context.Background(), "lifeng").Result()
			println(result)
		})
	}
}
