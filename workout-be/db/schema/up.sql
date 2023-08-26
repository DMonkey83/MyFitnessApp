-- Equipment Type Enum
CREATE TYPE EquipmentType AS ENUM (
    'Barbell',
    'Dumbbell',
    'Machine',
    'Bodyweight',
    'Other'
);
--
-- Completion Type Enum
CREATE TYPE CompletionEnum AS ENUM (
  'Completed',
  'Incomplete',
  'NotStarted'
);

-- Weight Unit Enum
CREATE TYPE WeightUnit AS ENUM (
    'kg',  -- Kilograms
    'lb'   -- Pounds
);

-- Create the muscle_group_enum type if it doesn't exist
CREATE TYPE MuscleGroupEnum AS ENUM (
    'Chest',
    'Back',
    'Legs',
    'Shoulders',
    'Arms',
    'Abs',
    'Cardio'
    -- Add more muscle groups as needed
);

-- User Table
CREATE TABLE IF NOT EXISTS "User" (
    username VARCHAR(255) NOT NULL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    password_changed_at timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00z',
    "created_at" timestamptz NOT NULL DEFAULT (now()),

  -- Add a unique constraint to the email column
    CONSTRAINT unique_email UNIQUE (email)
);

-- UserProfile Table
CREATE TABLE IF NOT EXISTS UserProfile (
    user_profile_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) REFERENCES "User"(username) UNIQUE NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(10) NOT NULL,
    height_cm INT NOT NULL,
    height_ft_in VARCHAR(20) NOT NULL,
    preferred_unit WeightUnit NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- Equipment Table
CREATE TABLE IF NOT EXISTS Equipment (
    equipment_name VARCHAR(255) NOT NULL PRIMARY KEY,
    description VARCHAR(255) NOT NULL,
    equipment_type EquipmentType NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- Workout Table
CREATE TABLE IF NOT EXISTS Workout (
    workout_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) REFERENCES "User"(username) NOT NULL,
    workout_date timestamptz NOT NULL DEFAULT (now()),
    workout_duration VARCHAR(8) NOT NULL DEFAULT (''),
    notes VARCHAR(255) NOT NULL DEFAULT(''),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- Exercise Table
CREATE TABLE IF NOT EXISTS Exercise (
    exercise_id BIGSERIAL PRIMARY KEY,
    workout_id BIGINT REFERENCES Workout(workout_id) NOT NULL,
    equipment_name VARCHAR(255) REFERENCES Equipment(equipment_name) NOT NULL,
    exercise_name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL DEFAULT (''),
    muscle_group_name MuscleGroupEnum NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),

    CONSTRAINT unique_equipment_name_per_exercise_id UNIQUE (equipment_name, exercise_id)
);

-- Set Table
CREATE TABLE IF NOT EXISTS Set (
    set_id BIGSERIAL PRIMARY KEY,
    exercise_id BIGINT REFERENCES Exercise(exercise_id) NOT NULL,
    set_number INT NOT NULL DEFAULT (1),
    weight int NOT NULL DEFAULT (1),
    rest_duration VARCHAR(8) NOT NULL DEFAULT (''),
    notes VARCHAR(255) NOT NULL DEFAULT(''),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- Rep Table
CREATE TABLE IF NOT EXISTS Rep (
    rep_id BIGSERIAL PRIMARY KEY,
    set_id BIGINT REFERENCES Set(set_id) NOT NULL,
    rep_number INT NOT NULL,
    completion_status CompletionEnum NOT NULL,
    notes  VARCHAR(255) NOT NULL DEFAULT(''),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);


-- WeightEntry Table
CREATE TABLE IF NOT EXISTS WeightEntry (
    weight_entry_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) REFERENCES "User"(username) NOT NULL,
    entry_date timestamptz NOT NULL DEFAULT(now()),
    weight_kg INT NOT NULL DEFAULT(0),
    weight_lb INT NOT NULL DEFAULT(0),
    notes VARCHAR(255) NOT NULL DEFAULT (''),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- MaxRepGoal Table
CREATE TABLE IF NOT EXISTS MaxRepGoal (
    goal_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) REFERENCES "User"(username) NOT NULL,
    exercise_id BIGINT REFERENCES Exercise(exercise_id) UNIQUE NOT NULL,
    goal_reps INT NOT NULL,
    notes VARCHAR(255) NOT NULL DEFAULT (''),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- MaxWeightGoal Table
CREATE TABLE IF NOT EXISTS MaxWeightGoal (
    goal_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) REFERENCES "User"(username) NOT NULL,
    exercise_id BIGINT REFERENCES Exercise(exercise_id) UNIQUE NOT NULL,
    goal_weight INT NOT NULL,
    notes VARCHAR(255) NOT NULL DEFAULT (''),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

