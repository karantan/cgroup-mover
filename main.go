package main

import (
	"cgroup-mover/logger"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

var log = logger.New("cgroup-mover")

const (
	CGROUP_PATH  = "/sys/fs/cgroup"
	CGROUP_PROCS = "cgroup.procs"
)

func main() {
	var cgroupOld, user, cgroupNew string
	flag.StringVar(&cgroupOld, "old", "", "Cgroup FROM all child processes will be moved")
	flag.StringVar(&user, "user", "", "User of which all processes will be moved")
	flag.StringVar(&cgroupNew, "new", "", "Cgroup TO which all child processes will be moved")
	flag.Parse()

	ticker := time.NewTicker(2 * time.Second)
	var pids []int
	for ; true; <-ticker.C {
		if cgroupOld != "" {
			pids = append(pids, findChildProcesses(path.Join(CGROUP_PATH, cgroupOld, CGROUP_PROCS))...)
		}
		if user != "" {
			pids = append(pids, findUserProcesses(user)...)
		}

		if err := addToCgroup(pids, path.Join(CGROUP_PATH, cgroupNew, CGROUP_PROCS)); err != nil {
			log.Errorf("Error trying to add pids to cgroup (%s)", cgroupNew)
		} else {
			for _, p := range pids {
				log.Infof("%d -> %s", p, cgroupNew)
			}
		}
	}
}

func findUserProcesses(user string) (childPids []int) {
	cmd := exec.Command("pgrep", "--uid", user)
	out, err := cmd.Output()
	if err != nil {
		log.Warnf("Problems finding processes for user %s", user)
		log.Error(err)
	}
	pidsRaw := strings.Split(strings.TrimSpace(string(out)), "\n")
	for _, p := range pidsRaw {
		i, _ := strconv.Atoi(p)
		childPids = append(childPids, i)
	}
	return
}

func findChildProcesses(cgroupProcsFile string) (childPids []int) {
	allPidsRaw, err := os.ReadFile(cgroupProcsFile)
	if err != nil {
		log.Fatalln("error opening file", cgroupProcsFile, err)
		return
	}
	// 1 pid is the master process which we don't want to move
	pidsRaw := strings.Split(strings.TrimSpace(string(allPidsRaw)), "\n")[1:]
	for _, p := range pidsRaw {
		i, _ := strconv.Atoi(p)
		childPids = append(childPids, i)
	}
	return childPids
}

func addToCgroup(pids []int, cgroupProcsFile string) error {
	f, err := os.OpenFile(cgroupProcsFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Error(err)
		return err
	}
	defer f.Close()

	for _, pid := range pids {
		if _, err := f.WriteString(fmt.Sprintf("%d\n", pid)); err != nil {
			log.Errorw("Couldn't write pid to the groupc.procs file", "pid", pid, "err", err.Error())
			return err
		}
	}
	return nil
}
