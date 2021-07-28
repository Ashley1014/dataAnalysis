package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

type LogEvent struct {
	Id orm.CharField `orm:"size(100)"`
	Cps int
	UserId int `orm:"pk"`
	CreateTime orm.DateTimeField `orm:"null"`
	DeviceId string
	EventId int
	RegCps int
	Money int
	Scene int
	DataStr int
	FromReg int
	Version orm.CharField `orm:"size(100)"`
}

type EventDetail struct {
	EventId int
	EventName string `orm:"pk"`
	Version orm.CharField `orm:"size(100)"`
}

type StepInfo struct {
	Num int `form:"num"`
}

func fetchConfig(key string) string {
	config, _ := web.AppConfig.String(key)
	return config
}

func init() {
	driver := fetchConfig("driverName")
	sqlUser := fetchConfig("mysqlUser")
	sqlPass := fetchConfig("mysqlPass")
	sqlUrl := fetchConfig("mysqlUrl")
	dbname := fetchConfig("dbname")
	attr := fetchConfig("attr")
	config := sqlUser + ":" + sqlPass + "@tcp(" + sqlUrl + ")/" + dbname + attr
	err := orm.RegisterDataBase("default", driver, config)
	if err != nil {
		fmt.Println("failed to connect!")
		return
	}
	orm.RegisterModel(new(LogEvent), new(EventDetail))
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		err.Error()
		return
	}
}

func GetEvents(o orm.Ormer) *[]EventDetail {
	var events []EventDetail
	//var maps []orm.Params
	_, err := o.QueryTable("event_detail").All(&events)
	if err != nil {
		err.Error()
	}
	return &events
}

func CountRows(o orm.Ormer, fieldName string, eventId int) (int64, error) {
	var eventLogs []LogEvent
	sql := fmt.Sprintf("SELECT DISTINCT %s FROM log_event WHERE event_id = ?", fieldName)
	num, err := o.Raw(sql, eventId).QueryRows(&eventLogs)
	return num, err
}

func GetEventId(o orm.Ormer, eventName string) (int64, error) {
	var id int64
	var maps []orm.Params
	_, err := o.QueryTable("event_detail").Filter("EventName",eventName).Values(&maps)
	if err != nil {
		return 0, err
	}
	for _, m := range maps {
		if eid, ok := m["EventId"].(int64); ok {
			id = eid
		}
		break
	}

	return id, nil
}

func GetEventName(o orm.Ormer, eventId string) (string, error) {
	var name string
	var maps []orm.Params
	_, err := o.QueryTable("event_detail").Values(&maps)
	if err != nil {
		return "", err
	}
	for _, m := range maps {
		if m["EventId"] == eventId {
			if ename, ok := m["EventName"].(string); ok {
				name = ename
			}
			break
		}
	}
	return name, nil
}
