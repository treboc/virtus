package habit

type TrackingType string

const (
	Daily    TrackingType = "daily"
	Weekly   TrackingType = "weekly"
	Monthly  TrackingType = "monthly"
	Interval TrackingType = "interval"
)

type Habit struct {
	ID             int64
	Title          string
	Note           string
	TrackingType   TrackingType
	TrackingConfig string
	Active         bool
}
