package cmd

import (
	`fmt`
	"strconv"
	`strings`
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	labelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#38bdf8")).Bold(true)
	valueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#a3e635")).Bold(true)
	errStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0087")).Bold(true)
)

var timeCmd = &cobra.Command{
	Use:   "time <in-format> <value> [out-format]",
	Short: "Convert between time formats",
	Args:  cobra.RangeArgs(2, 3),
	RunE: func(cmd *cobra.Command, args []string) error {
		inFormat := args[0]
		val := args[1]
		outFormat := ""
		if len(args) == 3 {
			outFormat = args[2]
		}

		var t time.Time
		var err error
		var sec, nsec int64

		switch inFormat {
		case "rfc":
			t, err = time.Parse(time.RFC3339, val)
		case "pg":
			t, err = time.Parse("2006-01-02 15:04:05.000000-07", val)
			if err != nil {
				t, err = time.Parse("2006-01-02 15:04:05.000000", val)
			}
		case "s":
			sec, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				break
			}
			t = time.Unix(sec, 0)
		case "ms":
			ms, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				break
			}
			t = time.Unix(0, ms*int64(time.Millisecond))
		case "us":
			us, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				break
			}
			t = time.Unix(0, us*int64(time.Microsecond))
		case "ns":
			ns, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				break
			}
			t = time.Unix(0, ns)
		case "pb":
			parts := strings.Split(val, ",")
			if len(parts) != 2 {
				err = fmt.Errorf("pb input must be in 'seconds,nanos' format")
				break
			}
			sec, err = strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
			if err != nil {
				break
			}
			nsec, err = strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
			if err != nil {
				break
			}
			t = time.Unix(sec, nsec)
		default:
			return fmt.Errorf("unsupported input format: %s", inFormat)
		}
		if err != nil {
			fmt.Println(errStyle.Render("Error parsing input:"), err)
			return nil
		}

		outputs := map[string]func(){
			"rfc": func() {
				fmt.Printf("%s %s\n", labelStyle.Render("RFC3339:     "), valueStyle.Render(t.UTC().Format(time.RFC3339)))
			},
			"pg": func() {
				fmt.Printf("%s %s\n", labelStyle.Render("Postgres:    "), valueStyle.Render(t.Format("2006-01-02 15:04:05.000000-07")))
			},
			"s": func() {
				fmt.Printf("%s %s\n", labelStyle.Render("Seconds:     "), valueStyle.Render(strconv.FormatInt(t.Unix(), 10)))
			},
			"ms": func() {
				fmt.Printf("%s %s\n", labelStyle.Render("Milliseconds:"), valueStyle.Render(strconv.FormatInt(t.UnixMilli(), 10)))
			},
			"us": func() {
				fmt.Printf("%s %s\n", labelStyle.Render("Microseconds:"), valueStyle.Render(strconv.FormatInt(t.UnixMicro(), 10)))
			},
			"ns": func() {
				fmt.Printf("%s %s\n", labelStyle.Render("Nanoseconds: "), valueStyle.Render(strconv.FormatInt(t.UnixNano(), 10)))
			},
			"pb": func() {
				fmt.Printf("%s %s\n", labelStyle.Render("timestamppb: "), valueStyle.Render(fmt.Sprintf("timestamppb.Timestamp{Seconds: %d, Nanos: %d}", t.Unix(), t.Nanosecond())))
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

func init() {
	rootCmd.AddCommand(timeCmd)
}
