package chromedp

import (
	"context"
	"io"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/css"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/runtime"
)

type QueryAction Action

type Selector struct {
	sel           any
	fromNode      *cdp.Node
	retryInterval time.Duration
	exp           int
	by            func(context.Context, *cdp.Node) ([]cdp.NodeID, error)
	wait          func(context.Context, *cdp.Frame, runtime.ExecutionContextID, ...cdp.NodeID) ([]*cdp.Node, error)
	after         []func(context.Context, runtime.ExecutionContextID, ...*cdp.Node) error
}

func Query(sel any, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func (s *Selector) Do(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Selector) selAsString() string { _ = "STUB: not implemented"; return "" }

func (s *Selector) waitReady(check func(context.Context, runtime.ExecutionContextID, *cdp.Node) error) func(context.Context, *cdp.Frame, runtime.ExecutionContextID, ...cdp.NodeID) ([]*cdp.Node, error) {
	_ = "STUB: not implemented"
	return nil
}

func QueryAfter(sel any, f func(context.Context, runtime.ExecutionContextID, ...*cdp.Node) error, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

type QueryOption = func(*Selector)

func FromNode(node *cdp.Node) QueryOption { _ = "STUB: not implemented"; return *new(QueryOption) }

func ByFunc(f func(context.Context, *cdp.Node) ([]cdp.NodeID, error)) QueryOption {
	_ = "STUB: not implemented"
	return *new(QueryOption)
}

func ByQuery(s *Selector) { _ = "STUB: not implemented"; return }

func ByQueryAll(s *Selector) { _ = "STUB: not implemented"; return }

func ByID(s *Selector) { _ = "STUB: not implemented"; return }

func BySearch(s *Selector) { _ = "STUB: not implemented"; return }

func ByJSPath(s *Selector) { _ = "STUB: not implemented"; return }

func ByNodeID(s *Selector) { _ = "STUB: not implemented"; return }

func WaitFunc(wait func(context.Context, *cdp.Frame, runtime.ExecutionContextID, ...cdp.NodeID) ([]*cdp.Node, error)) QueryOption {
	_ = "STUB: not implemented"
	return *new(QueryOption)
}

func NodeReady(s *Selector) { _ = "STUB: not implemented"; return }

func callFunctionOnNode(ctx context.Context, node *cdp.Node, function string, res any, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func NodeVisible(s *Selector) { _ = "STUB: not implemented"; return }

func NodeNotVisible(s *Selector) { _ = "STUB: not implemented"; return }

func NodeEnabled(s *Selector) { _ = "STUB: not implemented"; return }

func NodeSelected(s *Selector) { _ = "STUB: not implemented"; return }

func NodeNotPresent(s *Selector) { _ = "STUB: not implemented"; return }

func AtLeast(n int) QueryOption { _ = "STUB: not implemented"; return *new(QueryOption) }

func RetryInterval(interval time.Duration) QueryOption {
	_ = "STUB: not implemented"
	return *new(QueryOption)
}

func After(f func(context.Context, runtime.ExecutionContextID, ...*cdp.Node) error) QueryOption {
	_ = "STUB: not implemented"
	return *new(QueryOption)
}

func Populate(depth int64, pierce bool, opts ...PopulateOption) QueryOption {
	_ = "STUB: not implemented"
	return *new(QueryOption)
}

type PopulateOption = func(*time.Duration)

func PopulateWait(wait time.Duration) PopulateOption {
	_ = "STUB: not implemented"
	return *new(PopulateOption)
}

func WaitReady(sel any, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func WaitVisible(sel any, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func WaitNotVisible(sel any, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func WaitEnabled(sel any, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func WaitSelected(sel any, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func WaitNotPresent(sel any, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func Nodes(sel any, nodes *[]*cdp.Node, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func NodeIDs(sel any, ids *[]cdp.NodeID, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func Focus(sel any, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func Blur(sel any, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func Dimensions(sel any, model **dom.BoxModel, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func Text(sel any, text *string, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func TextContent(sel any, text *string, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func Clear(sel any, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func Value(sel any, value *string, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func SetValue(sel any, value string, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func Attributes(sel any, attributes *map[string]string, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func AttributesAll(sel any, attributes *[]map[string]string, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func SetAttributes(sel any, attributes map[string]string, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func AttributeValue(sel any, name string, value *string, ok *bool, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func SetAttributeValue(sel any, name, value string, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func RemoveAttribute(sel any, name string, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func JavascriptAttribute(sel any, name string, res any, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func SetJavascriptAttribute(sel any, name, value string, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func OuterHTML(sel any, html *string, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func InnerHTML(sel any, html *string, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func Click(sel any, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func DoubleClick(sel any, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func SendKeys(sel any, v string, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func SetUploadFiles(sel any, files []string, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func Submit(sel any, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func Reset(sel any, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func ComputedStyle(sel any, style *[]*css.ComputedStyleProperty, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func MatchedStyle(sel any, style **css.GetMatchedStylesForNodeReturns, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func ScrollIntoView(sel any, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func DumpTo(sel any, w io.Writer, prefix, indent string, nodeIDs bool, depth int64, pierce bool, wait time.Duration, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}

func Dump(sel any, w io.Writer, opts ...QueryOption) QueryAction {
	_ = "STUB: not implemented"
	return *new(QueryAction)
}
