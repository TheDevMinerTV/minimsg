package minimsg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.minekube.com/common/minecraft/color"
	c "go.minekube.com/common/minecraft/component"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		output any
		desc   string
		input  string
	}{
		{
			desc:  "Colors - Hex",
			input: "<#ffffff>Foo</#ffffff> <color:#ffffff>Bar</color:#ffffff> <c:#ffffff>Baz</c:#ffffff>",
			output: &c.Text{
				Extra: []c.Component{
					&c.Text{
						Content: "Foo",
						S: c.Style{
							Color: &color.RGB{1, 1, 1},
						},
					},
					&c.Text{
						Content: "Bar",
						S: c.Style{
							Color: &color.RGB{1, 1, 1},
						},
					},
					&c.Text{
						Content: "Baz",
						S: c.Style{
							Color: &color.RGB{1, 1, 1},
						},
					},
				},
			},
		},
		{
			desc:  "Colors - Name",
			input: "<gray>Foo</gray> <color:gray>Bar</color:gray> <c:gray>Baz</c:gray>",
			output: &c.Text{
				Extra: []c.Component{
					&c.Text{
						Content: "Foo",
						S: c.Style{
							Color: color.Gray,
						},
					},
					&c.Text{
						Content: "Bar",
						S: c.Style{
							Color: color.Gray,
						},
					},
					&c.Text{
						Content: "Baz",
						S: c.Style{
							Color: color.Gray,
						},
					},
				},
			},
		},
		{
			desc:  "Styling - Bold",
			input: "<bold>Foo</bold> <b>Bar</b>",
			output: &c.Text{
				Extra: []c.Component{
					&c.Text{
						Content: "Foo",
						S: c.Style{
							Color: color.White,
							Bold:  c.True,
						},
					},
					&c.Text{
						Content: "Bar",
						S: c.Style{
							Color: color.White,
							Bold:  c.True,
						},
					},
				},
			},
		},
		{
			desc:  "Styling - Underline",
			input: "<underline>Foo</underline> <underlined>Bar</underlined> <u>Baz</u>",
			output: &c.Text{
				Extra: []c.Component{
					&c.Text{
						Content: "Foo",
						S: c.Style{
							Color:      color.White,
							Underlined: c.True,
						},
					},
					&c.Text{
						Content: "Bar",
						S: c.Style{
							Color:      color.White,
							Underlined: c.True,
						},
					},
					&c.Text{
						Content: "Baz",
						S: c.Style{
							Color:      color.White,
							Underlined: c.True,
						},
					},
				},
			},
		},
		{
			desc:  "Styling - Strikethrough",
			input: "<strikethrough>Foo</strikethrough> <st>Bar</st>",
			output: &c.Text{
				Extra: []c.Component{
					&c.Text{
						Content: "Foo",
						S: c.Style{
							Color:         color.White,
							Strikethrough: c.True,
						},
					},
					&c.Text{
						Content: "Bar",
						S: c.Style{
							Color:         color.White,
							Strikethrough: c.True,
						},
					},
				},
			},
		},
		{
			desc:  "Styling - Obfuscated",
			input: "<obfuscated>Foo</obfuscated> <obfuscate>Bar</obfuscate> <obf>Baz</obf>",
			output: &c.Text{
				Extra: []c.Component{
					&c.Text{
						Content: "Foo",
						S: c.Style{
							Color:      color.White,
							Obfuscated: c.True,
						},
					},
					&c.Text{
						Content: "Bar",
						S: c.Style{
							Color:      color.White,
							Obfuscated: c.True,
						},
					},
					&c.Text{
						Content: "Baz",
						S: c.Style{
							Color:      color.White,
							Obfuscated: c.True,
						},
					},
				},
			},
		},
		{
			desc:  "Gradient - Hex",
			input: "<gradient:#ffffff:#000000>Foo</gradient>",
			output: &c.Text{
				Extra: []c.Component{
					&c.Text{
						Content: "",
						S:       c.Style{},
						Extra: []c.Component{
							&c.Text{
								Content: "F",
								S: c.Style{
									Color: &color.RGB{1, 1, 1},
								},
							},
							&c.Text{
								Content: "o",
								S: c.Style{
									Color: &color.RGB{0.6666666666666666, 0.6666666666666666, 0.6666666666666666},
								},
							},
							&c.Text{
								Content: "o",
								S: c.Style{
									Color: &color.RGB{0.3333333333333333, 0.3333333333333333, 0.3333333333333333},
								},
							},
						},
					},
				},
			},
		},
		{
			desc:  "Gradient - By Name",
			input: "<gradient:white:black>Foo</gradient>",
			output: &c.Text{
				Extra: []c.Component{
					&c.Text{
						Content: "",
						S:       c.Style{},
						Extra: []c.Component{
							&c.Text{
								Content: "F",
								S: c.Style{
									Color: &color.RGB{1, 1, 1},
								},
							},
							&c.Text{
								Content: "o",
								S: c.Style{
									Color: &color.RGB{0.6666666666666666, 0.6666666666666666, 0.6666666666666666},
								},
							},
							&c.Text{
								Content: "o",
								S: c.Style{
									Color: &color.RGB{0.3333333333333333, 0.3333333333333333, 0.3333333333333333},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			parsed := Parse(tC.input)

			assert.Equal(t, tC.output, parsed)
		})
	}
}
