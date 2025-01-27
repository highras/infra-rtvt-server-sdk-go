package rtvt

import (
	"sync"

	"github.com/highras/fpnn-sdk-go/src/fpnn"
)

const VERSION = "0.1.2"

type RTVTClient struct {
	client         *fpnn.TCPClient
	processor      *rtvtServerQuestProcessor
	logger         RTVTLogger
	pid            int32
	loginTimestamp int64
	loginToken     string
	isLogin        bool
	isClose        bool
	mutex          sync.Mutex
}

type AudioCodec int

const (
	PCM AudioCodec = iota
	OPUS
)

func CreateRTVTClient(endpoints string, callbacks RTVTCallback, logger RTVTLogger) *RTVTClient {
	rtvtClient := &RTVTClient{}
	rtvtClient.client = fpnn.NewTCPClient(endpoints)
	rtvtClient.logger = logger
	rtvtClient.processor = &rtvtServerQuestProcessor{}
	rtvtClient.processor.callbacks = callbacks
	rtvtClient.processor.logger = logger
	rtvtClient.client.SetLogger(rtvtClient.logger)
	rtvtClient.client.SetKeepAlive(true)
	rtvtClient.client.SetQuestProcessor(rtvtClient.processor)
	rtvtClient.client.SetOnClosedCallback(func(connId uint64, endpoint string) {
		rtvtClient.isClose = true
		rtvtClient.isLogin = false
		rtvtClient.pid = 0
		rtvtClient.loginTimestamp = 0
		rtvtClient.loginToken = ""
	})
	rtvtClient.client.SetOnConnectedCallback(func(connId uint64, endpoint string, connected bool) {
		rtvtClient.isClose = connected
	})
	return rtvtClient
}

func (client *RTVTClient) Login(pid int32, timestamp int64, token string) (bool, error) {
	client.pid = pid
	client.loginTimestamp = timestamp
	client.loginToken = token
	quest := fpnn.NewQuest("login")
	quest.Param("pid", pid)
	quest.Param("token", token)
	quest.Param("ts", timestamp)
	quest.Param("version", "rtvt-go-client-"+VERSION)
	answer, err := client.client.SendQuest(quest)
	if err != nil {
		client.logger.Println("login failed err:", err)
		return false, err
	}
	successed, ok := answer.GetBool("successed")
	if !ok || !successed {
		client.logger.Println("login failed err: invalid token")
		return false, ErrFooInvalidToken
	}
	client.isLogin = true
	return true, nil
}

func (client *RTVTClient) voiceStart(asrResult bool, tempResult bool, transResult bool, ttsResult bool, srcLanguage string, srcAltLanguage []string, destLanguage string, userId string, ttsSpeaker string, vadSlienceTime int64, codec AudioCodec) (int64, error) {
	quest := fpnn.NewQuest("voiceStart")
	quest.Param("asrResult", asrResult)
	quest.Param("asrTempResult", tempResult)
	quest.Param("transResult", transResult)
	quest.Param("ttsResult", ttsResult)
	quest.Param("srcLanguage", srcLanguage)
	quest.Param("srcAltLanguage", srcAltLanguage)
	quest.Param("destLanguage", destLanguage)
	quest.Param("userId", userId)
	quest.Param("ttsSpeaker", ttsSpeaker)
	quest.Param("vadSlienceTime", vadSlienceTime)
	quest.Param("codec", int(codec))
	answer, err := client.client.SendQuest(quest)
	if err != nil {
		return 0, err
	}
	streamId, ok := answer.GetInt64("streamId")
	if ok {
		return streamId, nil
	} else {
		code, _ := answer.GetInt64("code")
		client.logger.Println("voice start failed, code:", code)
		if code == 800105 {
			return 0, ErrFooUnsupportedLanguage
		} else if code == 800107 {
			return 0, ErrFooStreamTooMany
		} else {
			return 0, ErrFooStartStreamFailed
		}
	}
}

func (client *RTVTClient) voiceData(streamId int64, data []byte, seq int64, timestamp int64) error {
	quest := fpnn.NewQuest("voiceData")
	quest.Param("streamId", streamId)
	quest.Param("seq", seq)
	quest.Param("data", data)
	quest.Param("ts", timestamp)
	answer, err := client.client.SendQuest(quest)
	if err != nil {
		return err
	}
	code, _ := answer.GetInt64("code")
	if code != fpnn.FPNN_EC_OK {
		client.logger.Println("send voice data failed code:", code)
		if code == 800200 {
			return ErrFooStreamNotExist
		} else {
			return ErrFooSendVoiceDataFailed
		}
	}
	return nil
}

func (client *RTVTClient) voiceEnd(streamId int64) error {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	quest := fpnn.NewQuest("voiceEnd")
	quest.Param("streamId", streamId)
	answer, err := client.client.SendQuest(quest)
	if err != nil {
		return err
	}
	code, _ := answer.GetInt64("code")
	if code != fpnn.FPNN_EC_OK {
		client.logger.Println("voice end failed code:", code)
		if code == 800200 {
			return ErrFooStreamNotExist
		} else {
			return ErrFooEndStreamFailed
		}
	}
	return nil
}

func (client *RTVTClient) StartTranslate(asrResult bool, tempResult bool, transResult bool, srcLanguage string, srcAltLanguage []string, destLanguage string, userId string, vadSlienceTime int64, codec AudioCodec) (int64, error) {
	if vadSlienceTime > 2000 || (vadSlienceTime < 20 && vadSlienceTime != -1) {
		return 0, ErrFooInvalidVadSlienceTime
	}
	return client.voiceStart(asrResult, tempResult, transResult, false, srcLanguage, srcAltLanguage, destLanguage, userId, "", vadSlienceTime, codec)
}

func (client *RTVTClient) SendData(streamId int64, data []byte, seq int64, timestamp int64) error {
	return client.voiceData(streamId, data, seq, timestamp)
}

func (client *RTVTClient) EndTranslate(streamId int64) error {
	return client.voiceEnd(streamId)
}

// func (client *RTVTClient) StartTranslateWithTTS(asrResult bool, tempResult bool, transResult bool, srcLanguage string, destLanguage string, userId string, ttsSpeaker string) (int64, error) {
// 	return client.voiceStart(asrResult, tempResult, transResult, true, srcLanguage, destLanguage, userId, ttsSpeaker)
// }
