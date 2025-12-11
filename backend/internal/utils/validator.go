package utils

import (
	"fmt"
	"regexp"
	"time"
)

// ValidateTCNumber validates Turkish Republic identification number (11 digits)
func ValidateTCNumber(tcNumber string) error {
	if tcNumber == "" {
		return fmt.Errorf("TC number is required")
	}

	// Check if it's exactly 11 digits
	matched, err := regexp.MatchString(`^\d{11}$`, tcNumber)
	if err != nil {
		return fmt.Errorf("error validating TC number: %v", err)
	}

	if !matched {
		return fmt.Errorf("TC number must be exactly 11 digits")
	}

	return nil
}

// ValidateEnum validates that a value is one of the allowed enum values
func ValidateEnum(value string, allowedValues []string, fieldName string) error {
	if value == "" {
		return fmt.Errorf("%s is required", fieldName)
	}

	for _, allowed := range allowedValues {
		if value == allowed {
			return nil
		}
	}

	return fmt.Errorf("invalid %s: must be one of %v", fieldName, allowedValues)
}

// ValidateDateRange validates that start date is before or equal to end date
func ValidateDateRange(startDate, endDate time.Time, fieldPrefix string) error {
	if startDate.IsZero() || endDate.IsZero() {
		return nil // Skip validation if dates are not provided
	}

	if startDate.After(endDate) {
		return fmt.Errorf("%s start date must be before or equal to end date", fieldPrefix)
	}

	return nil
}

// ValidateDateNotFuture validates that a date is not in the future
func ValidateDateNotFuture(date time.Time, fieldName string) error {
	if date.IsZero() {
		return nil // Skip validation if date is not provided
	}

	if date.After(time.Now()) {
		return fmt.Errorf("%s cannot be in the future", fieldName)
	}

	return nil
}

// ValidatePositiveInteger validates that a value is a positive integer
func ValidatePositiveInteger(value int, fieldName string) error {
	if value <= 0 {
		return fmt.Errorf("%s must be a positive integer", fieldName)
	}

	return nil
}

// ValidateBloodPressure validates that systolic is greater than diastolic
func ValidateBloodPressure(systolic, diastolic *float64) error {
	if systolic == nil || diastolic == nil {
		return nil // Skip validation if values are not provided
	}

	if *systolic <= *diastolic {
		return fmt.Errorf("systolic blood pressure must be greater than diastolic blood pressure")
	}

	return nil
}
