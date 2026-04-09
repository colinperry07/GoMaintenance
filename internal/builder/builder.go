package builder

import (
	"fmt"
	"strings"
)

type Option func(*Style)

type Style struct {
	foreground    string
	background    string
	bold          string
	dim           string
	italic        string
	underline     string
	blinking      string
	reverse       string
	hidden        string
	strikethrough string
}

func WithForeground(color int) Option {
	return func(s *Style) {
		s.foreground = fmt.Sprintf("%d", color)
	}
}

func WithBackground(color int) Option {
	return func(s *Style) {
		s.background = fmt.Sprintf("%d", color)
	}
}

func WithBold() Option {
	return func(s *Style) {
		s.bold = "1m"
	}
}

func WithDim() Option {
	return func(s *Style) {
		s.dim = "2m"
	}
}

func WithItalic() Option {
	return func(s *Style) {
		s.italic = "3m"
	}
}

func WithUnderline() Option {
	return func(s *Style) {
		s.underline = "4m"
	}
}

func WithBlinking() Option {
	return func(s *Style) {
		s.blinking = "5m"
	}
}

func WithReverse() Option {
	return func(s *Style) {
		s.reverse = "7m"
	}
}

func WithHidden() Option {
	return func(s *Style) {
		s.hidden = "8m"
	}
}

func WithStrikethrough() Option {
	return func(s *Style) {
		s.strikethrough = "9m"
	}
}

func Build(textToModify string, opts ...Option) string {
	s := &Style{}
	for _, opt := range opts {
		opt(s)
	}

	var codes []string

	if s.foreground != "" {
		codes = append(codes, s.foreground)
	}

	if s.background != "" {
		codes = append(codes, s.background)
	}

	if s.bold != "" {
		codes = append(codes, s.bold)
	}

	if s.dim != "" {
		codes = append(codes, s.dim)
	}

	if s.italic != "" {
		codes = append(codes, s.italic)
	}

	if s.underline != "" {
		codes = append(codes, s.underline)
	}

	if s.blinking != "" {
		codes = append(codes, s.blinking)
	}

	if s.reverse != "" {
		codes = append(codes, s.reverse)
	}

	if s.hidden != "" {
		codes = append(codes, s.hidden)
	}

	if s.strikethrough != "" {
		codes = append(codes, s.strikethrough)
	}

	if len(codes) == 0 {
		return textToModify
	}

	return fmt.Sprintf("\x1b[%sm%s\x1b[0m", strings.Join(codes, ";"), textToModify)

}
