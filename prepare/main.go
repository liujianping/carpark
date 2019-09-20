package prepare

import (
	"context"
	"log"
	"os"

	"github.com/liujianping/carpark/orm/model"
	"github.com/liujianping/carpark/svy21"
	"github.com/spf13/viper"
	"github.com/x-mod/cmd"
	"github.com/x-mod/errors"
)

//Main for prepare
func Main(c *cmd.Command, args []string) error {
	viper.SetConfigName("config")
	viper.AddConfigPath("etc")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return errors.Annotate(err, "read config")
	}

	if viper.GetString("csv-file") == "" {
		return errors.New("csv file required")
	}

	csv, err := os.OpenFile(viper.GetString("csv-file"), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return errors.Annotate(err, "open file")
	}
	defer csv.Close()

	items, err := Parse(csv)
	if err != nil {
		return err
	}

	model.MySQLSetup(&model.MySQLConfig{
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetInt("mysql.port"),
		UserName: viper.GetString("mysql.username"),
		Password: viper.GetString("mysql.password"),
		Database: viper.GetString("mysql.database"),
	})

	for _, item := range items {
		if err := store(item); err != nil {
			log.Println("ERR: csv item storing failed:", err)
		}
	}
	return nil
}

func store(item *CsvItem) error {
	obj := model.CarParkMgr.NewCarPark()
	obj.CarParkNo = item.CarParkNo
	obj.Address = item.Address
	//location svy21 convert
	lat, long := svy21.ToLatLon(item.Xcoord, item.Ycoord)
	obj.Latitude = lat
	obj.Longitude = long

	obj.CarParkType = item.CarParkType
	obj.TypeOfParkingSystem = item.TypeOfParkingSystem

	//available parking term parse
	obj.ShortTermParking = item.ShortTermParking
	from, to, err := parkingPeriod(item.ShortTermParking)
	if err != nil {
		return err
	}
	obj.ShortTermParkingFrom = int64(from.Seconds())
	obj.ShortTermParkingTo = int64(to.Seconds())

	//free parking to boolean
	obj.FreeParking = item.FreeParking
	if item.NightParking == "YES" {
		obj.NightParking = true
	}
	obj.CarParkDecks = item.CarParkDecks
	obj.GantryHeight = item.GantryHeight
	//basement to boolean
	if item.CarParkBasement != "N" {
		obj.CarParkBasement = true
	}

	tx, err := model.MySQL().BeginTx(context.TODO())
	if err != nil {
		return err
	}
	defer tx.Close()

	pk := model.CarParkMgr.NewPrimaryKey()
	pk.CarParkNo = item.CarParkNo
	exist, err := model.CarParkDBMgr(tx).Exist(pk)
	if err != nil {
		return errors.Annotatef(err, "%s exist", obj.CarParkNo)
	}
	if exist {
		if _, err := model.CarParkDBMgr(tx).Update(obj); err != nil {
			return errors.Annotatef(err, "%s update", obj.CarParkNo)
		}
	} else {
		if _, err := model.CarParkDBMgr(tx).Create(obj); err != nil {
			return errors.Annotatef(err, "%s create", obj.CarParkNo)
		}
	}
	return nil
}

func init() {
	c := cmd.Add(
		cmd.Path("/prepare"),
		cmd.Short("preprocess sg car park dataset"),
		cmd.Main(Main),
	)
	c.Flags().StringP("csv-file", "f", "", "csv file path")
	if err := viper.BindPFlags(c.Flags()); err != nil {
		log.Println("WARN: flag binding failed ", err)
	}
}
