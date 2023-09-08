package api

import (
	"github.com/DMonkey83/MyFitnessApp/util"
	"github.com/go-playground/validator/v10"
)

var ValidCompletion validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if completion, ok := fieldLevel.Field().Interface().(util.Completionenum); ok {
		return util.IsSupportedCompletion(completion)
	}
	return false
}

var ValidDifficulty validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if difficulty, ok := fieldLevel.Field().Interface().(util.Difficulty); ok {
		return util.IsSupportedDifficulty(difficulty)
	}
	return false
}

var ValidEquipment validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if eq, ok := fieldLevel.Field().Interface().(util.Equipmenttype); ok {
		return util.IsSupportedEquipment(eq)
	}
	return false
}

var ValidFatigueLevel validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if fLevel, ok := fieldLevel.Field().Interface().(util.Fatiguelevel); ok {
		return util.IsSupportedFatigueLevel(fLevel)
	}
	return false
}

var ValidMuscleGroup validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if mGroup, ok := fieldLevel.Field().Interface().(util.Musclegroupenum); ok {
		return util.IsSupportedMuscleGroup(mGroup)
	}
	return false
}

var ValidRating validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if rating, ok := fieldLevel.Field().Interface().(util.Rating); ok {
		return util.IsSupportedRating(rating)
	}
	return false
}

var ValidWeightUnit validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if wUnit, ok := fieldLevel.Field().Interface().(util.Weightunit); ok {
		return util.IsSupportedWeightUnit(wUnit)
	}
	return false
}

var ValidGoal validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if goal, ok := fieldLevel.Field().Interface().(util.Workoutgoalenum); ok {
		return util.IsSupportedGoal(util.Workoutgoalenum(goal))
	}
	return false
}

var ValidVisibility validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if visibility, ok := fieldLevel.Field().Interface().(util.Visibility); ok {
		return util.IsSupportedVisibility(util.Visibility(visibility))
	}
	return false
}
