package cmd

import (
	"github.com/qiniu/qshell/v2/iqshell/common/data"
	"github.com/qiniu/qshell/v2/iqshell/storage/object/upload/operations"
	"github.com/spf13/cobra"
)

var uploadCmdBuilder = func() *cobra.Command {
	info := operations.BatchUploadInfo{}
	cmd := &cobra.Command{
		Use:   "qupload <quploadConfigFile>",
		Short: "Batch upload files to the qiniu bucket",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				cfg.UploadConfigFile = args[0]
			}
			loadConfig()
			operations.BatchUpload(info)
		},
	}
	cmd.Flags().StringVarP(&info.GroupInfo.SuccessExportFilePath, "success-list", "s", "", "upload success (all) file list")
	cmd.Flags().StringVarP(&info.GroupInfo.FailExportFilePath, "failure-list", "f", "", "upload failure file list")
	cmd.Flags().StringVarP(&info.GroupInfo.OverrideExportFilePath, "overwrite-list", "w", "", "upload success (overwrite) file list")
	cmd.Flags().IntVarP(&info.GroupInfo.WorkCount, "worker", "c", 1, "worker count")
	cmd.Flags().StringVarP(&cfg.CmdCfg.Up.Policy.CallbackURL, "callback-urls", "l", "", "upload callback urls, separated by comma")
	cmd.Flags().StringVarP(&cfg.CmdCfg.Up.Policy.CallbackHost, "callback-host", "T", "", "upload callback host")
	return cmd
}

var upload2CmdBuilder = func() *cobra.Command {
	var (
		resumableAPIV2 = false
		ignoreDir      = false
		overwrite      = false
		checkExists    = false
		checkHash      = false
		checkSize      = false
		rescanLocal    = false
	)
	info := operations.BatchUploadInfo{}
	cmd := &cobra.Command{
		Use:   "qupload2",
		Short: "Batch upload files to the qiniu bucket",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				cfg.UploadConfigFile = args[0]
			}
			cfg.CmdCfg.Up.ResumableAPIV2 = data.GetBoolString(resumableAPIV2)
			cfg.CmdCfg.Up.IgnoreDir = data.GetBoolString(ignoreDir)
			cfg.CmdCfg.Up.Overwrite = data.GetBoolString(overwrite)
			cfg.CmdCfg.Up.CheckExists = data.GetBoolString(checkExists)
			cfg.CmdCfg.Up.CheckHash = data.GetBoolString(checkHash)
			cfg.CmdCfg.Up.CheckSize = data.GetBoolString(checkSize)
			cfg.CmdCfg.Up.RescanLocal = data.GetBoolString(rescanLocal)
			loadConfig()
			operations.BatchUpload(info)
		},
	}
	cmd.Flags().StringVar(&info.GroupInfo.SuccessExportFilePath, "success-list", "", "upload success file list")
	cmd.Flags().StringVar(&info.GroupInfo.FailExportFilePath, "failure-list", "", "upload failure file list")
	cmd.Flags().StringVar(&info.GroupInfo.OverrideExportFilePath, "overwrite-list", "", "upload success (overwrite) file list")
	cmd.Flags().IntVar(&info.GroupInfo.WorkCount, "thread-count", 1, "multiple thread count")

	cmd.Flags().BoolVarP(&resumableAPIV2, "resumable-api-v2", "", false, "use resumable upload v2 APIs to upload")
	cmd.Flags().BoolVar(&ignoreDir, "ignore-dir", false, "ignore the dir in the dest file key")
	cmd.Flags().BoolVar(&overwrite, "overwrite", false, "overwrite the file of same key in bucket")
	cmd.Flags().BoolVar(&checkExists, "check-exists", false, "check file key whether in bucket before upload")
	cmd.Flags().BoolVar(&checkHash, "check-hash", false, "check hash")
	cmd.Flags().BoolVar(&checkSize, "check-size", false, "check file size")
	cmd.Flags().BoolVar(&rescanLocal, "rescan-local", false, "rescan local dir to upload newly add files")

	cmd.Flags().Int64Var(&cfg.CmdCfg.Up.ResumableAPIV2PartSize, "resumable-api-v2-part-size", data.BLOCK_SIZE, "the part size when use resumable upload v2 APIs to upload")
	cmd.Flags().StringVar(&cfg.CmdCfg.Up.SrcDir, "src-dir", "", "src dir to upload")
	cmd.Flags().StringVar(&cfg.CmdCfg.Up.FileList, "file-list", "", "file list to upload")
	cmd.Flags().StringVar(&cfg.CmdCfg.Up.Bucket, "bucket", "", "bucket")
	cmd.Flags().Int64Var(&cfg.CmdCfg.Up.PutThreshold, "put-threshold", 0, "chunk upload threshold")
	cmd.Flags().StringVar(&cfg.CmdCfg.Up.KeyPrefix, "key-prefix", "", "key prefix prepended to dest file key")
	cmd.Flags().StringVar(&cfg.CmdCfg.Up.SkipFilePrefixes, "skip-file-prefixes", "", "skip files with these file prefixes")
	cmd.Flags().StringVar(&cfg.CmdCfg.Up.SkipPathPrefixes, "skip-path-prefixes", "", "skip files with these relative path prefixes")
	cmd.Flags().StringVar(&cfg.CmdCfg.Up.SkipFixedStrings, "skip-fixed-strings", "", "skip files with the fixed string in the name")
	cmd.Flags().StringVar(&cfg.CmdCfg.Up.SkipSuffixes, "skip-suffixes", "", "skip files with these suffixes")
	cmd.Flags().StringVar(&cfg.CmdCfg.Up.UpHost, "up-host", "", "upload host")
	cmd.Flags().StringVar(&cfg.CmdCfg.Up.BindUpIp, "bind-up-ip", "", "upload host ip to bind")
	cmd.Flags().StringVar(&cfg.CmdCfg.Up.BindRsIp, "bind-rs-ip", "", "rs host ip to bind")
	cmd.Flags().StringVar(&cfg.CmdCfg.Up.BindNicIp, "bind-nic-ip", "", "local network interface card to bind")
	cmd.Flags().StringVar(&cfg.CmdCfg.Up.LogFile, "log-file", "", "log file")
	cmd.Flags().StringVar(&cfg.CmdCfg.Up.LogLevel, "log-level", "info", "log level")
	cmd.Flags().IntVar(&cfg.CmdCfg.Up.LogRotate, "log-rotate", 1, "log rotate days")
	cmd.Flags().IntVar(&cfg.CmdCfg.Up.FileType, "file-type", 0, "set storage file type")
	cmd.Flags().StringVarP(&cfg.CmdCfg.Up.Policy.CallbackURL, "callback-urls", "l", "", "upload callback urls, separated by comma")
	cmd.Flags().StringVarP(&cfg.CmdCfg.Up.Policy.CallbackHost, "callback-host", "T", "", "upload callback host")
	return cmd
}

var syncCmdBuilder = func() *cobra.Command {
	resumeAPIV2 := false
	info := operations.UploadInfo{}
	cmd := &cobra.Command{
		Use:   "sync <SrcResUrl> <Buckets> [-k <Key>]",
		Short: "Sync big file to qiniu bucket",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			cfg.CmdCfg.Up.ResumableAPIV2 = data.GetBoolString(resumeAPIV2)
			if len(args) > 0 {
				info.FilePath = args[0]
				info.Bucket = args[1]
			}
			loadConfig()
			operations.UploadFile(info)
		},
	}
	cmd.Flags().StringVarP(&info.Key, "key", "k", "", "save as <key> in bucket")
	cmd.Flags().BoolVarP(&resumeAPIV2, "resumable-api-v2", "", false, "use resumable upload v2 APIs to upload")
	cmd.Flags().StringVarP(&cfg.CmdCfg.Up.UpHost, "uphost", "u", "", "upload host")
	return cmd
}

var formUploadCmdBuilder = func() *cobra.Command {
	overwrite := false
	info := operations.UploadInfo{}
	cmd := &cobra.Command{
		Use:   "fput <Bucket> <Key> <LocalFile>",
		Short: "Form upload a local file",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			cfg.CmdCfg.Up.DisableResume = data.TrueString
			cfg.CmdCfg.Up.Overwrite = data.GetBoolString(overwrite)
			if len(args) > 2 {
				info.Bucket = args[0]
				info.Key = args[1]
				info.FilePath = args[2]
			}
			loadConfig()
			operations.UploadFile(info)
		},
	}
	//cmd.Flags().IntVarP(&info.w, "worker", "c", 16, "worker count")
	cmd.Flags().BoolVar(&overwrite, "overwrite", false, "overwrite the file of same key in bucket")
	cmd.Flags().StringVarP(&info.MimeType, "mimetype", "t", "", "file mime type")
	cmd.Flags().IntVarP(&cfg.CmdCfg.Up.FileType, "storage", "s", 0, "storage type")
	cmd.Flags().StringVarP(&cfg.CmdCfg.Up.UpHost, "up-host", "u", "", "uphost")
	cmd.Flags().StringVarP(&cfg.CmdCfg.Up.Policy.CallbackURL, "callback-urls", "l", "", "upload callback urls, separated by comma")
	cmd.Flags().StringVarP(&cfg.CmdCfg.Up.Policy.CallbackHost, "callback-host", "T", "", "upload callback host")

	return cmd
}

var resumeUploadCmdBuilder = func() *cobra.Command {
	var (
		resumeAPIV2 = false
		overwrite   = false
	)
	info := operations.UploadInfo{}
	cmd := &cobra.Command{
		Use:   "rput <Bucket> <Key> <LocalFile>",
		Short: "Resumable upload a local file",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			cfg.CmdCfg.Up.DisableForm = data.TrueString
			cfg.CmdCfg.Up.ResumableAPIV2 = data.GetBoolString(resumeAPIV2)
			cfg.CmdCfg.Up.Overwrite = data.GetBoolString(overwrite)
			if len(args) > 2 {
				info.Bucket = args[0]
				info.Key = args[1]
				info.FilePath = args[2]
			}
			loadConfig()
			operations.UploadFile(info)
		},
	}

	cmd.Flags().StringVarP(&info.MimeType, "mimetype", "t", "", "file mime type")
	cmd.Flags().BoolVarP(&resumeAPIV2, "resumable-api-v2", "", false, "use resumable upload v2 APIs to upload")
	cmd.Flags().BoolVar(&overwrite, "overwrite", false, "overwrite the file of same key in bucket")

	cmd.Flags().Int64VarP(&cfg.CmdCfg.Up.ResumableAPIV2PartSize, "v2-part-size", "", data.BLOCK_SIZE, "the part size when use resumable upload v2 APIs to upload, default 4M")
	cmd.Flags().IntVarP(&cfg.CmdCfg.Up.FileType, "storage", "s", 0, "storage type")
	cmd.Flags().IntVarP(&cfg.CmdCfg.Up.WorkerCount, "worker", "c", 16, "worker count")
	cmd.Flags().StringVarP(&cfg.CmdCfg.Up.UpHost, "up-host", "u", "", "uphost")
	cmd.Flags().StringVarP(&cfg.CmdCfg.Up.Policy.CallbackURL, "callback-urls", "l", "", "upload callback urls, separated by comma")
	cmd.Flags().StringVarP(&cfg.CmdCfg.Up.Policy.CallbackHost, "callback-host", "T", "", "upload callback host")
	return cmd
}

func init() {
	rootCmd.AddCommand(
		uploadCmdBuilder(),
		upload2CmdBuilder(),
		syncCmdBuilder(),
		formUploadCmdBuilder(),
		resumeUploadCmdBuilder(),
	)
}
