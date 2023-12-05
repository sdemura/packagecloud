package cmd

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/sdemura/packagecloud/pkgcloud"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push user/repo[/distro/version] /path/to/package_file",
	Short: "Push a package",
	Long: `Push a package. When --overwrite is set, the existing package will be yanked
and replaced.`,
	Args: cobra.ExactArgs(2),
	RunE: execPushCmd,
}

const flagOverwrite = "overwrite"

func init() {
	pushCmd.Flags().Bool(flagOverwrite, false, "When set, yank the package if it already exists")
	rootCmd.AddCommand(pushCmd)
}

func execPushCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	overwrite, err := cmd.Flags().GetBool(flagOverwrite)
	if err != nil {
		return err
	}

	target, err := pkgcloud.NewTarget(args[0])
	if err != nil {
		return err
	}

	pkg := args[1]

	client, err := pkgcloud.NewClient("")
	if err != nil {
		return err
	}

	const maxRetries = 5

	for i := 0; i < maxRetries; i++ {
		err = doPushPackage(ctx, client, target, pkg, overwrite)
		if err == nil {
			fmt.Printf("OK: %s:%s\n", target.String(), filepath.Base(pkg))
			return nil
		}

		if err == pkgcloud.ErrBadGateway || strings.Contains(err.Error(), "timeout") {
			// We can retry
			time.Sleep(time.Millisecond * 200 * time.Duration(i))
			continue
		}

		break
	}

	return err
}

func doPushPackage(ctx context.Context, client *pkgcloud.Client, target *pkgcloud.Target, pkg string, overwrite bool) error {
	err := client.CreatePackage(ctx, target.Repo, target.Distro, pkg)
	if err == nil {
		return nil
	}

	if err == pkgcloud.ErrFilenameAlreadyTaken && overwrite {
		if err = client.Destroy(ctx, target.String(), pkg); err != nil {
			return fmt.Errorf("failed to yank existing package {%s}: %w", pkg, err)
		}

		return client.CreatePackage(ctx, target.Repo, target.Distro, pkg)
	}

	return err
}
