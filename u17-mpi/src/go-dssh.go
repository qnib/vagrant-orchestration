package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	var size, rank int
	pEnv := []string{}
	for _, e := range os.Environ() {
		kv := strings.Split(e, "=")
		k := kv[0]
		v := kv[1]
		switch k {
		case "OMPI_COMM_WORLD_SIZE":
			size, _ = strconv.Atoi(v)
		case "OMPI_COMM_WORLD_RANK":
			rank, _ = strconv.Atoi(v)
		default:
			if strings.HasPrefix(e, "OMPI_") {
				pEnv = append(pEnv, e)
			}
		}
	}
	//fmt.Printf("Env: %v\n", pEnv)
	fmt.Printf("Rank: %d/%d\n", rank, size-1)
	cmd := exec.Command(os.Args[1], "")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Run()
}
