package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestParseListQuery(t *testing.T) {
	request := httptest.NewRequest("GET", "/api/songs?id=16&sort=title&order=desc", nil)
	query, err := parseListQuery(request, "id", "asc", map[string]string{
		"id": "songs.id", "title": "songs.title",
	})
	if err != nil {
		t.Fatalf("parseListQuery() error = %v", err)
	}
	if query.ID == nil || *query.ID != 16 || query.SortValue != "songs.title" || query.Order != "desc" {
		t.Fatalf("parseListQuery() = %+v", query)
	}
}

func TestParseListQueryRejectsInvalidValues(t *testing.T) {
	allowedSorts := map[string]string{"id": "songs.id"}
	for _, target := range []string{"/api/songs?id=abc", "/api/songs?sort=unknown", "/api/songs?order=random"} {
		request := httptest.NewRequest("GET", target, nil)
		if _, err := parseListQuery(request, "id", "asc", allowedSorts); err == nil {
			t.Errorf("parseListQuery(%q) returned nil error", target)
		}
	}
}
