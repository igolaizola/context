package context

import (
	"context"
	"sync"
	"time"
)

// Context is a context.Context implementation with an updatable deadline
type Context interface {
	context.Context
	SetDeadline(t time.Time) error
}

type deadlineCtx struct {
	base    context.Context
	timeout context.Context
	cancel  context.CancelFunc
	done    chan struct{}
	mu      sync.Mutex
}

// WithDeadline returns a new context with an updatable deadline
func WithDeadline(ctx context.Context) (Context, context.CancelFunc) {
	done := make(chan struct{})
	timeout, cancel := context.WithCancel(ctx)
	d := &deadlineCtx{
		base:    ctx,
		timeout: timeout,
		cancel:  cancel,
		done:    done,
		mu:      sync.Mutex{},
	}
	go func() {
		defer close(d.done)
		for {
			d.mu.Lock()
			t := d.timeout
			d.mu.Unlock()
			<-t.Done()
			if err := d.Err(); err != nil {
				return
			}
		}
	}()
	c := func() {
		d.mu.Lock()
		d.cancel()
		d.mu.Unlock()
	}
	return d, c
}

// SetDeadline sets a new deadline value
func (d *deadlineCtx) SetDeadline(t time.Time) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if err := d.timeout.Err(); err != nil {
		return err
	}
	d.cancel()
	d.timeout, d.cancel = context.WithDeadline(d.base, t)
	return nil
}

// Done implements context.Context.Done
func (d *deadlineCtx) Done() <-chan struct{} {
	return d.done
}

// Deadline implements context.Context.Deadline
func (d *deadlineCtx) Deadline() (deadline time.Time, ok bool) {
	d.mu.Lock()
	t, ok := d.timeout.Deadline()
	d.mu.Unlock()
	return t, ok
}

// Err implements context.Context.Err
func (d *deadlineCtx) Err() error {
	d.mu.Lock()
	err := d.timeout.Err()
	d.mu.Unlock()
	return err
}

// Value implements context.Context.Value
func (d *deadlineCtx) Value(key interface{}) interface{} {
	return d.base.Value(key)
}
