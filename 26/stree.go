package stree

type RuleTwoType int
const (
	_ = iota
	new_son
	split
)

type LastPosType int
const (
	_ = iota
	last_char_in_edge
	other_char
)

type SkipType int
const (
	_ = iota
	skip
	no_skip
)

type Node struct {
	sons *Node
	right_sibling *Node
	left_sibling *Node
	father *Node
	suffix_link *Node
	path_position int
	edge_label_start int
	edge_label_end int
}

type SuffixTree struct {
	e int
	tree_string string
	length int
	root *Node
}

type Path struct {
	begin int
	end int	
}

type Pos struct {
	node *Node
	edge_pos int
}

func create_node(father *Node, start int, end int, position int) (node *Node) {
	node = new(Node)
	node.father = father
	node.path_position = position
	node.edge_label_start = start
	node.edge_label_end = end
	return node
}

func (tree SuffixTree) find_son(node *Node, character byte) (*Node) {
	node = node.sons
	for ; node != nil && tree.tree_string[node.edge_label_start] != character; node = node.right_sibling {
	}
	return node
}

func (tree SuffixTree) get_node_label_end(node *Node) (int) {
	if node.sons == nil {
		return tree.e
	}
	return node.edge_label_end
}

func (tree SuffixTree) get_node_label_length(node *Node) (int) {
	return tree.get_node_label_end(node) - node.edge_label_start + 1
}

func (tree SuffixTree) is_last_char_in_edge(node *Node, edge_pos int) (bool) {
	if edge_pos == tree.get_node_label_length(node) - 1 {
		return true
	}
	return false
}

func connect_siblings(left_sib *Node, right_sib *Node) {
	if left_sib != nil {
		left_sib.right_sibling = right_sib
	}
	if right_sib != nil {
		right_sib.left_sibling = left_sib
	}
}

func apply_extension_rule_2(node *Node, edge_label_begin int, edge_label_end int, path_pos int, edge_pos int, _type RuleTwoType) (*Node) {
	var new_leaf *Node
	var new_internal *Node
	var son *Node
	
	if _type == new_son {
		new_leaf = create_node(node, edge_label_begin, edge_label_end, path_pos)
		son = node.sons
		for ; son.right_sibling != nil; {
			son = son.right_sibling
		}
		connect_siblings(son, new_leaf)
		return new_leaf
	}

	new_internal = create_node(node.father, node.edge_label_start, node.edge_label_start + edge_pos, node.path_position)
	node.edge_label_start += edge_pos + 1
	new_leaf = create_node(new_internal, edge_label_begin, edge_label_end, path_pos)

	connect_siblings(node.left_sibling, new_internal)
	connect_siblings(new_internal, node.right_sibling)
	node.left_sibling = nil

	if new_internal.father.sons == node {
		new_internal.father.sons = new_internal
	}

	new_internal.sons = node
	node.father = new_internal
	connect_siblings(node, new_leaf)
	
	return new_internal
}

func (tree SuffixTree) trace_single_edge (str Path, _type SkipType) (node *Node, edge_pos int, chars_found int, search_done bool) {

	var cont_node *Node
	var length int
	var str_len int

	search_done = false
	edge_pos = 0

	cont_node = tree.find_son(node, tree.tree_string[str.begin])
	if cont_node == nil {
		edge_pos = tree.get_node_label_length(node) - 1
		chars_found = 0
		return node, edge_pos, chars_found, search_done
	}

	node = cont_node
	length = tree.get_node_label_length(node)
	str_len = str.end - str.begin + 1 

	if _type == skip {
		if length <= str_len {
			chars_found = length
			edge_pos = length - 1
			if length < str_len {
				search_done = false
			}
		} else {
			chars_found = str_len
			edge_pos = str_len - 1
		}
		return node, edge_pos, chars_found, search_done
	} else {
		if str_len < length {
			length = str_len
		}
		
		edge_pos = 1
		chars_found = 1
		for ; edge_pos < length ; {			
			chars_found++
			edge_pos++
			if tree.tree_string[node.edge_label_start+edge_pos] != tree.tree_string[str.begin+edge_pos] {
				edge_pos--
				return node, edge_pos, chars_found, search_done
			}
		}
	}
	
	edge_pos--
	if chars_found < str_len {
		search_done = false
	}

	return node, edge_pos, chars_found, search_done
}
		
func (tree SuffixTree) trace_string(str Path, _type SkipType) (node *Node, edge_pos int, chars_found int) {

	search_done := false
	edge_chars_found := 0
	
	for ; search_done == false ; {
		edge_pos = 0
		edge_chars_found = 0
		node, edge_pos, edge_chars_found, search_done = tree.trace_single_edge(str, _type)
		str.begin += edge_chars_found
		chars_found += edge_chars_found
	}
	return node, edge_pos, chars_found
}

func (tree SuffixTree) ST_FindSubstring(W string, P int) (path_pos int, err error) {
	node := tree.find_son(tree.root, W[0])
	k := 0
	j := 0
	node_label_end := 0
	
	for ; node != nil ; {
		k = node.edge_label_start
		node_label_end := tree.get_node_label_end(node)
		
		for ; j < P && k <= node_label_end && tree.tree_string[k] == W[j]; {
			j++
			k++
		}
		
		if j == P {
			return node.path_position, err
		} else if k > node_label_end {
			node = tree.find_son(node, W[j])
		} else {
			// return ST_ERROR
			return path_pos, err
		}
	}
	//return ST_ERROR
	return path_pos, err
}

func (tree SuffixTree) follow_suffix_link(pos *Pos) {
	var gama Path
	chars_found := 0

	if pos.node == tree.root {
		return
	}
	
	if pos.node.suffix_link == nil ||
		tree.is_last_char_in_edge(pos.node, pos.edge_pos) == false {
		if pos.node.father == tree.root {
			pos.node = tree.root
			return
		}
		
		gama.begin = pos.node.edge_label_start
		gama.end = pos.node.edge_label_start + pos.edge_pos
		pos.node = pos.node.father.suffix_link
		pos.node, pos.edge_pos, chars_found = tree.trace_string(gama, skip)
	} else {
		pos.node = pos.node.suffix_link
		pos.edge_pos = tree.get_node_label_length(pos.node) - 1 
	}	
}


func create_suffic_link(node *Node, link *Node) {
	node.suffix_link = link
}

func (tree SuffixTree) SEA(pos *Pos, str Path, after_rule_3 bool) (rule_applied int) {
	chars_found := 0
	path_pos := str.begin
	var tmp Node
	
	if after_rule_3 == false {
		tree.follow_suffix_link(pos)
	}

	if pos.node == tree.root {
		pos.node, pos.edge_pos, chars_found = tree.trace_string(tree.root, str, no_skip)
	} else {
		str.begin = str.end
		chars_found = 0
		
		if tree.is_last_char_in_edge(pos.node, pos.edge_pos) {
			tmp = tree.find_son(pos.node, tree.tree_string[str.end])
			if tmp != nil {
				pos.node = tmp
				pos.edge_pos = 0
				chars_found = 1
			}
		} else {
			if tree.tree_string[pos.node.edge_label_start+pos.edge_pos+1] = tree.tree_string[str.end] {
				pos.edge_pos++
				chars_found = 1
			}
		}
	}	

	if chars_found == str.end - str.begin + 1 {
		rule_applied = 3

		if suffixless != 0 {
			create_suffix_link(sufixless, pos.node.father)
			suffixless = 0
		}
		return
	}

	if tree.is_last_char_in_edge(pos.node, pos.edge_pos) || pos.node == tree.root
	{
		if pos.node.sons != nil
		{
			apply_extension_rule_2(pos.node, str.begin+chars_found, str.end, path_pos, 0, new_son)
			rule_applied = 2
			if suffixless != 0 {
				create_suffix_link
		}