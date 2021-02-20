package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"k8s-test-backend/conf"
	"k8s-test-backend/internal/server"
	"k8s-test-backend/pkg/client"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
	"strings"
)

/*
For quick and easily to parse command and envs of application start.
*/

var showVersionFlag bool

func startCmd() *cobra.Command {
	cobra.OnInitialize(initConf)
	cmd := &cobra.Command{
		Use:   conf.ApplicationConfig.CApplicationName,
		Short: conf.ApplicationConfig.CApplicationName + " is test for kubernetes, but also can run everywhere",
		Long: conf.ApplicationConfig.CApplicationName + `
 is test for kubernetes, but also can run everywhere, application is packing with the docker.
But can run in some platform by build with go build.
The application provide http bin handler for restful request mapping, also support istio feature.
For debug, use debug mode.

run : ./application
All of the flags can set to ENV, and prefix is "KTS_"
`,
		PreRun: func(cmd *cobra.Command, args []string) {
			// update global config
			conf.ApplicationConfig.Port = viper.GetString("server.port")
			conf.ApplicationConfig.Mode = viper.GetString("server.mode")
			conf.ApplicationConfig.LogPath = viper.GetString("server.log_path")
			conf.ApplicationConfig.UseCommonHttp = viper.GetBool("common.enable_common_http")
			conf.ApplicationConfig.UseCacheHttp = viper.GetBool("common.enable_cache_http")
			conf.ApplicationConfig.UseKubernetesFeature = viper.GetBool("k8s.enable_kubernetes_feature")
			conf.ApplicationConfig.KubernetesConfigPath = viper.GetString("k8s.kubernetes_config_path")
			conf.ApplicationConfig.IsInCluster = viper.GetBool("k8s.is_in_kubernetes")
			conf.ApplicationConfig.ServiceIp = viper.GetString("env_service_ip")
			conf.ApplicationConfig.ServiceName = viper.GetString("env_service_name")
			conf.ApplicationConfig.ServiceNamespace = viper.GetString("env_service_namespace")

			// copy from build
			if len(Version) != 0 {
				conf.ApplicationConfig.Version = Version
			}
			conf.ApplicationConfig.BuildStamp = BuildStamp
			conf.ApplicationConfig.BuildPlatform = BuildPlatform

			// init logrus
			// What the mean of 0666?
			file, err := os.OpenFile(conf.ApplicationConfig.LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			writers := []io.Writer{
				file,
				os.Stdout,
			}
			fileAndStdoutWriter := io.MultiWriter(writers...)
			if err == nil {
				logrus.SetOutput(fileAndStdoutWriter)
			} else {
				logrus.Infoln("fail to log to file")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			// show version
			if showVersionFlag {
				showVersion()
				return
			}

			// init k8s client
			if conf.ApplicationConfig.UseKubernetesFeature {
				err := client.InitClient()
				if err != nil {
					logrus.Warnf("Can't use kube feature: %s", err)
					conf.ApplicationConfig.UseKubernetesFeature = false
				}
			}

			server.Start()
		},
	}

	return cmd
}

func showVersion() {
	fmt.Println("The version is :", conf.ApplicationConfig.Version)
	fmt.Println("The build from :", BuildPlatform)
	fmt.Println("The build stamp:", BuildStamp)
}

var configPath = ""

func InitCmd() *cobra.Command {
	viper.SetEnvPrefix(conf.ApplicationConfig.CEnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// init cobra
	cmd := startCmd()
	cmd.PersistentFlags().BoolVarP(&showVersionFlag, "version", "v", false, "show the version of application")
	cmd.PersistentFlags().StringVar(&configPath, "config_path", "", "the application run config file path")
	cmd.PersistentFlags().StringVar(&conf.ApplicationConfig.Port, "port", conf.ApplicationConfig.Port, "the application start port, default is :3000")
	cmd.PersistentFlags().StringVar(&conf.ApplicationConfig.Mode, "mode", conf.ApplicationConfig.Mode, "the application run mode, in 'debug', 'release', 'test'")
	cmd.PersistentFlags().StringVar(&conf.ApplicationConfig.LogPath, "log_path", conf.ApplicationConfig.LogPath, "the file of log output")
	cmd.PersistentFlags().BoolVar(&conf.ApplicationConfig.UseCommonHttp, "enable_common_http", conf.ApplicationConfig.UseCommonHttp, "whether use common http handle, default is true")
	cmd.PersistentFlags().BoolVar(&conf.ApplicationConfig.UseCacheHttp, "enable_cache_http", conf.ApplicationConfig.UseCacheHttp, "whether use cache http handle, default is true")
	cmd.PersistentFlags().BoolVar(&conf.ApplicationConfig.UseKubernetesFeature, "enable_kubernetes_feature", conf.ApplicationConfig.UseKubernetesFeature, "whether enable kubernetes "+
		"feature, default is false")
	if home := homedir.HomeDir(); home != "" {
		cmd.PersistentFlags().StringVar(&conf.ApplicationConfig.KubernetesConfigPath, "kubernetes_config_path", filepath.Join(home, ".kube", "config"), "the config path of kubernetes")
	} else {
		cmd.PersistentFlags().StringVar(&conf.ApplicationConfig.KubernetesConfigPath, "kubernetes_config_path", conf.ApplicationConfig.KubernetesConfigPath, "the config path of kubernetes")
	}
	cmd.PersistentFlags().BoolVar(&conf.ApplicationConfig.IsInCluster, "is_in_kubernetes", conf.ApplicationConfig.IsInCluster, "whether the application in kubernetes cluster"+
		" as the pods")

	// set flags bind
	_ = viper.BindPFlag("server.port", cmd.PersistentFlags().Lookup("port"))
	_ = viper.BindPFlag("server.mode", cmd.PersistentFlags().Lookup("mode"))
	_ = viper.BindPFlag("server.log_path", cmd.PersistentFlags().Lookup("log_path"))
	_ = viper.BindPFlag("common.enable_common_http", cmd.PersistentFlags().Lookup("enable_common_http"))
	_ = viper.BindPFlag("common.enable_cache_http", cmd.PersistentFlags().Lookup("enable_cache_http"))
	_ = viper.BindPFlag("k8s.enable_kubernetes_feature", cmd.PersistentFlags().Lookup("enable_kubernetes_feature"))
	_ = viper.BindPFlag("k8s.kubernetes_config_path", cmd.PersistentFlags().Lookup("kubernetes_config_path"))
	_ = viper.BindPFlag("k8s.is_in_kubernetes", cmd.PersistentFlags().Lookup("is_in_kubernetes"))

	// set env bind
	_ = viper.BindEnv("server.port")
	_ = viper.BindEnv("server.mode")
	_ = viper.BindEnv("server.log_path")
	_ = viper.BindEnv("common.enable_common_http")
	_ = viper.BindEnv("common.enable_cache_http")
	_ = viper.BindEnv("k8s.enable_kubernetes_feature")
	_ = viper.BindEnv("k8s.kubernetes_config_path")
	_ = viper.BindEnv("k8s.is_in_kubernetes")
	_ = viper.BindEnv("env_service_ip")
	_ = viper.BindEnv("env_service_name")
	_ = viper.BindEnv("env_service_namespace")

	// init viper of application config
	viper.SetDefault("server.port", conf.ApplicationConfig.Port)
	viper.SetDefault("server.mode", conf.ApplicationConfig.Mode)
	viper.SetDefault("server.log_path", conf.ApplicationConfig.LogPath)
	viper.SetDefault("common.enable_common_http", conf.ApplicationConfig.UseCommonHttp)
	viper.SetDefault("common.enable_cache_http", conf.ApplicationConfig.UseCacheHttp)
	viper.SetDefault("k8s.enable_kubernetes_feature", conf.ApplicationConfig.UseKubernetesFeature)
	viper.SetDefault("k8s.kubernetes_config_path", conf.ApplicationConfig.KubernetesConfigPath)
	viper.SetDefault("k8s.is_in_kubernetes", conf.ApplicationConfig.IsInCluster)
	viper.SetDefault("env_service_ip", "can't find in env")
	viper.SetDefault("env_service_name", "can't find in env")
	viper.SetDefault("env_service_namespace", "can't find in env")
	return cmd
}

func initConf() {
	if configPath == "" {
		viper.SetConfigFile(conf.ApplicationConfig.CDefaultConfigFile)
	} else {
		viper.SetConfigFile(configPath)
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		logrus.Infof("Using config file:", viper.ConfigFileUsed())
	}
}
