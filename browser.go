package chromedp

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/chromedp/cdproto"
	"github.com/chromedp/cdproto/target"
	jsonv2 "github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

var (
	DefaultUnmarshalOptions = jsonv2.JoinOptions(
		jsonv2.DefaultOptionsV2(),
		jsontext.AllowInvalidUTF8(true),
	)

	DefaultMarshalOptions = jsonv2.JoinOptions(
		jsonv2.DefaultOptionsV2(),
		jsontext.AllowInvalidUTF8(true),
	)
)

type Browser struct {
	next int64

	LostConnection chan struct{}

	closingGracefully chan struct{}

	dialTimeout time.Duration

	pages map[target.SessionID]*Target

	listenersMu sync.Mutex
	listeners   []cancelableListener

	conn Transport

	newTabQueue chan *Target

	cmdQueue chan *cdproto.Message

	logf func(string, ...any)
	errf func(string, ...any)
	dbgf func(string, ...any)

	process *os.Process

	userDataDir string
}

func NewBrowser(ctx context.Context, urlstr string, opts ...BrowserOption) (*Browser, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Browser) Process() *os.Process { _ = "STUB: not implemented"; return nil }

func (b *Browser) newExecutorForTarget(ctx context.Context, targetID target.ID, sessionID target.SessionID) (*Target, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Browser) Execute(ctx context.Context, method string, params, res any) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Browser) execute(ctx context.Context, method string, params, res any) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Browser) run(ctx context.Context) { _ = "STUB: not implemented"; return }

type BrowserOption = func(*Browser)

func WithBrowserLogf(f func(string, ...any)) BrowserOption {
	_ = "STUB: not implemented"
	return *new(BrowserOption)
}

func WithBrowserErrorf(f func(string, ...any)) BrowserOption {
	_ = "STUB: not implemented"
	return *new(BrowserOption)
}

func WithBrowserDebugf(f func(string, ...any)) BrowserOption {
	_ = "STUB: not implemented"
	return *new(BrowserOption)
}

func WithConsolef(f func(string, ...any)) BrowserOption {
	_ = "STUB: not implemented"
	return *new(BrowserOption)
}

func WithDialTimeout(d time.Duration) BrowserOption {
	_ = "STUB: not implemented"
	return *new(BrowserOption)
}
