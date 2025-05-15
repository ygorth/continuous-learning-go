package main

import (
	"os"
	"time"

	clockface "github.com/BrazilITRemote/bir-oficinago-2024-t4/estudantes/ygorth/16-maths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
