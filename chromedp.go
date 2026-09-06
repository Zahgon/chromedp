package chromedp

import (
	"context"
	"sync"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/target"
)

type Context struct {
	Allocator Allocator

	Browser *Browser

	Target *Target

	targetID target.ID

	createBrowserContextParams *target.CreateBrowserContextParams

	browserContextOwner bool

	BrowserContextID cdp.BrowserContextID

	browserListeners []cancelableListener
	targetListeners  []cancelableListener

	browserOpts []BrowserOption

	cancel func()

	first bool

	closedTarget sync.WaitGroup

	allocated chan struct{}

	cancelErr error
}

func NewContext(parent context.Context, opts ...ContextOption) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

type contextKey struct{}

func FromContext(ctx context.Context) *Context { _ = "STUB: not implemented"; return nil }

func Cancel(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func initContextBrowser(ctx context.Context) (*Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Run(ctx context.Context, actions ...Action) error { _ = "STUB: not implemented"; return nil }

func (c *Context) newTarget(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *Context) attachTarget(ctx context.Context, targetID target.ID) error {
	_ = "STUB: not implemented"
	return nil
}

type ContextOption = func(*Context)

func WithTargetID(id target.ID) ContextOption {
	_ = "STUB: not implemented"
	return *new(ContextOption)
}

type CreateBrowserContextOption = func(*target.CreateBrowserContextParams) *target.CreateBrowserContextParams

func WithNewBrowserContext(options ...CreateBrowserContextOption) ContextOption {
	_ = "STUB: not implemented"
	return *new(ContextOption)
}

func WithExistingBrowserContext(id cdp.BrowserContextID) ContextOption {
	_ = "STUB: not implemented"
	return *new(ContextOption)
}

func WithLogf(f func(string, ...any)) ContextOption {
	_ = "STUB: not implemented"
	return *new(ContextOption)
}

func WithErrorf(f func(string, ...any)) ContextOption {
	_ = "STUB: not implemented"
	return *new(ContextOption)
}

func WithDebugf(f func(string, ...any)) ContextOption {
	_ = "STUB: not implemented"
	return *new(ContextOption)
}

func WithBrowserOption(opts ...BrowserOption) ContextOption {
	_ = "STUB: not implemented"
	return *new(ContextOption)
}

func RunResponse(ctx context.Context, actions ...Action) (*network.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func responseAction(resp **network.Response, actions ...Action) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

func Targets(ctx context.Context) ([]*target.Info, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Action interface {
	Do(context.Context) error
}

type ActionFunc func(context.Context) error

func (f ActionFunc) Do(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

type Tasks []Action

func (t Tasks) Do(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func Sleep(d time.Duration) Action { _ = "STUB: not implemented"; return *new(Action) }

func sleepContext(ctx context.Context, d time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func retryWithSleep(ctx context.Context, d time.Duration, f func(ctx context.Context) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

type cancelableListener struct {
	ctx context.Context
	fn  func(ev any)
}

func ListenBrowser(ctx context.Context, fn func(ev any)) { _ = "STUB: not implemented"; return }

func ListenTarget(ctx context.Context, fn func(ev any)) { _ = "STUB: not implemented"; return }

func WaitNewTarget(ctx context.Context, fn func(*target.Info) bool) <-chan target.ID {
	_ = "STUB: not implemented"
	return nil
}
