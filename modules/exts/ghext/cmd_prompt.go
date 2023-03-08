package ghext

import (
	"github.com/manifoldco/promptui"
	"github.com/virzz/logger"
)

const (
	ReturnErr = iota
	ReturnBack
	ReturnComplete
)

var commitTypeItems []string

type promptCommitMsg struct {
	Type                         MsgType
	Scope, Subject, Body, Footer string
}

var cmsg = promptCommitMsg{}

func promptInput(label string, dst *string) {
	prompt := promptui.Prompt{
		Label: label,
	}
	res, err := prompt.Run()
	if err != nil {
		logger.ErrorF("Prompt failed %v\n", err)
		return
	}
	*dst = res
}
func promptMsg() int {
	prompt := promptui.Select{
		Label: "Select Commit Type",
		Items: []string{"Complete", "Back", "Scope", "Subject", "Body", "Footer"},
	}
	for {
		_, result, err := prompt.Run()
		if err != nil {
			logger.ErrorF("Prompt failed %v\n", err)
			return ReturnErr
		}
		switch result {
		case "Complete":
			return ReturnComplete
		case "Back":
			return ReturnBack
		case "Scope":
			promptInput(result, &cmsg.Scope)
		case "Subject":
			promptInput(result, &cmsg.Subject)
		case "Body":
			promptInput(result, &cmsg.Body)
		case "Footer":
			promptInput(result, &cmsg.Footer)
		}
	}
}

func prompt() error {
	prompt := promptui.Select{
		Label: "Select Commit Type",
		Items: commitTypeItems,
	}
	for {
		_, t, err := prompt.Run()
		if err != nil {
			logger.ErrorF("Prompt failed %v\n", err)
			return err
		}
		for i := 1; i < len(_MsgType_index); i++ {
			if MsgType(i).String() == t {
				cmsg.Type = MsgType(i)
				break
			}
		}
		r := promptMsg()
		if r == ReturnComplete {
			return nil
		}
	}
}
