package rtvt

type IRTVTCallback interface {
	PushRecognizedResult(streamId int64, startTs int64, endTs int64, taskId int64, result string)
	PushRecognizedTempResult(streamId int64, startTs int64, endTs int64, taskId int64, result string)
	PushTranslatedResult(streamId int64, startTs int64, endTs int64, taskId int64, result string)
	PushTranslatedTempResult(streamId int64, startTs int64, endTs int64, taskId int64, result string)
}
