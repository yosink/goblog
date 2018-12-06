package data

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	connectStr       = "root:root@(127.0.0.1)/default?charset=utf8&parseTime=True&loc=Local"
	TIME_FORMAT_TIME = "2006-01-02 15:04:05"
	TIME_FORMAT_DATE = "2006-01-02"
)

var (
	DB *gorm.DB
)

type LocalTime time.Time

// MarshalJSON 重写格式化为json的方法
func (l LocalTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(l).Format(TIME_FORMAT_TIME))
	return []byte(stamp), nil
}

type LocalDate time.Time

func (l LocalDate) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(l).Format(TIME_FORMAT_DATE))
	return []byte(stamp), nil
}

func init() {
	var err error
	DB, err = gorm.Open("mysql", connectStr)
	if err != nil {
		log.Fatal(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "go_" + defaultTableName
	}
}
