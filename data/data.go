package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	//connectStr       = "root:root@(127.0.0.1)/blog?charset=utf8&parseTime=True&loc=Local"
	TIME_FORMAT_TIME = "2006-01-02 15:04:05"
	TIME_FORMAT_DATE = "2006-01-02"
)

var (
	DB *gorm.DB
)

type DBConfig struct {
	User string	`json:"db_user"`
	Pass string	`json:"db_pass"`
	Addr string `json:"db_addr"`
	Port int `json:"db_port"`
	Name string `json:"db_name"`
	Args string `json:"connect_args"`
}

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
	v := DBConfig{}
	err1 := LoadConf("./db.json",&v)
	if err1 != nil {
		log.Fatal(err1)
	}
	//fmt.Printf("%#v",v)
	str := GenerateConnectStr(v)
	var err error
	DB, err = gorm.Open("mysql", str)
	if err != nil {
		log.Fatal(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "go_" + defaultTableName
	}
}

func LoadConf(fi string,v interface{}) error {
	data,err := ioutil.ReadFile(fi)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data,v)
	if err != nil {
		return err
	}
	return nil
}

func GenerateConnectStr(c DBConfig) string {
	return fmt.Sprint(c.User ,":", c.Pass , "@(",c.Addr , ":",c.Port , ")/" + c.Name + "?"+c.Args)
}
