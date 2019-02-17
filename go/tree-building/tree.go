package tree

import (
	"errors"
	"sort"
)

// Record : Record is a struct that holds its ID and its parent's ID
type Record struct {
	ID     int
	Parent int
}

// Node : Node is a struct that holds its ID and a slice of pointers
// to its children
type Node struct {
	ID       int
	Children []*Node
}

// Build : Build inputs a slice of records and returns the root node
// of a tree
func Build(records []Record) (*Node, error) {
	// Sort the slice
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	rootExists := false
	var root Node

	for index, record := range records {
		// Check if records are valid
		switch {
		case record.ID == 0 && record.Parent != 0:
			return nil, errors.New("Root node has parent")
		case record.ID > 0 && record.ID <= record.Parent:
			return nil, errors.New("Node's ID is smaller than, or equal to parent's ID")
		case index+1 != len(records) && record.ID+1 != records[index+1].ID:
			return nil, errors.New("Nodes are not continuous")
		case record.ID == 0:
			rootExists = true
		default:
			if rootExists == false {
				return nil, errors.New("Root node does not exist")
			}

			// If the records are valid, start building the tree.
			// A breadth first search is done to find the parent node
			queue := []*Node{&root}

			for len(queue) > 0 {
				// If the parent node is found, add the child and break
				if record.Parent == queue[0].ID {
					newNode := Node{}
					newNode.ID = record.ID
					queue[0].Children = append(queue[0].Children, &newNode)

					queue = []*Node{}
					break
					// If parent is not found, add children to queue and move
					// to next item in the queue
				} else {
					queue = append(queue, queue[0].Children...)
					queue = queue[1:]
				}
			}
		}
	}
	if len(records) == 0 {
		return nil, nil
	}
	return &root, nil
}
