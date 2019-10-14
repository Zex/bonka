package lvl

import (
  "fmt"
  "github.com/syndtr/goleveldb/leveldb"
  "github.com/syndtr/goleveldb/leveldb/util"
	"github.com/hyperledger/fabric/common/flogging"
)

var (
  logger = flogging.MustGetLogger("store.lvl")
)

type Lvl struct {
  db *leveldb.DB
}

func NewLvl(path string) *Lvl {
  db, err := leveldb.OpenFile(path, nil)
  if err != nil {
    logger.Fatal(err)
    return nil
  }
  return &Lvl{db: db,}
}

func (s *Lvl) Close() {
  s.db.Close()
}

func (s *Lvl) Get(key []byte) []byte {
  if key == nil {
    return nil
  }

  val, err := s.db.Get(key, nil)
  if err != nil {
    return nil
  }
  return val
}

func (s *Lvl) Put(key, val []byte) error {
  if key == nil {
    return fmt.Errorf("[store] invalid key")
  }
  return s.db.Put(key, val, nil)
}

func (s *Lvl) Has(key []byte) (bool, error) {
  return s.db.Has(key, nil)
}

func (s *Lvl) Delete(key []byte) error {
  if key == nil {
    return nil
  }
  return s.db.Delete(key, nil)
}

func (s *Lvl) ListKeys(prev []byte) ([][]byte, error){
  var ret [][]byte

  it := s.db.NewIterator(util.BytesPrefix(prev), nil)
  for it.Next() {
    key := make([]byte, len(it.Key()))
    copy(key, it.Key())
    ret = append(ret, key)
  }

  it.Release()
  if it.Error() != nil {
    return nil, it.Error()
  }

  return ret, nil
}
