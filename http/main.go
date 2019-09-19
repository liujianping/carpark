package http

import (
	"context"
	"time"

	"github.com/liujianping/carpark/orm/model"
	"github.com/spf13/viper"
	"github.com/x-mod/cmd"
	"github.com/x-mod/errors"
	"github.com/x-mod/httpserver"
	"github.com/x-mod/routine"
)

func init() {
	cmd.Add(
		cmd.Path("/http"),
		cmd.Short("http service"),
		cmd.Main(Main),
	)
}

//Main for http
func Main(c *cmd.Command, args []string) error {
	viper.SetConfigName("config")
	viper.AddConfigPath("etc")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return errors.Annotate(err, "read config")
	}

	model.MySQLSetup(&model.MySQLConfig{
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetInt("mysql.port"),
		UserName: viper.GetString("mysql.username"),
		Password: viper.GetString("mysql.password"),
		Database: viper.GetString("mysql.database"),
	})

	srv := NewCarParkServer(
		httpserver.NewServer(
			httpserver.Address(":8080"),
		),
	)
	if err := srv.Open(); err != nil {
		return errors.Annotate(err, "srv open")
	}
	return routine.Main(
		routine.ExecutorFunc(srv.Serve),
		routine.Interrupts(routine.DefaultCancelInterruptors...),
		routine.Cleanup(
			routine.ExecutorFunc(func(ctx context.Context) error {
				//graceful shutdown MaxTime 15s
				tmctx, cancel := context.WithTimeout(ctx, 15*time.Second)
				defer cancel()
				return srv.Shutdown(tmctx)
			})),
	)
}
