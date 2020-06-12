package main

// QuadTree data structure for collision detection
// Based on the quadtree implementation by Magnus Lallassu
// https://github.com/Lallassu/gizmo/blob/master/quadtree.go

type QuadTree struct {
	level   int        // The level of this node
	box     Bounds     // The Bounds of this quadtree in world space
	objects []*Bounds  // Objects located within the Bounds of this quadtree node
	nodes   []QuadTree // Child nodes
	count   int        // The number of objects in this tree and its children
}

// Representation of an entity's hit box
type Bounds struct {
	X, Y, Width, Height float64
	entity              Entity
}

func (b *Bounds) IsPoint() bool {
	return b.Height == 0 && b.Width == 0
}

// TODO implement collisions separate from drawing to reduce multiplication
func (b *Bounds) Intersects(cmp *Bounds) bool {
	thisMaxX := b.X + b.Width
	thisMaxY := b.Y + b.Height
	cmpMaxX := cmp.X + cmp.Width
	cmpMaxY := cmp.Y + cmp.Height

	if thisMaxX < cmp.X || b.X > cmpMaxX || thisMaxY < cmp.Y || b.Y > cmpMaxY {
		return false
	}

	return true
}

// Creates an instance of a new quad tree, returns the root node
func MakeQuadTree(w, h float64) QuadTree {
	return QuadTree{
		level: 0,
		box: Bounds{
			X:      0,
			Y:      0,
			Width:  w,
			Height: h,
		},
		objects: make([]*Bounds, 0),
		nodes:   make([]QuadTree, 0, 4),
		count:   0,
	}
}

// Add an object into the QuadTree
// If this node is a leaf, we can try to add it to the list of objects
// If the target node is full, we must split the node into four new nodes and distribute objects before adding
func (q *QuadTree) Insert(b *Bounds) {
	if len(q.nodes) > 0 {
		index := q.getIndex(b)
		if index != -1 {
			q.nodes[index].Insert(b)
			q.count += 1
			return
		}
		// TODO handle large objects
	}

	q.objects = append(q.objects, b)
	q.count += 1

	if len(q.objects) > qtMaxObjects && q.level < qtMaxLevels && len(q.nodes) == 0 {
		q.redistribute()
	}
}

func (q *QuadTree) Remove(b *Bounds) {
	q.removeAux(q, b)
}

// Get all boxes that will intersect with a given box
func (q *QuadTree) GetIntersections(b *Bounds) []*Bounds {
	var found []*Bounds
	potential := q.retrieve(b)
	for _, p := range potential {
		if p.Intersects(b) {
			found = append(found, p)
		}
	}
	return found
}

// Remove all objects from the tree
func (q *QuadTree) Clear() {
	q.objects = []*Bounds{}

	if len(q.nodes)-1 > 0 {
		for i := 0; i < len(q.nodes); i += 1 {
			q.nodes[i].Clear()
		}
	}

	q.nodes = []QuadTree{}
}

// Recursively remove the box from the QuadTree node.  If it was removed, return true
func (q *QuadTree) removeAux(qt *QuadTree, b *Bounds) bool {
	index := -1
	// Find index of the object
	for i := range qt.objects {
		if qt.objects[i] == b {
			index = i
			break
		}
	}
	// If the object is in this node, remove it
	if index != -1 {
		qt.objects[index] = nil
		qt.objects = append(qt.objects[:index], qt.objects[index+1:]...)
		return true
	}
	for i := range qt.nodes {
		if q.removeAux(&qt.nodes[i], b) == true {
			q.count -= 1
			break
		}
	}

	return false
}

// Get the index of the QuadTree node that this box should be added to.  Returns a number 0 - 3
func (q *QuadTree) getIndex(b *Bounds) int {
	// Get midpoint for this node's box.  This corresponds to the boundary of each child node
	vMidPoint := q.box.X + q.box.Width/2
	hMidPoint := q.box.Y + q.box.Height/2

	// Determine if the box can fit entirely within a hemisphere
	top := b.Y < hMidPoint && b.Y+b.Height < hMidPoint
	bot := b.Y > hMidPoint
	left := b.X < vMidPoint && b.X+b.Width < vMidPoint
	right := b.X > vMidPoint

	if top {
		if right {
			return 0
		} else if left {
			return 1
		}
	} else if bot {
		if left {
			return 2
		} else if right {
			return 3
		}
	}

	// If -1 is returned, this object does not fix into any box and therefore it should be added to the current node
	return -1
}

// Initialize new nodes
func (q *QuadTree) split() {
	if len(q.nodes) != 0 {
		// oops
		return
	}

	nextLevel := q.level + 1
	width := q.box.Width / 2
	height := q.box.Height / 2

	// top right
	q.nodes = append(q.nodes, QuadTree{
		level: nextLevel,
		box: Bounds{
			X:      q.box.X + width,
			Y:      q.box.Y + height,
			Width:  width,
			Height: height,
		},
		objects: make([]*Bounds, 0),
		nodes:   make([]QuadTree, 0, 4),
		count:   0,
	})
	// top left
	q.nodes = append(q.nodes, QuadTree{
		level: nextLevel,
		box: Bounds{
			X:      q.box.X,
			Y:      q.box.Y + height,
			Width:  width,
			Height: height,
		},
		objects: make([]*Bounds, 0),
		nodes:   make([]QuadTree, 0, 4),
		count:   0,
	})
	// bottom left
	q.nodes = append(q.nodes, QuadTree{
		level: nextLevel,
		box: Bounds{
			X:      q.box.X,
			Y:      q.box.Y,
			Width:  width,
			Height: height,
		},
		objects: make([]*Bounds, 0),
		nodes:   make([]QuadTree, 0, 4),
		count:   0,
	})
	// bottom right
	q.nodes = append(q.nodes, QuadTree{
		level: nextLevel,
		box: Bounds{
			X:      q.box.X + width,
			Y:      q.box.Y,
			Width:  width,
			Height: height,
		},
		objects: make([]*Bounds, 0),
		nodes:   make([]QuadTree, 0, 4),
		count:   0,
	})
}

// Redistribute the objects in the current node into its child nodes
// If the object can't fit into a single node, then it is left in its current node
func (q *QuadTree) redistribute() {
	q.split()
	for i := 0; i < len(q.objects); i += 1 {
		idx := q.getIndex(q.objects[i])
		if idx != -1 {
			box := q.objects[i]
			q.objects = append(q.objects[:i], q.objects[i+1:]...)
			q.nodes[idx].Insert(box)
		}
	}
}

// Get all boxes that could collide with the given box.
// This can be defined as all boxes within a single quadrant or multiple quadrants
// Returns all boxes in a quadrant, or multiple quadrants
func (q *QuadTree) retrieve(b *Bounds) []*Bounds {
	index := q.getIndex(b)
	objs := q.objects

	if len(q.nodes) > 0 {
		if index != -1 {
			objs = append(objs, q.nodes[index].retrieve(b)...)
		} else {
			// If the box doesn't fit into a single node's box, add objects from all quadrants
			for i := 0; i < len(q.nodes); i += 1 {
				objs = append(objs, q.nodes[i].retrieve(b)...)
			}
		}
	}

	return objs
}

// Get all point boxes that do collide with a given box
func (q *QuadTree) retrievePoints(find *Bounds) []*Bounds {
	var found []*Bounds
	potential := q.retrieve(find)
	for _, box := range potential {
		if box.X == float64(find.X) && box.Y == float64(find.Y) && box.IsPoint() {
			found = append(found, find)
		}
	}
	return found
}
