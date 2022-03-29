package app

import "fmt"

type Operator interface {
	Next() error
	Pause() error
	Play() error
	PlayPause() error
	Previouse() error
}

type OperatorService struct {
	Operator Operator
}

type Operation int

const (
	OperationNext Operation = iota
	OperationPause
	OperationPlay
	OperationPlayPause
	OperationPreviouse
)

var OperationMap map[string]Operation = map[string]Operation{
	"next":      OperationNext,
	"pause":     OperationPause,
	"play":      OperationPause,
	"playpause": OperationPlayPause,
	"previouse": OperationPreviouse,
}

func (s *OperatorService) Operate(operation Operation) error {
	switch operation {
	case OperationNext:
		if err := s.Operator.Next(); err != nil {
			return fmt.Errorf("failed to operate next: %w", err)
		}

	case OperationPause:
		if err := s.Operator.Pause(); err != nil {
			return fmt.Errorf("failed to operate pause: %w", err)
		}

	case OperationPlay:
		if err := s.Operator.Play(); err != nil {
			return fmt.Errorf("failed to operate play: %w", err)
		}

	case OperationPlayPause:
		if err := s.Operator.PlayPause(); err != nil {
			return fmt.Errorf("failed to operate play/pause: %w", err)
		}

	case OperationPreviouse:
		if err := s.Operator.Previouse(); err != nil {
			return fmt.Errorf("failed to operate Previouse: %w", err)
		}
	}

	return nil
}
