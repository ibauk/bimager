package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

const PROGRAMVERSION = "BImager v1.0 Copyright (c) 2025 Bob Stammers"

var dir = flag.String("smf", "", "Path to root of ScoreMaster folder")
var fld = flag.String("fld", "", "Path to image folder itself")
var ptn = flag.String("bre", "\\d+", "RE to capture bonusid")
var img = flag.String("img", "jpg|png", "RE to capture filetype")
var re *regexp.Regexp

func main() {

	fmt.Println(PROGRAMVERSION)

	flag.Parse()
	if *dir == "" && *fld == "" {
		fmt.Println("You must specify the ScoreMaster root folder using -smf or the image folder itself using -fld")
		return
	}
	fldr := ""
	if *fld != "" {
		fldr = *fld
	} else {
		fldr = filepath.Join(*dir, "sm", "images", "bonuses")
	}

	fmt.Printf("Changing to %v\n", fldr)

	err := os.Chdir(fldr)
	if err != nil {
		panic(err)
	}
	re = regexp.MustCompile(fmt.Sprintf(`(?i)(%v).*(%v)$`, *ptn, *img))

	filepath.WalkDir(".", walk)

}

func walk(s string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !d.IsDir() {
		subs := re.FindAllStringSubmatch(s, 1)
		if len(subs) > 0 {
			newname := subs[0][1] + "." + subs[0][2]
			if newname != s {
				err := os.Rename(s, newname)
				if err != nil {
					panic(err)
				}
				fmt.Printf("%v ==> %v\n", s, newname)
			}
		}
	}
	return nil
}
