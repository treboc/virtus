-- name: GetHabitsDueToday :many
WITH
  habit_status AS (
    SELECT
      h.id,
      h.title,
      h.tracking_type,
      h.tracking_config,
      (
        SELECT
          CASE
            WHEN date(hc.completed_at) = date('now') THEN 1
            ELSE 0
          END
        FROM
          habit_completions hc
        WHERE
          hc.habit_id = h.id
        ORDER BY
          completed_at DESC
        LIMIT
          1
      ) as completed_today
    FROM
      habits h
    WHERE
      h.active = true
  )
SELECT
  id,
  title,
  tracking_type,
  tracking_config,
  completed_today
FROM
  habit_status
WHERE
  -- Daily habits are always due
  (tracking_type = 'daily')
  OR
  -- Weekly habits: Check if current weekday bit is set
  (
    tracking_type = 'weekly'
    AND substr (
      tracking_config,
      cast(strftime ('%w', 'now') as integer) + 1,
      1
    ) = '1'
  )
  OR
  -- Monthly habits: Check if current day is in the JSON array
  (
    tracking_type = 'monthly'
    AND CAST(strftime ('%d', 'now') as integer) IN (
      SELECT
        CAST(value as integer)
      FROM
        json_each (tracking_config)
    )
  )
  OR
  -- Interval habits: Check days since last completion
  (
    tracking_type = 'interval'
    AND (
      SELECT
        CAST(
          julianday ('now') - julianday (max(completed_at)) as integer
        )
      FROM
        habit_completions
      WHERE
        habit_id = habit_status.id
    ) >= CAST(tracking_config as integer)
  );

-- name: CreateHabit :one
INSERT INTO
  habits (title, note, tracking_type, tracking_config)
VALUES
  (?, ?, ?, ?) RETURNING *;

-- name: CompleteHabit :one
INSERT INTO
  habit_completions (habit_id, completed_at)
VALUES
  (?, CURRENT_TIMESTAMP) RETURNING *;

-- name: GetHabits :many
SELECT
  id,
  title,
  note,
  tracking_type,
  tracking_config,
  created_at,
  active
FROM
  habits;

-- name: GetHabit :one
SELECT
  id,
  title,
  note,
  tracking_type,
  tracking_config,
  created_at,
  active
FROM
  habits
WHERE
  id = ?;
