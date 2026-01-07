package cmd

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var colorCmd = &cobra.Command{
	Use:   "color <in-format> <value> [out-format]",
	Short: "Convert between color formats and display colors",
	Args:  cobra.RangeArgs(2, 3),
	RunE: func(cmd *cobra.Command, args []string) error {
		inFormat := args[0]
		val := args[1]
		outFormat := ""
		if len(args) == 3 {
			outFormat = args[2]
		}

		var r, g, b int
		var err error

		switch inFormat {
		case "hex":
			r, g, b, err = parseHex(val)
		case "rgb":
			r, g, b, err = parseRGB(val)
		case "hsl":
			r, g, b, err = parseHSL(val)
		default:
			return fmt.Errorf("unsupported input format: %s (use hex, rgb, hsl)", inFormat)
		}

		if err != nil {
			fmt.Println(errStyle.Render("Error parsing input:"), err)
			return nil
		}

		// Create color display block
		colorBlock := lipgloss.NewStyle().
			Background(lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", r, g, b))).
			Render("     ")

		h, s, l := rgbToHSL(r, g, b)

		outputs := map[string]func(){
			"hex": func() {
				fmt.Printf("%s %s %s\n", labelStyle.Render("Hex:        "), valueStyle.Render(fmt.Sprintf("#%02x%02x%02x", r, g, b)), colorBlock)
			},
			"rgb": func() {
				fmt.Printf("%s %s %s\n", labelStyle.Render("RGB:        "), valueStyle.Render(fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)), colorBlock)
			},
			"hsl": func() {
				fmt.Printf("%s %s %s\n", labelStyle.Render("HSL:        "), valueStyle.Render(fmt.Sprintf("hsl(%d, %d%%, %d%%)", h, s, l)), colorBlock)
			},
		}

		if outFormat != "" {
			if fn, ok := outputs[outFormat]; ok {
				fn()
			} else {
				fmt.Println(errStyle.Render("Unknown output format:"), outFormat)
			}
		} else {
			for _, fn := range outputs {
				fn()
			}
		}
		return nil
	},
}

func parseHex(hex string) (int, int, int, error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) == 3 {
		hex = string(hex[0]) + string(hex[0]) + string(hex[1]) + string(hex[1]) + string(hex[2]) + string(hex[2])
	}
	if len(hex) != 6 {
		return 0, 0, 0, fmt.Errorf("invalid hex color format")
	}

	r, err := strconv.ParseInt(hex[0:2], 16, 0)
	if err != nil {
		return 0, 0, 0, err
	}
	g, err := strconv.ParseInt(hex[2:4], 16, 0)
	if err != nil {
		return 0, 0, 0, err
	}
	b, err := strconv.ParseInt(hex[4:6], 16, 0)
	if err != nil {
		return 0, 0, 0, err
	}

	return int(r), int(g), int(b), nil
}

func parseRGB(rgb string) (int, int, int, error) {
	rgb = strings.TrimPrefix(rgb, "rgb(")
	rgb = strings.TrimSuffix(rgb, ")")
	parts := strings.Split(rgb, ",")
	if len(parts) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid RGB format, use rgb(r,g,b)")
	}

	r, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, 0, err
	}
	g, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, 0, err
	}
	b, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil {
		return 0, 0, 0, err
	}

	return r, g, b, nil
}

func parseHSL(hsl string) (int, int, int, error) {
	hsl = strings.TrimPrefix(hsl, "hsl(")
	hsl = strings.TrimSuffix(hsl, ")")
	parts := strings.Split(hsl, ",")
	if len(parts) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid HSL format, use hsl(h,s%%,l%%)")
	}

	h, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return 0, 0, 0, err
	}
	s, err := strconv.ParseFloat(strings.TrimSpace(strings.TrimSuffix(parts[1], "%")), 64)
	if err != nil {
		return 0, 0, 0, err
	}
	l, err := strconv.ParseFloat(strings.TrimSpace(strings.TrimSuffix(parts[2], "%")), 64)
	if err != nil {
		return 0, 0, 0, err
	}

	r, g, b := hslToRGB(h, s/100, l/100)
	return r, g, b, nil
}

func hslToRGB(h, s, l float64) (int, int, int) {
	h = h / 360
	var r, g, b float64

	if s == 0 {
		r = l
		g = l
		b = l
	} else {
		hue2rgb := func(p, q, t float64) float64 {
			if t < 0 {
				t += 1
			}
			if t > 1 {
				t -= 1
			}
			if t < 1.0/6 {
				return p + (q-p)*6*t
			}
			if t < 1.0/2 {
				return q
			}
			if t < 2.0/3 {
				return p + (q-p)*(2.0/3-t)*6
			}
			return p
		}

		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q

		r = hue2rgb(p, q, h+1.0/3)
		g = hue2rgb(p, q, h)
		b = hue2rgb(p, q, h-1.0/3)
	}

	return int(math.Round(r * 255)), int(math.Round(g * 255)), int(math.Round(b * 255))
}

func rgbToHSL(r, g, b int) (int, int, int) {
	rNorm := float64(r) / 255.0
	gNorm := float64(g) / 255.0
	bNorm := float64(b) / 255.0

	max := math.Max(rNorm, math.Max(gNorm, bNorm))
	min := math.Min(rNorm, math.Min(gNorm, bNorm))

	var h, s, l float64
	l = (max + min) / 2

	if max == min {
		h = 0
		s = 0
	} else {
		d := max - min
		if l > 0.5 {
			s = d / (2 - max - min)
		} else {
			s = d / (max + min)
		}

		switch max {
		case rNorm:
			h = (gNorm-bNorm)/d + (func() float64 {
				if gNorm < bNorm {
					return 6
				}
				return 0
			})()
		case gNorm:
			h = (bNorm-rNorm)/d + 2
		case bNorm:
			h = (rNorm-gNorm)/d + 4
		}
		h /= 6
	}

	return int(math.Round(h * 360)), int(math.Round(s * 100)), int(math.Round(l * 100))
}

func init() {
	rootCmd.AddCommand(colorCmd)
}
