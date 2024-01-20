package main

import (
	"context"
	"log"
	"os"
	"os/exec"
)

func ExecuteThis(filepath string, env map[string]string) {
	cmd := exec.CommandContext(context.Background(), "sh", filepath)
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	for k, v := range env {
		cmd.Env = append(cmd.Env, k+"="+v)
	}

	out, err := cmd.Output()
	if err != nil {
		log.Println("error:", err.Error())
		return
	}
	log.Println(string(out))

}
