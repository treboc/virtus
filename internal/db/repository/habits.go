package repository

import (
	"context"
	"database/sql"

	db "github.com/treboc/virtus/internal/db/generated"
	"github.com/treboc/virtus/internal/habit"
)

type HabitRepository interface {
	Create(ctx context.Context, habit *habit.Habit) error
	Get(ctx context.Context, id int64) (*habit.Habit, error)
	List(ctx context.Context) ([]*habit.Habit, error)
	GetDueToday(ctx context.Context) ([]*habit.Habit, error)
	Complete(ctx context.Context, habitID int64) error
}

func (r *Repository) Create(ctx context.Context, habit *habit.Habit) error {
	params := db.CreateHabitParams{
		Title:          habit.Title,
		Note:           sql.NullString{String: habit.Note, Valid: habit.Note != ""},
		TrackingType:   string(habit.TrackingType),
		TrackingConfig: sql.NullString{String: habit.TrackingConfig, Valid: habit.TrackingConfig != ""},
	}

	result, err := r.q.CreateHabit(ctx, params)
	if err != nil {
		return err
	}

	habit.ID = result.ID
	return nil
}

func (r *Repository) Get(ctx context.Context, id int64) (*habit.Habit, error) {
	h, err := r.q.GetHabit(ctx, id)
	if err != nil {
		return nil, err
	}

	return &habit.Habit{
		ID:             h.ID,
		Title:          h.Title,
		Note:           h.Note.String,
		TrackingType:   habit.TrackingType(h.TrackingType),
		TrackingConfig: h.TrackingConfig.String,
		Active:         h.Active.Bool,
	}, nil
}

func (r *Repository) List(ctx context.Context) ([]*habit.Habit, error) {
	habits, err := r.q.GetHabits(ctx)
	if err != nil {
		return nil, err
	}

	var result []*habit.Habit
	for _, h := range habits {
		result = append(result, &habit.Habit{
			ID:             h.ID,
			Title:          h.Title,
			Note:           h.Note.String,
			TrackingType:   habit.TrackingType(h.TrackingType),
			TrackingConfig: h.TrackingConfig.String,
			Active:         h.Active.Bool,
		})
	}

	return result, nil
}

func (r *Repository) GetDueToday(ctx context.Context) ([]*habit.Habit, error) {
	habits, err := r.q.GetHabitsDueToday(ctx)
	if err != nil {
		return nil, err
	}

	var result []*habit.Habit
	for _, h := range habits {
		result = append(result, &habit.Habit{
			ID:             h.ID,
			Title:          h.Title,
			Note:           h.Note.String,
			TrackingType:   habit.TrackingType(h.TrackingType),
			TrackingConfig: h.TrackingConfig.String,
			Active:         h.Active.Bool,
		})
	}

	return result, nil
}

func (r *Repository) Complete(ctx context.Context, habitID int64) error {
	_, err := r.q.CompleteHabit(ctx, habitID)
	return err
}
