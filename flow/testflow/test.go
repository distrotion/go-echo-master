package test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	memList01 = []string{}
	memList02 = []string{}
)

func TESTPOST(c echo.Context) error {
	fmt.Println(`--TESTPOST--`)

	input := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&input)
	if err != nil {
		log.Error("empty json body")
		return nil
	}
	return c.JSON(200, input)
}

func TESTPOSTBSON(c echo.Context) error {
	fmt.Println(`--TESTPOSTBSON--`)

	input := make(bson.M)
	err := json.NewDecoder(c.Request().Body).Decode(&input)
	if err != nil {
		log.Error("empty json body")
		return nil
	}
	return c.JSON(200, input)
}

func TESTcrash(c echo.Context) error {
	fmt.Println(`--TESTPOST--`)
	//------------------------------------------------------
	input := toBSON(c)
	//
	inputUID := toDATABODY(input[`UID`])
	output := `NOK`
	//------------------------------------------------------------------------- airbag
	if stringInSlice(inputUID.Str, memList01) == false {
		memList01 = append(memList01, inputUID.Str)
		output = `OK`
		//----------------------------------------------------------------------- airbag

		time.Sleep(20 * time.Second)
		fmt.Println("Sleep Over.....")

		//----------------------------------------------------------------------- airbag
		list := remove(memList01, inputUID.Str)
		memList01 = list
	}

	// list := remove(memList01, inputUID.Str)
	// memList01 = list
	//------------------------------------------------------------------------- airbag
	//------------------------------------------------------
	return c.JSON(200, output)
}

func TESTcrashCLEAR_SUPER(c echo.Context) error {
	fmt.Println(`--TESTPOST--`)
	//------------------------------------------------------

	//
	memList01 = []string{}
	memList02 = []string{}

	output := `OK`
	//------------------------------------------------------------------------- airbag

	// list := remove(memList01, inputUID.Str)
	// memList01 = list
	//------------------------------------------------------------------------- airbag
	//------------------------------------------------------
	return c.JSON(200, output)
}

func UseSubroute(group *echo.Group) {
	group.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "TESTFLOW")
	})
	group.POST("/test", TESTPOST)
	group.POST("/testBSON", TESTPOSTBSON)
	group.POST("/TESTcrash", TESTcrash)
	group.POST("/TESTcrashCLEAR_SUPER", TESTcrashCLEAR_SUPER)
}

//insClinicProB := maindbv2.UpdateArchive(ctx, authDB, authDB_clinic_profile_model, bson.M{"SUID": sUID}, ClinicProB)
//insClinicProMoB := maindbv2.Insertdb(ctx, authDB, authDB_clinic_profile_model, ClinicProMoB)
//dbs1S := maindbv2.Finddb(ctx, staffDB, staffDB_profile_model, bson.M{"SUID": UIDSinput}, "_id", 1, 0, 0)
//SSessionUIDSA := jwt.GenerateToken(fmt.Sprintf("%v", RsCaregiver_At_Patientbuffer[j])+`<>`+strconv.FormatInt(t1, 10), 7)
// SessionUIDS := jwt.ParseToken(input.SSessionUIDs, 7)
// t1 := time.Now().Unix()
// t2 := time.Now().UnixNano()
//String >> fmt.Sprintf("%v", listPatient[i][`SFirstName`])
//int32 >>  int(listPatient[i][`NImgAvatarID`].(int32))

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func remove(items []string, item string) []string {
	newitems := []string{}

	for _, i := range items {
		if i != item {
			newitems = append(newitems, i)
		}
	}

	return newitems
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		b[i] = letterRunes[r1.Intn(len(letterRunes))]
	}
	return string(b)
}

func intoliin(t interface{}) []int {
	var output []int
	if t != nil {
		switch reflect.TypeOf(t).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				// fmt.Println(s.Index(i))
				output = append(output, int(s.Index(i).Interface().(int32)))
			}
		}
	} else {
		return output
	}

	return output
}

func intolistr(t interface{}) []string {
	var output []string
	if t != nil {
		switch reflect.TypeOf(t).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				// fmt.Println(s.Index(i))
				output = append(output, fmt.Sprintf("%v", s.Index(i).Interface()))
			}
		}
	} else {
		return output
	}

	return output
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func removeDuplicates(strList []string) []string {
	list := []string{}
	if strList != nil {
		for _, item := range strList {
			fmt.Println(item)
			if contains(list, item) == false {
				list = append(list, item)
			}
		}
	} else {
		return list
	}

	return list
}

func intolistBsonM(t interface{}) []bson.M {
	var output []bson.M
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)
		for i := 0; i < s.Len(); i++ {
			// fmt.Println(s.Index(i))
			//------------------------------
			bs, err := json.Marshal(s.Index(i).Interface())
			if err != nil {
				panic(err)
			}
			var o bson.M
			if err := json.Unmarshal(bs, &o); err != nil {
				panic(err)
			}
			//------------------------------
			output = append(output, o)
		}
	}
	return output
}

func primitivetoBsonM(input interface{}, my string) []bson.M {
	bs, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	var o bson.M
	if err := json.Unmarshal(bs, &o); err != nil {
		panic(err)
	}

	if o[my] == nil {
		var output []bson.M
		return output

	} else {
		output := intolistBsonM(o[my])
		return output
	}

}

func jsonToMap(jsonStr string) map[string]interface{} {
	result := make(map[string]interface{})
	json.Unmarshal([]byte(jsonStr), &result)
	return result
}

func UNIXstrToDayOfWeek(input string) time.Duration {
	i, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	// fmt.Println(tm)
	return GetStartDayOfWeek(tm)
}

func GetStartDayOfWeek(tm time.Time) time.Duration { //get monday 00:00:00
	weekday := time.Duration(tm.Weekday())
	// if weekday == 0 {
	// 	weekday = 7
	// }
	// year, month, day := tm.Date()
	// currentZeroDay := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return weekday
}

func UNIXstrToYear(input string) time.Duration {
	i, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	// fmt.Println(tm)
	return time.Duration(tm.Year())
}

type DATABODY struct {
	Str   string `json:"str"`
	Inte  int64  `json:"inte"`
	Bools bool   `json:"bools"`
}

func toDATABODY(c interface{}) DATABODY {
	var msg DATABODY

	itf := intoBsonM(c)

	intVar, _ := strconv.ParseInt(fmt.Sprintf("%v", itf[`inte`]), 0, 64)
	boolVar := strings.ToUpper(fmt.Sprintf("%v", itf[`bools`]))
	msg.Str = fmt.Sprintf("%v", itf[`str`])
	msg.Inte = intVar
	fmt.Println(boolVar)
	if boolVar == `TRUE` {
		msg.Bools = true
	} else {
		msg.Bools = false
	}

	return msg
}

func intoBsonM(t interface{}) bson.M {
	var output bson.M
	switch reflect.TypeOf(t).Kind() {
	case reflect.Map:
		s := reflect.ValueOf(t)
		bs, err := json.Marshal(s.Interface())
		if err != nil {
			panic(err)
		}
		var o bson.M
		if err := json.Unmarshal(bs, &o); err != nil {
			panic(err)
		}
		//------------------------------
		output = o

	}
	return output
}

func toBSON(c echo.Context) bson.M {

	input := make(bson.M)
	err := json.NewDecoder(c.Request().Body).Decode(&input)
	if err != nil {
		log.Error("empty json body")
	}

	return input
}

func toINTERFACE(c echo.Context) map[string]interface{} {

	input := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&input)
	if err != nil {
		log.Error("empty json body")
	}

	return input
}
