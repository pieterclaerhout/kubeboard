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

func (kb KubeBoard) startProxy() {

	log.Info("Starting kubectl proxy")

	kb.kubectlCmd = exec.Command("kubectl", "proxy", "--port=8001")
	kb.kubectlCmd.Stdout = os.Stdout
	kb.kubectlCmd.Stderr = os.Stderr
	if err := kb.kubectlCmd.Start(); err != nil {
		log.Fatalf("Failed to start kubectl proxy: %v", err)
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
