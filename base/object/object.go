package object

import (
	uuid "github.com/satori/go.uuid"
)

type Cobject struct {
	name string
	id   uuid.UUID
}

func (o *Cobject) Init(name string) {
	o.name = name
	o.id = uuid.NewV4()
}
