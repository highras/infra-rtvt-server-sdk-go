package rtvt

import (
	"errors"

	"github.com/highras/fpnn-sdk-go/src/fpnn"
)

type rtvtServerQuestProcessor struct {
	logger    RTVTLogger
	callbacks IRTVTCallback
}

func (processor *rtvtServerQuestProcessor) Process(method string) func(*fpnn.Quest) (*fpnn.Answer, error) {
	switch method {
	case "ping":
		return processor.processPing
	case "recognizedResult":
		return processor.processRecognizedResult
	case "recognizedTempResult":
		return processor.processRecognizedTempResult
	case "translatedResult":
		return processor.processTranslatedResult
	case "translatedTempResult":
		return processor.processTranslatedTempResult
	default:
		return nil
	}
}

func (processor *rtvtServerQuestProcessor) processPing(quest *fpnn.Quest) (*fpnn.Answer, error) {
	return fpnn.NewAnswer(quest), nil
}

func (processor *rtvtServerQuestProcessor) processRecognizedResult(quest *fpnn.Quest) (*fpnn.Answer, error) {
	streamId, ok := quest.GetInt64("streamId")
	if !ok {
		processor.logger.Println("invalid recognized result.")
		return fpnn.NewAnswer(quest), errors.New("invalid recognized result")
	}
	startTs, ok := quest.GetInt64("startTs")
	if !ok {
		processor.logger.Println("invalid recognized result.")
		return fpnn.NewAnswer(quest), errors.New("invalid recognized result")
	}
	endTs, ok := quest.GetInt64("endTs")
	if !ok {
		processor.logger.Println("invalid recognized result.")
		return fpnn.NewAnswer(quest), errors.New("invalid recognized result")
	}
	result, ok := quest.GetString("asr")
	if !ok {
		processor.logger.Println("invalid recognized result.")
		return fpnn.NewAnswer(quest), errors.New("invalid recognized result")
	}
	taskId, ok := quest.GetInt64("taskId")
	if !ok {
		processor.logger.Println("invalid recognized result.")
		return fpnn.NewAnswer(quest), errors.New("invalid recognized result")
	}
	processor.callbacks.PushRecognizedResult(streamId, startTs, endTs, taskId, result)
	return fpnn.NewAnswer(quest), nil
}

func (processor *rtvtServerQuestProcessor) processRecognizedTempResult(quest *fpnn.Quest) (*fpnn.Answer, error) {
	streamId, ok := quest.GetInt64("streamId")
	if !ok {
		processor.logger.Println("invalid temp recognized result.")
		return fpnn.NewAnswer(quest), errors.New("invalid temp recognized result")
	}
	startTs, ok := quest.GetInt64("startTs")
	if !ok {
		processor.logger.Println("invalid temp recognized result.")
		return fpnn.NewAnswer(quest), errors.New("invalid temp recognized result")
	}
	endTs, ok := quest.GetInt64("endTs")
	if !ok {
		processor.logger.Println("invalid temp recognized result.")
		return fpnn.NewAnswer(quest), errors.New("invalid temp recognized result")
	}
	result, ok := quest.GetString("asr")
	if !ok {
		processor.logger.Println("invalid temp recognized result.")
		return fpnn.NewAnswer(quest), errors.New("invalid temp recognized result")
	}
	taskId, ok := quest.GetInt64("taskId")
	if !ok {
		processor.logger.Println("invalid recognized result.")
		return fpnn.NewAnswer(quest), errors.New("invalid recognized result")
	}
	processor.callbacks.PushRecognizedTempResult(streamId, startTs, endTs, taskId, result)
	return fpnn.NewAnswer(quest), nil
}

func (processor *rtvtServerQuestProcessor) processTranslatedResult(quest *fpnn.Quest) (*fpnn.Answer, error) {
	streamId, ok := quest.GetInt64("streamId")
	if !ok {
		processor.logger.Println("invalid translated result.")
		return fpnn.NewAnswer(quest), errors.New("invalid translated result")
	}
	startTs, ok := quest.GetInt64("startTs")
	if !ok {
		processor.logger.Println("invalid translated result.")
		return fpnn.NewAnswer(quest), errors.New("invalid translated result")
	}
	endTs, ok := quest.GetInt64("endTs")
	if !ok {
		processor.logger.Println("invalid translated result.")
		return fpnn.NewAnswer(quest), errors.New("invalid translated result")
	}
	result, ok := quest.GetString("trans")
	if !ok {
		processor.logger.Println("invalid translated result.")
		return fpnn.NewAnswer(quest), errors.New("invalid translated result")
	}
	taskId, ok := quest.GetInt64("taskId")
	if !ok {
		processor.logger.Println("invalid recognized result.")
		return fpnn.NewAnswer(quest), errors.New("invalid recognized result")
	}
	processor.callbacks.PushTranslatedResult(streamId, startTs, endTs, taskId, result)
	return fpnn.NewAnswer(quest), nil
}

func (processor *rtvtServerQuestProcessor) processTranslatedTempResult(quest *fpnn.Quest) (*fpnn.Answer, error) {
	streamId, ok := quest.GetInt64("streamId")
	if !ok {
		processor.logger.Println("invalid temp translated result.")
		return fpnn.NewAnswer(quest), errors.New("invalid temp translated result")
	}
	startTs, ok := quest.GetInt64("startTs")
	if !ok {
		processor.logger.Println("invalid temp translated result.")
		return fpnn.NewAnswer(quest), errors.New("invalid temp translated result")
	}
	endTs, ok := quest.GetInt64("endTs")
	if !ok {
		processor.logger.Println("invalid temp translated result.")
		return fpnn.NewAnswer(quest), errors.New("invalid temp translated result")
	}
	result, ok := quest.GetString("trans")
	if !ok {
		processor.logger.Println("invalid temp translated result.")
		return fpnn.NewAnswer(quest), errors.New("invalid temp translated result")
	}
	taskId, ok := quest.GetInt64("taskId")
	if !ok {
		processor.logger.Println("invalid recognized result.")
		return fpnn.NewAnswer(quest), errors.New("invalid recognized result")
	}
	processor.callbacks.PushTranslatedTempResult(streamId, startTs, endTs, taskId, result)
	return fpnn.NewAnswer(quest), nil
}
