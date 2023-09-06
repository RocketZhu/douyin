package config

import (
	"strings"

	"github.com/BurntSushi/toml"
)

type Server struct {
	Ip   string
	Port int
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
	if _, err := toml.DecodeFile("/Users/rocketzhu/CS/Projects/douyin/config/config.toml", &ServerConfig); err != nil {
		panic(err)
	}
	//去除左右的空格
	ServerConfig.Server.Ip = strings.Trim(ServerConfig.Server.Ip, " ")
	ServerConfig.FfmpegPath = strings.Trim(ServerConfig.FfmpegPath, " ")
	ServerConfig.StaticSourcePath = strings.Trim(ServerConfig.StaticSourcePath, " ")
}
