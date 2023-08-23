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
    user_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

-- UserProfile Table
CREATE TABLE IF NOT EXISTS UserProfile (
    user_profile_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES "User"(user_id) UNIQUE NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(10) NOT NULL,
    height_cm FLOAT NOT NULL,
    height_ft_in VARCHAR(20),
    preferred_unit WeightUnit NOT NULL
);

-- MuscleGroup Table
CREATE TABLE IF NOT EXISTS MuscleGroup (
    muscle_group_id BIGSERIAL PRIMARY KEY,
    muscle_group_name VARCHAR(255) NOT NULL
);

-- Equipment Table
CREATE TABLE IF NOT EXISTS Equipment (
    equipment_id BIGSERIAL PRIMARY KEY,
    equipment_name VARCHAR(255) NOT NULL,
    description TEXT,
    equipment_type EquipmentType NOT NULL
);

-- Exercise Table
CREATE TABLE IF NOT EXISTS Exercise (
    exercise_id BIGSERIAL PRIMARY KEY,
    exercise_name VARCHAR(255) NOT NULL,
    muscle_group muscle_group_enum not null,
    description TEXT,
    equipment_id BIGINT REFERENCES Equipment(equipment_id)
);

-- Workout Table
CREATE TABLE IF NOT EXISTS Workout (
    workout_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES "User"(user_id) NOT NULL,
    workout_date DATE NOT NULL,
    workout_duration INTERVAL,
    notes TEXT
);

-- Set Table
CREATE TABLE IF NOT EXISTS Set (
    set_id BIGSERIAL PRIMARY KEY,
    workout_id BIGINT REFERENCES Workout(workout_id) NOT NULL,
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
    user_id BIGINT REFERENCES "User"(user_id) NOT NULL,
    program_name VARCHAR(255) NOT NULL,
    description TEXT
);

-- ProgramWorkout Table
CREATE TABLE IF NOT EXISTS ProgramWorkout (
    program_workout_id BIGSERIAL PRIMARY KEY,
    program_id BIGINT REFERENCES WorkoutProgram(program_id) NOT NULL,
    workout_id BIGINT REFERENCES Workout(workout_id) NOT NULL,
    day_of_week INT NOT NULL,
    notes TEXT
);

-- WeightEntry Table
CREATE TABLE IF NOT EXISTS WeightEntry (
    weight_entry_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES "User"(user_id) NOT NULL,
    entry_date DATE NOT NULL,
    weight_kg FLOAT,
    weight_lb FLOAT,
    notes TEXT
);

-- MaxRepGoal Table
CREATE TABLE IF NOT EXISTS MaxRepGoal (
    goal_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES "User"(user_id) NOT NULL,
    exercise_id BIGINT REFERENCES Exercise(exercise_id) NOT NULL,
    goal_reps INT NOT NULL,
    notes TEXT
);

-- MaxWeightGoal Table
CREATE TABLE IF NOT EXISTS MaxWeightGoal (
    goal_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES "User"(user_id) NOT NULL,
    exercise_id BIGINT REFERENCES Exercise(exercise_id) NOT NULL,
    goal_weight FLOAT NOT NULL,
    notes TEXT
);

