package structures

/*

			7
		3		5
	  1	  2   6   8
	9
*/

type node struct {
	val   int
	left  *node
	right *node
}

type Btree struct {
	root *node
}

func (t Btree) Create([]int) *Btree {

	return nil
}
