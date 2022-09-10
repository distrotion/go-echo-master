package maindb01

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var server = "mongodb://127.0.0.1:15000" mongodb+srv://testdb:testdb@cluster0.2xffgea.mongodb.net/?retryWrites=true&w=majority
var server = "mongodb+srv://testdb:testdb@cluster0.2xffgea.mongodb.net/?retryWrites=true&w=majority"
var server2 = "mongodb+srv://testdb:testdb@cluster0.2xffgea.mongodb.net/?retryWrites=true&w=majority"
var server3 = "mongodb+srv://testdb:testdb@cluster0.2xffgea.mongodb.net/?retryWrites=true&w=majority"

func Updateonly(ctx context.Context, db_mongo_u string, collec_u string, input1 bson.M, input2 bson.M) string {

	time := time.Now().Add(time.Minute).Unix()
	clientOptions_u := options.Client().ApplyURI(server)
	client_u, err := mongo.Connect(ctx, clientOptions_u)
	if err != nil {
		log.Fatal(err)
	}
	err = client_u.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client_u.Disconnect(ctx)

	collection_u := client_u.Database(db_mongo_u).Collection(collec_u)
	// collection_i := client_u.Database(db_mongo_u).Collection(collec_u + `_Archive`)

	//cur, currErr := collection.Find(ctx, bson.M{})
	cur, currErr := collection_u.Find(ctx, input1)
	if currErr != nil {
		panic(currErr)
	}
	// fmt.Println(cur)

	var msg []bson.M

	if err = cur.All(ctx, &msg); err != nil {
		panic(err)
	}
	// fmt.Println(msg)

	delete(msg[0], "_id")

	// res_i, insertErr := collection_i.InsertOne(ctx, msg[0])
	// if insertErr != nil {
	// 	return `nok`
	// }
	// fmt.Println(res_i)
	input2[`timestamp`] = time
	res_u, insertErr := collection_u.UpdateOne(ctx, input1, bson.M{"$set": input2})
	if insertErr != nil {
		client_u.Disconnect(ctx)
		return `nok`
	}
	fmt.Println(res_u)
	client_u.Disconnect(ctx)
	return `ok`
}

func UpdateArchive(ctx context.Context, db_mongo_u string, collec_u string, input1 bson.M, input2 bson.M) string {

	time := time.Now().Add(time.Minute).Unix()
	clientOptions_u := options.Client().ApplyURI(server)
	client_u, err := mongo.Connect(ctx, clientOptions_u)
	if err != nil {
		log.Fatal(err)
	}
	err = client_u.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client_u.Disconnect(ctx)
	//---------------------------
	clientOptions_u2 := options.Client().ApplyURI(server2)
	client_u2, err2 := mongo.Connect(ctx, clientOptions_u2)
	if err2 != nil {
		log.Fatal(err2)
	}
	err = client_u2.Ping(ctx, nil)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer client_u2.Disconnect(ctx)
	//-----------------------------

	collection_u := client_u.Database(db_mongo_u).Collection(collec_u)
	collection_i := client_u2.Database(db_mongo_u + `_Archive`).Collection(collec_u + `_Archive`)

	//cur, currErr := collection.Find(ctx, bson.M{})
	cur, currErr := collection_u.Find(ctx, input1)
	if currErr != nil {
		panic(currErr)
	}
	// fmt.Println(cur)

	var msg []bson.M

	if err = cur.All(ctx, &msg); err != nil {
		panic(err)
	}
	// fmt.Println(msg)

	delete(msg[0], "_id")

	res_i, insertErr := collection_i.InsertOne(ctx, msg[0])
	if insertErr != nil {
		client_u.Disconnect(ctx)
		client_u2.Disconnect(ctx)
		return `nok`
	}
	fmt.Println(res_i)
	input2[`timestamp`] = time
	res_u, insertErr := collection_u.UpdateOne(ctx, input1, bson.M{"$set": input2})
	if insertErr != nil {
		client_u.Disconnect(ctx)
		client_u2.Disconnect(ctx)
		return `nok`
	}
	fmt.Println(res_u)
	client_u.Disconnect(ctx)
	client_u2.Disconnect(ctx)
	return `ok`

}

func Insertdb(ctx context.Context, db_mongo_u string, collec_u string, input1 bson.M) string {

	// db_mongo_u = "auth_main_demo2_07_21"
	// collec_u = "users_main_demo2_07_21"
	time := time.Now().Add(time.Minute).Unix()
	// var ctx = context.TODO()
	clientOptions_u := options.Client().ApplyURI(server)
	client_u, err := mongo.Connect(ctx, clientOptions_u)
	if err != nil {
		log.Fatal(err)
	}
	err = client_u.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client_u.Disconnect(ctx)

	collection_u := client_u.Database(db_mongo_u).Collection(collec_u)
	input1[`timestamp`] = time
	res_u, insertErr := collection_u.InsertOne(ctx, input1)
	if insertErr != nil {
		client_u.Disconnect(ctx)
		return `nok`
	}
	fmt.Println(res_u)
	client_u.Disconnect(ctx)
	return `ok`
}

func Finddb(ctx context.Context, db_mongo_u string, collec_u string, input1 bson.M, sortby string, sortorder int, limmit int64, skip int64) []bson.M {

	// db_mongo_u = "auth_main_demo2_07_21"
	// collec_u = "users_main_demo2_07_21"
	// var ctx = context.TODO()
	clientOptions_u := options.Client().ApplyURI(server)
	client_u, err := mongo.Connect(ctx, clientOptions_u)
	if err != nil {
		log.Fatal(err)
	}
	err = client_u.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client_u.Disconnect(ctx)

	collection_u := client_u.Database(db_mongo_u).Collection(collec_u)

	opts := options.Find()
	opts.SetSort(bson.D{{sortby, sortorder}}).SetLimit(limmit).SetSkip(skip)
	cur, err := collection_u.Find(ctx, input1, opts)
	//cur, err := collection.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		var msg2 []bson.M
		client_u.Disconnect(ctx)
		return msg2
	}
	var msg []bson.M
	if err = cur.All(ctx, &msg); err != nil {
		panic(err)
	}
	client_u.Disconnect(ctx)
	return msg
}

func FinddbST(ctx context.Context, db_mongo_u string, collec_u string, input1 bson.M, sortby string, sortorder int, limmit int64, skip int64, Pro bson.M) []bson.M {

	// db_mongo_u = "auth_main_demo2_07_21"
	// collec_u = "users_main_demo2_07_21"
	// var ctx = context.TODO()
	clientOptions_u := options.Client().ApplyURI(server)
	client_u, err := mongo.Connect(ctx, clientOptions_u)
	if err != nil {
		log.Fatal(err)
	}
	err = client_u.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client_u.Disconnect(ctx)

	collection_u := client_u.Database(db_mongo_u).Collection(collec_u)

	opts := options.Find()
	opts.SetSort(bson.D{{sortby, sortorder}}).SetLimit(limmit).SetSkip(skip).SetProjection(Pro)
	cur, err := collection_u.Find(ctx, input1, opts)
	//cur, err := collection.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		var msg2 []bson.M
		client_u.Disconnect(ctx)
		return msg2
	}
	var msg []bson.M
	if err = cur.All(ctx, &msg); err != nil {
		panic(err)
	}
	client_u.Disconnect(ctx)
	return msg
}

func UpdatePushArray(ctx context.Context, db_mongo_u string, collec_u string, input1 bson.M, input2 bson.M, input3 string) string {

	// Month := time.Now().Month()
	// Year := time.Now().Year()
	clientOptions_u := options.Client().ApplyURI(server)
	client_u, err := mongo.Connect(ctx, clientOptions_u)
	if err != nil {
		log.Fatal(err)
	}
	err = client_u.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client_u.Disconnect(ctx)

	collection_u := client_u.Database(db_mongo_u).Collection(collec_u)

	//----------------------------------------------------------------------

	cur, err := collection_u.Find(ctx, input1)
	var msg []bson.M
	if err = cur.All(ctx, &msg); err != nil {
		panic(err)
	}

	if len(msg) > 0 {
		fmt.Println(`have exited`)
	} else {
		res_ins, insertErr := collection_u.InsertOne(ctx, input1)
		if insertErr != nil {
			client_u.Disconnect(ctx)
			return `nok`
		}
		fmt.Println(res_ins)
	}

	Month := int(time.Now().Month())
	Year := time.Now().Year()
	setdata := input3 + `.` + strconv.Itoa(Month) + `-` + strconv.Itoa(Year)

	fmt.Println(setdata)
	res_u, insertErr := collection_u.UpdateOne(ctx, input1, bson.M{"$push": bson.M{setdata: input2}})
	if insertErr != nil {
		client_u.Disconnect(ctx)
		return `nok`
	}
	fmt.Println(res_u)
	client_u.Disconnect(ctx)
	return `ok`
}

func UpdatePushArraycus(ctx context.Context, db_mongo_u string, collec_u string, input1 bson.M, input2 bson.M, input3 string) string {

	// Month := time.Now().Month()
	// Year := time.Now().Year()
	clientOptions_u := options.Client().ApplyURI(server)
	client_u, err := mongo.Connect(ctx, clientOptions_u)
	if err != nil {
		log.Fatal(err)
	}
	err = client_u.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client_u.Disconnect(ctx)

	collection_u := client_u.Database(db_mongo_u).Collection(collec_u)

	//----------------------------------------------------------------------

	cur, err := collection_u.Find(ctx, input1)
	var msg []bson.M
	if err = cur.All(ctx, &msg); err != nil {
		panic(err)
	}

	if len(msg) > 0 {
		fmt.Println(`have exited`)
	} else {
		res_ins, insertErr := collection_u.InsertOne(ctx, input1)
		if insertErr != nil {
			client_u.Disconnect(ctx)
			return `nok`
		}
		fmt.Println(res_ins)
	}

	// Month := int(time.Now().Month())
	// Year := time.Now().Year()
	// setdata := strconv.Itoa(Month) + `-` + strconv.Itoa(Year)

	// fmt.Println(setdata)
	res_u, insertErr := collection_u.UpdateOne(ctx, input1, bson.M{"$push": bson.M{input3: input2}})
	if insertErr != nil {
		client_u.Disconnect(ctx)
		return `nok`
	}
	fmt.Println(res_u)
	client_u.Disconnect(ctx)
	return `ok`
}

func UpdatePushArraycusARC(ctx context.Context, db_mongo_u string, collec_u string, input1 bson.M, input2 bson.M, input3 string) string {

	// Month := time.Now().Month()
	// Year := time.Now().Year()
	clientOptions_u := options.Client().ApplyURI(server)
	client_u, err := mongo.Connect(ctx, clientOptions_u)
	if err != nil {
		log.Fatal(err)
	}
	err = client_u.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client_u.Disconnect(ctx)

	collection_u := client_u.Database(db_mongo_u).Collection(collec_u)

	//----------------------------------------------------------------------

	clientOptions_u2 := options.Client().ApplyURI(server2)
	client_u2, err2 := mongo.Connect(ctx, clientOptions_u2)
	if err2 != nil {
		log.Fatal(err2)
	}
	err = client_u2.Ping(ctx, nil)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer client_u2.Disconnect(ctx)

	collection_i := client_u2.Database(db_mongo_u + `_Archive`).Collection(collec_u + `_Archive`)
	//-----------------------------

	cur, err := collection_u.Find(ctx, input1)
	var msg []bson.M
	if err = cur.All(ctx, &msg); err != nil {
		panic(err)
	}

	if len(msg) > 0 {
		fmt.Println(`have exited`)
	} else {
		res_ins, insertErr := collection_u.InsertOne(ctx, input1)

		if insertErr != nil {
			client_u.Disconnect(ctx)
			client_u2.Disconnect(ctx)
			return `nok`
		}
		fmt.Println(res_ins)
	}

	// Month := int(time.Now().Month())
	// Year := time.Now().Year()
	// setdata := strconv.Itoa(Month) + `-` + strconv.Itoa(Year)

	// fmt.Println(setdata)
	delete(msg[0], "_id")

	res_i, insertErr := collection_i.InsertOne(ctx, msg[0])
	if insertErr != nil {

		client_u.Disconnect(ctx)
		client_u2.Disconnect(ctx)
		return `nok`
	}

	fmt.Println(res_i)

	res_u, insertErr := collection_u.UpdateOne(ctx, input1, bson.M{"$push": bson.M{input3: input2}})
	if insertErr != nil {
		client_u.Disconnect(ctx)
		client_u2.Disconnect(ctx)
		return `nok`
	}
	fmt.Println(res_u)
	client_u.Disconnect(ctx)
	client_u2.Disconnect(ctx)
	return `ok`
}

func Findonly(ctx context.Context, db_mongo_u string, collec_u string, input1 bson.M, key string) []bson.M {

	// db_mongo_u = "auth_main_demo2_07_21"
	// collec_u = "users_main_demo2_07_21"
	// var ctx = context.TODO()
	clientOptions_u := options.Client().ApplyURI(server)
	client_u, err := mongo.Connect(ctx, clientOptions_u)
	if err != nil {
		log.Fatal(err)
	}
	err = client_u.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client_u.Disconnect(ctx)

	collection_u := client_u.Database(db_mongo_u).Collection(collec_u)

	opts := options.Find().SetProjection(bson.M{key: 1, "_id": 0})
	// opts.SetSort(bson.D{{sortby, sortorder}}).SetLimit(limmit).SetSkip(skip)

	cur, err := collection_u.Find(ctx, input1, opts)
	//cur, err := collection.Find(ctx, bson.D{{}}, opts)

	fmt.Println(cur)
	if err != nil {
		var msg2 []bson.M
		client_u.Disconnect(ctx)
		return msg2
	}
	var msg []bson.M
	if err = cur.All(ctx, &msg); err != nil {
		panic(err)
	}
	client_u.Disconnect(ctx)
	return msg
}

func Findmutikey(ctx context.Context, db_mongo_u string, collec_u string, input1 bson.M, key []string) []bson.M {

	// db_mongo_u = "auth_main_demo2_07_21"
	// collec_u = "users_main_demo2_07_21"
	// var ctx = context.TODO()
	clientOptions_u := options.Client().ApplyURI(server)
	client_u, err := mongo.Connect(ctx, clientOptions_u)
	if err != nil {
		log.Fatal(err)
	}
	err = client_u.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client_u.Disconnect(ctx)

	collection_u := client_u.Database(db_mongo_u).Collection(collec_u)

	dataneed := make(bson.M)
	dataneed[`_id`] = 0

	for i := 0; i < len(key); i++ {
		dataneed[key[i]] = 1
	}

	fmt.Println(dataneed)

	opts := options.Find().SetProjection(dataneed)
	// opts.SetSort(bson.D{{sortby, sortorder}}).SetLimit(limmit).SetSkip(skip)

	cur, err := collection_u.Find(ctx, input1, opts)
	//cur, err := collection.Find(ctx, bson.D{{}}, opts)

	fmt.Println(cur)
	if err != nil {
		var msg2 []bson.M
		client_u.Disconnect(ctx)
		return msg2
	}
	var msg []bson.M
	if err = cur.All(ctx, &msg); err != nil {
		panic(err)
	}
	client_u.Disconnect(ctx)
	return msg
}

func MOVEdb(ctx context.Context, db_mongo_u string, collec_u string, collec_new string, input1 bson.M) string {

	clientOptions_u := options.Client().ApplyURI(server)
	client_u, err := mongo.Connect(ctx, clientOptions_u)
	if err != nil {
		log.Fatal(err)
	}
	err = client_u.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client_u.Disconnect(ctx)
	//---------------------------
	// clientOptions_u2 := options.Client().ApplyURI(server)
	// client_u2, err2 := mongo.Connect(ctx, clientOptions_u2)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
	// err = client_u2.Ping(ctx, nil)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
	// defer client_u2.Disconnect(ctx)
	//-----------------------------

	collection_u := client_u.Database(db_mongo_u).Collection(collec_u)
	collection_i := client_u.Database(db_mongo_u).Collection(collec_new)

	//cur, currErr := collection.Find(ctx, bson.M{})
	cur, currErr := collection_u.Find(ctx, input1)
	if currErr != nil {
		panic(currErr)
	}
	// fmt.Println(cur)

	var msg []bson.M

	if err = cur.All(ctx, &msg); err != nil {
		panic(err)
	}
	// fmt.Println(msg)

	delete(msg[0], "_id")

	res_i, insertErr := collection_i.InsertOne(ctx, msg[0])
	if insertErr != nil {
		client_u.Disconnect(ctx)
		// client_u2.Disconnect(ctx)
		return `nok`
	}
	fmt.Println(res_i)

	res_u, insertErr := collection_u.DeleteOne(ctx, input1)
	if insertErr != nil {
		client_u.Disconnect(ctx)
		// client_u2.Disconnect(ctx)
		return `nok`
	}
	fmt.Println(res_u)
	client_u.Disconnect(ctx)
	// client_u2.Disconnect(ctx)
	return `ok`

}

//-----------------------------------------

func UpdateonlySTATEDB(ctx context.Context, db_mongo_u string, collec_u string, input1 bson.M, input2 bson.M) string {

	time := time.Now().Add(time.Minute).Unix()
	clientOptions_u := options.Client().ApplyURI(server3)
	client_u, err := mongo.Connect(ctx, clientOptions_u)
	if err != nil {
		log.Fatal(err)
	}
	err = client_u.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client_u.Disconnect(ctx)

	collection_u := client_u.Database(db_mongo_u).Collection(collec_u)
	// collection_i := client_u.Database(db_mongo_u).Collection(collec_u + `_Archive`)

	//cur, currErr := collection.Find(ctx, bson.M{})
	cur, currErr := collection_u.Find(ctx, input1)
	if currErr != nil {
		panic(currErr)
	}
	// fmt.Println(cur)

	var msg []bson.M

	if err = cur.All(ctx, &msg); err != nil {
		panic(err)
	}
	// fmt.Println(msg)

	delete(msg[0], "_id")

	// res_i, insertErr := collection_i.InsertOne(ctx, msg[0])
	// if insertErr != nil {
	// 	return `nok`
	// }
	// fmt.Println(res_i)
	input2[`timestamp`] = time
	res_u, insertErr := collection_u.UpdateOne(ctx, input1, bson.M{"$set": input2})
	if insertErr != nil {
		client_u.Disconnect(ctx)
		return `nok`
	}
	fmt.Println(res_u)
	client_u.Disconnect(ctx)
	return `ok`
}

func InsertdbSTATEDB(ctx context.Context, db_mongo_u string, collec_u string, input1 bson.M) string {

	// db_mongo_u = "auth_main_demo2_07_21"
	// collec_u = "users_main_demo2_07_21"
	time := time.Now().Add(time.Minute).Unix()
	// var ctx = context.TODO()
	clientOptions_u := options.Client().ApplyURI(server3)
	client_u, err := mongo.Connect(ctx, clientOptions_u)
	if err != nil {
		log.Fatal(err)
	}
	err = client_u.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client_u.Disconnect(ctx)

	collection_u := client_u.Database(db_mongo_u).Collection(collec_u)
	input1[`timestamp`] = time
	res_u, insertErr := collection_u.InsertOne(ctx, input1)
	if insertErr != nil {
		client_u.Disconnect(ctx)
		return `nok`
	}
	fmt.Println(res_u)
	client_u.Disconnect(ctx)
	return `ok`
}

func FinddbSTATEDB(ctx context.Context, db_mongo_u string, collec_u string, input1 bson.M, sortby string, sortorder int, limmit int64, skip int64) []bson.M {

	// db_mongo_u = "auth_main_demo2_07_21"
	// collec_u = "users_main_demo2_07_21"
	// var ctx = context.TODO()
	clientOptions_u := options.Client().ApplyURI(server3)
	client_u, err := mongo.Connect(ctx, clientOptions_u)
	if err != nil {
		log.Fatal(err)
	}
	err = client_u.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client_u.Disconnect(ctx)

	collection_u := client_u.Database(db_mongo_u).Collection(collec_u)

	opts := options.Find()
	opts.SetSort(bson.D{{sortby, sortorder}}).SetLimit(limmit).SetSkip(skip)
	cur, err := collection_u.Find(ctx, input1, opts)
	//cur, err := collection.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		var msg2 []bson.M
		client_u.Disconnect(ctx)
		return msg2
	}
	var msg []bson.M
	if err = cur.All(ctx, &msg); err != nil {
		panic(err)
	}
	client_u.Disconnect(ctx)
	return msg
}

func FinddbSTSTATEDB(ctx context.Context, db_mongo_u string, collec_u string, input1 bson.M, sortby string, sortorder int, limmit int64, skip int64, Pro bson.M) []bson.M {

	// db_mongo_u = "auth_main_demo2_07_21"
	// collec_u = "users_main_demo2_07_21"
	// var ctx = context.TODO()
	clientOptions_u := options.Client().ApplyURI(server3)
	client_u, err := mongo.Connect(ctx, clientOptions_u)
	if err != nil {
		log.Fatal(err)
	}
	err = client_u.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client_u.Disconnect(ctx)

	collection_u := client_u.Database(db_mongo_u).Collection(collec_u)

	opts := options.Find()
	opts.SetSort(bson.D{{sortby, sortorder}}).SetLimit(limmit).SetSkip(skip).SetProjection(Pro)
	cur, err := collection_u.Find(ctx, input1, opts)
	//cur, err := collection.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		var msg2 []bson.M
		client_u.Disconnect(ctx)
		return msg2
	}
	var msg []bson.M
	if err = cur.All(ctx, &msg); err != nil {
		panic(err)
	}
	client_u.Disconnect(ctx)
	return msg
}
