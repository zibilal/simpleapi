package appctx

import (
	"bitbucket.org/kudoindonesia/terracotta_data_access_library"
	"bitbucket.org/kudoindonesia/terracotta_data_access_library/mongoconnector"
	"bitbucket.org/kudoindonesia/terracotta_data_access_library/persistence"
	"bitbucket.org/kudoindonesia/terracotta_data_access_library/persistence/mongodbpersistence"
	"bytes"
	"context"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

const (
	defaultConfigFlagVal = "configurations/App.yaml"
)

type Config struct {
	Address  string `yaml:"address"`
	Mode     string `yaml:"mode"`
	Database struct {
		ConnectionString string `yaml:"connection_string"`
		Timeout          int    `yaml:"timeout"`
		Name             string `yaml:"name"`
	} `yaml:"database"`
}

type AppContext struct {
	Config        *Config
	DataConnector connector.Connector
	Persistence   persistence.Persistence
}

func NewAppContext() *AppContext {
	ctx := new(AppContext)
	ctx.Config = new(Config)
	return ctx
}

func (c *AppContext) LoadAppContext(readers ...io.Reader) error {
	buff := bytes.NewBuffer([]byte{})
	for _, reader := range readers {
		dat, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}
		buff.Write(dat)
	}

	if err := yaml.Unmarshal(buff.Bytes(), c.Config); err != nil {
		return err
	}

	return nil
}

var (
	instance *AppContext
	once     sync.Once
)

func GetAppContext() *AppContext {
	once.Do(func() {
		instance = NewAppContext()

		file, err := os.Open(defaultConfigFlagVal)
		if err != nil {
			panic(err)
		}
		err = instance.LoadAppContext(file)

		if instance.Config.Mode != "mock" {

			if err != nil {
				panic(err)
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(instance.Config.Database.Timeout)*time.Second)
			defer cancel()

			// 1. Setup mongodb connector
			conn, err := mongoconnector.NewMongodbConnector(ctx, instance.Config.Database.ConnectionString)
			if err != nil {
				panic(err)
			}

			err = conn.Connect(ctx)
			if err != nil {
				panic(err)
			}

			instance.DataConnector = conn

			// 2. Setup mongodb persistence
			instance.Persistence = mongodbpersistence.NewMongoPersistence(conn.Context(), instance.Config.Database.Name)
		}


	})

	return instance
}
