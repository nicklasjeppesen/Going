package view

//
//------------------------------------------------------------------------
// 					View Helper
//------------------------------------------------------------------------
//
// This View Helper is responsible for providing custom functionality to the view templates.
// It allows you to define custom functions that can be used within your templates to manipulate data or perform specific tasks.
//
//

var GetCustomViewFunction = map[string]any{

	// name: func
	"example": func(exampleInput string) any {
		return "example"
	},
}
