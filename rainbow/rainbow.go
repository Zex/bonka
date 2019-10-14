package rainbow

import (
  "crypto/sha256"
  "fmt"
  "github.com/zex/bonka/store/lvl"
//  "bytes"
)

type Rainbow struct {
  rbt *lvl.Lvl
  start_point string
}

func NewRainbow() *Rainbow {
  return &Rainbow {
    rbt: lvl.NewLvl(".rainbow/lvl"),
  }
}

func (self *Rainbow) nextSrc(prev [32]byte) []byte {
  var ret []byte

  for _, xi := range fmt.Sprintf("%x", prev) {
    //if xi < '0' || xi > '9' {
    //  continue
    //}
    ret = append(ret, byte(xi))
    if len(string(ret)) == len(self.start_point) {
      break
    }
  }
  return ret
}

func (self *Rainbow) Start(s string) {
  self.start_point = s

  for {
    r := sha256.Sum256([]byte(s))
    x := fmt.Sprintf("%x", r)
    fmt.Println(x, s)

    if found, _ := self.rbt.Has(r[:]); found {
      break
    }

    self.rbt.Put(r[:], []byte(s))
    s = string(self.nextSrc(r))
  }
}

func (self *Rainbow) Stop() {
  if self.rbt != nil {
    self.rbt.Close()
  }
}
