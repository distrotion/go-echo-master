package test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
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

func UseSubroute(group *echo.Group) {
	group.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "TESTFLOW")
	})
	group.POST("/test", TESTPOST)
	group.POST("/testBSON", TESTPOSTBSON)
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
