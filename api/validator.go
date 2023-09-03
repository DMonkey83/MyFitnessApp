package api

import (
	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/go-playground/validator/v10"
)

var validCompletion validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if completion, ok := fieldLevel.Field().Interface().(util.Completionenum); ok {
		return util.IsSupportedCompletion(completion)
	}
	return false
}

var validDifficulty validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if difficulty, ok := fieldLevel.Field().Interface().(util.Difficulty); ok {
		return util.IsSupportedDifficulty(difficulty)
	}
	return false
}

var validEquipment validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if eq, ok := fieldLevel.Field().Interface().(util.Equipmenttype); ok {
		return util.IsSupportedEquipment(eq)
	}
	return false
}

var validFatigueLevel validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if fLevel, ok := fieldLevel.Field().Interface().(util.Fatiguelevel); ok {
		return util.IsSupportedFatigueLevel(fLevel)
	}
	return false
}

var validMuscleGroup validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if mGroup, ok := fieldLevel.Field().Interface().(util.Musclegroupenum); ok {
		return util.IsSupportedMuscleGroup(mGroup)
	}
	return false
}

var validRating validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if rating, ok := fieldLevel.Field().Interface().(util.Rating); ok {
		return util.IsSupportedRating(rating)
	}
	return false
}

var validWeightUnit validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if wUnit, ok := fieldLevel.Field().Interface().(util.Weightunit); ok {
		return util.IsSupportedWeightUnit(wUnit)
	}
	return false
}

var validGoal validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if goal, ok := fieldLevel.Field().Interface().(util.Workoutgoalenum); ok {
		return util.IsSupportedGoal(util.Workoutgoalenum(goal))
	}
	return false
}

var validVisibility validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if visibility, ok := fieldLevel.Field().Interface().(util.Visibility); ok {
		return util.IsSupportedVisibility(util.Visibility(visibility))
	}
	return false
}
