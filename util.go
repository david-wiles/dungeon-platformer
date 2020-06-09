// General utility functions that aren't specific to an interface or struct
// An attempt to make most of the code in this project reusable in other projects
package main

// Some convenience functions for working with bounds
func (b *Bounds) Top() float64    { return b.Y + (b.Height * global.gScale) }
func (b *Bounds) Bottom() float64 { return b.Y }
func (b *Bounds) Left() float64   { return b.X }
func (b *Bounds) Right() float64  { return b.X + (b.Width * global.gScale) }
func (b *Bounds) GetDrawCenter() (float64, float64) {
	return b.X + (b.Width*global.gScale)/2, b.Y + (b.Height*global.gScale)/2
}
