//go:generate stringer -type=MsgType
package ghext

import "fmt"

type MsgType int

const (
	Init MsgType = iota + 1
	Feat
	Fix
	Docs
	Style
	Perf
	Refactor
	Test
	Build
	CI
	Chore
	Revert
	Release
)

type MessageTemplate struct {
	Type MsgType
	Icon string
}

var commitTemplate = map[MsgType]MessageTemplate{
	Init:     {Icon: "🎉 ", Type: Init},
	Feat:     {Icon: "✨ ", Type: Feat},
	Fix:      {Icon: "🐞 ", Type: Fix},
	Docs:     {Icon: "📃 ", Type: Docs},
	Style:    {Icon: "🌈 ", Type: Style},
	Refactor: {Icon: "🦄 ", Type: Refactor},
	Perf:     {Icon: "🎈 ", Type: Perf},
	Test:     {Icon: "🧪 ", Type: Test},
	Build:    {Icon: "🔧 ", Type: Build},
	CI:       {Icon: "🐎 ", Type: CI},
	Chore:    {Icon: "🐳 ", Type: Chore},
	Revert:   {Icon: "↩ ", Type: Revert},
}

func CommitTemplate(typ int, scope, subject, body, footer string, isHideEmoji bool) (string, error) {
	msg := commitTemplate[MsgType(typ)]
	if subject == "" {
		if typ == int(Init) {
			subject = "Initial commit"
		} else {
			return "", fmt.Errorf("subject is empty")
		}
	}
	if isHideEmoji {
		msg.Icon = ""
	}
	line := fmt.Sprintf(
		"%s%s(%s): %s",
		msg.Icon, msg.Type, scope, subject,
	)
	if body != "" {
		line = fmt.Sprintf("%s\n%s", line, body)
	}
	if footer != "" {
		if body == "" {
			line = fmt.Sprintf("%s\n\n%s", line, footer)
		} else {
			line = fmt.Sprintf("%s\n%s", line, footer)
		}
	}
	return line, nil
}
