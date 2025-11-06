package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func usage(fatal ...bool) {
	log.Println("DIR test")
	log.Println("Usage: ./dir [case] {rid {dttm}}")
	if len(fatal) > 0 && !fatal[0] {
		os.Exit(0)
	}
	os.Exit(1)
}

func lookup(
	pth, sfx string,
	elms ...string,
) (
	srch, fnd string,
	err error,
) {
	spr := ""
	for _, elm := range elms {
		if elm != "" {
			srch = fmt.Sprintf("%v%v%v", srch, spr, elm)
			spr = "-"
		}
	}

	lst, err := os.ReadDir(pth)
	if err != nil {
		return
	}

	idx := -1
	errs := make([]error, 0)
	for j, itm := range lst {
		if strings.HasPrefix(itm.Name(), srch) && (sfx == "" || strings.HasSuffix(itm.Name(), sfx)) {
			if idx < 0 {
				idx = j
			} else {
				errs = append(errs, fmt.Errorf("%v", itm.Name()))
			}
		}
	}

	if len(errs) > 0 {
		err = fmt.Errorf("'%v' matches multiple files:", srch)
		if idx >= 0 {
			err = fmt.Errorf("%v\n %v", err, lst[idx].Name())
		}
		for _, e := range errs {
			err = fmt.Errorf("%v\n %v", err, e)
		}
		return
	}

	if idx < 0 {
		err = fmt.Errorf("'%v' matches nothing", srch)
		return
	}
	fnd = fmt.Sprintf("%v%c%v", pth, os.PathSeparator, lst[idx].Name())
	return
}

func main() {
	cid := ""
	rid := ""
	dtm := ""

	switch len(os.Args) {
	case 4:
		dtm = os.Args[3]
		fallthrough
	case 3:
		rid = os.Args[2]
		fallthrough
	case 2:
		switch os.Args[1] {
		case "-version":
			log.Println("DIR test")
			os.Exit(0)
		case "-about":
			usage(false)
		}
		cid = os.Args[1]
	default:
		usage()
	}

	srch, found, err := lookup("logs", "stat.test", cid, rid, dtm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Search: '%v' -> found: '%v'\n", srch, found)
}
