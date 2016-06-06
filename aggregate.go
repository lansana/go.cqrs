// Copyright 2016 Jet Basrawi. All rights reserved.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package ycq

//AggregateRoot is the interface that all aggregates should implement
type AggregateRoot interface {
	AggregateID() string
	Version() int
	IncrementVersion()
	Apply(events EventMessage, isNew bool)
	TrackChange(EventMessage)
	GetChanges() []EventMessage
	ClearChanges()
}

//AggregateBase is a type that can be embedded in an AggregateRoot
//implementation to handle common aggragate behaviour
//
//All required methods to implement an aggregate are here, to implement the
//Aggregate root interface your aggregate will need to implement the Apply
//method that will contain behaviour specific to your aggregate.
type AggregateBase struct {
	id      string
	version int
	changes []EventMessage
}

//NewAggregateBase contructs a new AggregateBase.
func NewAggregateBase(id string) *AggregateBase {
	return &AggregateBase{
		id:      id,
		changes: []EventMessage{},
	}
}

//AggregateID returns the AggregateID
func (a *AggregateBase) AggregateID() string {
	return a.id
}

//Version returns the version of the aggregate.
func (a *AggregateBase) Version() int {
	return a.version
}

//IncrementVersion increments the aggregate version number
func (a *AggregateBase) IncrementVersion() {
	a.version++
}

//TrackChange stores the EventMessage in the changes collection.
//
//Changes are new, unpersisted events that have been applied to the aggregate.
func (a *AggregateBase) TrackChange(event EventMessage) {
	a.changes = append(a.changes, event)
}

//GetChanges returns the collection of new unpersisted events that have
//been applied to the aggregate.
func (a *AggregateBase) GetChanges() []EventMessage {
	return a.changes
}

//ClearChanges removes all unpersisted events from the aggregate.
func (a *AggregateBase) ClearChanges() {
	a.changes = []EventMessage{}
}
