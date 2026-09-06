package chromedp

import (
	"context"
	"sync"

	"github.com/chromedp/cdproto"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/cdproto/target"
)

type Target struct {
	browser   *Browser
	SessionID target.SessionID
	TargetID  target.ID

	listenersMu sync.Mutex
	listeners   []cancelableListener

	messageQueue chan *cdproto.Message

	frameMu sync.RWMutex

	frames       map[cdp.FrameID]*cdp.Frame
	execContexts map[cdp.FrameID]runtime.ExecutionContextID

	cur cdp.FrameID

	logf, errf func(string, ...any)

	isWorker bool
}

func (t *Target) enclosingFrame(node *cdp.Node) cdp.FrameID {
	_ = "STUB: not implemented"
	return *new(cdp.FrameID)
}

func (t *Target) ensureFrame() (*cdp.Frame, *cdp.Node, runtime.ExecutionContextID, bool) {
	_ = "STUB: not implemented"
	return nil, nil, *new(runtime.ExecutionContextID), false
}

func (t *Target) run(ctx context.Context) { _ = "STUB: not implemented"; return }

func (t *Target) Execute(ctx context.Context, method string, params, res any) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Target) runtimeEvent(ev any) { _ = "STUB: not implemented"; return }

func (t *Target) documentUpdated(ctx context.Context) { _ = "STUB: not implemented"; return }

func (t *Target) pageEvent(ev any) { _ = "STUB: not implemented"; return }

func (t *Target) domEvent(ctx context.Context, ev any) { _ = "STUB: not implemented"; return }
