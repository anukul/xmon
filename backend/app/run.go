package app

import (
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

func (a *App) loop(wg *sync.WaitGroup) {
	for {
		select {
		case <-a.quitChannel:
			log.Debug().Msg("quitting")
			wg.Done()
		case <-time.After(time.Duration(a.refreshInterval) * time.Second):
			a.monitor.Refresh()
		}
	}
}

func (a *App) Run() error {
	a.monitor.Refresh()
	go a.server.Listen()
	var wg sync.WaitGroup
	wg.Add(1)
	go a.loop(&wg)
	wg.Wait()
	return nil
}
