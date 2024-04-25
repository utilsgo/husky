package main

import (
	"path"

	"github.com/spf13/cobra"
	"github.com/utilsgo/husky/cmd/husky/presets"
	"github.com/utilsgo/husky/pkg/husky"
)

func init() {
	cmdPreset := &cobra.Command{
		Use:   "preset",
		Short: "preset for setup",
	}

	CmdRoot.AddCommand(cmdPreset)

	for name := range presets.Presets {
		p := presets.Presets[name]

		c := &cobra.Command{
			Use: name,
			Run: func(cmd *cobra.Command, args []string) {
				for f, data := range p {
					err := husky.WriteFile(path.Join(projectRoot, f), data)
					if err != nil {
						panic(err)
					}
				}
			},
		}

		cmdPreset.AddCommand(c)
	}
}
