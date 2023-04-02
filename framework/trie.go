package framework

import (
	"errors"
	"strings"
)

type Tree struct {
	root *node
}

type node struct {
	isLast   bool                // 代表这个节点是否可以成为最终的路由规则
	segment  string              // 代表路由中某个段的字符串
	handlers []ControllerHandler // 代表这个节点中包含的中间件和控制器
	childs   []*node             // 代表这个节点下的子节点
	parent   *node
}

func NewTree() *Tree {
	return &Tree{&node{}}
}

// isGeneralSegment 是否是通用 segment（即以:开头）
func isGeneralSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

func (n *node) matchNode(uri string) *node {
	segments := strings.SplitN(uri, "/", 2)
	segment := segments[0]

	if !isGeneralSegment(segment) {
		segment = strings.ToUpper(segment)
	}

	// 匹配符合的下一层子节点
	cnodes := n.filterChildNodes(segment)
	// 如果没有符合的子节点，说明这个 uri 一定是之前不存在的
	if cnodes == nil || len(cnodes) == 0 {
		return nil
	}

	// 如果只有一个 segment，就看子节点里有没有 isLast == true 的
	if len(segments) == 1 {
		for _, cnode := range cnodes {
			if cnode.isLast {
				return cnode
			}
		}
		return nil
	}

	// 如果有两个 segment，则递归每个子节点查找
	for _, cnode := range cnodes {
		matchNode := cnode.matchNode(segments[1])
		if matchNode != nil {
			return matchNode
		}
	}
	return nil
}

// 过滤下一层满足 segment 规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}

	// 如果 segment 是通配符，则所有下一层子节点都符合
	if isGeneralSegment(segment) {
		return n.childs
	}

	cnodes := make([]*node, 0, len(n.childs))
	for _, child := range n.childs {
		if isGeneralSegment(child.segment) {
			// 如果子节点是通配符，则符合
			cnodes = append(cnodes, child)
		} else if child.segment == segment {
			cnodes = append(cnodes, child)
		}
	}
	return cnodes
}

func (t *Tree) AddRouter(uri string, handlers []ControllerHandler) error {
	n := t.root
	if n.matchNode(uri) != nil {
		return errors.New("router exist: " + uri)
	}

	segments := strings.Split(uri, "/")
	for index, segment := range segments {
		if !isGeneralSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == (len(segments) - 1)
		var objNode *node

		cnodes := n.filterChildNodes(segment)
		if len(cnodes) > 0 {
			for _, cnode := range cnodes {
				if cnode.segment == segment {
					// 跟子节点的 segment 相等，只可能有一个
					objNode = cnode
					break
				}
			}
		}

		if objNode == nil {
			objNode = &node{}
			objNode.segment = segment
			objNode.parent = n
			if isLast {
				objNode.isLast = true
				objNode.handlers = handlers
			}
			n.childs = append(n.childs, objNode)
		}

		n = objNode
	}

	return nil
}

func (n *node) parseParamsFromEndNode(uri string) map[string]string {
	ret := map[string]string{}
	segments := strings.Split(uri, "/")
	count := len(segments)
	cur := n
	for i := count - 1; i >= 0; i-- {
		if cur.segment == "" {
			break
		}
		if isGeneralSegment(cur.segment) {
			ret[cur.segment[1:]] = segments[i]
		}
		cur = cur.parent
	}
	return ret
}
