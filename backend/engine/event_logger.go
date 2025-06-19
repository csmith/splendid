package engine

import (
	"encoding/json"
	"fmt"
	"github.com/csmith/splendid/backend/model"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type EventLogger[G model.Game] interface {
	LogEvent(event model.Event[G]) error
}

type JSONLEventLogger[G model.Game] struct {
	dir       string
	sessionID string
	file      *os.File
	mutex     sync.Mutex
	timer     *time.Timer
	timeout   time.Duration
}

func NewJSONLEventLogger[G model.Game](dir, sessionID string) *JSONLEventLogger[G] {
	return &JSONLEventLogger[G]{
		dir:       dir,
		sessionID: sessionID,
		timeout:   5 * time.Minute,
	}
}

func (l *JSONLEventLogger[G]) LogEvent(event model.Event[G]) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if err := l.ensureFileOpen(); err != nil {
		return err
	}

	eventData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	_, err = l.file.Write(append(eventData, '\n'))
	if err != nil {
		return fmt.Errorf("failed to write event to file: %w", err)
	}

	if err := l.file.Sync(); err != nil {
		return fmt.Errorf("failed to sync file: %w", err)
	}

	l.resetTimer()
	return nil
}

func (l *JSONLEventLogger[G]) ensureFileOpen() error {
	if l.file != nil {
		return nil
	}

	if err := os.MkdirAll(l.dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", l.dir, err)
	}

	filename := filepath.Join(l.dir, l.sessionID+".jsonl")
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file %s: %w", filename, err)
	}

	l.file = file
	return nil
}

func (l *JSONLEventLogger[G]) resetTimer() {
	if l.timer != nil {
		l.timer.Stop()
	}

	l.timer = time.AfterFunc(l.timeout, func() {
		l.mutex.Lock()
		defer l.mutex.Unlock()
		l.closeFile()
	})
}

func (l *JSONLEventLogger[G]) closeFile() {
	if l.file != nil {
		l.file.Close()
		l.file = nil
	}
	if l.timer != nil {
		l.timer.Stop()
		l.timer = nil
	}
}

type NoopEventLogger[G model.Game] struct{}

func (n *NoopEventLogger[G]) LogEvent(_ model.Event[G]) error {
	return nil
}
