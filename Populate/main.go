package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/Ibukun-tech/try"
)

func main() {
	fmt.Println("---Already About to start sending to the database---")
	level := []string{"EASY", "MEDIUM", "HARD", "DIFFICULT"}
	message := []string{"This is easy", "This is medium", "This is hard", "This is difficult"}

	// send := try.Log{
	// 	Level:      level[rand.Intn(4)],
	// 	Message:    message[rand.Intn(4)],
	// 	ResourceId: randomString(4),
	// 	Timestamp:  time.Now(),
	// 	TraceId:    "trace-" + randomString(4),
	// 	SpanId:     "span-" + randomString(4),
	// 	Commit:     randomString(6),
	// 	Metadata: try.Metadata{
	// 		ParentResourceId: randomString(4),
	// 	}}
	// byteSend, err := json.Marshal(send)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// res, err := http.Post("http://localhost:4000/add", "Application/json", bytes.NewBuffer(byteSend))
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println("err: ", err)
	// var sy string
	// bit, _ := io.ReadAll(res.Body)
	// err = json.Unmarshal(bit, &sy)
	// if err != nil {
	// 	fmt.Println(err, "aee")
	// }
	// fmt.Println(sy)
	for i := 0; i < 100; i++ {
		for i := 0; i < 1000; i++ {

			send := try.Log{
				Level:      level[rand.Intn(4)],
				Message:    message[rand.Intn(4)],
				ResourceId: randomString(4),
				Timestamp:  time.Now(),
				TraceId:    "trace-" + randomString(4),
				SpanId:     "span-" + randomString(4),
				Commit:     randomString(6),
				Metadata: try.Metadata{
					ParentResourceId: randomString(4),
				}}
			byteSend, err := json.Marshal(send)
			if err != nil {
				fmt.Println(err)
				continue
			}
			res, err := http.Post("http://localhost:4000/add", "Application/json", bytes.NewBuffer(byteSend))
			if err != nil {
				log.Println(err)
			}
			log.Println("err: ", err)
			log.Println(res.StatusCode)
		}
	}
}

func randomString(n int) string {
	numbers := "1234567"
	b := make([]byte, n)
	for i := range b {
		b[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(b)
	// numbersMap := map[int]string{
	// 	0: "1",
	// 	1: "2",
	// 	2: "3",
	// 	3: "4",
	// 	4: "5",
	// 	5: "6",
	// }
	// b := make([]string, n)
	// for i := range b {
	// 	v, _ := numbersMap[rand.Intn(len(numbersMap))]
	// 	b[i] = v
	// }
	// return
}
