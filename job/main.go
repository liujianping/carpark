package job

import (
	"context"
	"log"

	"github.com/liujianping/carpark/orm/model"
	"github.com/spf13/viper"
	"github.com/x-mod/cmd"
	"github.com/x-mod/errors"
	"github.com/x-mod/routine"
)

func init() {
	c := cmd.Add(
		cmd.Path("/job"),
		cmd.Short("job to refresh carpark status"),
		cmd.Main(Main),
	)
	c.Flags().StringP("crontab", "c", "* * * * *", "crontab schedule")
	if err := viper.BindPFlags(c.Flags()); err != nil {
		log.Println("WARN: flag binding failed ", err)
	}
}

//Main for job
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

	//job
	job := NewJOB(nil)

	return routine.Main(
		routine.ExecutorFunc(func(ctx context.Context) error {
			//start
			if err := job.Execute(ctx); err != nil {
				return err
			}
			//job
			return <-routine.Go(ctx,
				routine.Crontab(
					viper.GetString("crontab"),
					routine.ExecutorFunc(job.Execute),
				),
			)
		}),
		routine.Interrupts(routine.DefaultCancelInterruptors...),
	)
}
