package gee

import "strings"

type node struct {
	pattern  string  // 设置叶子节点的值，后续用于判断该路由是否存在
	part     string  // 如果是模糊匹配，则会把当前模糊匹配对应的值赋值给part
	children []*node // 子节点
	isWild   bool    // ture则进行模糊匹配
}

// 找第一个匹配的子节点，该方法用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 找到所有匹配的子节点，该方法用于查询
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 插入一个新的子节点
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		// 说明当前已经完成了insert
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		// 没有该孩子，则创建
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// 查询节点是否存在，
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			//该节点不存在路由
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
