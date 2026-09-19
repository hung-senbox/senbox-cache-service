package cached

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
	"github.com/hung-senbox/senbox-cache-service/pkg/model/profile"
	goredis "github.com/redis/go-redis/v9"
)

// Intercept commands before network I/O to exercise the real cache wrapper and
// JSON decoding with deterministic Redis pages, misses, and failures.
type profileRedisHook struct {
	process goredis.ProcessHook
}

func (h profileRedisHook) DialHook(next goredis.DialHook) goredis.DialHook { return next }
func (h profileRedisHook) ProcessHook(goredis.ProcessHook) goredis.ProcessHook {
	return h.process
}
func (h profileRedisHook) ProcessPipelineHook(next goredis.ProcessPipelineHook) goredis.ProcessPipelineHook {
	return next
}

func TestGetStudentInformationsByOrgId(t *testing.T) {
	redisErr := errors.New("redis unavailable")
	for _, scenario := range []string{"matches", "empty input", "no matches", "scan error", "get error", "malformed JSON", "cancelled"} {
		t.Run(scenario, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			orgID := "org-1"
			if scenario == "empty input" {
				orgID = ""
			}
			if scenario == "no matches" {
				orgID = "org-absent"
			}
			if scenario == "cancelled" {
				cancel()
			}
			a, b := keys.StudentInformationCacheKey("a"), keys.StudentInformationCacheKey("b")
			other := keys.StudentInformationCacheKey("other")
			missing, null := keys.StudentInformationCacheKey("missing"), keys.StudentInformationCacheKey("null")
			values := map[string]string{
				a:     `{"id":"info-a","student_id":"a","organization_id":"org-1","gender":2}`,
				b:     `{"student_id":"b","organization_id":"org-1"}`,
				other: `{"student_id":"other","organization_id":"org-2"}`,
				null:  "null",
			}
			if scenario == "malformed JSON" {
				values[a] = "{"
			}
			scanCalls := 0
			gets := make(map[string]int)
			client := goredis.NewClient(&goredis.Options{Addr: "unused:6379"})
			t.Cleanup(func() { _ = client.Close() })
			client.AddHook(profileRedisHook{process: func(ctx context.Context, cmd goredis.Cmder) error {
				if scenario == "empty input" || scenario == "cancelled" {
					t.Fatal("unexpected Redis command")
				}
				switch cmd := cmd.(type) {
				case *goredis.ScanCmd:
					cursors := []uint64{0, 7, 9}
					if scanCalls >= len(cursors) {
						t.Fatal("scan did not stop at cursor zero")
					}
					wantArgs := []interface{}{"scan", cursors[scanCalls], "match", keys.StudentInformationCacheKey("*"), "count", int64(100)}
					if !reflect.DeepEqual(cmd.Args(), wantArgs) {
						t.Fatalf("scan args = %v, want %v", cmd.Args(), wantArgs)
					}
					scanCalls++
					if scenario == "scan error" && scanCalls == 2 {
						return redisErr
					}
					switch scanCalls {
					case 1:
						cmd.SetVal([]string{b, other}, 7)
					case 2:
						cmd.SetVal(nil, 9)
					case 3:
						cmd.SetVal([]string{b, a, missing, null}, 0)
					}
					return nil
				case *goredis.StringCmd:
					key := cmd.Args()[1].(string)
					gets[key]++
					if scenario == "get error" && key == a {
						return redisErr
					}
					value, ok := values[key]
					if !ok {
						return goredis.Nil
					}
					cmd.SetVal(value)
					return nil
				default:
					t.Fatalf("unexpected command: %v", cmd.Args())
					return nil
				}
			}})
			gateway := NewCachedProfileGateway(cache.NewRedisCache(client))
			got, err := gateway.GetStudentInformationsByOrgId(ctx, orgID)
			switch scenario {
			case "scan error", "get error", "malformed JSON", "cancelled":
				if err == nil || got != nil {
					t.Fatalf("got %v, %v; want nil result and error", got, err)
				}
				if (scenario == "scan error" || scenario == "get error") && !errors.Is(err, redisErr) {
					t.Fatalf("error = %v, want %v", err, redisErr)
				}
				if scenario == "cancelled" && !errors.Is(err, context.Canceled) {
					t.Fatalf("error = %v, want context.Canceled", err)
				}
			default:
				var want []profile.StudentInformation
				if scenario == "matches" {
					want = []profile.StudentInformation{
						{ID: "info-a", StudentId: "a", OrganizationId: "org-1", Gender: 2},
						{StudentId: "b", OrganizationId: "org-1"},
					}
				}
				if err != nil || !reflect.DeepEqual(got, want) {
					t.Fatalf("got %v, %v; want %v, nil", got, err, want)
				}
				if scenario != "empty input" && (scanCalls != 3 || gets[b] != 1 || gets[missing] != 1 || gets[null] != 1) {
					t.Fatalf("unexpected command counts: scans=%d gets=%v", scanCalls, gets)
				}
			}
		})
	}
}
