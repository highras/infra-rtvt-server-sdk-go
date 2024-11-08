package rtvt

import "errors"

var ErrFooSendQuestFailed = errors.New("send quest failed")
var ErrFooEndStreamFailed = errors.New("end stream failed")
var ErrFooSendVoiceDataFailed = errors.New("send voice data failed")
var ErrFooStartStreamFailed = errors.New("start stream failed")
var ErrFooUnsupportedLanguage = errors.New("start stream failed, unsupported language")
var ErrFooStreamTooMany = errors.New("start stream failed, start too many stream")
var ErrFooStreamNotExist = errors.New("stream id not exist")
var ErrFooInvalidToken = errors.New("login failed invalid token")
var ErrFooInvalidTimestamp = errors.New("invalid timestamp")
var ErrFooInvalidProjectID = errors.New("invalid pid")
var ErrFooInvalidVadSlienceTime = errors.New("invalid vad slience time")
