package version

import (
	"context"
	"fmt"

	"github.com/kubescape/go-logger"
	"github.com/kubescape/kubescape/v3/core/cautils"
	"github.com/spf13/cobra"
)

func GetVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Get current version",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			v := cautils.NewIVersionCheckHandler(ctx) // 创建版本检查处理器
			versionCheckRequest := cautils.NewVersionCheckRequest(cautils.BuildNumber, "", "", "version")
			v.CheckLatestVersion(ctx, versionCheckRequest) // 检查最新版本
			fmt.Fprintf(cmd.OutOrStdout(),
				"Your current version is: %s\n",
				versionCheckRequest.ClientVersion,
			)
			logger.L().Debug(fmt.Sprintf("git enabled in build: %t", isGitEnabled()))
			return nil
		},
	}
	return versionCmd
}
