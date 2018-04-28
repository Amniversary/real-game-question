package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

type DBInfo struct {
	User   string
	Pass   string
	Host   string
	DBName string
}

type Etcd struct {
	Address          []string
	RegisterTTL      int64
	RegisterInterval int64
}

type Config struct {
	Debug     bool
	Path      string
	Version   string
	IfShowSql bool

	ServerName string

	DBInfo
	Etcd
}

func NewConfig() *Config {
	c := new(Config)
	initFlag(c)

	if c.Path == "" {
		fmt.Println("server must run with config file, please check.")
		os.Exit(0)
	}

	cfg, err := ini.Load(c.Path)
	if err != nil {
		fmt.Printf("ini[%s] load error: %v\n", c.Path, err)
		os.Exit(1)
	}
	cfg.BlockMode = false
	err = cfg.MapTo(c)
	if err != nil {
		fmt.Printf("config MapTo error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("config: %+v\n", c)

	return c
}

func initFlag(c *Config) {
	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	v := fs.Bool("v", false, "Print version and exit")
	fs.StringVar(&c.Path, "c", "", "server config file.")

	fs.Parse(os.Args[1:])
	fs.Usage = func() {
		fmt.Println("Usage: " + os.Args[0] + " -c config.ini")
		fmt.Printf("\nglobal flags:\n")
		fs.PrintDefaults()
	}

	if *v {
		fmt.Println("version: 1.0.0")
		os.Exit(0)
	}
}
