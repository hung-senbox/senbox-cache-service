package cached

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
	"github.com/hung-senbox/senbox-cache-service/pkg/model/profile"
	goredis "github.com/redis/go-redis/v9"
)

type familyCachedRedisHook struct {
	process goredis.ProcessHook
}

func (h familyCachedRedisHook) DialHook(next goredis.DialHook) goredis.DialHook { return next }
func (h familyCachedRedisHook) ProcessHook(goredis.ProcessHook) goredis.ProcessHook {
	return h.process
}
func (h familyCachedRedisHook) ProcessPipelineHook(next goredis.ProcessPipelineHook) goredis.ProcessPipelineHook {
	return next
}

func TestGetFamily(t *testing.T) {
	now := time.Date(2026, time.September, 19, 10, 0, 0, 0, time.UTC)
	want := &profile.Family{
		ID:        "family-1",
		Father:    profile.FamilyMember{ID: "father-1"},
		Mother:    profile.FamilyMember{ID: "mother-1"},
		Children:  []profile.FamilyMember{{ID: "child-1"}},
		CreatedAt: now,
		UpdatedAt: now,
	}

	client := goredis.NewClient(&goredis.Options{Addr: "unused:6379"})
	t.Cleanup(func() { _ = client.Close() })
	client.AddHook(familyCachedRedisHook{process: func(ctx context.Context, cmd goredis.Cmder) error {
		stringCmd, ok := cmd.(*goredis.StringCmd)
		if !ok {
			t.Fatalf("unexpected Redis command: %T", cmd)
		}
		if got := stringCmd.Args()[1]; got != keys.FamilyCacheKey(want.ID) {
			t.Fatalf("cache key = %v, want %q", got, keys.FamilyCacheKey(want.ID))
		}
		stringCmd.SetVal(`{"id":"family-1","father":{"id":"father-1"},"mother":{"id":"mother-1"},"children":[{"id":"child-1"}],"created_at":"2026-09-19T10:00:00Z","updated_at":"2026-09-19T10:00:00Z"}`)
		return nil
	}})

	gateway := NewCachedProfileGateway(cache.NewRedisCache(client))
	got, err := gateway.GetFamily(context.Background(), want.ID)
	if err != nil {
		t.Fatalf("GetFamily() error = %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GetFamily() = %#v, want %#v", got, want)
	}
}

func TestGetFamilyMissAndValidation(t *testing.T) {
	redisErr := errors.New("redis unavailable")
	for _, tc := range []struct {
		name      string
		familyID  string
		value     string
		cacheErr  error
		wantError bool
		wantCalls int
	}{
		{name: "empty ID"},
		{name: "miss", familyID: "missing", cacheErr: goredis.Nil, wantCalls: 1},
		{name: "null", familyID: "null", value: "null", wantCalls: 1},
		{name: "malformed JSON", familyID: "bad", value: "{", wantError: true, wantCalls: 1},
		{name: "Redis error", familyID: "error", cacheErr: redisErr, wantError: true, wantCalls: 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			client := goredis.NewClient(&goredis.Options{Addr: "unused:6379"})
			t.Cleanup(func() { _ = client.Close() })
			client.AddHook(familyCachedRedisHook{process: func(ctx context.Context, cmd goredis.Cmder) error {
				calls++
				if tc.cacheErr != nil {
					return tc.cacheErr
				}
				cmd.(*goredis.StringCmd).SetVal(tc.value)
				return nil
			}})

			gateway := NewCachedProfileGateway(cache.NewRedisCache(client))
			got, err := gateway.GetFamily(context.Background(), tc.familyID)
			if (err != nil) != tc.wantError {
				t.Fatalf("GetFamily() error = %v, wantError %v", err, tc.wantError)
			}
			if got != nil {
				t.Fatalf("GetFamily() = %#v, want nil", got)
			}
			if calls != tc.wantCalls {
				t.Fatalf("Redis calls = %d, want %d", calls, tc.wantCalls)
			}
		})
	}
}
