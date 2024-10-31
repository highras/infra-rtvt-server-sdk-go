package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	rtvt "github.com/highras/infra-rtvt-server-sdk-go"
)

var pid int32 = 0
var secret string = "your secret"

type Callbacks struct {
}

func (callback Callbacks) PushRecognizedResult(streamId int64, startTs int64, endTs int64, taskId int64, result string) {
	fmt.Println("asr result:", result)
}

func (callback Callbacks) PushRecognizedTempResult(streamId int64, startTs int64, endTs int64, taskId int64, result string) {
	fmt.Println("asr temp result:", result)
}

func (callback Callbacks) PushTranslatedResult(streamId int64, startTs int64, endTs int64, taskId int64, result string) {
	fmt.Println("trans result:", result)
}

func (callback Callbacks) PushTranslatedTempResult(streamId int64, startTs int64, endTs int64, taskId int64, result string) {
	fmt.Println("trans temp result:", result)
}

type MyLogger struct {
}

func (logger MyLogger) Println(a ...any) {
	fmt.Println(a)
}

func (logger MyLogger) Printf(format string, a ...any) {
	fmt.Printf(format, a)
}

func genHMACToken(pid int32, ts int64, key string) string {
	content := strconv.FormatInt(int64(pid), 10) + ":" + strconv.FormatInt(int64(ts), 10)
	keyb, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		fmt.Printf("error:%s", err)
		return ""
	}
	hmacsha256 := hmac.New(func() hash.Hash {
		return sha256.New()
	}, keyb)
	hmacsha256.Write([]byte(content))
	res := make([]byte, 0)
	res = hmacsha256.Sum(res)
	return base64.StdEncoding.EncodeToString(res)
}

func main() {
	filePath := os.Args[1]
	callbacks := &Callbacks{}
	// logger := &MyLogger{}
	logger := log.New(os.Stdout, "RTVT SDK", log.LstdFlags)
	rtvtClient := rtvt.CreateRTVTClient("rtvt-bj.ilivedata.com:14001", callbacks, logger)
	ts := time.Now().Unix()
	token := genHMACToken(pid, ts, secret)
	if len(token) == 0 {
		return
	}

	succ := rtvtClient.Login(pid, ts, token)
	if !succ {
		return
	}
	streamId, _ := rtvtClient.StartTranslate(true, true, true, "zh", "en", "go test user", 2000, rtvt.PCM)
	ticker := time.NewTicker(20 * time.Millisecond)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(t *time.Ticker) {
		defer wg.Done()
		seq := int64(0)
		file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
		buffer := make([]byte, 640)
		file.Seek(44, 0)

		if err != nil {
			return
		}
		for {
			<-t.C
			ts := time.Now().UnixMilli()
			for i := range buffer {
				buffer[i] = 0
			}
			n, err := file.Read(buffer)
			if err != nil || n != 640 {
				rtvtClient.SendData(streamId, buffer, seq, ts)
				break
			} else {
				rtvtClient.SendData(streamId, buffer, seq, ts)
				seq++
			}
		}
	}(ticker)
	wg.Wait()
	rtvtClient.EndTranslate(streamId)
}
