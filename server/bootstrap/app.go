package bootstrap

import (
	"github/tekeoglan/discord-clone/mongo"
	"github/tekeoglan/discord-clone/redis"
)

type Application struct {
	Env   *Env
	Mongo mongo.Client
	Redis redis.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	app.Redis = NewCache(app.Env)

	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}

func (app *Application) CloseCacheConnection() {
	KillCacheClient(app.Env, app.Redis)
}
