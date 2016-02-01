package spiral

//Unravel traverses the elements of the matrix in a clockwise fashion from the outer edge.
func Unravel(matrix [][]uint) []uint {
	res := make([]uint, 0, len(matrix)*len(matrix[0]))

	bounds := &box{top: 0, bottom: len(matrix) - 1, left: 0, right: len(matrix[0]) - 1}

	side := top

	for bounds.hasContents() {
		switch side {
		case top:
			for i := bounds.left; i <= bounds.right; i++ {
				res = append(res, matrix[bounds.top][i])
			}
			bounds.top++
		case right:
			for i := bounds.top; i <= bounds.bottom; i++ {
				res = append(res, matrix[i][bounds.right])
			}
			bounds.right--
		case bottom:
			for i := bounds.right; i >= bounds.left; i-- {
				res = append(res, matrix[bounds.bottom][i])
			}
			bounds.bottom--
		case left:
			for i := bounds.bottom; i >= bounds.top; i-- {
				res = append(res, matrix[i][bounds.left])
			}
			bounds.left++
		}

		side = side.Next()
	}

	return res
}

type box struct {
	top, right, bottom, left int
}

func (b *box) hasContents() bool {
	return b.top <= b.bottom && b.left <= b.right
}
