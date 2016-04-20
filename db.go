// Copyright ©2016 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var (
	bucketMon = []byte("data-mon")
)

func openDB(name string) *bolt.DB {
	db, err := bolt.Open(name, 0644, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketMon)
		return err
	})

	return db
}

func (srv *server) monitor() {
	hfreq := time.NewTicker(5 * time.Second)
	defer hfreq.Stop()
	lfreq := time.NewTicker(5 * time.Minute)
	defer lfreq.Stop()

	srv.publishData()
	srv.storeMonData()
	for {
		select {
		case <-hfreq.C:
			srv.publishData()
		case <-lfreq.C:
			srv.publishData()
			srv.storeMonData()
		}
	}
}

func (srv *server) storeMonData() {
	db := openDB("mon.db")
	defer db.Close()
	mon := srv.histos.rows[len(srv.histos.rows)-1]
	db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketMon)
		if err != nil {
			return err
		}
		var buf [monDataLen]byte
		mon.write(buf[:])
		err = bucket.Put(buf[:8], buf[8:])
		return err
	})
}
