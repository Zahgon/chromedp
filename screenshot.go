package chromedp

import (
	"github.com/chromedp/cdproto/cdp"
)

func Screenshot(sel any, picbuf *[]byte, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func ScreenshotScale(sel any, scale float64, picbuf *[]byte, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func ScreenshotNodes(nodes []*cdp.Node, scale float64, picbuf *[]byte) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

func CaptureScreenshot(res *[]byte) Action { _ = "STUB: not implemented"; return *new(Action) }

func FullScreenshot(res *[]byte, quality int) EmulateAction {
	_ = "STUB: not implemented"
	return *new(EmulateAction)
}

func extents(m, n, o, p float64) (float64, float64) { _ = "STUB: not implemented"; return 0, 0 }
