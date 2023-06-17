package db

import (
	"bytes"
	"encoding/gob"
)

func (d *Database) Set(key string, value interface{}) error {
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	if err := enc.Encode(value); err != nil {
		return err
	}
	err := d.level.Put([]byte(d.getScopedKey(key)), data.Bytes(), nil)
	if err != nil {
		return err
	}
	return nil
}
