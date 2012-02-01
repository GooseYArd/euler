package stree

import (
	"fmt"
	"log"
)

type RuleTwoType int

const (
	_ = iota
	newSon
	split
)

type LastPosType int

const (
	_ = iota
	lastCharInEdge
	otherChar
)

type SkipType int

const (
	_ = iota
	skip
	noSkip
)

type Node struct {
	sons             *Node
	rightSibling    *Node
	leftSibling     *Node
	father           *Node
	suffixLink      *Node
	pathPosition    int
	edgeLabelStart int
	edgeLabelEnd   int
}

type SuffixTree struct {
	e           int
	treeString  []byte
	length      int
	root        *Node
}

type Path struct {
	begin int
	end   int
}

type Pos struct {
	node     *Node
	edgePos int
}

var suffixless *Node
var ST_ERROR int
var logger *log.Logger

func createNode(father *Node, start int, end int, position int) (node *Node) {
	node = new(Node)
	node.father = father
	node.pathPosition = position
	node.edgeLabelStart = start
	node.edgeLabelEnd = end
	return node
}

func (tree *SuffixTree) findSon(node *Node, character byte) (newnode *Node) {
	if node == nil {
		return nil
	}
	newnode = node.sons
	for ; newnode != nil && tree.treeString[node.edgeLabelStart] != character; newnode = newnode.rightSibling {
	}
	return newnode
}

func (tree *SuffixTree) getNodeLabelEnd(node *Node) int {
	//if node == nil {
	//fmt.Printf("getNodeLabelEnd: node is nil\n");
	//	return tree.e
	//}
	if node.sons == nil {
		return tree.e
	}
	return node.edgeLabelEnd
}

func (tree *SuffixTree) getNodeLabelLength(node *Node) int {
	if node == nil {
		return 1
	}
	return tree.getNodeLabelEnd(node) - node.edgeLabelStart + 1
}

func (tree *SuffixTree) isLastCharInEdge(node *Node, edgePos int) bool {
	if edgePos == tree.getNodeLabelLength(node)-1 {
		return true
	}
	return false
}

func connectSiblings(leftSib *Node, rightSib *Node) {
	if leftSib != nil {
		leftSib.rightSibling = rightSib
	}
	if rightSib != nil {
		rightSib.leftSibling = leftSib
	}
}

func applyExtensionRule2(node *Node, edgeLabelBegin int, edgeLabelEnd int, pathPos int, edgePos int, _type RuleTwoType) *Node {

	if _type == newSon {
		//DEBUG
		fmt.Printf("rule 2: new leaf (%v,%v)\n", edgeLabelBegin, edgeLabelEnd)
		//END

		newLeaf := createNode(node, edgeLabelBegin, edgeLabelEnd, pathPos)
		son := node.sons
		for son.rightSibling != nil {
			son = son.rightSibling
		}
		connectSiblings(son, newLeaf)
		return newLeaf
	}

	fmt.Printf("rule 2: split(%v, %v)\n", edgeLabelBegin, edgeLabelEnd)

	newInternal := createNode(node.father, node.edgeLabelStart, node.edgeLabelStart+edgePos, node.pathPosition)
	node.edgeLabelStart += edgePos + 1
	newLeaf := createNode(newInternal, edgeLabelBegin, edgeLabelEnd, pathPos)

	connectSiblings(node.leftSibling, newInternal)
	connectSiblings(newInternal, node.rightSibling)
	node.leftSibling = nil

	if newInternal.father.sons == node {
		newInternal.father.sons = newInternal
	}

	newInternal.sons = node
	node.father = newInternal
	connectSiblings(node, newLeaf)

	return newInternal
}

func (tree *SuffixTree) traceSingleEdge(node *Node, str Path, _type SkipType) (contNode *Node, edgePos int, charsFound int, searchDone bool) {

	var length int
	var strLen int

	searchDone = true
	edgePos = 0
	
	contNode = tree.findSon(node, tree.treeString[str.begin])
	if contNode == nil {
		edgePos = tree.getNodeLabelLength(node) - 1
		charsFound = 0
		return node, edgePos, charsFound, searchDone
	}

	length = tree.getNodeLabelLength(contNode)
	strLen = str.end - str.begin + 1

	if _type == skip {
		if length <= strLen {
			charsFound = length
			edgePos = length - 1
			if length < strLen {
				searchDone = false
			}
		} else {
			charsFound = strLen
			edgePos = strLen - 1
		}
		return contNode, edgePos, charsFound, searchDone
	} else {
		if strLen < length {
			length = strLen
		}

		edgePos = 1
		charsFound = 1
		for edgePos < length {
			charsFound++
			edgePos++
			if tree.treeString[contNode.edgeLabelStart+edgePos] != tree.treeString[str.begin+edgePos] {
				edgePos--
				return contNode, edgePos, charsFound, searchDone
			}
		}
	}

	edgePos--
	if charsFound < strLen {
		searchDone = false
	}

	return contNode, edgePos, charsFound, searchDone
}

func (tree *SuffixTree) traceString(node *Node, str Path, _type SkipType) (newnode *Node, edgePos int, charsFound int) {

	searchDone := false
	edgeCharsFound := 0
	newnode = node
	
	for searchDone == false {
		newnode, edgePos, edgeCharsFound, searchDone = tree.traceSingleEdge(newnode, str, _type)
		str.begin += edgeCharsFound
		charsFound += edgeCharsFound
	}

	return newnode, edgePos, charsFound

}

func (tree *SuffixTree) FindSubstring(W string, P int) (pathPos int, err error) {
	node := tree.findSon(tree.root, W[0])
	k := 0
	j := 0
	nodeLabelEnd := 0

	for node != nil {
		k = node.edgeLabelStart
		nodeLabelEnd = tree.getNodeLabelEnd(node)

		for j < P && k <= nodeLabelEnd && tree.treeString[k] == W[j] {
			j++
			k++
		}

		if j == P {
			return node.pathPosition, err
		} else if k > nodeLabelEnd {
			node = tree.findSon(node, W[j])
		} else {
			// return ST_ERROR
			return pathPos, err
		}
	}
	//return ST_ERROR
	return pathPos, err
}

func (tree *SuffixTree) followSuffixLink(pos *Pos) {
	var gama Path
	if pos.node == tree.root {
		return
	}

	if pos.node.suffixLink == nil ||
		tree.isLastCharInEdge(pos.node, pos.edgePos) == false {
		if pos.node.father == tree.root {
			pos.node = tree.root
			return
		}
		
		gama.begin = pos.node.edgeLabelStart
		gama.end = pos.node.edgeLabelStart + pos.edgePos
		pos.node = pos.node.father.suffixLink
		pos.node, pos.edgePos, _ = tree.traceString(pos.node, gama, skip)
	} else {
		pos.node = pos.node.suffixLink
		pos.edgePos = tree.getNodeLabelLength(pos.node) - 1
	}
	return
}

func createSuffixLink(node *Node, link *Node) {
	node.suffixLink = link
}

func (tree *SuffixTree) SEA(pos *Pos, str Path, afterRule3 bool) (ruleApplied int) {
	charsFound := 0
	pathPos := str.begin
	var tmp *Node

	// DEBUG
	tree.PrintTree()
	fmt.Printf("extension: %v  phase+1: %v",str.begin, str.end);
 	if afterRule3 == false {
 		fmt.Printf("   followed from (%v,%v | %v) ", pos.node.edgeLabelStart, tree.getNodeLabelEnd(pos.node), pos.edgePos);
 	} else {
 		fmt.Printf("   starting at (%v,%v | %v) ", pos.node.edgeLabelStart, tree.getNodeLabelEnd(pos.node), pos.edgePos);
 	}
	// END DEBUG

	if afterRule3 == false {
		tree.followSuffixLink(pos)
	}

	if pos.node == tree.root {
		_, pos.edgePos, charsFound = tree.traceString(pos.node, str, noSkip)
	} else {
		str.begin = str.end
		charsFound = 0

		if tree.isLastCharInEdge(pos.node, pos.edgePos) {
			tmp = tree.findSon(pos.node, tree.treeString[str.end])
			if tmp != nil {
				pos.node = tmp
				pos.edgePos = 0
				charsFound = 1
			}
		} else {
			if tree.treeString[pos.node.edgeLabelStart+pos.edgePos+1] == tree.treeString[str.end] {
				pos.edgePos++
				charsFound = 1
			}
		}
	}

	if charsFound == str.end-str.begin+1 {
		ruleApplied = 3
		if suffixless != nil {
			createSuffixLink(suffixless, pos.node.father)
			suffixless = nil
		}
		return ruleApplied
	}

	if tree.isLastCharInEdge(pos.node, pos.edgePos) || pos.node == tree.root {
		if pos.node.sons != nil {
			applyExtensionRule2(pos.node, str.begin+charsFound, str.end, pathPos, 0, newSon)
			ruleApplied = 2
			if suffixless != nil {
				createSuffixLink(suffixless, pos.node)
				suffixless = nil
			}
		}
	} else {
		tmp = applyExtensionRule2(
			pos.node,
			str.begin+charsFound,
			str.end,
			pathPos,
			pos.edgePos,
			split)
		if suffixless != nil {
			createSuffixLink(suffixless, tmp)
		}
		if tree.getNodeLabelLength(tmp) == 1 && tmp.father == tree.root {
			tmp.suffixLink = tree.root
			suffixless = nil
		} else {
			suffixless = tmp
		}

		pos.node = tmp
		ruleApplied = 2
	}

	// fmt.Printf("rule 3 (%v,%v)\n", str.begin, str.end)
	return ruleApplied
}

// Input :The tree, 
// * pos - the node and position in its incoming edge where extension begins, 
// * the phase number, 
// * the first extension number of that phase, 
// * a flag signaling whether the extension is the first of this phase, 
//   after the last phase ended with rule 3. If so - extension will be executed again in this phase, and thus its suffix link would not be followed.

func (tree *SuffixTree) SPA(pos *Pos, phase int, extension int, repeatedExtension bool) (int, bool) {
	ruleApplied := 0
	var str Path
	tree.e = phase + 1

	for extension <= phase+1 {
		str.begin = extension
		str.end = phase + 1

		ruleApplied = tree.SEA(pos, str, repeatedExtension)
		if ruleApplied == 3 {
			repeatedExtension = true
			break
		}
		repeatedExtension = false
		extension++
	}

	return extension, repeatedExtension
}

func CreateTree(str string) (tree *SuffixTree, err error) {

	pos := new(Pos)

	tree = new(SuffixTree)

	ST_ERROR = len(str) + 10	
	length := len(str)
	
	// real string starts at index 1
	tree.treeString = make([]byte, length + 2)
	tree.treeString[0] = 0
	tree.treeString[length] = '$'
	tree.length = length + 1

	for i,v := range str + "$" {
		tree.treeString[i+1] = byte(v)
	}

	tree.root = createNode(nil, 0, 0, 0)
	tree.root.suffixLink = nil

	repeatedExtension := false
	extension := 2
	phase := 2

	tree.root.sons = createNode(tree.root, 1, tree.length, 1)
	suffixless = nil
	pos.node = tree.root
	pos.edgePos = 0
	
	for phase < tree.length {
		extension, repeatedExtension = tree.SPA(pos, phase, extension, repeatedExtension)
		phase++
	}

	return tree, err
}

func DeleteSubTree(node *Node) {
	if node == nil {
		return
	}

	if node.rightSibling != nil {
		DeleteSubTree(node.rightSibling)
	}

	if node.sons != nil {
		DeleteSubTree(node.sons)
	}

}

func DeleteTree(tree *SuffixTree) {
	if tree == nil {
		return
	}
	DeleteSubTree(tree.root)
}

func (tree *SuffixTree) PrintNode(node1 *Node, depth int) {
	fmt.Printf("AT TOP, DEPTH IS %v\n", depth)
	if depth > 0 {
		for d:= depth; d > 1; d-- {
			fmt.Printf("|")
		}
		fmt.Printf("+")		
		end := tree.getNodeLabelEnd(node1)
		for i := node1.edgeLabelStart; i <= end; i++ {
			fmt.Printf("%c", tree.treeString[i])
		}
		//DEBUG
		fmt.Printf("  \t\t\t(%v,%v | %v)",node1.edgeLabelStart,end,node1.pathPosition);
		fmt.Println()
	}
	fmt.Printf("AT BOTTOM, DEPTH IS %v\n", depth)
	for node2 := node1.sons; node2 != nil; node2 = node2.rightSibling {
		fmt.Printf("recurring (depth: %v)...\n", (depth+1));
		tree.PrintNode(node2, depth+1)
	}
}

func (tree *SuffixTree) PrintFullNode(node *Node) {
	if node == nil {
		return
	}

	start := node.edgeLabelStart
	end := tree.getNodeLabelEnd(node)

	if node.father != tree.root {
		tree.PrintFullNode(node.father)
	}

	// TODO
	for start <= end {
		fmt.Printf("%v", tree.treeString[start])
		start++
	}
}

func (tree *SuffixTree) PrintTree() {
	fmt.Printf("\nroot\n")
	tree.PrintNode(tree.root, 0)
}
