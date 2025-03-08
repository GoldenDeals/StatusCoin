package share

import (
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func StringVarReplace(s string) string {
	s = strings.ReplaceAll(s, "$time", GetTimeFormatted())
	s = strings.ReplaceAll(s, "$date", GetDateFormatted())
	s = strings.ReplaceAll(s, "$version", viper.GetString("version"))
	s = strings.ReplaceAll(s, "$name", viper.GetString("name"))
	s = strings.ReplaceAll(s, "$uuid", uuid.NewString())

	return s
}
