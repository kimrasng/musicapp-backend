package handlers

import (
	"database/sql"
	"testing"
)

func TestNullStringValue(t *testing.T) {
	value := "image.webp"

	if got := nullStringValue(sql.NullString{String: value, Valid: true}); got == nil || *got != value {
		t.Fatalf("nullStringValue() = %v, want %q", got, value)
	}
	if got := nullStringValue(sql.NullString{}); got != nil {
		t.Fatalf("nullStringValue() = %v, want nil", got)
	}
}

func TestNullIntValue(t *testing.T) {
	if got := nullIntValue(sql.NullInt64{Int64: 2, Valid: true}); got == nil || *got != 2 {
		t.Fatalf("nullIntValue() = %v, want 2", got)
	}
	if got := nullIntValue(sql.NullInt64{}); got != nil {
		t.Fatalf("nullIntValue() = %v, want nil", got)
	}
}
