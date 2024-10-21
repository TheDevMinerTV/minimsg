// Package mini provides utilities for parsing and manipulating Minecraft text colors and styles.
// It includes functions for parsing strings with embedded style information, modifying styles,
// and creating gradient effects. It also provides functions for parsing color names and hex codes,
// and for linear interpolation of colors.
//
// Credits to the partial Go port of MiniMessage (https://docs.advntr.dev/minimessage/index.html) by
// https://github.com/emortalmc/GateProxy/blob/main/minimessage/minimessage.go.
package minimsg

import (
	"fmt"
	"math"
	"strings"

	"go.minekube.com/common/minecraft/color"
	c "go.minekube.com/common/minecraft/component"
)

// Parse takes a string as input and returns a `c.Text` object. It splits the input string by "<",
// then further splits each substring by ">". It modifies the style based on the key (the part before ">")
// and appends a new text component with the modified style and content (the part after ">").
func Parse(mini string) *c.Text {
	styleStack := []c.Style{{Color: color.White}}

	var components []c.Component

	// Check if the string has any < characters before splitting
	if !strings.Contains(mini, "<") {
		// If no < character, treat it as plain text
		return &c.Text{
			Content: mini,
			S:       styleStack[0],
		}
	}

	for _, s := range strings.Split(mini, "<") {
		if s == "" {
			continue
		}

		split := strings.Split(s, ">")

		// Ensure split has at least two elements before accessing split[1]
		if len(split) < 2 {
			fmt.Println("Warning: missing closing '>' in mini string:", s)
			continue
		}

		key := split[0]
		if strings.HasPrefix(key, "/") {
			// Pop last style from the stack
			styleStack = styleStack[:len(styleStack)-1]
			key = key[1:]
		} else {
			// Get last style to modify
			newStyle := styleStack[len(styleStack)-1]

			// Push new style onto stack
			styleStack = append(styleStack, newStyle)

			newText := modify(key, split[1], &styleStack[len(styleStack)-1])
			if newText != nil { // Add a check to avoid nil components
				components = append(components, newText)
			}
		}
	}

	return &c.Text{
		Extra: components,
	}
}

// modify takes a key, content, and style as input and returns a `c.Text` object. It modifies the style
// based on the key and returns a new text component with the modified style and content.
func modify(key string, content string, style *c.Style) *c.Text {
	newText := &c.Text{}

	switch {
	case strings.HasPrefix(key, "gradient"): // <gradient:light_purple:gold>
		colorKey := strings.Split(key, ":")
		colorNames := colorKey[1:]

		colors := make([]color.RGB, len(colorNames))
		for i, col := range colorNames {
			parsedColor, err := ParseColor(col)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			newColor, _ := color.Make(parsedColor)
			colors[i] = *newColor
		}

		newText = Gradient(content, *style, colors...)

	case strings.HasPrefix(key, "color:") || strings.HasPrefix(key, "colour:") || strings.HasPrefix(key, "c:"): // <color:light_purple>, <colour:light_purple>, <c:light_purple>
		key = strings.Split(key, ":")[1]
		fallthrough
	default:
		switch key {
		case "b":
			fallthrough
		case "bold":
			style.Bold = c.True

		case "i":
			fallthrough
		case "em":
			fallthrough
		case "italic":
			style.Italic = c.True

		case "u":
			fallthrough
		case "underline":
			fallthrough
		case "underlined":
			style.Underlined = c.True

		case "st":
			fallthrough
		case "strikethrough":
			style.Strikethrough = c.True

		case "obf":
			fallthrough
		case "obfuscate":
			fallthrough
		case "obfuscated":
			style.Obfuscated = c.True

		default: // <light_purple>
			parsed, err := ParseColor(key)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			style.Color = parsed
		}

		newText.Content = content
		newText.S = *style
	}

	return newText
}

// ParseColor takes a string as input and returns a `color.Color` object. It checks if the input string
// starts with "#". If it does, it tries to parse it as a hex color. If it doesn't, it tries to find a
// named color that matches the input string.
func ParseColor(name string) (color.Color, error) {
	if strings.HasPrefix(name, "#") {
		newColor, err := color.Hex(name)
		if err != nil {
			return nil, err
		}
		return newColor, nil
	} else {
		return FromName(name)
	}
}

// FromName takes a string as input and returns a `color.Color` object.
// It iterates over the named colors and returns the one that matches the input string.
func FromName(name string) (color.Color, error) {
	col, ok := color.Names[name]
	if ok {
		return col, nil
	}
	for _, a := range color.Names {
		if strings.EqualFold(a.String(), name) {
			return a, nil
		}
	}
	return nil, fmt.Errorf("unknown color name: %s", name)
}

// Gradient takes a string, a style, and a variable number of colors as input and returns a `c.Text` object.
// It creates a gradient effect by interpolating between the input colors based on their position in the input string.
func Gradient(content string, style c.Style, colors ...color.RGB) *c.Text {
	var component []c.Component
	for id, i := range strings.Split(content, "") {
		t := float64(id) / float64(len(content))
		hex, _ := color.Hex(LerpColor(t, colors...).Hex())

		style.Color = hex
		component = append(component, &c.Text{
			Content: string(i),
			S:       style,
		})
	}

	return &c.Text{
		Extra: component,
	}
}

// LerpColor takes a float and a variable number of colors as input and returns a `color.Color` object.
// It interpolates between the input colors based on the input float.
func LerpColor(t float64, colors ...color.RGB) color.Color {
	t = math.Min(t, 1)

	if t == 1 {
		return &colors[len(colors)-1]
	}

	colorT := t * float64(len(colors)-1)
	newT := colorT - math.Floor(colorT)
	lastColor := colors[int(colorT)]
	nextColor := colors[int(colorT+1)]

	return &color.RGB{
		R: lerpInt(newT, nextColor.R, lastColor.R),
		G: lerpInt(newT, nextColor.G, lastColor.G),
		B: lerpInt(newT, nextColor.B, lastColor.B),
	}
}

// lerpInt takes three floats as input and returns a float. It performs linear interpolation between the
// second and third input floats based on the first input float.
func lerpInt(t float64, a float64, b float64) float64 {
	return a*t + b*(1-t)
}
