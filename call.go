package chromedp

import (
	"context"

	"github.com/chromedp/cdproto/runtime"
)

type CallAction Action

func CallFunctionOn(functionDeclaration string, res any, opt CallOption, args ...any) CallAction {
	_ = "STUB: not implemented"
	return *new(CallAction)
}

func callFunctionOn(ctx context.Context, functionDeclaration string, res any, opt CallOption, args ...any) (*runtime.RemoteObject, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CallOption = func(params *runtime.CallFunctionOnParams) *runtime.CallFunctionOnParams

type errAppender struct {
	args []*runtime.CallArgument
	err  error
}

func (ea *errAppender) append(v any) { _ = "STUB: not implemented"; return }
