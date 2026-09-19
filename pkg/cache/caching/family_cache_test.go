package caching

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
	"github.com/hung-senbox/senbox-cache-service/pkg/model/profile"
	goredis "github.com/redis/go-redis/v9"
)

type familyRedisHook struct {
	process goredis.ProcessHook
}

func (h familyRedisHook) DialHook(next goredis.DialHook) goredis.DialHook { return next }
func (h familyRedisHook) ProcessHook(goredis.ProcessHook) goredis.ProcessHook {
	return h.process
}
func (h familyRedisHook) ProcessPipelineHook(next goredis.ProcessPipelineHook) goredis.ProcessPipelineHook {
	return next
}

func TestSetFamily(t *testing.T) {
	family := &profile.Family{
		ID:       "family-1",
		Father:   profile.FamilyMember{ID: "father-1"},
		Mother:   profile.FamilyMember{ID: "mother-1"},
		Children: []profile.FamilyMember{{ID: "child-1"}},
	}
	calls := 0
	client := goredis.NewClient(&goredis.Options{Addr: "unused:6379"})
	t.Cleanup(func() { _ = client.Close() })
	client.AddHook(familyRedisHook{process: func(ctx context.Context, cmd goredis.Cmder) error {
		calls++
		args := cmd.Args()
		if len(args) != 5 || args[0] != "set" || args[1] != keys.FamilyCacheKey(family.ID) || args[3] != "ex" || args[4] != int64(60) {
			t.Fatalf("SET args = %#v", args)
		}
		data, ok := args[2].([]byte)
		if !ok {
			t.Fatalf("cached value type = %T, want []byte", args[2])
		}
		var got profile.Family
		if err := json.Unmarshal(data, &got); err != nil {
			t.Fatalf("decode cached family: %v", err)
		}
		if got.ID != family.ID || got.Father.ID != family.Father.ID || len(got.Children) != 1 {
			t.Fatalf("cached family = %#v, want %#v", got, family)
		}
		cmd.(*goredis.StatusCmd).SetVal("OK")
		return nil
	}})

	service := NewCachingProfileService(cache.NewRedisCache(client), 60)
	if err := service.SetFamily(context.Background(), family.ID, family); err != nil {
		t.Fatalf("SetFamily() error = %v", err)
	}
	if calls != 1 {
		t.Fatalf("Redis calls = %d, want 1", calls)
	}
}

func TestSetFamilyValidationAndError(t *testing.T) {
	redisErr := errors.New("redis unavailable")
	for _, tc := range []struct {
		name      string
		familyID  string
		family    *profile.Family
		wantError bool
		wantCalls int
	}{
		{name: "empty ID", family: &profile.Family{ID: "family-1"}},
		{name: "nil family", familyID: "family-1"},
		{name: "Redis error", familyID: "family-1", family: &profile.Family{ID: "family-1"}, wantError: true, wantCalls: 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			client := goredis.NewClient(&goredis.Options{Addr: "unused:6379"})
			t.Cleanup(func() { _ = client.Close() })
			client.AddHook(familyRedisHook{process: func(ctx context.Context, cmd goredis.Cmder) error {
				calls++
				return redisErr
			}})

			service := NewCachingProfileService(cache.NewRedisCache(client), 60)
			err := service.SetFamily(context.Background(), tc.familyID, tc.family)
			if (err != nil) != tc.wantError {
				t.Fatalf("SetFamily() error = %v, wantError %v", err, tc.wantError)
			}
			if calls != tc.wantCalls {
				t.Fatalf("Redis calls = %d, want %d", calls, tc.wantCalls)
			}
		})
	}
}
