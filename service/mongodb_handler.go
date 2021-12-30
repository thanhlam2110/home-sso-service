package service

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var CNX = Connection()

//connect MongoDB
func Connection() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://casuser:Mellon@192.168.0.106:27017/?authSource=cas")
	// Connect to MongoDB
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOptions)
	//defer client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

//check User assign
type ArrayUserAssignThing struct {
	UserAssign []string
}

func CheckUserAssignThing(thingid string) (err error, response []string) {
	collection := CNX.Database("users").Collection("things")
	//fmt.Println(collection)
	var arrayUserAssignThing ArrayUserAssignThing
	ctx := context.Background()
	err = collection.FindOne(ctx, bson.M{"thingid": thingid}).Decode(&arrayUserAssignThing)
	if err != nil {
		fmt.Println(err)
		return err, response
	}
	//fmt.Println(arrayUserAssignThing.UserAssign)
	return nil, arrayUserAssignThing.UserAssign
}

//check user role
type ArrayUserRole struct {
	Userole []string
}

func CheckUserRole(userid string) (err error, response []string) {
	collection := CNX.Database("users").Collection("user_role")
	//fmt.Println(collection)
	var arrayUserRole ArrayUserRole
	ctx := context.Background()
	err = collection.FindOne(ctx, bson.M{"userid": userid}).Decode(&arrayUserRole)
	if err != nil {
		fmt.Println(err)
		return err, response
	}
	//fmt.Println(arrayUserRole)
	return nil, arrayUserRole.Userole
}

//check User Status
type Userstatus struct {
	Userstatus string
}

func CheckUserStatus(username string) (err error, status string) {
	collection := CNX.Database("users").Collection("users")
	//fmt.Println(collection)
	var userstatus Userstatus
	ctx := context.Background()
	err = collection.FindOne(ctx, bson.M{"username": username}).Decode(&userstatus)
	if err != nil {
		fmt.Println(err)
		return err, ""
	}
	//fmt.Println(userstatus.Userstatus)
	return nil, userstatus.Userstatus
}
