package fgee

import (
	"strings"
)

type node struct {
	pattern  string  // 待匹配路由, 例如 /p/:lang (会在路由创建的时候保存在最后一个节点)
	part     string  // 路由中的一部分, 例如 :lang
	children []*node // 子节点, 构建树, 例如[doc, r, intro]
	isWild   bool    // 是否模糊匹配, part含有 : 或 * 为true
}

// 插入
// todo 路由冲突问题
func (n *node) insert(pattern string, parts []string, h int) {
	partsLen := len(parts)
	if partsLen == h {
		n.pattern = pattern
		return
	}
	part := parts[h]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: getWild(part)}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, h+1)
}

// 查找
func (n *node) search(parts []string, h int) *node {
	if len(parts) == h || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[h]
	childs := n.matchChilds(part)
	for _, child := range childs {
		res := child.search(parts, h+1)
		if res != nil {
			return res
		}
	}
	return nil
}

// 通过part与isWild字段匹配第一个节点, 用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 通过part与isWild字段找到所有符合条件的节点, 用于查找
func (n *node) matchChilds(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 获取当前字符串是否为模糊匹配
func getWild(part string) bool {
	return part[0] == ':' || part[0] == '*'
}
