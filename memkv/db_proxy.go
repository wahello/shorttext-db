package memkv

type DBProxy struct {
	dbs      map[int]MemDB
	dbCount  int
	maxCount uint64
}

func NewDBProxy(dbCount int, maxCount uint64) (KVClient, error) {
	instance := &DBProxy{}
	instance.dbs = make(map[int]MemDB)
	var err error
	var db MemDB
	instance.dbCount = dbCount
	instance.maxCount = maxCount

	for i := 1; i <= dbCount; i++ {
		db, err = Open(":memory:")
		if err != nil {
			return nil, err
		}
		instance.dbs[i] = db
	}
	return instance, err
}
func (d *DBProxy) Put(item *DbItem) (err error) {
	db := d.dbs[1]
	err = db.Put(item)
	return err
}

func (d *DBProxy) Get(key Key) (val *DbItem) {
	db := d.dbs[1]
	val = db.Get(key)
	return val
}

func (d *DBProxy) Delete(key Key) (err error) {
	db := d.dbs[1]
	return db.Delete(key)
}

//func (d *DBProxy)Range(start,stop Key)[]*DbItem{
//	db := d.dbs[1]
//	result :=db.Range(start,stop)
//	return result
//}

func (d *DBProxy) Close() {
	db := d.dbs[1]
	db.Close()
}

func (d *DBProxy) NewIterator(key Key) Iterator {
	db := d.dbs[1]
	stop := mvccEncode(key, 0)
	start := mvccEncode(key, lockVer)

	data := db.Scan(start, stop)
	iter := NewListIterator(data, false)
	return iter
}

func (d *DBProxy) scan(db MemDB, startKey Key, endKey Key) []*DbItem {
	var start, stop Key
	if len(startKey) > 0 {
		start = mvccEncode(startKey, lockVer)
	}
	if len(endKey) > 0 {
		stop = mvccEncode(endKey, lockVer)
	}
	data := db.Scan(start, stop)
	return data
}

func (d *DBProxy) NewScanIterator(startKey Key, endKey Key) Iterator {
	db := d.dbs[1]
	data := d.scan(db, startKey, endKey)
	iter := NewListIterator(data, false)
	return iter
}

func (d *DBProxy) NewDescendIterator(startKey Key, endKey Key) Iterator {
	db := d.dbs[1]
	data := d.scan(db, startKey, endKey)
	iter := NewListIterator(data, true)
	return iter
}

func (d *DBProxy) Write(batch *Batch) error {
	var err error
	for _, added := range batch.addedBuf {
		added.key = mvccEncode(added.key, added.ts)
		err = d.Put(added)
		if err != nil {
			return err
		}
	}
	for _, deleted := range batch.deletedBuf {
		deleted.key = mvccEncode(deleted.key, deleted.ts)
		err = d.Delete(deleted.key)
		if err != nil {
			return err
		}
	}
	return err
}
