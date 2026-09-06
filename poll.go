package chromedp

import (
	"context"
	"time"

	"github.com/chromedp/cdproto/cdp"
)

type PollAction Action

type pollTask struct {
	frame     *cdp.Node
	predicate string
	polling   string
	interval  time.Duration
	timeout   time.Duration
	args      []any
	res       any
}

func (p *pollTask) Do(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func Poll(expression string, res any, opts ...PollOption) PollAction {
	_ = "STUB: not implemented"
	return *new(PollAction)
}

func PollFunction(pageFunction string, res any, opts ...PollOption) PollAction {
	_ = "STUB: not implemented"
	return *new(PollAction)
}

func poll(predicate string, res any, opts ...PollOption) PollAction {
	_ = "STUB: not implemented"
	return *new(PollAction)
}

type PollOption = func(task *pollTask)

func WithPollingInterval(interval time.Duration) PollOption {
	_ = "STUB: not implemented"
	return *new(PollOption)
}

func WithPollingMutation() PollOption { _ = "STUB: not implemented"; return *new(PollOption) }

func WithPollingTimeout(timeout time.Duration) PollOption {
	_ = "STUB: not implemented"
	return *new(PollOption)
}

func WithPollingInFrame(frame *cdp.Node) PollOption {
	_ = "STUB: not implemented"
	return *new(PollOption)
}

func WithPollingArgs(args ...any) PollOption { _ = "STUB: not implemented"; return *new(PollOption) }
