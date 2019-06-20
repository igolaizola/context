package context

import (
	"context"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	ctx, cancel := WithDeadline(context.Background())
	defer cancel()
	if err := ctx.SetDeadline(time.Now().Add(100 * time.Millisecond)); err != nil {
		t.Fatal(err)
	}
	if err := ctx.Err(); err != nil {
		t.Fatal("expecter: nil, got %v", err)
	}
	select {
	case <-ctx.Done():
	case <-time.After(150 * time.Millisecond):
		t.Fatal("context has not timed out")
	}
	if err := ctx.Err(); err != context.DeadlineExceeded {
		t.Fatalf("expected: deadline exceeded, got: %v", err)
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := WithDeadline(context.Background())
	if err := ctx.SetDeadline(time.Now().Add(100 * time.Millisecond)); err != nil {
		t.Fatal(err)
	}
	cancel()
	select {
	case <-ctx.Done():
	case <-time.After(50 * time.Millisecond):
		t.Fatal("context has not been canceled")
	}
	if err := ctx.Err(); err != context.Canceled {
		t.Fatalf("expected: context canceled, got: %v", err)
	}
}
