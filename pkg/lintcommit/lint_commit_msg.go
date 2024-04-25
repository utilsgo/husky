package lintcommit

import (
	"github.com/utilsgo/husky/pkg/conventionalcommit"
)

func LintCommitMsg(commitMsg string) error {
	_, err := conventionalcommit.ParseCommitMsg(commitMsg)
	if err != nil {
		return err
	}
	return nil
}
