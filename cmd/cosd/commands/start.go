package commands

import (
	"fmt"
	"github.com/coschain/cobra"
	ctrl "github.com/coschain/contentos-go/app"
	"github.com/coschain/contentos-go/app/plugins"
	"github.com/coschain/contentos-go/common"
	"github.com/coschain/contentos-go/common/logging"
	"github.com/coschain/contentos-go/common/pprof"
	"github.com/coschain/contentos-go/config"
	"github.com/coschain/contentos-go/db/storage"
	"github.com/coschain/contentos-go/iservices"
	"github.com/coschain/contentos-go/node"
	"github.com/coschain/contentos-go/rpc"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var StartCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start cosd node",
		Run:   startNode,
	}
	cmd.Flags().StringVarP(&cfgName, "name", "n", "", "node name (default is cosd)")
	return cmd
}

func makeNode() (*node.Node, node.Config) {
	var cfg node.Config
	if cfgName == "" {
		cfg.Name = ClientIdentifier
	} else {
		cfg.Name = cfgName
	}
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	confdir := filepath.Join(config.DefaultDataDir(), cfg.Name)
	viper.AddConfigPath(confdir)
	err := viper.ReadInConfig()
	if err == nil {
		viper.Unmarshal(&cfg)
	} else {
		fmt.Printf("fatal: not be initialized (do `init` first)\n")
		os.Exit(1)
	}
	if cfg.DataDir != "" {
		dir, err := filepath.Abs(cfg.DataDir)
		if err != nil {
			common.Fatalf("DataDir in cfg cannot be converted to absolute path")
		}
		cfg.DataDir = dir
	}
	app, err := node.New(&cfg)
	if err != nil {
		fmt.Println("Fatal: ", err)
		os.Exit(1)
	}
	return app, cfg

}

// NO OTHER CONFIGS HERE EXCEPT NODE CONFIG
func startNode(cmd *cobra.Command, args []string) {
	logging.Init("logs", "debug", 0)

	pprof.StartPprof()

	// _ is cfg as below process has't used
	_, _ = cmd, args
	app, _ := makeNode()
	//app.Register("timer", func(ctx *node.ServiceContext) (node.Service, error) {
	//	return timer.New(ctx, ctx.Config().Timer)
	//})
	//app.Register("printer", func(ctx *node.ServiceContext) (node.Service, error) {
	//	return printer.New(ctx)
	//})

	//app.Register(iservices.P2P_SERVER_NAME, func(ctx *node.ServiceContext) (node.Service, error) {
	//		return p2p.NewServer(ctx)
	//})

	app.Register(iservices.DB_SERVER_NAME, func(ctx *node.ServiceContext) (node.Service, error) {
		return storage.New(ctx, "./db/")
	})

	app.Register(iservices.CTRL_SERVER_NAME, func(ctx *node.ServiceContext) (node.Service, error) {
		return ctrl.NewController(ctx)
	})

	app.Register(iservices.RPC_SERVER_NAME, func(ctx *node.ServiceContext) (node.Service, error) {
		return rpc.NewGRPCServer(ctx, ctx.Config().GRPC)
	})
	app.Register( plugins.FOLLOW_SERVICE_NAME, func(ctx *node.ServiceContext) (node.Service, error) {
		return plugins.NewFollowService(ctx)
	})
	app.Register( plugins.POST_SERVICE_NAME, func(ctx *node.ServiceContext) (node.Service, error) {
		return plugins.NewPostService(ctx)
	})
	app.Register( plugins.DEMO_SERVICE_NAME, func(ctx *node.ServiceContext) (node.Service, error) {
		return plugins.NewDemoService(ctx)
	})

	if err := app.Start(); err != nil {
		common.Fatalf("start node failed, err: %v\n", err)
	}

	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigc)
		<-sigc
		logging.CLog().Infoln("Got interrupt, shutting down...")
		go app.Stop()
	}()

	app.Wait()
}
