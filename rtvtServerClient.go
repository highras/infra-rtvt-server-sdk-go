package rtvt

import (
	"sync"

	"github.com/highras/fpnn-sdk-go/src/fpnn"
)

type RTVTServerClient struct {
	client                    *fpnn.TCPClient
	processor                 *rtvtServerQuestProcessor
	logger                    rtvtLogger
	pid                       int32
	secretKey                 string
	regressiveState           *RtmRegressiveState
	isClose                   bool
	defaultRegressiveStrategy *RTMRegressiveConnectStrategy
	regressiveConnectStrategy *RTMRegressiveConnectStrategy
	mutex                     sync.Mutex
}
