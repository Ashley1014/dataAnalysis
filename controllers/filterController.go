package controllers

import (
	"dataAnalysis/models"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type FilterController struct {
	beego.Controller
}

func (c *FilterController) Get() {
	c.Data["num"] = 2
	c.TplName = "index.html"
	o := orm.NewOrm()
	c.Data["Events"] = models.GetEvents(o)
	c.Data["IsGet"] = 1
}

func (c *FilterController) Post() {
	c.TplName = "index.html"
	identifier := c.GetString("identifier")
	c.Data["IsGet"] = 0
	c.Data["numList"] = c.BatchGetData(identifier)
}

func (c *FilterController) BatchGetData(identifier string) *[]int {
	num, _ := c.GetInt("num")
	steps := make([]string, num)
	m := make(map[string]string)
	for i := 0; i <= num-1; i++ {
		fieldName := "step" + strconv.Itoa(i+1)
		eventName := "events" + strconv.Itoa(i+1)
		m[fieldName] = c.GetString(eventName)
		steps[i] = c.GetString(eventName)
	}
	numList := getNumList(identifier, steps)
	c.Data["num"] = num
	c.Data["eventMap"] = m
	return &numList

}

func getNumList(identifier string, steps []string) []int {
	o := orm.NewOrm()
	numList := make([]int, len(steps))
	for i, step := range steps {
		eventId, _ := models.GetEventId(o, step)
		num, _ := models.CountRows(o, identifier, int(eventId))
		numList[i] = int(num)
	}
	return numList
}


