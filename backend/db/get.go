package db

import (
	"bytes"
	"encoding/gob"
)

func (d *Database) Get(key string, result any) error {
	data, err := d.level.Get([]byte(d.getScopedKey(key)), nil)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(result); err != nil {
		return err
	}
	return nil
}
