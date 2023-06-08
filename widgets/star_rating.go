package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// TODO: move those to themed icons maybe?
var star = &fyne.StaticResource{
	StaticName: "star",
	StaticContent: []byte(
		"<?xml version=\"1.0\" encoding=\"utf-8\"?><svg width=\"800px\" height=\"800px\" viewBox=\"2 2 22 22\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\">\r\n<path d=\"M17.2 20.7501C17.0776 20.7499 16.9573 20.7189 16.85 20.6601L12 18.1101L7.14999 20.6601C7.02675 20.7262 6.88746 20.7566 6.74786 20.7478C6.60825 20.7389 6.47391 20.6912 6.35999 20.6101C6.24625 20.5267 6.15796 20.4133 6.10497 20.2826C6.05199 20.1519 6.03642 20.0091 6.05999 19.8701L6.99999 14.4701L3.05999 10.6501C2.96124 10.5512 2.89207 10.4268 2.86027 10.2907C2.82846 10.1547 2.83529 10.0124 2.87999 9.88005C2.92186 9.74719 3.00038 9.62884 3.10652 9.53862C3.21266 9.4484 3.34211 9.38997 3.47999 9.37005L8.89999 8.58005L11.33 3.67005C11.3991 3.55403 11.4973 3.45795 11.6147 3.39123C11.7322 3.32451 11.8649 3.28943 12 3.28943C12.1351 3.28943 12.2678 3.32451 12.3853 3.39123C12.5027 3.45795 12.6008 3.55403 12.67 3.67005L15.1 8.58005L20.52 9.37005C20.6579 9.38997 20.7873 9.4484 20.8935 9.53862C20.9996 9.62884 21.0781 9.74719 21.12 9.88005C21.1647 10.0124 21.1715 10.1547 21.1397 10.2907C21.1079 10.4268 21.0387 10.5512 20.94 10.6501L17 14.4701L17.93 19.8701C17.9536 20.0091 17.938 20.1519 17.885 20.2826C17.832 20.4133 17.7437 20.5267 17.63 20.6101C17.5034 20.6976 17.3539 20.7463 17.2 20.7501ZM12 16.5201C12.121 16.5215 12.2403 16.5488 12.35 16.6001L16.2 18.6001L15.47 14.3101C15.4502 14.1897 15.4589 14.0664 15.4953 13.9501C15.5318 13.8337 15.595 13.7275 15.68 13.6401L18.8 10.6401L14.49 10.0001C14.3708 9.98109 14.2578 9.93401 14.1605 9.86271C14.0631 9.79141 13.9841 9.69795 13.93 9.59005L12 5.69005L10.07 9.60005C10.0159 9.70795 9.9369 9.80141 9.83952 9.87271C9.74214 9.94401 9.62918 9.99109 9.50999 10.0101L5.19999 10.6401L8.31999 13.6401C8.40493 13.7275 8.46817 13.8337 8.50464 13.9501C8.54111 14.0664 8.54979 14.1897 8.52999 14.3101L7.79999 18.6301L11.65 16.6301C11.7573 16.5683 11.8767 16.5308 12 16.5201Z\" fill=\"#000000\"/>\r\n</svg>",
	),
}
var starHalf = &fyne.StaticResource{
	StaticName: "star_half",
	StaticContent: []byte(
		"<?xml version=\"1.0\" encoding=\"utf-8\"?><svg fill=\"#000000\" width=\"800px\" height=\"800px\" viewBox=\"0 0 36 36\" version=\"1.1\"  preserveAspectRatio=\"xMidYMid meet\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"><title>half-star-solid</title><path class=\"clr-i-solid clr-i-solid-path-1\" d=\"M34,16.78a2.22,2.22,0,0,0-1.29-4l-9-.34a.23.23,0,0,1-.2-.15L20.4,3.89a2.22,2.22,0,0,0-4.17,0l-3.1,8.43a.23.23,0,0,1-.2.15l-9,.34a2.22,2.22,0,0,0-1.29,4l7.06,5.55a.23.23,0,0,1,.08.24L7.35,31.21a2.22,2.22,0,0,0,3.38,2.45l7.46-5a.22.22,0,0,1,.25,0l7.46,5a2.2,2.2,0,0,0,2.55,0,2.2,2.2,0,0,0,.83-2.4l-2.45-8.64a.22.22,0,0,1,.08-.24ZM24.9,23.11l2.45,8.64A.22.22,0,0,1,27,32l-7.46-5a2.21,2.21,0,0,0-1.24-.38h0V4.44h0a.2.2,0,0,1,.21.15L21.62,13a2.22,2.22,0,0,0,2,1.46l9,.34a.22.22,0,0,1,.13.4l-7.06,5.55A2.21,2.21,0,0,0,24.9,23.11Z\"></path></svg>",
	),
}
var starFull = &fyne.StaticResource{
	StaticName: "star_full",
	StaticContent: []byte(
		"<?xml version=\"1.0\" encoding=\"utf-8\"?><svg width=\"800px\" height=\"800px\" viewBox=\"2 2 22 22\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\">\r\n<path d=\"M21.12 9.88005C21.0781 9.74719 20.9996 9.62884 20.8935 9.53862C20.7873 9.4484 20.6579 9.38997 20.52 9.37005L15.1 8.58005L12.67 3.67005C12.6008 3.55403 12.5027 3.45795 12.3853 3.39123C12.2678 3.32451 12.1351 3.28943 12 3.28943C11.8649 3.28943 11.7322 3.32451 11.6147 3.39123C11.4973 3.45795 11.3991 3.55403 11.33 3.67005L8.89999 8.58005L3.47999 9.37005C3.34211 9.38997 3.21266 9.4484 3.10652 9.53862C3.00038 9.62884 2.92186 9.74719 2.87999 9.88005C2.83529 10.0124 2.82846 10.1547 2.86027 10.2907C2.89207 10.4268 2.96124 10.5512 3.05999 10.6501L6.99999 14.4701L6.06999 19.8701C6.04642 20.0091 6.06199 20.1519 6.11497 20.2826C6.16796 20.4133 6.25625 20.5267 6.36999 20.6101C6.48391 20.6912 6.61825 20.7389 6.75785 20.7478C6.89746 20.7566 7.03675 20.7262 7.15999 20.6601L12 18.1101L16.85 20.6601C16.9573 20.7189 17.0776 20.7499 17.2 20.7501C17.3573 20.7482 17.5105 20.6995 17.64 20.6101C17.7537 20.5267 17.842 20.4133 17.895 20.2826C17.948 20.1519 17.9636 20.0091 17.94 19.8701L17 14.4701L20.93 10.6501C21.0305 10.5523 21.1015 10.4283 21.1351 10.2922C21.1687 10.1561 21.1634 10.0133 21.12 9.88005Z\" fill=\"#000000\"/>\r\n</svg>",
	),
}

type StarRating struct {
	widget.BaseWidget
	full      int
	half      bool
	container *fyne.Container
}

func NewStarRating(stars int) *StarRating {
	if stars > 5 {
		stars = 5
	}

	if stars < 0 {
		stars = 0
	}
	s := &StarRating{full: stars}
	s.ExtendBaseWidget(s)
	return s
}

func NewStarRatingFromFloat(rating float32) *StarRating {
	full := int(rating)
	mantissa := rating - float32(full)
	half := false
	if mantissa >= .5 {
		half = true
	}
	return NewStarRatingWithHalf(full, half)
}

func NewStarRatingWithHalf(stars int, half bool) *StarRating {
	if stars >= 5 {
		half = false
		stars = 5
	}
	s := &StarRating{full: stars, half: half}
	s.ExtendBaseWidget(s)
	return s
}

func (s *StarRating) Refresh() {
	s.updateContainer()
	s.BaseWidget.Refresh()
}

func (s *StarRating) SetValues(rating float32) {
	full := int(rating)
	mantissa := rating - float32(full)
	half := false
	if mantissa >= .5 {
		half = true
	}
	s.full = full
	s.half = half
	s.Refresh()
}

func (s *StarRating) updateContainer() {
	full := s.full
	halves := 0
	if s.half {
		halves = 1
	}
	empty := 5 - (halves + s.full)
	if empty < 0 {
		empty = 0
	}
	if s.container == nil {
		s.container = container.NewHBox()
	} else {
		s.container.RemoveAll()
	}
	for i := 0; i < full; i++ {
		s.container.Add(widget.NewIcon(theme.NewThemedResource(starFull)))
	}
	if s.half {
		s.container.Add(widget.NewIcon(theme.NewThemedResource(starHalf)))
	}
	for i := 0; i < empty; i++ {
		s.container.Add(widget.NewIcon(theme.NewThemedResource(star)))
	}
}

func (s *StarRating) CreateRenderer() fyne.WidgetRenderer {
	s.updateContainer()
	return widget.NewSimpleRenderer(
		container.NewCenter(s.container),
	)
}
