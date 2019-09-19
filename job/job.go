package job

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/liujianping/carpark/orm/model"
	"github.com/x-mod/errors"
	"github.com/x-mod/httpclient"
)

//ParkInfo struct
type ParkInfo struct {
	TotalLots     string `json:"total_lots"`
	LotType       string `json:"lot_type"`
	LotsAvailable string `json:"lots_available"`
}

//Park struct
type Park struct {
	Infos      []*ParkInfo `json:"carpark_info"`
	ParkNo     string      `json:"carpark_number"`
	UpdateTime string      `json:"update_datetime"`
}

//QueryItem struct
type QueryItem struct {
	Timestamp string  `json:"timestamp"`
	Parks     []*Park `json:"carpark_data"`
}

//Result struct
type Result struct {
	Items []*QueryItem `json:"items"`
}

//JOB define
type JOB struct {
	*httpclient.Client
	sg *time.Location
}

//NewJOB create a new job
func NewJOB(c *httpclient.Client) *JOB {
	client := httpclient.New()
	if c != nil {
		client = c
	}
	sg, err := time.LoadLocation("Asia/Singapore")
	if err != nil {
		panic(err)
	}
	return &JOB{
		Client: client,
		sg:     sg,
	}
}

//Execute job execute once
func (job *JOB) Execute(ctx context.Context) error {
	//Host timezone set Asia/Singapore
	req, err := httpclient.MakeRequest(
		httpclient.Method("GET"),
		httpclient.URL(
			httpclient.Schema("https"),
			httpclient.Host("api.data.gov.sg"),
			httpclient.URI("/v1/transport/carpark-availability"),
		),
		httpclient.Query("date_time", time.Now().Format(time.RFC3339)),
	)
	if err != nil {
		return errors.Annotate(err, "make request")
	}
	return job.Client.Execute(ctx, req, job)
}

//Process job process http.Response
func (job *JOB) Process(ctx context.Context, rsp *http.Response) error {
	if rsp.StatusCode != http.StatusOK {
		return errors.New(rsp.Status)
	}
	defer rsp.Body.Close()

	var result Result
	if err := json.NewDecoder(rsp.Body).Decode(&result); err != nil {
		return errors.Annotate(err, "json decode")
	}
	log.Println("Parks count:", len(result.Items[0].Parks))
	for _, item := range result.Items {
		for _, park := range item.Parks {
			if err := job.updateParkStatus(ctx, park); err != nil {
				log.Println("WARN: park status refresh failed: ", err)
				continue
			}
		}
	}
	return nil
}

func (job *JOB) updateParkStatus(ctx context.Context, park *Park) error {
	tx, err := model.MySQL().BeginTx(ctx)
	if err != nil {
		return errors.Annotate(err, "begin tx")
	}
	defer tx.Close()

	pk := model.CarParkStatusMgr.NewPrimaryKey()
	pk.CarParkNo = park.ParkNo

	exist, err := model.CarParkStatusDBMgr(tx).Exist(pk)
	if err != nil {
		return errors.Annotatef(err, "%s exist", pk.CarParkNo)
	}

	obj := model.CarParkStatusMgr.NewCarParkStatus()

	obj.CarParkNo = park.ParkNo
	t1, err := job.parseSGTime(park.UpdateTime)
	if err != nil {
		return errors.Annotate(err, "parse sg time")
	}
	obj.ReportAt = t1.Unix()
	total := 0
	avail := 0
	for _, info := range park.Infos {
		t1, err := strconv.Atoi(info.TotalLots)
		if err != nil {
			return errors.Annotatef(err, "strconv %s", info.TotalLots)
		}
		t2, err := strconv.Atoi(info.LotsAvailable)
		if err != nil {
			return errors.Annotatef(err, "strconv %s", info.TotalLots)
		}
		total += t1
		avail += t2
	}
	obj.TotalLots = int32(total)
	obj.AvailableLots = int32(avail)
	//create
	if !exist {
		obj.CreatedAt = time.Now().Unix()
		if _, err := model.CarParkStatusDBMgr(tx).Create(obj); err != nil {
			return errors.Annotatef(err, "%s create", obj.CarParkNo)
		}
		return nil
	}

	//update
	if _, err := model.CarParkStatusDBMgr(tx).UpdateBySQL(
		"total_lots = ?, available_lots = ?, report_at = ?",
		"car_park_no = ? AND report_at < ?",
		obj.TotalLots, obj.AvailableLots, obj.ReportAt,
		obj.CarParkNo, obj.ReportAt,
	); err != nil {
		return errors.Annotatef(err, "%s update", obj.CarParkNo)
	}
	return nil
}

func (job *JOB) parseSGTime(dest string) (time.Time, error) {
	layout := "2006-01-02T15:04:05"
	return time.ParseInLocation(layout, dest, job.sg)
}
