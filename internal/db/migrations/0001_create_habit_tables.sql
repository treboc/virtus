-- Habits table to store habit definitions
CREATE TABLE habits (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title VARCHAR(255) NOT NULL,
  note TEXT,
  tracking_type TEXT NOT NULL CHECK (
    tracking_type IN ('daily', 'weekly', 'monthly', 'interval')
  ),
  -- For weekly: Bitmap of days (e.g., '0111110' means Mon-Fri)
  -- For monthly: Days of month as JSON array (e.g., '[1,15]')
  -- For interval: Number of days between repetitions
  tracking_config TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  active BOOLEAN DEFAULT true
);

-- Habit completions to track when habits are done
CREATE TABLE habit_completions (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  habit_id INTEGER NOT NULL,
  completed_at TIMESTAMP NOT NULL,
  FOREIGN KEY (habit_id) REFERENCES habits (id)
);

-- Daily mood tracking
CREATE TABLE daily_moods (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  date DATE UNIQUE NOT NULL, -- Ensures one mood entry per day
  score INTEGER NOT NULL CHECK (
    score >= 1
    AND score <= 10
  ),
  note TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better query performance
CREATE INDEX idx_habit_completions_habit_id ON habit_completions (habit_id);

CREATE INDEX idx_habit_completions_completed_at ON habit_completions (completed_at);

CREATE INDEX idx_daily_moods_date ON daily_moods (date);
