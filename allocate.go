package chromedp

import (
	"context"
	"io"
	"os/exec"
	"sync"
	"time"
)

type Allocator interface {
	Allocate(context.Context, ...BrowserOption) (*Browser, error)

	Wait()
}

func setupExecAllocator(opts ...ExecAllocatorOption) *ExecAllocator {
	_ = "STUB: not implemented"
	return nil
}

var DefaultExecAllocatorOptions = [...]ExecAllocatorOption{
	NoFirstRun,
	NoDefaultBrowserCheck,
	Headless,

	Flag("disable-background-networking", true),
	Flag("enable-features", "NetworkService,NetworkServiceInProcess"),
	Flag("disable-background-timer-throttling", true),
	Flag("disable-backgrounding-occluded-windows", true),
	Flag("disable-breakpad", true),
	Flag("disable-client-side-phishing-detection", true),
	Flag("disable-default-apps", true),
	Flag("disable-dev-shm-usage", true),
	Flag("disable-extensions", true),
	Flag("disable-features", "site-per-process,Translate,BlinkGenPropertyTrees"),
	Flag("disable-hang-monitor", true),
	Flag("disable-ipc-flooding-protection", true),
	Flag("disable-popup-blocking", true),
	Flag("disable-prompt-on-repost", true),
	Flag("disable-renderer-backgrounding", true),
	Flag("disable-sync", true),
	Flag("force-color-profile", "srgb"),
	Flag("metrics-recording-only", true),
	Flag("safebrowsing-disable-auto-update", true),
	Flag("enable-automation", true),
	Flag("password-store", "basic"),
	Flag("use-mock-keychain", true),
}

func NewExecAllocator(parent context.Context, opts ...ExecAllocatorOption) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

type ExecAllocatorOption = func(*ExecAllocator)

type ExecAllocator struct {
	execPath  string
	initFlags map[string]any
	initEnv   []string

	wsURLReadTimeout time.Duration

	modifyCmdFunc func(cmd *exec.Cmd)

	wg sync.WaitGroup

	combinedOutputWriter io.Writer
}

var allocTempDir string

func (a *ExecAllocator) Allocate(ctx context.Context, opts ...BrowserOption) (*Browser, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readOutput(rc io.ReadCloser, forward io.Writer) (wsURL string, _ func(), _ error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (a *ExecAllocator) Wait() { _ = "STUB: not implemented"; return }

func ExecPath(path string) ExecAllocatorOption {
	_ = "STUB: not implemented"
	return *new(ExecAllocatorOption)
}

func findExecPath() string { _ = "STUB: not implemented"; return "" }

func Flag(name string, value any) ExecAllocatorOption {
	_ = "STUB: not implemented"
	return *new(ExecAllocatorOption)
}

func Env(vars ...string) ExecAllocatorOption {
	_ = "STUB: not implemented"
	return *new(ExecAllocatorOption)
}

func ModifyCmdFunc(f func(cmd *exec.Cmd)) ExecAllocatorOption {
	_ = "STUB: not implemented"
	return *new(ExecAllocatorOption)
}

func UserDataDir(dir string) ExecAllocatorOption {
	_ = "STUB: not implemented"
	return *new(ExecAllocatorOption)
}

func ProxyServer(proxy string) ExecAllocatorOption {
	_ = "STUB: not implemented"
	return *new(ExecAllocatorOption)
}

func IgnoreCertErrors(a *ExecAllocator) { _ = "STUB: not implemented"; return }

func WindowSize(width, height int) ExecAllocatorOption {
	_ = "STUB: not implemented"
	return *new(ExecAllocatorOption)
}

func UserAgent(userAgent string) ExecAllocatorOption {
	_ = "STUB: not implemented"
	return *new(ExecAllocatorOption)
}

func NoSandbox(a *ExecAllocator) { _ = "STUB: not implemented"; return }

func NoFirstRun(a *ExecAllocator) { _ = "STUB: not implemented"; return }

func NoDefaultBrowserCheck(a *ExecAllocator) { _ = "STUB: not implemented"; return }

func Headless(a *ExecAllocator) { _ = "STUB: not implemented"; return }

func DisableGPU(a *ExecAllocator) { _ = "STUB: not implemented"; return }

func CombinedOutput(w io.Writer) ExecAllocatorOption {
	_ = "STUB: not implemented"
	return *new(ExecAllocatorOption)
}

func WSURLReadTimeout(t time.Duration) ExecAllocatorOption {
	_ = "STUB: not implemented"
	return *new(ExecAllocatorOption)
}

func NewRemoteAllocator(parent context.Context, url string, opts ...RemoteAllocatorOption) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

type RemoteAllocatorOption = func(*RemoteAllocator)

type RemoteAllocator struct {
	wsURL         string
	modifyURLFunc func(ctx context.Context, wsURL string) (string, error)

	wg sync.WaitGroup
}

func (a *RemoteAllocator) Allocate(ctx context.Context, opts ...BrowserOption) (*Browser, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *RemoteAllocator) Wait() { _ = "STUB: not implemented"; return }

func NoModifyURL(a *RemoteAllocator) { _ = "STUB: not implemented"; return }
