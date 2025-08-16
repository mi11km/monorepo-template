// entrypoint for the sample application
package main

import (
	"context"

	"github.com/mi11km/monorepo-template/go/apps/sample/internal/api"
)

func main() {
	api.Run(context.Background())
}
