package cmd

// This file is auto-generated!
// Do not edit.

import (
	"time"

	"github.com/spf13/cast"
)

func decodeBool(in, out interface{}) {
	*out.(*bool) = cast.ToBool(in)
}

// GetBool reads a boolean from the configuration file
func GetBool(key string, val *bool) {
	getConfig(key, val, decodeBool)
}

func decodeFloat64(in, out interface{}) {
	*out.(*float64) = cast.ToFloat64(in)
}

// GetFloat64 reads a 64-bit float from the configuration file
func GetFloat64(key string, val *float64) {
	getConfig(key, val, decodeFloat64)
}

func decodeInt(in, out interface{}) {
	*out.(*int) = cast.ToInt(in)
}

// GetInt reads an integer from the configuration file
func GetInt(key string, val *int) {
	getConfig(key, val, decodeInt)
}

func decodeIntSlice(in, out interface{}) {
	*out.(*[]int) = cast.ToIntSlice(in)
}

// GetIntSlice reads an integer array from the configuration file
func GetIntSlice(key string, val *[]int) {
	getConfig(key, val, decodeIntSlice)
}

func decodeString(in, out interface{}) {
	*out.(*string) = cast.ToString(in)
}

// GetString reads a string from the configuration file
func GetString(key string, val *string) {
	getConfig(key, val, decodeString)
}

func decodeStringMapBool(in, out interface{}) {
	*out.(*map[string]bool) = cast.ToStringMapBool(in)
}

// GetStringMapBool reads a boolean dictionary from the configuration file
func GetStringMapBool(key string, val *map[string]bool) {
	getConfig(key, val, decodeStringMapBool)
}

func decodeStringMapInt(in, out interface{}) {
	*out.(*map[string]int) = cast.ToStringMapInt(in)
}

// GetStringMapInt reads an int dictionary from the configuration file
func GetStringMapInt(key string, val *map[string]int) {
	getConfig(key, val, decodeStringMapInt)
}

func decodeStringMapString(in, out interface{}) {
	*out.(*map[string]string) = cast.ToStringMapString(in)
}

// GetStringMapString reads a string dictionary from the configuration file
func GetStringMapString(key string, val *map[string]string) {
	getConfig(key, val, decodeStringMapString)
}

func decodeStringMapStringSlice(in, out interface{}) {
	*out.(*map[string][]string) = cast.ToStringMapStringSlice(in)
}

// GetStringMapStringSlice reads a string array dictionary from the configuration file
func GetStringMapStringSlice(key string, val *map[string][]string) {
	getConfig(key, val, decodeStringMapStringSlice)
}

func decodeStringSlice(in, out interface{}) {
	*out.(*[]string) = cast.ToStringSlice(in)
}

// GetStringSlice reads a string array from the configuration file
func GetStringSlice(key string, val *[]string) {
	getConfig(key, val, decodeStringSlice)
}

func decodeTime(in, out interface{}) {
	*out.(*time.Time) = cast.ToTime(in)
}

// GetTime reads a time from the configuration file
func GetTime(key string, val *time.Time) {
	getConfig(key, val, decodeTime)
}

func decodeDuration(in, out interface{}) {
	*out.(*time.Duration) = cast.ToDuration(in)
}

// GetDuration reads a duration from the configuration file
func GetDuration(key string, val *time.Duration) {
	getConfig(key, val, decodeDuration)
}
