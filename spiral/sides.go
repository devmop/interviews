package spiral

import "container/ring"

var top *ring.Ring
var right *ring.Ring
var bottom *ring.Ring
var left *ring.Ring

func init() {
	top = ring.New(4)
	right = top.Next()
	bottom = right.Next()
	left = bottom.Next()
}
