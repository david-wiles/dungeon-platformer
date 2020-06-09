// General utility functions that aren't specific to an interface or struct
// An attempt to make most of the code in this project reusable in other projects
package main

func (b *Bounds) GetDrawCenter() (float64, float64) {
	return b.X + (b.Width*global.gScale)/2, b.Y + (b.Height*global.gScale)/2
}
