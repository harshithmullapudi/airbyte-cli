package common

import "github.com/spf13/viper"

func GetFullApiURL(path string) string {
	api_url := viper.GetString("api_url")
	return api_url + path
}
