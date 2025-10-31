package fetcher

import (
	"tmbot/client"
)

type IncomeUpdatesPool = <-chan []client.Update
