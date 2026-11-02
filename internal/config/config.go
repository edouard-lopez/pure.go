// Package config reads Pure settings from PURE_* environment variables.
// Names match Fish's pure_* variables in upper case, e.g. PURE_SYMBOL_PROMPT.
package config

import "os"

// Get returns the value of the PURE_* variable name, or fallback when it is unset.
// An empty value still overrides the fallback, like a Fish global set to an empty string.
func Get(name string, fallback string) string {
	if value, isSet := os.LookupEnv(name); isSet {
		return value
	}
	return fallback
}
