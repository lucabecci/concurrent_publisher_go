package entities

import (
	"math/rand"
	"time"
)

type Operation struct {
	Id          string
	Status      Status
	Value       int
	CreatedAt   time.Time
	StartedAt   time.Time
	CompletedAt time.Time
}

func NewOperation(id string) *Operation {
	return &Operation{
		Id:        id,
		Status:    Waiting,
		CreatedAt: time.Now(),
	}
}

func (o *Operation) GenerateValue() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 1000
	randomNumber := rand.Intn(max-min+1) + min
	o.Value = randomNumber
}

func (o *Operation) GetStatus() Status {
	return o.Status
}

func (o *Operation) Start() {
	o.Status = Started
	o.StartedAt = time.Now()
}

func (o *Operation) Complete() {
	o.Status = Completed
	o.CompletedAt = time.Now()
}

func (o *Operation) Block() {
	o.Status = Blocked
}
