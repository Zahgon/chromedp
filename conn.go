package chromedp

import (
	"context"
	"io"
	"net"

	"github.com/chromedp/cdproto"
	"github.com/go-json-experiment/json/jsontext"
	"github.com/gobwas/ws/wsutil"
)

type Transport interface {
	Read(context.Context, *cdproto.Message) error
	Write(context.Context, *cdproto.Message) error
	io.Closer
}

type Conn struct {
	conn net.Conn

	reader wsutil.Reader
	writer wsutil.Writer

	decoder jsontext.Decoder
	encoder jsontext.Encoder

	debugf func(string, ...any)
}

func DialContext(ctx context.Context, urlstr string, opts ...DialOption) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *Conn) Read(_ context.Context, msg *cdproto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) Write(_ context.Context, msg *cdproto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

type DialOption = func(*Conn)

func WithConnDebugf(f func(string, ...any)) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}
