package fetcher

import (
	cl "tmbot/client"
)

type fetcher struct {
}

func NewFetcher(client cl.Client) fetcher {
	return fetcher{}
}

func (f *fetcher) Fetch() Channel {
	in := make(chan []string)

	in <- []string{"s"}

	return in
}
