package common

import (
	"context"
	firebase "firebase.google.com/go/v4"
)

func NewFirebaseApp() *firebase.App {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	return app
}
