package chromedp

import (
	"github.com/chromedp/cdproto/page"
)

type NavigateAction Action

func Navigate(urlstr string) NavigateAction { _ = "STUB: not implemented"; return *new(NavigateAction) }

func NavigationEntries(currentIndex *int64, entries *[]*page.NavigationEntry) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

func NavigateToHistoryEntry(entryID int64) NavigateAction {
	_ = "STUB: not implemented"
	return *new(NavigateAction)
}

func NavigateBack() NavigateAction { _ = "STUB: not implemented"; return *new(NavigateAction) }

func NavigateForward() NavigateAction { _ = "STUB: not implemented"; return *new(NavigateAction) }

func Reload() NavigateAction { _ = "STUB: not implemented"; return *new(NavigateAction) }

func Stop() Action { _ = "STUB: not implemented"; return *new(Action) }

func Location(urlstr *string) Action { _ = "STUB: not implemented"; return *new(Action) }

func Title(title *string) Action { _ = "STUB: not implemented"; return *new(Action) }
