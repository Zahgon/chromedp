package chromedp

import (
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp/device"
)

type EmulateAction Action

func EmulateViewport(width, height int64, opts ...EmulateViewportOption) EmulateAction {
	_ = "STUB: not implemented"
	return *new(EmulateAction)
}

type EmulateViewportOption = func(*emulation.SetDeviceMetricsOverrideParams, *emulation.SetTouchEmulationEnabledParams)

func EmulateScale(scale float64) EmulateViewportOption {
	_ = "STUB: not implemented"
	return *new(EmulateViewportOption)
}

func EmulateOrientation(orientation emulation.OrientationType, angle int64) EmulateViewportOption {
	_ = "STUB: not implemented"
	return *new(EmulateViewportOption)
}

func EmulateLandscape(p1 *emulation.SetDeviceMetricsOverrideParams, p2 *emulation.SetTouchEmulationEnabledParams) {
	_ = "STUB: not implemented"
	return
}

func EmulatePortrait(p1 *emulation.SetDeviceMetricsOverrideParams, p2 *emulation.SetTouchEmulationEnabledParams) {
	_ = "STUB: not implemented"
	return
}

func EmulateMobile(p1 *emulation.SetDeviceMetricsOverrideParams, p2 *emulation.SetTouchEmulationEnabledParams) {
	_ = "STUB: not implemented"
	return
}

func EmulateTouch(p1 *emulation.SetDeviceMetricsOverrideParams, p2 *emulation.SetTouchEmulationEnabledParams) {
	_ = "STUB: not implemented"
	return
}

func ResetViewport() EmulateAction { _ = "STUB: not implemented"; return *new(EmulateAction) }

type Device interface {
	Device() device.Info
}

func Emulate(device Device) EmulateAction { _ = "STUB: not implemented"; return *new(EmulateAction) }

func EmulateReset() EmulateAction { _ = "STUB: not implemented"; return *new(EmulateAction) }
