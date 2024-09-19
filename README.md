# RTVT SDK GO

This is RTVT SDK in golang.

## Usage

Create a rtvt client with method CreateRTVTClient. You will also need two stuct implement the RTVTCallback and RTVTLogger interface.

```go
type Callbacks struct {
}

func (callback Callbacks) PushRecognizedResult(streamId int64, startTs int64, endTs int64, taskId int64, result string) {
  // your code here
}

func (callback Callbacks) PushRecognizedTempResult(streamId int64, startTs int64, endTs int64, taskId int64, result string) {
  // your code here
}

func (callback Callbacks) PushTranslatedResult(streamId int64, startTs int64, endTs int64, taskId int64, result string) {
  // your code here
}

func (callback Callbacks) PushTranslatedTempResult(streamId int64, startTs int64, endTs int64, taskId int64, result string) {
  // your code here
}

callbacks := &Callbacks{}
```

```go
type MyLogger struct {
}

func (logger MyLogger) Println(a ...any) {
	fmt.Println(a)
}

func (logger MyLogger) Printf(format string, a ...any) {
	fmt.Printf(format, a)
}

logger := &MyLogger{}
```

Or you can use the log package.

```go
logger := log.New(os.Stdout, "RTVT SDK", log.LstdFlags)
```

```go
rtvtClient := rtvt.CreateRTVTClient("rtvt.ilivedata.com:14001", callbacks, logger)
```

Then login with method Login and your pid. You need generate a token using the key you get from your project. Token generation see:[https://docs.ilivedata.com/rtm/features/authentication/](https://docs.ilivedata.com/rtm/features/authentication/), you need to remove uid to generate the token.

```go
ts := time.Now().Unix()
yourpid := 0
token := genHMACToken(yourpid, ts, "your key")
if len(token) == 0 {
	return
}

succ := rtvtClient.Login(yourpid, ts, token)
if !succ {
	return
}
```

After login, you can start realtime translate and recognize with StartTranslate. You will get a streamId from StartTranslate. The streamId is a unique identifier to one stream.

```go
streamId, _ := rtvtClient.StartTranslate(true, true, true, "zh", "en", "user id")
```

Use SendData to send PCM data in realtime, the PCM data's samplerate must be 16000Hz, sampledepth must be 16bit and length of the data must be 20ms. That means one frame of the audio must be 640 bytes, if you using SendData with other length of data, it will failed, and return an error.

```go
seq := int64(0)
for {
      // read PCM data to buffer.
      ts := time.Now().UnixMilli()
      rtvtClient.SendData(streamId, buffer, seq, ts)
      seq++
}
```

When you finish your translate, you can use EndTranslate to close the stream.

```go
rtvtClient.EndTranslate(streamId)
```

The examples is in the examples directory.

## Reference

[interface.md](docs/interface.md)
