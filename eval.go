package chromedp

import (
	"github.com/chromedp/cdproto/runtime"
)

type EvaluateAction Action

func Evaluate(expression string, res any, opts ...EvaluateOption) EvaluateAction {
	_ = "STUB: not implemented"
	return *new(EvaluateAction)
}

func parseRemoteObject(v *runtime.RemoteObject, res any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func EvaluateAsDevTools(expression string, res any, opts ...EvaluateOption) EvaluateAction {
	_ = "STUB: not implemented"
	return *new(EvaluateAction)
}

type EvaluateOption = func(*runtime.EvaluateParams) *runtime.EvaluateParams

func EvalObjectGroup(objectGroup string) EvaluateOption {
	_ = "STUB: not implemented"
	return *new(EvaluateOption)
}

func EvalWithCommandLineAPI(p *runtime.EvaluateParams) *runtime.EvaluateParams {
	_ = "STUB: not implemented"
	return nil
}

func EvalIgnoreExceptions(p *runtime.EvaluateParams) *runtime.EvaluateParams {
	_ = "STUB: not implemented"
	return nil
}

func EvalAsValue(p *runtime.EvaluateParams) *runtime.EvaluateParams {
	_ = "STUB: not implemented"
	return nil
}
