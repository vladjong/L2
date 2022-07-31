package event

import (
	"errors"
	"fmt"
	"math"
	"time"
)

type EventCache struct {
	events *[]Event
}

func NewEventCache() *EventCache {
	return &EventCache{
		events: new([]Event),
	}
}

func (cache *EventCache) Create(event Event) {
	*cache.events = append(*cache.events, event)
}

func (cache *EventCache) Update(event Event) error {
	for i, ev := range *cache.events {
		if ev.ID == event.ID {
			fmt.Println(ev)
			(*cache.events)[i] = event
			fmt.Println(ev)
			return nil
		}
	}
	return errors.New("the event does not exist")
}

func (cache *EventCache) Delete(event Event) error {
	for _, ev := range *cache.events {
		if ev.ID == event.ID {
			ev = (*cache.events)[len(*cache.events)-1]
			*cache.events = (*cache.events)[:len(*cache.events)-1]
			return nil
		}
	}
	return errors.New("the event does not exist")
}

func (cache *EventCache) GetEventByDay(date int) ([]Event, error) {
	var events []Event
	for _, val := range *cache.events {
		_, _, day := val.Date.Date()
		if day == date {
			events = append(events, val)
		}
	}
	if len(events) == 0 {
		return nil, errors.New("the event does not exist")
	}
	return events, nil
}

func (cache *EventCache) GetEventByMonth(date time.Month) ([]Event, error) {
	var events []Event
	for _, val := range *cache.events {
		_, month, _ := val.Date.Date()
		if month == date {
			events = append(events, val)
		}
	}
	if len(events) == 0 {
		return nil, errors.New("the event does not exist")
	}
	return events, nil
}

func (cache *EventCache) GetEventByWeek(date time.Time) ([]Event, error) {
	var events []Event
	for _, val := range *cache.events {
		year, month, day := val.Date.Date()
		if year == date.Year() && month == date.Month() {
			if math.Abs(float64(date.Day()-day)) >= 0 && math.Abs(float64(date.Day()-day)) < 7 {
				events = append(events, val)
			}
		}
	}
	if len(events) == 0 {
		return nil, errors.New("the event does not exist")
	}
	return events, nil
}
