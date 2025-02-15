package kafka

import (
	"context"
	kafkago "github.com/segmentio/kafka-go"
	"github.com/tx7do/kratos-transport/broker"
)

type publication struct {
	topic  string
	err    error
	m      *broker.Message
	ctx    context.Context
	reader *kafkago.Reader
	km     kafkago.Message
}

func (p *publication) Topic() string {
	return p.topic
}

func (p *publication) Message() *broker.Message {
	p.m.Body = p.km
	return p.m
}

func (p *publication) Ack() error {
	return p.reader.CommitMessages(p.ctx, p.km)
}

func (p *publication) Error() error {
	return p.err
}

func (p *publication) KafKaMessage() kafkago.Message {
	return p.km
}
