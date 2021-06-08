package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"time"
)

type Config struct {
	istagged bool
}

type Build struct {
	BuildDate   string
	BuildNumber string
}

type version struct {
	Number       string `json:Number`
	CommitId     string `json:CommitId`
	TagLink      string `json:TagLink`
	ReleaseNotes string `json:ReleaseNotes`
	BuildNumber  string `json:BuildNumber`
}

type project struct {
	Name           string  `json:Name`
	Description    string  `json:Description`
	Repository     string  `json:Repository`
	CurrentVersion string  `json:CurrentVersion`
	Details        version `json:Details`
}

func genbuildnum() string {
	year := strconv.Itoa(int(time.Now().Year() % 100))
	counter := strconv.Itoa(int((time.Now().Hour()*60 + time.Now().Minute()) / 2))
	var day string
	var month string

	if int(time.Now().Month()) < 10 {
		month = "0" + strconv.Itoa(int(time.Now().Month()))
	} else {
		month = strconv.Itoa(int(time.Now().Month()))
	}

	if int(time.Now().Day()) < 10 {
		day = "0" + strconv.Itoa(int(time.Now().Day()))
	} else {
		day = strconv.Itoa(int(time.Now().Day()))
	}

	return year + month + day + counter

}

func initProject() project {

	var ver version

	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Project Name: ")
	input.Scan()
	name := input.Text()

	fmt.Print("Project description: ")
	input.Scan()
	desc := input.Text()

	fmt.Print("Project repository: ")
	input.Scan()
	repo := input.Text()

	fmt.Print("Current or First version: ")
	input.Scan()
	number := input.Text()

	var hash string
	var commitId string

	for {
		fmt.Print("Log commit hash linked to this release (yes/no): ")
		input.Scan()
		if (input.Text() == "yes") || (input.Text() == "no") {
			hash = input.Text()
			if hash == "yes" {
				commit, err := exec.Command("git", "rev-list", "-1", "HEAD").Output()
				if err != nil {
					println("hello: ", err.Error())
				}
				commitId = string(commit[:])

			} else if hash == "no" {
				commit := "Not disclosed"
				commitId = commit

			}
			commitId = strings.TrimSuffix(commitId, "\n")
			break
		}

	}

	fmt.Print("Tag link: ")
	input.Scan()
	tag := input.Text()

	fmt.Print("Link to release notes: ")
	input.Scan()
	releasenotes := input.Text()

	buildnumber := genbuildnum()

	ver = version{number, commitId, tag, releasenotes, buildnumber}
	currentver := ver.Number + "-" + ver.BuildNumber
	return project{name, desc, repo, currentver, ver}

}

func newRelease() {
	var Project project

	if len(os.Args) < 3 {
		fmt.Println("Wrong!!")
		return
	}
	if _, err := os.Stat("./autover.json"); err == nil {
		fmt.Println("Found Auto Version config!")
		file, err := os.Open("autover.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		byteValue, _ := ioutil.ReadAll(file)
		json.Unmarshal(byteValue, &Project)

		switch os.Args[2] {
		case "patch":
			release := strings.Split(Project.Details.Number, ".")
			patch, _ := strconv.Atoi(release[2])
			release[2] = strconv.Itoa(patch + 1)
			bumped := strings.Join(release, ".")
			Project.CurrentVersion = bumped + "-" + genbuildnum()
			Project.Details.Number = bumped
			Project.Details.BuildNumber = genbuildnum()
			saveConfig(Project)
		case "minor":
			release := strings.Split(Project.Details.Number, ".")
			patch, _ := strconv.Atoi(release[1])
			release[1] = strconv.Itoa(patch + 1)
			bumped := strings.Join(release, ".")
			Project.CurrentVersion = bumped + "-" + genbuildnum()
			Project.Details.Number = bumped
			Project.Details.BuildNumber = genbuildnum()
			saveConfig(Project)
		case "major":
			release := strings.Split(Project.Details.Number, ".")
			patch, _ := strconv.Atoi(release[0])
			release[0] = strconv.Itoa(patch + 1)
			bumped := strings.Join(release, ".")
			Project.CurrentVersion = bumped + "-" + genbuildnum()
			Project.Details.Number = bumped
			Project.Details.BuildNumber = genbuildnum()
			saveConfig(Project)
		case "build":
			Project.CurrentVersion = Project.Details.Number + "-" + genbuildnum()
			Project.Details.BuildNumber = genbuildnum()
			saveConfig(Project)
		}

	} else if os.IsNotExist(err) {
		fmt.Println("No Config Found, please use 'autover init' to create you version file")
	}

}

func help() {
	fmt.Printf("Usage: %s [options] <csvFile>\nOptions:\n", os.Args[0])
}

func rollBack() {
	var Project project

	if len(os.Args) < 3 {
		fmt.Println("Wrong!!")
		return
	}
	if _, err := os.Stat("./autover.json"); err == nil {
		fmt.Println("Found Auto Version config!")
		file, err := os.Open("autover.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		byteValue, _ := ioutil.ReadAll(file)
		json.Unmarshal(byteValue, &Project)

		switch os.Args[2] {
		case "patch":
			release := strings.Split(Project.Details.Number, ".")
			patch, _ := strconv.Atoi(release[2])
			release[2] = strconv.Itoa(patch - 1)
			bumped := strings.Join(release, ".")
			Project.CurrentVersion = bumped + "-" + genbuildnum()
			Project.Details.Number = bumped
			Project.Details.BuildNumber = genbuildnum()
			saveConfig(Project)
		case "minor":
			release := strings.Split(Project.Details.Number, ".")
			patch, _ := strconv.Atoi(release[1])
			release[1] = strconv.Itoa(patch - 1)
			bumped := strings.Join(release, ".")
			Project.CurrentVersion = bumped + "-" + genbuildnum()
			Project.Details.Number = bumped
			Project.Details.BuildNumber = genbuildnum()
			saveConfig(Project)
		case "major":
			release := strings.Split(Project.Details.Number, ".")
			patch, _ := strconv.Atoi(release[0])
			release[0] = strconv.Itoa(patch - 1)
			bumped := strings.Join(release, ".")
			Project.CurrentVersion = bumped + "-" + genbuildnum()
			Project.Details.Number = bumped
			Project.Details.BuildNumber = genbuildnum()
			saveConfig(Project)
		}

	} else if os.IsNotExist(err) {
		fmt.Println("No Config Found, please use 'autover init' to create you version file")
	}

}

func saveConfig(Project project) {

	JSON, err := json.MarshalIndent(&Project, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = ioutil.WriteFile("autover.json", JSON, 0644)

}

func main() {
	switch os.Args[1] {
	case "help":
		help()

	case "init":
		var Project project
		fmt.Println("Looking for previous Auto Version config.")
		time.Sleep(1 * time.Second)
		fmt.Println("Looking for previous Auto Version config..")
		time.Sleep(1 * time.Second)
		fmt.Println("Looking for previous Auto Version config...")
		time.Sleep(1 * time.Second)
		if _, err := os.Stat("./autover.json"); err == nil {
			fmt.Println("Found Auto Version config!")
			file, err := os.Open("autover.json")
			if err != nil {
				fmt.Println(err)
				return
			}
			byteValue, _ := ioutil.ReadAll(file)
			json.Unmarshal(byteValue, &Project)
			mp, err := json.MarshalIndent(&Project, "", "  ")

			input := bufio.NewScanner(os.Stdin)

			fmt.Println(string(mp))
			fmt.Print("Overwrite This configuration (yes/no) (default no): ")
			input.Scan()
			switch input.Text() {
			case "yes":
				Project := initProject()
				saveConfig(Project)
			case "no":
				break
			default:
				break
			}

		} else if os.IsNotExist(err) {
			Project := initProject()

			saveConfig(Project)
		} else {
			Project := initProject()

			saveConfig(Project)
		}

	case "gen":
		genbuildnum()

	case "release":
		newRelease()

	case "rollback":
		rollBack()

	}

}
