package repositories

import (
	"github.com/google/wire"
)

var RepositoriesSet = wire.NewSet(
	NewFileRepository,
	wire.Bind(new(IFilesRepository), new(*FileRepository)),
	NewSubsRepository,
	wire.Bind(new(ISubsRepository), new(*SubsRepository)),
)
