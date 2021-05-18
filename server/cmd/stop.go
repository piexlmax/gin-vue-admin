package cmd

import (
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

func NewStopCommand() *cobra.Command {
	sc := &cobra.Command{
		Use:   "stop [-f]",
		Short: "stop server",
		Run:   stopCommandFunc,
	}
	sc.Flags().BoolP("force", "f", false, "stop server force")
	return sc
}

func stopCommandFunc(cmd *cobra.Command, args []string) {
	pidMap := make(map[string]int)
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	content, err := ioutil.ReadFile(dir + "/pid")
	if err != nil {
		cmd.PrintErr("error: ", err)
		return
	}

	if err := yaml.Unmarshal(content, pidMap); err != nil {
		cmd.PrintErr("error: ", err)
		return
	}

	cmd.Println("pid:" + strconv.Itoa(pidMap["pid"]))
	cmd.Println("stopping ...")
	force, _ := cmd.Flags().GetBool("force")

	if force {
		if err := syscall.Kill(pidMap["pid"], syscall.SIGKILL); err != nil {
			cmd.PrintErr("error: ", err)
			return
		}
	}

	if err := syscall.Kill(pidMap["pid"], syscall.SIGTERM); err != nil {
		cmd.PrintErr("error: ", err)
		return
	}

	cmd.Println("server stopped successes")
}
