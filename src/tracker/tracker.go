package tracker

import (
	"net/http"
	"os"
//	"time"
	"tracker/config"
	"tracker/db"
	"tracker/feed"
	"tracker/index"
	"tracker/log"
)

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		start := time.Now()

		targetMux.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr

		log.Infof(
//			"\"%s %s\" - %s [%v]",
			"\"%s %s\" - %s",
			r.Method,
			r.RequestURI,
			requesterIP,
//			time.Since(start),
		)
	})
}

func Run() {
	fname := "default.ini"
	if len(os.Args) > 1 {
		fname = os.Args[1]
	}
	log.SetLevel("info")
	cfg := new(config.Config)
	err := cfg.Load(fname)

	if err != nil {
		log.Fatalf("%s", err)
	}

	log.SetLevel(cfg.Log.Level)

	idx := index.New(&cfg.Index)

	idx.DB, err = db.NewPostgres(&cfg.DB)
	idx.Cfg_scrape = &cfg.Scrape

	if err != nil {
		log.Fatalf("%s", err)
	}
	err = idx.DB.Init()
	if err != nil {
		log.Fatalf("%s", err)
	}
	if cfg.Feeds.Enabled {
		fetcher := feed.NewFetcher(cfg.Feeds, idx.DB)
		go fetcher.Run(cfg.Feeds.Jobs)
	}
	addr := cfg.Index.Addr

	log.Infof("torrent.ano running on http://%s/", addr)

	//log.Infof( "%s" , cfg.Log.Level )

	if cfg.Log.Level == "debug" {
		err = http.ListenAndServe(addr, RequestLogger(idx))
	} else {
		err = http.ListenAndServe(addr, idx)
	}
	if err != nil {
		log.Fatalf("serve start err %s", err)
	}
}
