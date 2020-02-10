/**
 * @auther:  zu1k
 * @date:    2020/2/10
 */
package main

import "github.com/syndtr/goleveldb/leveldb"

var db *leveldb.DB

func init() {
	dbtmp, err := leveldb.OpenFile("data", nil)
	if err != nil {
		panic(err.Error())
	}
	db = dbtmp
}

func get(key string) (value string, err error) {
	data, err := db.Get([]byte(key), nil)
	if err == nil {
		value = string(data)
		return
	}
	return
}

func set(key, value string) (err error) {
	err = db.Put([]byte(key), []byte(value), nil)
	return
}
