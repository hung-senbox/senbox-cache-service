package keys

import "testing"

func TestFamilyCacheKey(t *testing.T) {
	const familyID = "68cab123"
	const want = "profile-service:family:" + familyID

	if got := FamilyCacheKey(familyID); got != want {
		t.Fatalf("FamilyCacheKey(%q) = %q, want %q", familyID, got, want)
	}
}
