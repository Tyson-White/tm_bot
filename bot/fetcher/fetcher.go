package fetcher

import (
	cl "tmbot/client"
)

type fetcher struct {
}

func New(client cl.Client) fetcher {
	return fetcher{}
}

func (f *fetcher) Fetch() IncomeUpdatesPool {
	pool := make(chan []cl.Update)

	return pool
}
