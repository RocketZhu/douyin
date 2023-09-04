package config

import (
	"github.com/BurntSushi/toml"
	"strings"
)

type Server struct {
	ip   string
	port int
}
type Path struct {
	FfmpegPath       string `toml:"ffmpeg_path"`
	StaticSourcePath string `toml:"static_source_path"`
}
type Config struct {
	Server `toml:"server"`
	Path   `toml:"path"`
}

var ServerConfig Config

func init() {
	if _, err := toml.DecodeFile("/Users/rocketzhu/Downloads/douyin/douyin/config/config.toml", &ServerConfig); err != nil {
		panic(err)
	}
	//去除左右的空格
	strings.Trim(ServerConfig.Server.ip, " ")
	strings.Trim(ServerConfig.FfmpegPath, " ")
	strings.Trim(ServerConfig.StaticSourcePath, " ")
}
