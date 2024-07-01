package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"

	"github.com/ClickHouse/clickhouse-go/v2"
)

var dialCount int = 0

// Get conn
func Click() (driver.Conn, error) {

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		DialContext: func(ctx context.Context, addr string) (net.Conn, error) {
			dialCount++
			var d net.Dialer
			return d.DialContext(ctx, "tcp", addr)
		},
		Debug: true,
		Debugf: func(format string, v ...any) {
			fmt.Printf(format+"\n", v...)
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		DialTimeout:          time.Second * 30,
		MaxOpenConns:         5,
		MaxIdleConns:         5,
		ConnMaxLifetime:      time.Duration(10) * time.Minute,
		ConnOpenStrategy:     clickhouse.ConnOpenInOrder,
		BlockBufferSize:      10,
		MaxCompressionBuffer: 10240,
		ClientInfo: clickhouse.ClientInfo{ // optional, please see Client info section in the README.md
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "my-app", Version: "0.1"},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return conn, conn.Ping(context.Background())

}

// Errorprone Get conn
type ClickhouseCloseFunc func() error

func ErrorproneGetConnect() (driver.Conn, ClickhouseCloseFunc, error) {
	conn, err := Click()
	return conn, ClickhouseCloseFunc(conn.Close), err
}
func main() {
	// Get conn
	conn, close, err := ErrorproneGetConnect()
	defer func() {
		err := close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		panic(err)
	}
	// Send to clickhouse
	BatchSend(conn)

	// Get from clickhouse

	QueryOne(conn)

	// Get a lot from clickhouse

	QueryALot(conn)

}
func BatchSend(conn driver.Conn) {
	batch, err := conn.PrepareBatch(context.Background(), "INSERT INTO history.history_order(client_name,time_placed)")

	if err != nil {
		panic(err)
	}

	for i := range 5 {
		//err = batch.Append("test"+strconv.Itoa(i), time.Now().UnixNano())
		time.Sleep(time.Millisecond * 10)
		err = batch.AppendStruct(&struct {
			Client_name string    `ch:"client_name"`
			Time_placed time.Time `ch:"time_placed"`
		}{
			Client_name: "test" + strconv.Itoa(i),
			Time_placed: time.Now(),
		})
		if err != nil {
			panic(err)
		}
	}
	err = batch.Send()
	if err != nil {
		panic(err)
	}
}

// Using DateTime64(3) specific
func QueryOne(conn driver.Conn) {
	chCtx := clickhouse.Context(context.Background(), clickhouse.WithParameters(clickhouse.Parameters{
		"time": "2188-07-24 06:12:15.036",
	}))
	row := conn.QueryRow(chCtx,
		`SELECT client_name AS cname
		 FROM history.history_order 
		 WHERE time_placed = {time:DateTime64(3)} `)
	if err := row.Err(); err != nil {
		log.Println("__Error", err)
	}
	var wtf struct {
		Client_name string `ch:"cname"`
	}
	err := row.ScanStruct(&wtf)
	if err != nil {
		log.Println("__Error", err)
	}
	log.Println("get row", row)
}
func QueryALot(conn driver.Conn) {
	rows, err := conn.Query(context.Background(), "SELECT client_name FROM history.history_order")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	clientNames := make([]struct {
		Client_name string `ch:"client_name"`
	}, 0)
	for i := 0; rows.Next(); i++ {
		clientNames = append(clientNames, struct {
			Client_name string `ch:"client_name"`
		}{})
		rows.ScanStruct(&clientNames[i])
	}
	log.Println(clientNames)
}
