package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

// future me - a list of goals
type Fme struct {
	goals []string
}

// create a future me - goals list
func mk_Fme() Fme {
	return Fme{goals: make([]string, 0)}
}

// add a goal
func (fp *Fme) add(g string) {
	fp.goals = append(fp.goals, g)
}

func headerfn(s string) header_renderer {
	return func(f Fme) string {
		var horiz = strings.Repeat("-", len(s))
		return fmt.Sprintf("%s\n%s\n%s", horiz, s, horiz)
	}
}
func items(f Fme) string {
	var ss = f.goals
	var br strings.Builder
	for _, line := range ss {
		fmt.Fprintf(&br, "* %s\n", line)
	}
	return br.String()
}

type header_renderer func(Fme) string
type items_renderer func(Fme) string

func (f Fme) render(hr header_renderer, ir items_renderer) string {
	return fmt.Sprintf("%s\n%s", hr(f), ir(f))
}

func (f Fme) String() string {
	return f.render(headerfn("Hello, these are my best goals"), items)
}

// wtite fm to a  file with given name
func (f Fme) toFile(file_name string) {
	file, err := os.OpenFile(file_name, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, line := range f.goals {
		if _, err := file.WriteString(line + "\n"); err != nil {
			log.Fatal(err)
		}
	}
}

func (f *Fme) fromFile(file_name string) {
	file, err := os.Open(file_name)
	if err != nil {
		log.Printf("Error while opening file: %v", file_name)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		f.add(text)

		if err := scanner.Err(); err != nil {
		}
	}
}

const prompt = "Your goal (q quits): "
const default_filename = "betterme.txt"
const envname = "FMEDIR"

func loop(fme *Fme) {
	s := bufio.NewScanner(os.Stdin)
	fmt.Println(prompt)
	for s.Scan() {
		line := s.Text()
		if line == "q" {
			break
		}
		fme.add(line)
	}
}

// Return the output file path located in directory pointed to
// by envname environment variable or in current working directory
func filename(default_filename string) string {
	if dirname, ok := os.LookupEnv(envname); ok {
		return path.Join(dirname, default_filename)
	} else {
		log.Printf("Could not find %s env variable, will use %s", envname, default_filename)
		return default_filename
	}

}
func futureme() {
	fme := mk_Fme()
	file := filename(default_filename)
	fme.fromFile(file)
	fmt.Printf("From file : \n%v\n", fme)
	loop(&fme)

	fmt.Printf("With appended lines : \n%v\n", fme)
	fme.toFile(file)
	fmt.Printf("Data was  written to %v\n", file)

}
