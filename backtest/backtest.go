package backtest

import (
	"fmt"

	"github.com/spf13/viper"
)

// write the info back to the yaml file and download the data
func btnDownload_Clicked(instr, bd, ed, sub, cal string) {
	// write the info back to the yaml file
	// download the data
}

// make a string slice to a string, separated by "," but not end with ","
func SS2S(ss []string) (s string) {
	for _, v := range ss {
		s = s + v + ","
	}
	// if s is not empty, remove the last ","
	if s != "" {
		s = s[:len(s)-1]
	}

	return
}

// read the yaml file from dir and output the info
func BtnReadConf_Clicked(dir string) (instr, bd, ed, inds, sub, cal string) {
	// read the yaml file from dir

	viper.SetConfigName("BackTest") // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name

	// optionally look for config in the working directory
	viper.AddConfigPath(dir)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// get the default.sinstrnames from the yaml file to instr
	instr = SS2S(viper.GetStringSlice("default.sinstrnames"))
	bd = viper.GetString("default.begindate")
	ed = viper.GetString("default.enddate")
	inds = SS2S(viper.GetStringSlice("default.sindinames"))
	sub = SS2S(viper.GetStringSlice("default.scsvdatafields"))
	cal = SS2S(viper.GetStringSlice("default.sadfields"))
	// output the info
	return
}
