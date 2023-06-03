package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Fborder struct {
	top    fyne.CanvasObject
	bottom fyne.CanvasObject
	left   fyne.CanvasObject
	right  fyne.CanvasObject
}

// Returns a FluentBorder
func NewFborder() *Fborder {
	return &Fborder{
		nil,
		nil,
		nil,
		nil,
	}
}

// Sets the Top of the Border
func (b *Fborder) Top(c fyne.CanvasObject) *Fborder {
	b.top = c
	return b
}

// Sets the Bottom of the Border
func (b *Fborder) Bottom(c fyne.CanvasObject) *Fborder {
	b.bottom = c
	return b
}

// Sets the Left of the Border
func (b *Fborder) Left(c fyne.CanvasObject) *Fborder {
	b.left = c
	return b
}

// Sets the Right of the Border
func (b *Fborder) Right(c fyne.CanvasObject) *Fborder {
	b.right = c
	return b
}

// Gets the Border Container with additional content on the center
func (b *Fborder) Get(c ...fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(
		b.top,
		b.bottom,
		b.left,
		b.right,
		c...,
	)
}
