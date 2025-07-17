// cmd/intellisensearchitectframeworknext/main.go
package main

import (
"flag"
"log"
"os"

"intellisensearchitectframeworknext/internal/intellisensearchitectframeworknext"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := intellisensearchitectframeworknext.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
