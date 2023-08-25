-- Equipment Type Enum
CREATE TYPE EquipmentType AS ENUM (
    'Barbell',
    'Dumbbell',
    'Machine',
    'Bodyweight',
    'Other'
);

-- Weight Unit Enum
CREATE TYPE WeightUnit AS ENUM (
    'kg',  -- Kilograms
    'lb'   -- Pounds
);

-- Create the muscle_group_enum type if it doesn't exist
CREATE TYPE muscle_group_enum AS ENUM (
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
    password_changed_at timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00z'
);

-- UserProfile Table
CREATE TABLE IF NOT EXISTS UserProfile (
    user_profile_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) REFERENCES "User"(username) UNIQUE NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(10) NOT NULL,
    height_cm FLOAT NOT NULL,
    height_ft_in VARCHAR(20),
    preferred_unit WeightUnit NOT NULL
);

-- Equipment Table
CREATE TABLE IF NOT EXISTS Equipment (
    equipment_id BIGSERIAL PRIMARY KEY,
    equipment_name VARCHAR(255) NOT NULL,
    description TEXT,
    equipment_type EquipmentType NOT NULL
);

-- Workout Table
CREATE TABLE IF NOT EXISTS Workout (
    workout_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) REFERENCES "User"(username) UNIQUE NOT NULL,
    workout_date DATE NOT NULL,
    workout_duration INTERVAL,
    notes TEXT
);


-- Exercise Table
CREATE TABLE IF NOT EXISTS Exercise (
    exercise_id BIGSERIAL PRIMARY KEY,
    workout_id BIGINT REFERENCES Workout(workout_id) NOT NULL,
    exercise_name VARCHAR(255) NOT NULL,
    description TEXT,
    equipment_id BIGINT REFERENCES Equipment(equipment_id)
);

-- MuscleGroup Table
CREATE TABLE IF NOT EXISTS MuscleGroup (
    muscle_group_id BIGSERIAL PRIMARY KEY,
    exercise_id BIGINT REFERENCES Exercise(exercise_id) NOT NULL,
    muscle_group_name VARCHAR(255) NOT NULL
);

-- Set Table
CREATE TABLE IF NOT EXISTS Set (
    set_id BIGSERIAL PRIMARY KEY,
    exercise_id BIGINT REFERENCES Exercise(exercise_id) NOT NULL,
    set_number INT NOT NULL,
    weight FLOAT,
    rest_duration INTERVAL,
    notes TEXT
);

-- Rep Table
CREATE TABLE IF NOT EXISTS Rep (
    rep_id BIGSERIAL PRIMARY KEY,
    set_id BIGINT REFERENCES Set(set_id) NOT NULL,
    rep_number INT NOT NULL,
    completed BOOLEAN,
    notes TEXT
);

-- WorkoutProgram Table
CREATE TABLE IF NOT EXISTS WorkoutProgram (
    program_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) REFERENCES "User"(username) UNIQUE NOT NULL,
    program_name VARCHAR(255) NOT NULL,
    description TEXT
);

-- WeightEntry Table
CREATE TABLE IF NOT EXISTS WeightEntry (
    weight_entry_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) REFERENCES "User"(username) UNIQUE NOT NULL,
    entry_date DATE NOT NULL,
    weight_kg FLOAT,
    weight_lb FLOAT,
    notes TEXT
);

-- MaxRepGoal Table
CREATE TABLE IF NOT EXISTS MaxRepGoal (
    goal_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) REFERENCES "User"(username) UNIQUE NOT NULL,
    exercise_id BIGINT REFERENCES Exercise(exercise_id) NOT NULL,
    goal_reps INT NOT NULL,
    notes TEXT
);

-- MaxWeightGoal Table
CREATE TABLE IF NOT EXISTS MaxWeightGoal (
    goal_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) REFERENCES "User"(username) UNIQUE NOT NULL,
    exercise_id BIGINT REFERENCES Exercise(exercise_id) NOT NULL,
    goal_weight FLOAT NOT NULL,
    notes TEXT
);

