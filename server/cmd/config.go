package cmd

import (
	"gin-vue-admin/global"
	"gin-vue-admin/utils"
	"github.com/spf13/cobra"
)

// NewConfigCommand return a config subcommand of rootCmd
func NewConfigCommand() *cobra.Command {
	conf := &cobra.Command{
		Use:   "config <subcommand>",
		Short: "tune pd configs",
	}
	conf.AddCommand(NewShowConfigCommand())
	return conf
}

// NewShowConfigCommand return a show subcommand of configCmd
func NewShowConfigCommand() *cobra.Command {
	sc := &cobra.Command{
		Use:   "show [replication|label-property|all]",
		Short: "show replication and schedule config of PD",
		//Run:   showConfigCommandFunc,
	}
	sc.AddCommand(NewShowAllConfigCommand())
	return sc
}

// NewShowAllConfigCommand return a show all subcommand of show subcommand
func NewShowAllConfigCommand() *cobra.Command {
	sc := &cobra.Command{
		Use:   "all",
		Short: "show all config of redissyncer-cli",
		Run:   showAllConfigCommandFunc,
	}
	return sc
}

// NewShowScheduleConfigCommand return a show all subcommand of show subcommand
func NewShowScheduleConfigCommand() *cobra.Command {
	sc := &cobra.Command{
		Use:   "schedule",
		Short: "show schedule config of PD",
		Run:   showScheduleConfigCommandFunc,
	}
	return sc
}

func showConfigCommandFunc(cmd *cobra.Command, args []string) {
	cmd.Println(cmd.Args)
}

func showScheduleConfigCommandFunc(cmd *cobra.Command, args []string) {
	cmd.Println(cmd.Args)
}

func showAllConfigCommandFunc(cmd *cobra.Command, args []string) {
	configs, err := utils.MapToYamlString(global.GVA_VP.AllSettings())
	if err != nil {
		cmd.PrintErrln(err)
		return
	}
	cmd.Println(configs)
}
