package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func suspend(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("/usr/lib/systemd/systemd-sleep", "suspend")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP)
	go func() {
		<- sigs
		os.Exit(0)
	}()
}
func main() {

	http.HandleFunc("/suspend", suspend)     //设置访问的路由
	err := http.ListenAndServe(":9999", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
