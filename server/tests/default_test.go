package test

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
	"time"
)

func init() {
}

type User struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}

func TestRe(t *testing.T) {
	var reName = `description\(([\S^\)]+?)\)`
	re := regexp.MustCompile(reName)
	// -1代表取全部
	results := re.FindAllStringSubmatch("description(你好呀);type(text)", -1)
	fmt.Printf(results[0][1])
}

func TestTag(t *testing.T) {
	rv := reflect.ValueOf(User{})
	typ := reflect.TypeOf(User{})
	for i := 0; i < rv.NumField(); i++ {
		field := typ.Field(i)
		typeStr := fmt.Sprintf("%v", field.Type)
		fmt.Printf("类型为:%v \n", typeStr)
	}
}
