package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// Use struct tags much like the standard JSON library,
// you can embed anonymous structs too!
type widget struct {
	UserID int       // Hash key, a.k.a. partition key
	Time   time.Time // Range key, a.k.a. sort key

	Msg       string              `dynamo:"Message"`    // Change name in the database
	Count     int                 `dynamo:",omitempty"` // Omits if zero value
	Children  []widget            // List of maps
	Friends   []string            `dynamo:",set"` // Sets
	Set       map[string]struct{} `dynamo:",set"` // Map sets, too!
	SecretKey string              `dynamo:"-"`    // Ignored
}


func main() {
	fmt.Println("start")
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{
		Region: aws.String("us-west-1"),
		// dockerで動かす場合はhost.docker.internal
		Endpoint: aws.String("http://localhost:4566"),
		// DisableSSL: aws.Bool(true),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	table := db.Table("Widgets")
	fmt.Printf("table: %v\n", table)

	// // put item
	// fmt.Printf("put item\n")
	w := widget{UserID: 613, Time: time.Now(), Msg: "hello"}
	// err := table.Put(w).Run()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("put item: %v\n", w)

	// get the same item
	fmt.Printf("get item\n")
	var result widget
	err := table.Get("UserID", w.UserID).
		Range("Time", dynamo.Equal, w.Time).
		One(&result)
	if err != nil {
		panic(err)
	}

	// get all items
	var results []widget
	err = table.Scan().All(&results)

	// use placeholders in filter expressions (see Expressions section below)
	var filtered []widget
	err = table.Scan().Filter("'Count' > ?", 10).All(&filtered)
}