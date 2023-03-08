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
	Init:     MessageTemplate{Icon: "🎉 ", Type: Init},
	Feat:     MessageTemplate{Icon: "✨ ", Type: Feat},
	Fix:      MessageTemplate{Icon: "🐞 ", Type: Fix},
	Docs:     MessageTemplate{Icon: "📃 ", Type: Docs},
	Style:    MessageTemplate{Icon: "🌈 ", Type: Style},
	Refactor: MessageTemplate{Icon: "🦄 ", Type: Refactor},
	Perf:     MessageTemplate{Icon: "🎈 ", Type: Perf},
	Test:     MessageTemplate{Icon: "🧪 ", Type: Test},
	Build:    MessageTemplate{Icon: "🔧 ", Type: Build},
	CI:       MessageTemplate{Icon: "🐎 ", Type: CI},
	Chore:    MessageTemplate{Icon: "🐳 ", Type: Chore},
	Revert:   MessageTemplate{Icon: "↩ ", Type: Revert},
}

func CommitTemplate(typ int, scope, subject, body, footer string, isHideEmoji bool) (string, error) {
	msg := commitTemplate[MsgType(typ)]
	if subject == "" {
		if typ == int(Init) {
			subject = "Initial commit "
		} else {
			return "", fmt.Errorf("subject is empty")
		}
	}
	if isHideEmoji {
		msg.Icon = ""
	}
	line := fmt.Sprintf(
		"%s%s(%s): %s ",
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
