package chromedp

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

func forceIP(ctx context.Context, urlstr string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func resolveHost(ctx context.Context, host string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func modifyURL(ctx context.Context, urlstr string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func runListeners(list []cancelableListener, ev any) []cancelableListener {
	_ = "STUB: not implemented"
	return nil
}

type frameOp func(*cdp.Frame)

func frameAttached(id cdp.FrameID) frameOp { _ = "STUB: not implemented"; return *new(frameOp) }

func frameDetached(f *cdp.Frame) { _ = "STUB: not implemented"; return }

func frameStartedLoading(f *cdp.Frame) { _ = "STUB: not implemented"; return }

func frameStoppedLoading(f *cdp.Frame) { _ = "STUB: not implemented"; return }

func setFrameState(f *cdp.Frame, fs cdp.FrameState) { _ = "STUB: not implemented"; return }

func clearFrameState(f *cdp.Frame, fs cdp.FrameState) { _ = "STUB: not implemented"; return }

type nodeOp func(*cdp.Node)

func walk(m map[cdp.NodeID]*cdp.Node, n *cdp.Node) { _ = "STUB: not implemented"; return }

func setChildNodes(m map[cdp.NodeID]*cdp.Node, nodes []*cdp.Node) nodeOp {
	_ = "STUB: not implemented"
	return *new(nodeOp)
}

func attributeModified(name, value string) nodeOp { _ = "STUB: not implemented"; return *new(nodeOp) }

func attributeRemoved(name string) nodeOp { _ = "STUB: not implemented"; return *new(nodeOp) }

func inlineStyleInvalidated(ids []cdp.NodeID) nodeOp {
	_ = "STUB: not implemented"
	return *new(nodeOp)
}

func characterDataModified(characterData string) nodeOp {
	_ = "STUB: not implemented"
	return *new(nodeOp)
}

func childNodeCountUpdated(count int64) nodeOp { _ = "STUB: not implemented"; return *new(nodeOp) }

func childNodeInserted(m map[cdp.NodeID]*cdp.Node, prevID cdp.NodeID, c *cdp.Node) nodeOp {
	_ = "STUB: not implemented"
	return *new(nodeOp)
}

func childNodeRemoved(m map[cdp.NodeID]*cdp.Node, id cdp.NodeID) nodeOp {
	_ = "STUB: not implemented"
	return *new(nodeOp)
}

func shadowRootPushed(m map[cdp.NodeID]*cdp.Node, c *cdp.Node) nodeOp {
	_ = "STUB: not implemented"
	return *new(nodeOp)
}

func shadowRootPopped(m map[cdp.NodeID]*cdp.Node, id cdp.NodeID) nodeOp {
	_ = "STUB: not implemented"
	return *new(nodeOp)
}

func pseudoElementAdded(m map[cdp.NodeID]*cdp.Node, c *cdp.Node) nodeOp {
	_ = "STUB: not implemented"
	return *new(nodeOp)
}

func pseudoElementRemoved(m map[cdp.NodeID]*cdp.Node, id cdp.NodeID) nodeOp {
	_ = "STUB: not implemented"
	return *new(nodeOp)
}

func distributedNodesUpdated(nodes []*cdp.BackendNode) nodeOp {
	_ = "STUB: not implemented"
	return *new(nodeOp)
}

func scrollableFlagUpdated(m map[cdp.NodeID]*cdp.Node, id cdp.NodeID) nodeOp {
	_ = "STUB: not implemented"
	return *new(nodeOp)
}

func insertNode(n []*cdp.Node, prevID cdp.NodeID, c *cdp.Node) []*cdp.Node {
	_ = "STUB: not implemented"
	return nil
}

func removeNode(n []*cdp.Node, id cdp.NodeID) []*cdp.Node { _ = "STUB: not implemented"; return nil }

func isCouldNotComputeBoxModelError(err error) bool { _ = "STUB: not implemented"; return false }
