package try

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	level   = []string{"EASY", "MEDIUM", "HARD", "DIFFICULT"}
	message = []string{"This is easy", "This is medium", "This is hard", "This is difficult"}
)

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

func randomString(n int) string {
	num := "123456"
	a := make([]byte, n)
	for i := range a {
		a[i] = num[rand.Intn(len(num))]
	}
	return string(a)
}
func TestNewMongoClient(t *testing.T) {
	as := assert.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer func() {
		cancel()
	}()
	_, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://ibk:secret@localhost:27017"))
	if err != nil {
		as.Nil(err)
	}
}

func TestConnectionAdd(t *testing.T) {
	as := assert.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer func() {
		cancel()
	}()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://ibk:secret@localhost:27017"))
	if err != nil {
		as.Nil(err)
	}
	cl := NewMongoClient(client)
	send := Log{
		Level:      level[rand.Intn(4)],
		Message:    message[rand.Intn(4)],
		ResourceId: randomString(4),
		Timestamp:  time.Now(),
		TraceId:    "trace-" + randomString(4),
		SpanId:     "span-" + randomString(4),
		Commit:     randomString(6),
		Metadata: Metadata{
			ParentResourceId: randomString(4),
		}}

	value, err := cl.Add(send)
	if err != nil {
		as.Nil(err)
	}
	as.NotEqual("", value)
	as.Equal("It has been inserted already", value)
}
