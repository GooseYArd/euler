package stree

import (
	"testing"
//	"fmt"
)

func TestCreateTree(t *testing.T) {	
	tree, err := CreateTree("mmm")
	if err != nil {
		t.Errorf("tree constructor returned non-nil")
	}
	tree.PrintTree()
	
}