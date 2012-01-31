package stree

import (
	"fmt"
	"log"
//	"os"
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
	treeString  string
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

func (tree SuffixTree) findSon(node *Node, character byte) *Node {
	if node == nil {
		return nil
	}
	node = node.sons
	for ; node != nil && tree.treeString[node.edgeLabelStart] != character; node = node.rightSibling {
	}
	return node
}

func (tree SuffixTree) getNodeLabelEnd(node *Node) int {
	if node == nil {
		return tree.e
	}
	if node.sons == nil {
		return tree.e
	}
	return node.edgeLabelEnd
}

func (tree SuffixTree) getNodeLabelLength(node *Node) int {
	if node == nil {
		return 1
	}
	return tree.getNodeLabelEnd(node) - node.edgeLabelStart + 1
}

func (tree SuffixTree) isLastCharInEdge(node *Node, edgePos int) bool {
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
	var newLeaf *Node
	var newInternal *Node
	var son *Node

	//DEBUG
	fmt.Printf("rule 2: new leaf (%v,%v)\n", edgeLabelBegin, edgeLabelEnd)
	//END

	if _type == newSon {
		newLeaf = createNode(node, edgeLabelBegin, edgeLabelEnd, pathPos)
		son = node.sons
		for son.rightSibling != nil {
			son = son.rightSibling
		}
		connectSiblings(son, newLeaf)
		return newLeaf
	}

	fmt.Printf("rule 2: split(%v, %v)\n", edgeLabelBegin, edgeLabelEnd)

	newInternal = createNode(node.father, node.edgeLabelStart, node.edgeLabelStart+edgePos, node.pathPosition)
	node.edgeLabelStart += edgePos + 1
	newLeaf = createNode(newInternal, edgeLabelBegin, edgeLabelEnd, pathPos)

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

func (tree SuffixTree) traceSingleEdge(str Path, _type SkipType) (node *Node, edgePos int, charsFound int, searchDone bool) {

	var contNode *Node
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

	node = contNode
	length = tree.getNodeLabelLength(node)
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
		return node, edgePos, charsFound, searchDone
	} else {
		if strLen < length {
			length = strLen
		}

		edgePos = 1
		charsFound = 1
		for edgePos < length {
			charsFound++
			edgePos++
			if tree.treeString[node.edgeLabelStart+edgePos] != tree.treeString[str.begin+edgePos] {
				edgePos--
				return node, edgePos, charsFound, searchDone
			}
		}
	}

	edgePos--
	if charsFound < strLen {
		searchDone = false
	}

	return node, edgePos, charsFound, searchDone
}

func (tree SuffixTree) traceString(str Path, _type SkipType) (node *Node, edgePos int, charsFound int) {

	searchDone := false
	edgeCharsFound := 0
		
	for searchDone == false {
		node, edgePos, edgeCharsFound, searchDone = tree.traceSingleEdge(str, _type)
		str.begin += edgeCharsFound
		charsFound += edgeCharsFound
	}

	return node, edgePos, charsFound

}

func (tree SuffixTree) FindSubstring(W string, P int) (pathPos int, err error) {
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

func (tree SuffixTree) followSuffixLink(pos *Pos) {
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
		pos.node, pos.edgePos, _ = tree.traceString(gama, skip)
	} else {
		pos.node = pos.node.suffixLink
		pos.edgePos = tree.getNodeLabelLength(pos.node) - 1
	}
}

func createSuffixLink(node *Node, link *Node) {
	node.suffixLink = link
}

func (tree SuffixTree) SEA(pos *Pos, str Path, afterRule3 bool) (ruleApplied int) {
	charsFound := 0
	pathPos := str.begin
	var tmp *Node

	// DEBUG
	tree.PrintTree()
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
		fmt.Println("pos.node == root, calling traceString")
		_, pos.edgePos, charsFound = tree.traceString(str, noSkip)
	} else {
		fmt.Println("pos.node != root, calling traceString")
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

	fmt.Printf("charsFound: %v", charsFound)

	if charsFound == str.end-str.begin+1 {
		ruleApplied = 3
		if suffixless != nil {
			createSuffixLink(suffixless, pos.node.father)
			suffixless = nil
		}
		fmt.Println("returning inside block A")
		return
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

	fmt.Printf("rule 3 (%v,%v)\n", str.begin, str.end)
	return ruleApplied
}

// Input :The tree, 
// * pos - the node and position in its incoming edge where extension begins, 
// * the phase number, 
// * the first extension number of that phase, 
// * a flag signaling whether the extension is the first of this phase, 
//   after the last phase ended with rule 3. If so - extension will be executed again in this phase, and thus its suffix link would not be followed.

func (tree SuffixTree) SPA(pos *Pos, phase int, extension int, repeatedExtension bool) (int) {
	ruleApplied := 0
	var str Path
	tree.e = phase + 1

	fmt.Println("in SPA")

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

	return extension
}

func CreateTree(str string) (tree *SuffixTree, err error) {

	pos := new(Pos)
	tree = new(SuffixTree)

	//length := len(str)

	//logf, err := os.Create("andy.log")
	if err != nil {
		panic("couldn't open log file")
	}

	//fmt = log.New(logf, "", 1)
	// fmt, _ = os.Create("andy.log")

	ST_ERROR = len(str) + 10

	tree.treeString = str + "$"
	tree.length = len(tree.treeString)

	//tree.treeString = make([]byte, length+1)
	//copy(tree.treeString[:], []byte("hello world")) 

	//tree.treeString[length] = '$'
	
	tree.root = createNode(nil, 0, 0, 0)
	tree.root.suffixLink = nil

	repeatedExtension := false
	extension := 2
	phase := 2

	tree.root.sons = createNode(tree.root, 1, tree.length, 1)
	suffixless = nil
	pos.node = tree.root
	pos.edgePos = 0

	for ; phase < tree.length; phase++ {
		extension = tree.SPA(pos, phase, extension, repeatedExtension)
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

func (tree SuffixTree) PrintNode(node1 *Node, depth int) {
	node2 := node1.sons
	d := depth
	start := node1.edgeLabelStart
	end := tree.getNodeLabelEnd(node1)

	if depth > 0 {
		for ; d > 1; d-- {
			fmt.Printf("|")
		}
		fmt.Printf("+")
		for ; start <= end; start++ {
			fmt.Printf("%c", tree.treeString[start])
		}
		//DEBUG
		fmt.Printf("  \t\t\t(%v,%v | %v)",node1.edgeLabelStart,end,node1.pathPosition);
		fmt.Println()
	}
	for node2 != nil {
		tree.PrintNode(node2, depth+1)
		node2 = node2.rightSibling
	}
}

func (tree SuffixTree) PrintFullNode(node *Node) {
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

func (tree SuffixTree) PrintTree() {
	fmt.Printf("\nroot\n")
	tree.PrintNode(tree.root, 0)
}
