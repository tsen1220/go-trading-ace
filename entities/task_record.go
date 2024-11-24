package entities

import "time"

type TaskRecord struct {
	ID           int64      `db:"id"`            // SERIAL PRIMARY KEY
	Address      string     `db:"address"`       // VARCHAR(255) NOT NULL
	TaskID       int64      `db:"task_id"`       // INT NOT NULL REFERENCES tasks(id)
	RewardPoints int64      `db:"reward_points"` // BIGINT NOT NULL
	Amount       *int64     `db:"amount"`        // BIGINT NULL
	CompletedAt  *time.Time `db:"completed_at"`  // TIMESTAMP NULL
	CreatedAt    time.Time  `db:"created_at"`    // TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	UpdatedAt    time.Time  `db:"updated_at"`    // TIMESTAMP DEFAULT CURRENT_TIMESTAMP
}
