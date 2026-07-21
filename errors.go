package chromedp

type Error string

func (err Error) Error() string { _ = "STUB: not implemented"; return "" }

const (
	ErrInvalidWebsocketMessage Error = "invalid websocket message"

	ErrInvalidDimensions Error = "invalid dimensions"

	ErrNoResults Error = "no results"

	ErrHasResults Error = "has results"

	ErrNotVisible Error = "not visible"

	ErrVisible Error = "visible"

	ErrDisabled Error = "disabled"

	ErrNotSelected Error = "not selected"

	ErrInvalidBoxModel Error = "invalid box model"

	ErrChannelClosed Error = "channel closed"

	ErrInvalidTarget Error = "invalid target"

	ErrInvalidContext Error = "invalid context"

	ErrPollingTimeout Error = "waiting for function failed: timeout"

	ErrJSUndefined Error = "encountered an undefined value"

	ErrJSNull Error = "encountered a null value"
)
