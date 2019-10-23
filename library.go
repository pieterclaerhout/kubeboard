package kubeboard

import (
	"os"
	"os/exec"

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
	return kb.startWebUI()
}

func (kb KubeBoard) Stop() {
	kb.kubectlCmd.Process.Kill()

}

func (kb KubeBoard) startProxy() {

	log.Info("Starting kubectl proxy")

	file, _ := os.Create("/Users/pclaerhout/Desktop/kubectl.log")
	defer file.Close()

	os.Setenv("PATH", os.Getenv("PATH")+":/usr/local/bin")
	file.WriteString(os.Getenv("PATH") + "\n")

	kb.kubectlCmd = exec.Command("/usr/local/bin/kubectl", "proxy", "--port=8001")
	kb.kubectlCmd.Stdout = file
	kb.kubectlCmd.Stderr = file
	if err := kb.kubectlCmd.Start(); err != nil {
		file.WriteString(err.Error() + "\n")
	}

}

func (kb KubeBoard) startWebUI() error {
	log.Info("Starting web UI")
	return webview.Open(
		"kubeboard",
		"http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/login",
		1280,
		960,
		true,
	)
}
