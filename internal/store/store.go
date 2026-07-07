package store

import (
	"errors"

	xapi "github.com/jasony87/simple-exercise-app/generated"
)

var ErrLogEntryNotFound = errors.New("log entry not found")
var ErrLogEntryIdRequired = errors.New("log entry id is required")

type MemoryStore struct {
	xlogs map[string][]xapi.LogItem
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		xlogs: make(map[string][]xapi.LogItem),
	}
}

func (s *MemoryStore) Add(date string, logItem xapi.LogItem) error {
	s.xlogs[date] = append(s.xlogs[date], logItem)
	return nil
}

func (s *MemoryStore) Update(date string, logItem xapi.LogItem) error {
	// ensure id is set
	if logItem.Id == nil || *logItem.Id == "" {
		return ErrLogEntryIdRequired
	}
	// check entry ID exists and update
	for i, item := range s.xlogs[date] {
		if item.Id != nil && *item.Id == *logItem.Id {
			s.xlogs[date][i] = logItem
			return nil
		}
	}
	return ErrLogEntryNotFound
}

func (s *MemoryStore) GetByDate(date string) ([]xapi.LogItem, error) {
	xlogsEntry, ok := s.xlogs[date]
	if !ok {
		return nil, ErrLogEntryNotFound
	}
	return xlogsEntry, nil
}
