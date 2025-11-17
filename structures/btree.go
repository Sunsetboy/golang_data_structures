package structures

/*

			7
		3		5
	  1	  2   6   8
	9
*/

type node[T any] struct {
	val   T
	left  *node[T]
	right *node[T]
}

type Btree[T any] struct {
	root *node[T]
}

func NewBtree[T any](val T) Btree[T] {
	return Btree[T]{&node[T]{val: val}}
}

func CreateBtreeFromSlice[T any](elements []T) (*Btree[T], error) {
	var queue Queue[node[T]]

	root := node[T]{elements[0], nil, nil}
	queue.Add(root)
	currentIndex := 0
	totalElementsCount := len(elements)

	for {
		currentNode, err := queue.Pop()
		if err != nil {
			return nil, err
		}

		currentIndex++
		if currentIndex == totalElementsCount {
			break
		}
		left := node[T]{elements[currentIndex], nil, nil}
		queue.Add(left)

		currentIndex++
		if currentIndex == totalElementsCount {
			break
		}
		right := node[T]{elements[currentIndex], nil, nil}
		queue.Add(right)

		currentNode.left = &left
		currentNode.right = &right
	}

	return &Btree[T]{&root}, nil
}

func InOrderTraversal[T any](*node[T]) ([]T, error) {
	// left, root, right

	return nil, nil
}

func PreOrderTraversal[T any](*node[T]) ([]T, error) {
	// root, left, right

	return nil, nil
}

func PostOrderTraversal[T any](*node[T]) ([]T, error) {
	// left, right, root

	return nil, nil
}
