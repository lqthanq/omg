package migrate

import (
	"reflect"

	"github.com/lqthanq/omg-simpler/app"
	"github.com/lqthanq/omg-simpler/graph/model"
	"github.com/lunny/log"
)

func Migrate() {
	var models = []interface{}{
		model.User{},
	}

	for _, m := range models {
		log.Info("Migrate:", reflect.TypeOf(m))
		err := app.DB.AutoMigrate(m)
		if err != nil {
			log.Info(err)
		}
	}
}
