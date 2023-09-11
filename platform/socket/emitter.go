package socket

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/redis/go-redis/v9"
	"github.com/vmihailenco/msgpack"
)

const (
	event = 2
)

// Emitter ... Socket.io Emitter
type Emitter struct {
	_key   string
	_flags map[string]string
	_rooms map[string]bool
	_pool  *redis.Client
}

// New ... Creates new Emitter using options
func New() *Emitter {
	emitter := Emitter{}

	initRedisConnPool(&emitter)

	//if opts.Key != "" {
	//	emitter._key = opts.Key
	//} else {
	// emitter._key = "socket.io#/#123#"
	//}
	fmt.Print(emitter._key)
	emitter._rooms = make(map[string]bool)
	emitter._flags = make(map[string]string)

	return &emitter
}

// In ... Limit emission to a certain `room`.`
func (emitter *Emitter) In(room string) *Emitter {
	if _, ok := emitter._rooms[room]; !ok {
		emitter._rooms[room] = true
	}
	emitter._key = fmt.Sprintf("socket.io#/#%s#", room)
	return emitter
}

// To ... Limit emission to a certain `room`.
func (emitter *Emitter) To(room string) *Emitter {
	return emitter.In(room)
}

// Of ... To Limit emission to certain `namespace`.
func (emitter *Emitter) Of(nsp string) *Emitter {
	emitter._flags["nsp"] = nsp
	return emitter
}

// Emit ... Send the packet.
func (emitter *Emitter) Emit(args ...interface{}) bool {
	packet := make(map[string]interface{})
	extras := make(map[string]interface{})

	//if ok := emitter.hasBin(args); ok {
	//	packet["type"] = binaryEvent
	//} else {
	//	packet["type"] = event
	//}
	packet["type"] = event
	packet["data"] = args

	if value, ok := emitter._flags["nsp"]; ok {
		packet["nsp"] = value
		delete(emitter._flags, "nsp")
	} else {
		packet["nsp"] = "/"
	}

	if ok := len(emitter._rooms); ok > 0 {
		extras["rooms"] = getKeys(emitter._rooms)
	} else {
		extras["rooms"] = make([]string, 0)
	}

	if ok := len(emitter._flags); ok > 0 {
		extras["flags"] = ""
	} else {
		extras["flags"] = "" //make(map[string]string)
	}

	//Pack & Publish
	b, err := msgpack.Marshal([]interface{}{"socket_id", packet, extras})
	if err != nil {
		panic(err)
	} else {
		publish(emitter, emitter._key, b)
	}

	emitter._rooms = make(map[string]bool)
	emitter._flags = make(map[string]string)

	return true
}

func initRedisConnPool(emitter *Emitter) {
	emitter._pool = newPool()
}

func newPool() *redis.Client {
	_url := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	_option := &redis.Options{
		Addr:     _url,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
	return redis.NewClient(_option)

}

func publish(emitter *Emitter, channel string, value interface{}) {
	if err := emitter._pool.Publish(context.TODO(), channel, value).Err(); err != nil {
		log.Fatal(err.Error())
	}
}

func getKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	return keys
}
