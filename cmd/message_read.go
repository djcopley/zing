package cmd

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"

    "github.com/djcopley/zing/internal/api"
    "github.com/djcopley/zing/internal/client"
    "github.com/djcopley/zing/internal/pager"
    "github.com/spf13/cobra"
)

var (
    pageSize  int32
    pageToken string
    noColor   bool
    termWidth int
)

var messageReadCmd = &cobra.Command{
	Use:   "read",
	Short: "Read the messages sent to you",
	RunE: func(cmd *cobra.Command, args []string) error {
		token := config.Token()
		if token == "" {
			return fmt.Errorf("authentication token is required; please login first")
		}
		ctx := client.AddAuthMetadata(cmd.Context(), token)

		addr := config.ServerAddr()
		client, err := client.NewClient(addr, config.Insecure(), config.Plaintext())
		if err != nil {
			return fmt.Errorf("failed to connect to server: %s", err)
		}
		defer client.Close()

  request := &api.ListMessagesRequest{PageSize: pageSize, PageToken: pageToken}
  r, err := client.ListMessages(ctx, request)
  if err != nil {
      return fmt.Errorf("failed to connect to server: %s", err)
  }

  // determine terminal width (for wrapping)
  if termWidth <= 0 {
      termWidth = detectWidth()
  }

  var b strings.Builder
  if len(r.Messages) == 0 {
      b.WriteString("No new messages.\n")
  } else {
      for i, message := range r.Messages {
          b.WriteString(prettyMessageCard(message, i+1, len(r.Messages)))
          if i < len(r.Messages)-1 {
              b.WriteString("\n")
          }
      }
      if r.NextPageToken != "" {
          b.WriteString(colorDim(strings.Repeat("\u2500", maxInt(20, minInt(termWidth-2, 60)))))
          b.WriteString("\n")
          b.WriteString(colorDim(fmt.Sprintf("More messages available. Fetch next page with: zing message read --page-token %s --page-size %d\n", r.NextPageToken, pageSize)))
      }
  }

		p := pager.NewPager(cmd.OutOrStdout(), cmd.ErrOrStderr())
		if err := p.Page(b.String()); err != nil {
			// Fall back to plain stdout if pager is unavailable
			_, _ = fmt.Fprint(cmd.OutOrStdout(), b.String())
		}
		return nil
	},
}

func init() {
    messageCmd.AddCommand(messageReadCmd)
    messageReadCmd.Flags().Int32Var(&pageSize, "page-size", 0, "Number of messages per page (default 50, max 1000)")
    messageReadCmd.Flags().StringVar(&pageToken, "page-token", "", "Page token from a previous response to fetch the next page")
    messageReadCmd.Flags().BoolVar(&noColor, "no-color", false, "Disable colorized output")
    messageReadCmd.Flags().IntVar(&termWidth, "width", 0, "Wrap output to this width (0 = auto)")
}

// === Pretty printing utilities ===

// ANSI colors (kept minimal). Disabled when noColor is true or NO_COLOR env var is set.
const (
    ansiReset  = "\x1b[0m"
    ansiBold   = "\x1b[1m"
    ansiDim    = "\x1b[2m"
    ansiCyan   = "\x1b[36m"
    ansiGreen  = "\x1b[32m"
    ansiYellow = "\x1b[33m"
    ansiGrey   = "\x1b[90m"
)

func colorsEnabled() bool {
    if noColor {
        return false
    }
    if os.Getenv("NO_COLOR") != "" { // https://no-color.org/
        return false
    }
    return true
}

func colorize(s, code string) string {
    if !colorsEnabled() {
        return s
    }
    return code + s + ansiReset
}

func colorBold(s string) string  { return colorize(s, ansiBold) }
func colorCyan(s string) string  { return colorize(s, ansiCyan) }
func colorGreen(s string) string { return colorize(s, ansiGreen) }
func colorYellow(s string) string { return colorize(s, ansiYellow) }
func colorDim(s string) string   { return colorize(s, ansiGrey) }

func detectWidth() int {
    if v := os.Getenv("COLUMNS"); v != "" {
        if n, err := strconv.Atoi(v); err == nil && n > 20 {
            return n
        }
    }
    // default width
    return 80
}

func maxInt(a, b int) int { if a > b { return a }; return b }
func minInt(a, b int) int { if a < b { return a }; return b }

// prettyMessageCard renders a single message inside a unicode box with metadata and colors.
func prettyMessageCard(m *api.Message, idx int, total int) string {
    width := termWidth
    if width <= 0 {
        width = detectWidth()
    }
    // inner width excluding borders and a space padding on each side
    inner := maxInt(20, width-2)

    // Metadata extraction
    var ts time.Time
    var tsStr string
    if m.GetMetadata() != nil && m.GetMetadata().GetTimestamp() != nil {
        ts = m.GetMetadata().GetTimestamp().AsTime().Local()
        tsStr = ts.Format("2006-01-02 15:04:05 MST")
    } else {
        tsStr = "(no timestamp)"
    }
    from := "unknown"
    to := "unknown"
    if md := m.GetMetadata(); md != nil {
        if u := md.GetFrom(); u != nil && u.GetUsername() != "" {
            from = u.GetUsername()
        }
        if u := md.GetTo(); u != nil && u.GetUsername() != "" {
            to = u.GetUsername()
        }
    }
    id := ""
    if md := m.GetMetadata(); md != nil {
        id = md.GetId()
    }

    title := fmt.Sprintf("Message %d/%d", idx, total)
    top := "\u250C" + strings.Repeat("\u2500", inner-2) + "\u2510\n" // ┌────┐
    // Header line with centered title-ish
    header := padLine(colorBold(colorCyan(" "+title+" ")), inner)

    // Metadata lines
    meta1 := padLine(" From: "+colorGreen(from)+"    To: "+colorGreen(to), inner)
    meta2 := padLine(" Time: "+colorYellow(tsStr), inner)
    meta3 := padLine("   ID: "+colorDim(id), inner)
    sep := "\u251C" + strings.Repeat("\u2500", inner-2) + "\u2524\n" // ├───┤

    // Content, wrapped
    content := m.GetContent()
    if strings.TrimSpace(content) == "" {
        content = "(no content)"
    }
    wrapped := wrapText(content, inner-2)
    var body strings.Builder
    for _, line := range strings.Split(wrapped, "\n") {
        body.WriteString(padLine(" "+line, inner))
    }

    bottom := "\u2514" + strings.Repeat("\u2500", inner-2) + "\u2518\n" // └───┘

    // Compose with borders
    var b strings.Builder
    b.WriteString(top)
    b.WriteString(header)
    b.WriteString(meta1)
    b.WriteString(meta2)
    b.WriteString(meta3)
    b.WriteString(sep)
    b.WriteString(body.String())
    b.WriteString(bottom)
    return b.String()
}

// padLine wraps a single line inside box borders and ensures it fills the inner width with spaces.
func padLine(s string, inner int) string {
    // strip ANSI for width calc by naively removing \x1b[ ... m sequences
    vis := stripANSI(s)
    pad := inner - 2 - visualLen(vis)
    if pad < 0 { pad = 0 }
    return "\u2502" + s + strings.Repeat(" ", pad) + "\u2502\n" // │ line │
}

// wrapText performs a simple rune-aware wrap at the given width.
func wrapText(text string, width int) string {
    if width <= 10 {
        width = 10
    }
    var out strings.Builder
    for i, line := range strings.Split(text, "\n") {
        runes := []rune(line)
        for len(runes) > width {
            out.WriteString(string(runes[:width]))
            out.WriteString("\n")
            runes = runes[width:]
        }
        out.WriteString(string(runes))
        if i < len(strings.Split(text, "\n"))-1 {
            out.WriteString("\n")
        }
    }
    return out.String()
}

// visualLen returns length excluding ANSI sequences.
func visualLen(s string) int { return len([]rune(stripANSI(s))) }

func stripANSI(s string) string {
    // simple state machine to drop ESC[ ... m
    out := make([]rune, 0, len(s))
    inEsc := false
    for _, r := range s {
        if inEsc {
            if r == 'm' { // end of sequence
                inEsc = false
            }
            continue
        }
        if r == 0x1b { // ESC
            inEsc = true
            continue
        }
        out = append(out, r)
    }
    return string(out)
}
