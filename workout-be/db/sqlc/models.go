// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Completionenum string

const (
	CompletionenumCompleted  Completionenum = "Completed"
	CompletionenumIncomplete Completionenum = "Incomplete"
	CompletionenumNotStarted Completionenum = "NotStarted"
)

func (e *Completionenum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Completionenum(s)
	case string:
		*e = Completionenum(s)
	default:
		return fmt.Errorf("unsupported scan type for Completionenum: %T", src)
	}
	return nil
}

type NullCompletionenum struct {
	Completionenum Completionenum `json:"completionenum"`
	Valid          bool           `json:"valid"` // Valid is true if Completionenum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCompletionenum) Scan(value interface{}) error {
	if value == nil {
		ns.Completionenum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Completionenum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCompletionenum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Completionenum), nil
}

type Difficulty string

const (
	DifficultyBeginner     Difficulty = "Beginner"
	DifficultyIntermediate Difficulty = "Intermediate"
	DifficultyAdvanced     Difficulty = "Advanced"
)

func (e *Difficulty) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Difficulty(s)
	case string:
		*e = Difficulty(s)
	default:
		return fmt.Errorf("unsupported scan type for Difficulty: %T", src)
	}
	return nil
}

type NullDifficulty struct {
	Difficulty Difficulty `json:"difficulty"`
	Valid      bool       `json:"valid"` // Valid is true if Difficulty is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullDifficulty) Scan(value interface{}) error {
	if value == nil {
		ns.Difficulty, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Difficulty.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullDifficulty) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Difficulty), nil
}

type Equipmenttype string

const (
	EquipmenttypeBarbell    Equipmenttype = "Barbell"
	EquipmenttypeDumbbell   Equipmenttype = "Dumbbell"
	EquipmenttypeMachine    Equipmenttype = "Machine"
	EquipmenttypeBodyweight Equipmenttype = "Bodyweight"
	EquipmenttypeOther      Equipmenttype = "Other"
)

func (e *Equipmenttype) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Equipmenttype(s)
	case string:
		*e = Equipmenttype(s)
	default:
		return fmt.Errorf("unsupported scan type for Equipmenttype: %T", src)
	}
	return nil
}

type NullEquipmenttype struct {
	Equipmenttype Equipmenttype `json:"equipmenttype"`
	Valid         bool          `json:"valid"` // Valid is true if Equipmenttype is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullEquipmenttype) Scan(value interface{}) error {
	if value == nil {
		ns.Equipmenttype, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Equipmenttype.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullEquipmenttype) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Equipmenttype), nil
}

type Fatiguelevel string

const (
	FatiguelevelVeryLight Fatiguelevel = "Very Light"
	FatiguelevelLight     Fatiguelevel = "Light"
	FatiguelevelModerate  Fatiguelevel = "Moderate"
	FatiguelevelHeavy     Fatiguelevel = "Heavy"
	FatiguelevelVeryHeavy Fatiguelevel = "Very Heavy"
)

func (e *Fatiguelevel) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Fatiguelevel(s)
	case string:
		*e = Fatiguelevel(s)
	default:
		return fmt.Errorf("unsupported scan type for Fatiguelevel: %T", src)
	}
	return nil
}

type NullFatiguelevel struct {
	Fatiguelevel Fatiguelevel `json:"fatiguelevel"`
	Valid        bool         `json:"valid"` // Valid is true if Fatiguelevel is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFatiguelevel) Scan(value interface{}) error {
	if value == nil {
		ns.Fatiguelevel, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Fatiguelevel.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullFatiguelevel) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Fatiguelevel), nil
}

type Musclegroupenum string

const (
	MusclegroupenumChest     Musclegroupenum = "Chest"
	MusclegroupenumBack      Musclegroupenum = "Back"
	MusclegroupenumLegs      Musclegroupenum = "Legs"
	MusclegroupenumShoulders Musclegroupenum = "Shoulders"
	MusclegroupenumArms      Musclegroupenum = "Arms"
	MusclegroupenumAbs       Musclegroupenum = "Abs"
	MusclegroupenumCardio    Musclegroupenum = "Cardio"
)

func (e *Musclegroupenum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Musclegroupenum(s)
	case string:
		*e = Musclegroupenum(s)
	default:
		return fmt.Errorf("unsupported scan type for Musclegroupenum: %T", src)
	}
	return nil
}

type NullMusclegroupenum struct {
	Musclegroupenum Musclegroupenum `json:"musclegroupenum"`
	Valid           bool            `json:"valid"` // Valid is true if Musclegroupenum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullMusclegroupenum) Scan(value interface{}) error {
	if value == nil {
		ns.Musclegroupenum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Musclegroupenum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullMusclegroupenum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Musclegroupenum), nil
}

type Rating string

const (
	Rating1 Rating = "1"
	Rating2 Rating = "2"
	Rating3 Rating = "3"
	Rating4 Rating = "4"
	Rating5 Rating = "5"
)

func (e *Rating) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Rating(s)
	case string:
		*e = Rating(s)
	default:
		return fmt.Errorf("unsupported scan type for Rating: %T", src)
	}
	return nil
}

type NullRating struct {
	Rating Rating `json:"rating"`
	Valid  bool   `json:"valid"` // Valid is true if Rating is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRating) Scan(value interface{}) error {
	if value == nil {
		ns.Rating, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Rating.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRating) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Rating), nil
}

type Visibility string

const (
	VisibilityPublic  Visibility = "Public"
	VisibilityPrivate Visibility = "Private"
)

func (e *Visibility) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Visibility(s)
	case string:
		*e = Visibility(s)
	default:
		return fmt.Errorf("unsupported scan type for Visibility: %T", src)
	}
	return nil
}

type NullVisibility struct {
	Visibility Visibility `json:"visibility"`
	Valid      bool       `json:"valid"` // Valid is true if Visibility is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVisibility) Scan(value interface{}) error {
	if value == nil {
		ns.Visibility, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Visibility.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVisibility) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Visibility), nil
}

type Weightunit string

const (
	WeightunitKg Weightunit = "kg"
	WeightunitLb Weightunit = "lb"
)

func (e *Weightunit) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Weightunit(s)
	case string:
		*e = Weightunit(s)
	default:
		return fmt.Errorf("unsupported scan type for Weightunit: %T", src)
	}
	return nil
}

type NullWeightunit struct {
	Weightunit Weightunit `json:"weightunit"`
	Valid      bool       `json:"valid"` // Valid is true if Weightunit is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullWeightunit) Scan(value interface{}) error {
	if value == nil {
		ns.Weightunit, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Weightunit.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullWeightunit) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Weightunit), nil
}

type Workoutgoalenum string

const (
	WorkoutgoalenumBuildMuscle      Workoutgoalenum = "Build Muscle"
	WorkoutgoalenumLoseWeight       Workoutgoalenum = "Lose Weight"
	WorkoutgoalenumImproveEndurance Workoutgoalenum = "Improve Endurance"
	WorkoutgoalenumMaintainFitness  Workoutgoalenum = "Maintain Fitness"
	WorkoutgoalenumToneBody         Workoutgoalenum = "Tone Body"
	WorkoutgoalenumCustom           Workoutgoalenum = "Custom"
)

func (e *Workoutgoalenum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Workoutgoalenum(s)
	case string:
		*e = Workoutgoalenum(s)
	default:
		return fmt.Errorf("unsupported scan type for Workoutgoalenum: %T", src)
	}
	return nil
}

type NullWorkoutgoalenum struct {
	Workoutgoalenum Workoutgoalenum `json:"workoutgoalenum"`
	Valid           bool            `json:"valid"` // Valid is true if Workoutgoalenum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullWorkoutgoalenum) Scan(value interface{}) error {
	if value == nil {
		ns.Workoutgoalenum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Workoutgoalenum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullWorkoutgoalenum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Workoutgoalenum), nil
}

type Availableplanexercise struct {
	ID           int64  `json:"id"`
	PlanID       int64  `json:"plan_id"`
	ExerciseName string `json:"exercise_name"`
	Sets         int32  `json:"sets"`
	RestDuration string `json:"rest_duration"`
	Notes        string `json:"notes"`
}

type Availableworkoutplan struct {
	PlanID          int64           `json:"plan_id"`
	PlanName        string          `json:"plan_name"`
	Description     string          `json:"description"`
	Goal            Workoutgoalenum `json:"goal"`
	Difficulty      Difficulty      `json:"difficulty"`
	IsPublic        Visibility      `json:"is_public"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	CreatorUsername string          `json:"creator_username"`
}

type Exercise struct {
	ExerciseName      string          `json:"exercise_name"`
	EquipmentRequired Equipmenttype   `json:"equipment_required"`
	Description       string          `json:"description"`
	MuscleGroupName   Musclegroupenum `json:"muscle_group_name"`
	CreatedAt         time.Time       `json:"created_at"`
}

type Exerciselog struct {
	ExerciseLogID        int64     `json:"exercise_log_id"`
	LogID                int64     `json:"log_id"`
	ExerciseName         string    `json:"exercise_name"`
	SetsCompleted        int32     `json:"sets_completed"`
	RepetitionsCompleted int32     `json:"repetitions_completed"`
	WeightLifted         int32     `json:"weight_lifted"`
	Notes                string    `json:"notes"`
	CreatedAt            time.Time `json:"created_at"`
}

type Exerciseset struct {
	SetID                int64 `json:"set_id"`
	ExerciseLogID        int64 `json:"exercise_log_id"`
	SetNumber            int32 `json:"set_number"`
	WeightLifted         int32 `json:"weight_lifted"`
	RepetitionsCompleted int32 `json:"repetitions_completed"`
}

type Maxrepgoal struct {
	GoalID       int64     `json:"goal_id"`
	Username     string    `json:"username"`
	ExerciseName string    `json:"exercise_name"`
	GoalReps     int32     `json:"goal_reps"`
	Notes        string    `json:"notes"`
	CreatedAt    time.Time `json:"created_at"`
}

type Maxweightgoal struct {
	GoalID       int64     `json:"goal_id"`
	Username     string    `json:"username"`
	ExerciseName string    `json:"exercise_name"`
	GoalWeight   int32     `json:"goal_weight"`
	Notes        string    `json:"notes"`
	CreatedAt    time.Time `json:"created_at"`
}

type Oneoffworkoutexercise struct {
	ID              int32           `json:"id"`
	WorkoutID       int64           `json:"workout_id"`
	ExerciseName    string          `json:"exercise_name"`
	Description     string          `json:"description"`
	MuscleGroupName Musclegroupenum `json:"muscle_group_name"`
	CreatedAt       time.Time       `json:"created_at"`
}

type Set struct {
	SetID         int64     `json:"set_id"`
	ExerciseName  string    `json:"exercise_name"`
	SetNumber     int32     `json:"set_number"`
	Weight        int32     `json:"weight"`
	RestDuration  string    `json:"rest_duration"`
	RepsCompleted int32     `json:"reps_completed"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"created_at"`
}

type User struct {
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	PasswordHash      string    `json:"password_hash"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type Userprofile struct {
	UserProfileID int64      `json:"user_profile_id"`
	Username      string     `json:"username"`
	FullName      string     `json:"full_name"`
	Age           int32      `json:"age"`
	Gender        string     `json:"gender"`
	HeightCm      int32      `json:"height_cm"`
	HeightFtIn    string     `json:"height_ft_in"`
	PreferredUnit Weightunit `json:"preferred_unit"`
	CreatedAt     time.Time  `json:"created_at"`
}

type Weightentry struct {
	WeightEntryID int64     `json:"weight_entry_id"`
	Username      string    `json:"username"`
	EntryDate     time.Time `json:"entry_date"`
	WeightKg      int32     `json:"weight_kg"`
	WeightLb      int32     `json:"weight_lb"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"created_at"`
}

type Workout struct {
	WorkoutID           int64        `json:"workout_id"`
	Username            string       `json:"username"`
	WorkoutDate         time.Time    `json:"workout_date"`
	WorkoutDuration     string       `json:"workout_duration"`
	FatigueLevel        Fatiguelevel `json:"fatigue_level"`
	Notes               string       `json:"notes"`
	TotalCaloriesBurned int32        `json:"total_calories_burned"`
	TotalDistance       int32        `json:"total_distance"`
	TotalRepetitions    int32        `json:"total_repetitions"`
	TotalSets           int32        `json:"total_sets"`
	TotalWeightLifted   int32        `json:"total_weight_lifted"`
	CreatedAt           time.Time    `json:"created_at"`
}

type Workoutlog struct {
	LogID               int64        `json:"log_id"`
	Username            string       `json:"username"`
	PlanID              int64        `json:"plan_id"`
	LogDate             time.Time    `json:"log_date"`
	Rating              Rating       `json:"rating"`
	FatigueLevel        Fatiguelevel `json:"fatigue_level"`
	OverallFeeling      string       `json:"overall_feeling"`
	Comments            string       `json:"comments"`
	WorkoutDuration     string       `json:"workout_duration"`
	TotalCaloriesBurned int32        `json:"total_calories_burned"`
	TotalDistance       int32        `json:"total_distance"`
	TotalRepetitions    int32        `json:"total_repetitions"`
	TotalSets           int32        `json:"total_sets"`
	TotalWeightLifted   int32        `json:"total_weight_lifted"`
	CreatedAt           time.Time    `json:"created_at"`
}

type Workoutplan struct {
	PlanID      int64           `json:"plan_id"`
	Username    string          `json:"username"`
	PlanName    string          `json:"plan_name"`
	Description string          `json:"description"`
	StartDate   time.Time       `json:"start_date"`
	EndDate     time.Time       `json:"end_date"`
	Goal        Workoutgoalenum `json:"goal"`
	Difficulty  Difficulty      `json:"difficulty"`
	IsPublic    Visibility      `json:"is_public"`
	CreatedAt   time.Time       `json:"created_at"`
}
