package operator

import (
   "fmt"

    "github.com/bahadley/esp/log"
)


func Window(ingest chan string) {

   for {
      msg := <-ingest
      log.Logoutput(log.InfoPrefix, fmt.Sprintf("Windowed: %s", msg))
   }
}
