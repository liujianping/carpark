package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/auto-program/db-orm/orm"
	"github.com/liujianping/carpark/orm/model"
	"github.com/x-mod/errors"
	"github.com/x-mod/httpserver"
	"github.com/x-mod/httpserver/render"
)

//CarParkServer define
type CarParkServer struct {
	*httpserver.Server
}

//NewCarParkServer create a new CarParkServer
func NewCarParkServer(srv *httpserver.Server) *CarParkServer {
	return &CarParkServer{
		Server: srv,
	}
}

//Open register route
func (srv *CarParkServer) Open() error {
	srv.Route(
		httpserver.Method("GET"),
		httpserver.Pattern("/carparks/nearest"),
		httpserver.Handler(srv.Nearest),
	)
	return nil
}

//Nearest handler
func (srv *CarParkServer) Nearest(ctx context.Context, wr http.ResponseWriter, req *http.Request) {
	slat := req.URL.Query().Get("latitude")
	slng := req.URL.Query().Get("longitude")
	sPage := req.URL.Query().Get("page")
	sPerPage := req.URL.Query().Get("per_page")
	if len(slat) == 0 {
		WARN(render.Error(errors.New("latitude required")).Response(wr, render.StatusCode(400)))
		return
	}
	if len(slng) == 0 {
		WARN(render.Error(errors.New("longitude required")).Response(wr, render.StatusCode(400)))
		return
	}
	lat, err := strconv.ParseFloat(slat, 64)
	if err != nil {
		WARN(render.Error(errors.Annotate(err, "latitude parse")).Response(wr, render.StatusCode(400)))
		return
	}
	lng, err := strconv.ParseFloat(slng, 64)
	if err != nil {
		WARN(render.Error(errors.Annotate(err, "longitude parse")).Response(wr, render.StatusCode(400)))
		return
	}
	page := 0
	if len(sPage) != 0 {
		num, ierr := strconv.ParseInt(sPage, 10, 64)
		if ierr != nil {
			WARN(render.Error(errors.Annotate(ierr, "page parse")).Response(wr, render.StatusCode(400)))
			return
		}
		page = int(num)
	}
	perPage := 10
	if len(sPerPage) != 0 {
		num, ierr := strconv.ParseInt(sPerPage, 10, 64)
		if ierr != nil {
			WARN(render.Error(errors.Annotate(ierr, "per_page parse")).Response(wr, render.StatusCode(400)))
			return
		}
		perPage = int(num)
	}

	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	stmt := fmt.Sprintf(`SELECT a.address, a.latitude, a.longitude, b.total_lots, b.available_lots
		  FROM carparks a inner join carpark_status b on a.car_park_no = b.car_park_no 
		  WHERE b.available_lots > 0 AND %d BETWEEN a.short_term_parking_from AND a.short_term_parking_to 
		  ORDER BY ST_Distance_Sphere(POINT(a.longitude,a.latitude), POINT(%f, %f)) ASC %s`,
		int64(now.Sub(start).Seconds()), lng, lat, orm.SQLOffsetLimit(page*perPage, perPage),
	)
	objs, err := model.NearestParkInfoDBMgr(model.MySQL()).QueryBySQL(stmt)
	if err != nil {
		WARN(render.Error(errors.Annotate(err, "query")).Response(wr, render.StatusCode(400)))
		return
	}
	// log.Println("stmt:", stmt)
	WARN(render.JSON(objs).Response(wr))
}

//WARN log warn error
func WARN(err error) {
	if err != nil {
		log.Println("WARN: ", err)
	}
}
