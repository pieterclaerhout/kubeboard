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

func NewKubeBoard() KubeBoard {
	return KubeBoard{}
}

func (kb KubeBoard) Start() error {
	go kb.startProxy()
	kb.startWebUI()
	return nil
}

func (kb KubeBoard) Stop() {
	log.Info("Closing down")
	if kb.kubectlCmd != nil && kb.kubectlCmd.Process != nil {
		kb.kubectlCmd.Process.Kill()
	}
}

func (kb KubeBoard) startProxy() {

	log.Info("Starting kubectl proxy")

	// TODO: kill the previous instance if any

	file, _ := os.Create("/Users/pclaerhout/Desktop/kubectl.log")
	defer file.Close()

	os.Setenv("PATH", os.Getenv("PATH")+":/usr/local/bin")
	file.WriteString(os.Getenv("PATH") + "\n")

	kb.kubectlCmd = exec.Command("/usr/local/bin/kubectl", "proxy", "--port=8001")
	kb.kubectlCmd.Stdout = file
	kb.kubectlCmd.Stderr = file
	kb.kubectlCmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := kb.kubectlCmd.Start(); err != nil {
		file.WriteString(err.Error() + "\n")
	}

}

func (kb KubeBoard) startWebUI() {
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

	// return webview.Open(
	// 	"kubeboard",
	// 	"http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/login",
	// 	1280,
	// 	960,
	// 	true,
	// )
}
