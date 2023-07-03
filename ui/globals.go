package ui

import "fyne.io/fyne/v2/data/binding"

// TODO:  this is a bit shit but works
var news, emails binding.UntypedList
var dateStr binding.String

// I made them globals to this package as Simulation needs to update the content of this page
