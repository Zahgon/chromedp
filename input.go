package chromedp

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/input"
)

type MouseAction Action

func MouseEvent(typ input.MouseType, x, y float64, opts ...MouseOption) MouseAction {
	_ = "STUB: not implemented"
	return *new(MouseAction)
}

func MouseClickXY(x, y float64, opts ...MouseOption) MouseAction {
	_ = "STUB: not implemented"
	return *new(MouseAction)
}

func MouseClickNode(n *cdp.Node, opts ...MouseOption) MouseAction {
	_ = "STUB: not implemented"
	return *new(MouseAction)
}

type MouseOption = func(*input.DispatchMouseEventParams) *input.DispatchMouseEventParams

func Button(btn string) MouseOption { _ = "STUB: not implemented"; return *new(MouseOption) }

func ButtonType(button input.MouseButton) MouseOption {
	_ = "STUB: not implemented"
	return *new(MouseOption)
}

func ButtonLeft(p *input.DispatchMouseEventParams) *input.DispatchMouseEventParams {
	_ = "STUB: not implemented"
	return nil
}

func ButtonMiddle(p *input.DispatchMouseEventParams) *input.DispatchMouseEventParams {
	_ = "STUB: not implemented"
	return nil
}

func ButtonRight(p *input.DispatchMouseEventParams) *input.DispatchMouseEventParams {
	_ = "STUB: not implemented"
	return nil
}

func ButtonNone(p *input.DispatchMouseEventParams) *input.DispatchMouseEventParams {
	_ = "STUB: not implemented"
	return nil
}

func ButtonModifiers(modifiers ...input.Modifier) MouseOption {
	_ = "STUB: not implemented"
	return *new(MouseOption)
}

func ClickCount(n int) MouseOption { _ = "STUB: not implemented"; return *new(MouseOption) }

type KeyAction Action

func KeyEvent(keys string, opts ...KeyOption) KeyAction {
	_ = "STUB: not implemented"
	return *new(KeyAction)
}

func KeyEventNode(n *cdp.Node, keys string, opts ...KeyOption) KeyAction {
	_ = "STUB: not implemented"
	return *new(KeyAction)
}

type KeyOption = func(*input.DispatchKeyEventParams) *input.DispatchKeyEventParams

func KeyModifiers(modifiers ...input.Modifier) KeyOption {
	_ = "STUB: not implemented"
	return *new(KeyOption)
}
