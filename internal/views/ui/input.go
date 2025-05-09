package ui

import (
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

// InputProps defines the properties for the Input component
type InputProps struct {
	CommonProps
	Type        string
	Placeholder string
	Name        string
	Id          string
	Required    bool
	Disabled    bool
	Value       string
	AutoFocus   bool
	Min         string
	Max         string
	Pattern     string
	Step        string
}

// InputBaseClasses defines the default styling classes for the Input component
const InputBaseClasses = "flex h-10 w-full rounded-md border border-input bg-transparent px-3 py-2 text-base shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 md:text-sm"

// Input creates an input element with the specified properties
func Input(props InputProps) g.Node {
	// Set a default type if none is provided
	inputType := props.Type
	if inputType == "" {
		inputType = "text"
	}

	attributes := []g.Node{
		html.Type(inputType),
		html.Class(mergeClasses(InputBaseClasses, props.Class)),
	}

	// Add optional attributes if they're set
	if props.Id != "" {
		attributes = append(attributes, html.ID(props.Id))
	}
	if props.Name != "" {
		attributes = append(attributes, html.Name(props.Name))
	}
	if props.Placeholder != "" {
		attributes = append(attributes, html.Placeholder(props.Placeholder))
	}
	if props.Value != "" {
		attributes = append(attributes, html.Value(props.Value))
	}
	if props.Required {
		attributes = append(attributes, html.Required())
	}
	if props.Disabled {
		attributes = append(attributes, html.Disabled())
	}
	if props.AutoFocus {
		attributes = append(attributes, html.AutoFocus())
	}
	if props.Min != "" {
		attributes = append(attributes, html.Min(props.Min))
	}
	if props.Max != "" {
		attributes = append(attributes, html.Max(props.Max))
	}
	if props.Pattern != "" {
		attributes = append(attributes, html.Pattern(props.Pattern))
	}
	if props.Step != "" {
		attributes = append(attributes, html.Step(props.Step))
	}

	return html.Input(attributes...)
}
