package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/edgeworx/packagecloud/pkgcloud"
	"github.com/spf13/cobra"
)

// yankCmd represents the yank command
var yankCmd = &cobra.Command{
	Use:   "yank user/repo[/distro/version] /path/to/packages",
	Short: "Yank (delete) a package file",
	RunE:  execYankCmd,
	Args:  cobra.ExactArgs(2),
}

func init() {
	rootCmd.AddCommand(yankCmd)
}

func execYankCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	target, err := pkgcloud.NewTarget(args[0])
	if err != nil {
		return err
	}
	pkg := args[1]

	client, err := pkgcloud.NewClient("")
	if err != nil {
		return err
	}

	err = client.Destroy(ctx, target.String(), pkg)
	if err != nil {
		return fmt.Errorf("%s:%s... %v", target.String(), filepath.Base(pkg), err)
	}

	fmt.Fprintf(os.Stdout, "YANKED: %s:%s\n", target.String(), filepath.Base(pkg))
	return nil
}
