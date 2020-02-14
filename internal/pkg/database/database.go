package database

import (
	"github.com/Infinity-OJ/Server/internal/pkg/models/users"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Options is  configuration of database
type Options struct {
	URL   string `yaml:"url"`
	Debug bool
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("db", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal db option error")
	}

	logger.Info("load database options success", zap.String("url", o.URL))

	return o, err
}

// Init 初始化数据库
func New(o *Options) (*gorm.DB, error) {
	var err error
	db, err := gorm.Open("postgres", "host=192.168.31.233 user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		return nil, errors.Wrap(err, "gorm open database connection error")
	}

	if o.Debug {
		db = db.Debug()
	}

	db.AutoMigrate(&users.Account{})

	return db, nil
}

var ProviderSet = wire.NewSet(New, NewOptions)