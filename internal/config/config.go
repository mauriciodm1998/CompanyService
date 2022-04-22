package config

import "flag"

var (
	DbPath     string
	OutputPath string
	HttpPort   string
)

func ParseConfiguration() {
	flag.StringVar(&DbPath, "db-path", "./", "Database file path")
	flag.StringVar(&HttpPort, "port", "8081", "Http port to listen")
	flag.StringVar(&OutputPath, "output", "~/", "Video output path")

	flag.Parse()
}
