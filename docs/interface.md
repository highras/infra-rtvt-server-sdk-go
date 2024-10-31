# CreateRTVTClient

```go
func CreateRTVTClient(endpoints string, callbacks RTVTCallback, logger RTVTLogger) *RTVTClient
```

| **param** | **description**                                     |
| --------------- | --------------------------------------------------------- |
| endpoints       | the address and port to rtvt eg. rtvt.ilivedata.com:14001 |
| callbacks       | callback interface                                        |
| logger          | internal logger interface                                 |

# Login

```go
func (client *RTVTClient) Login(pid int32, timestamp int64, token string) bool
```

| **param** | **description**                |
| --------------- | ------------------------------------ |
| pid             | the pid of your project eg. 81700001 |
| timestamp       | timestamp                            |
| token           | auth token                           |

# StartTranslate

```go
func (client *RTVTClient) StartTranslate(asrResult bool, tempResult bool, transResult bool, srcLanguage string, destLanguage string, userId string, vadSlienceTime int64, codec AudioCodec) (int64, error)
```

| **param** | **description**                                                                                                                                                    |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| asrResult       | if true, you can get the recognize result.                                                                                                                               |
| tempResult      | if both asrResult and temp Result are true. you can get the recognize temp result. If  both tempResult and transResult are true,  you can get the translate temp result. |
| transResult     | if true, you can get the translate result.                                                                                                                               |
| srcLanguage     | source language                                                                                                                                                          |
| destLanguage    | destination language                                                                                                                                                     |
| userId          | user id, can be empty                                                                                                                                                    |
| vadSlienceTime  | vad time. -1 means not use vad cut. 200-2000 means use the vad cut sentence. (miliseconds)                                                                              |
| codec           | the codec of audio data                                                                                                                                                  |

# SendData

```go
func (client *RTVTClient) SendData(streamId int64, data []byte, seq int64, timestamp int64) error
```

| **param** | **description**                    |
| --------------- | ---------------------------------------- |
| streamId        | the streamId you get from StartTranslate |
| data            | PCM data                                 |
| seq             | sequence of the package                  |
| timestamp       | timestamp of the package                 |

# EndTranslate

```go
func (client *RTVTClient) EndTranslate(streamId int64) error
```

| **param** | **description**                    |
| --------------- | ---------------------------------------- |
| streamId        | the streamId you get from StartTranslate |
