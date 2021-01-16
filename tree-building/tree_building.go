package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type nodeQueue []*Node

type nodeStack []*Node

func (nq *nodeQueue) dequeue() *Node {
	var retNode *Node
	n := len(*nq)

	if n >= 1 {
		retNode, *nq = ([]*Node)(*nq)[0], ([]*Node)(*nq)[1:]
	} else if n == 1 {
		retNode, *nq = ([]*Node)(*nq)[0], nodeQueue{}
	} else {
		retNode = nil
	}

	return retNode
}

func (nq *nodeQueue) enqueue(n *Node) {
	if *nq != nil {
		*nq = append([]*Node{n}, (([]*Node)(*nq))...)
	} else {
		*nq = nodeQueue{n}
	}
}

func (ns *nodeStack) push(n *Node) {
	if *ns != nil {
		*ns = append(*ns, n)
	} else {
		*ns = nodeStack{n}
	}
}

func (ns *nodeStack) isEmpty() bool {
	return len(*ns) == 0
}

func (ns *nodeStack) pop() *Node {
	var ret *Node

	ret, *ns = (([]*Node)(*ns))[len(*ns)-1], (*ns)[:len(*ns)-1]
	return ret
}

func removeRecord(records *[]Record, i int) {
	n := len(*records)
	if n > 1 {
		(*records)[i] = (*records)[len(*records)-1]
		(*records)[len(*records)-1] = Record{}
		*records = (*records)[:len(*records)-1]
	} else {
		*records = (*records)[:0]
	}
}

func addChild(parent *Node, n *Node) {
	parent.Children = append((*parent).Children, n)
}

func isQueueEmpty(nq *nodeQueue) bool {
	return len(*nq) == 0
}

func createNode(r Record) *Node {
	return &Node{ID: r.ID}
}

func startTree(records *[]Record) (*Node, error) {
	/* Look for the root record and remove it from records. */
	if len(*records) == 0 {
		return nil, nil
	}

	for i, record := range *records {
		if record.ID == 0 && record.Parent == 0 {
			removeRecord(records, i)
			return &Node{ID: 0, Children: nil}, nil
		}
	}

	return nil, errors.New("A valid root node must exist (no parent).")
}

func orderRecords(records *[]Record) {
	sort.Slice(*records, func(i, j int) bool {
		return (*records)[i].ID < (*records)[j].ID
	})
}

func indirectCycleErr(root *Node) error {
	var ns nodeStack
	var idMap map[int]bool
	idMap = make(map[int]bool)

	/* Visit root, add its ID to idMap, for each child push onto stack, pop one
	child from stack to visit, add child ID to map, recurisvely push and pop checking
	if one of the child ID's have been seen in the map. */
	ns.push(root)

	for !ns.isEmpty() {
		currNode := ns.pop()

		if idMap[currNode.ID] {
			return errors.New("Indirect cycle detected!")
		}

		idMap[currNode.ID] = true

		for _, child := range currNode.Children {
			ns.push(child)
		}
	}

	return nil
}

func buildTree(root *Node, records *[]Record) (*Node, error) {
	var frontier *nodeQueue = &nodeQueue{}
	frontier.enqueue(root)

	for !isQueueEmpty(frontier) {
		currentParent := frontier.dequeue()

		orderRecords(records)

		var recordsKeep []Record

		for _, record := range *records {
			if record.Parent == currentParent.ID {
				newNode := createNode(record)
				addChild(currentParent, newNode)
				frontier.enqueue(newNode)
			} else {
				recordsKeep = append(recordsKeep, record)
			}
		}

		if len(*records) > 0 {
			indirectCycleErr := indirectCycleErr(root)

			if indirectCycleErr != nil {
				return nil, indirectCycleErr
			}
		}

		*records = recordsKeep
	}

	if len(*records) > 0 {
		return nil, errors.New("There are records with non-existent parents.")
	}

	return root, nil
}

func duplicateErr(records []Record) error {
	recordMap := map[int]bool{}

	for _, r := range records {
		if !recordMap[r.ID] {
			recordMap[r.ID] = true
		} else {
			return errors.New("Duplicate records found.")
		}
	}

	return nil
}

func continunityErr(records []Record) error {
	orderRecords(&records)

	n := len(records)

	if n > 1 {
		for i, j := 0, 1; j <= len(records)-1; {
			if records[j].ID != records[i].ID+1 {
				return errors.New("Records are non-continuous.")
			}
			i++
			j++
		}
	}

	return nil
}

func directCycleErr(records []Record) error {
	for _, r := range records {
		if r.ID == r.Parent {
			return errors.New("Direct cycle detected.")
		}
	}

	return nil
}

func parentHigherThanChildErr(records []Record) error {
	for _, r := range records {
		if r.Parent > r.ID {
			return errors.New("Parent ID higher than ID.")
		}
	}

	return nil
}

func Build(records []Record) (*Node, error) {
	dupErr := duplicateErr(records)

	if dupErr != nil {
		return nil, dupErr
	}

	continunityErr := continunityErr(records)

	if continunityErr != nil {
		return nil, continunityErr
	}

	parentHigherThanChildErr := parentHigherThanChildErr(records)

	if parentHigherThanChildErr != nil {
		return nil, parentHigherThanChildErr
	}

	/* Declare tree. */
	var tree *Node

	root, startErr := startTree(&records)

	if startErr != nil {
		return nil, startErr
	}

	tree, buildErr := buildTree(root, &records)

	if buildErr != nil {
		return nil, buildErr
	}

	return tree, nil
}
