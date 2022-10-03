package redis

import (
	"github.com/go-redis/redis/v9"
	"reflect"
	"testing"
)

func TestRedisClient(t *testing.T) {
	tests := []struct {
		name string
		want *redis.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RedisClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RedisClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
