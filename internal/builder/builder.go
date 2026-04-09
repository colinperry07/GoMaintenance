package builder

import (
	"fmt"
	"strings"
)

type Option func (*Style)

type Style struct {
	foreground string
}

func WithForeground (color int) Option {
	return func (s *Style) {
		s.foreground = fmt.Sprintf("%d", color)
	}
}


func Build (textToModify string, opts ...Option) string {
	s := &Style{}
	for _, opt := range opts {
		opt(s)
	}

	var codes []string

	if s.foreground != "" {
		codes = append(codes, s.foreground)
	}
	
	if len(codes) == 0 {
		return textToModify
	}

	return fmt.Sprintf("\x1b[%sm%s\x1b[0m", strings.Join(codes, ";"), textToModify)
	
}

