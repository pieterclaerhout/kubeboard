package kubeboard

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/pieterclaerhout/go-log"
	"github.com/zserge/webview"
)

type KubeBoard struct {
	kubectlCmd *exec.Cmd
}

func NewKubeBoard() *KubeBoard {
	return &KubeBoard{}
}

func (kb *KubeBoard) Start() error {
	go kb.startProxy()
	kb.startWebUI()
	return nil
}

func (kb *KubeBoard) Stop() {
	log.Info("Closing down")
	log.InfoDump(kb.kubectlCmd, "cmd")
	if kb.kubectlCmd != nil && kb.kubectlCmd.Process != nil {
		log.Info("Killing subprocess")
		kb.kubectlCmd.Process.Kill()
	}
}

func (kb *KubeBoard) startProxy() {

	log.Info("Starting kubectl proxy")

	os.Setenv("PATH", os.Getenv("PATH")+":/usr/local/bin")
	log.Info(os.Getenv("PATH"))

	kb.kubectlCmd = exec.Command("/usr/local/bin/kubectl", "proxy", "--port=8001")
	kb.kubectlCmd.Stdout = log.Stdout
	kb.kubectlCmd.Stderr = log.Stderr
	kb.kubectlCmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := kb.kubectlCmd.Start(); err != nil {
		log.Error(err)
	}

}

func (kb *KubeBoard) startWebUI() {
	log.Info("Starting web UI")

	wv := webview.New(webview.Settings{
		Title:     "kubeboard",
		URL:       "http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/login",
		Width:     1280,
		Height:    960,
		Resizable: true,
	})
	defer func() {
		kb.Stop()
		wv.Exit()
	}()
	wv.Run()

}
