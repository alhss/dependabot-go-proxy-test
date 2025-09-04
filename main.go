package main

import (
    "k8s.io/client-go/kubernetes"
    "github.com/alhss/dependabot-go-proxy-test"
)

func main() {
    _ = kubernetes.NewForConfig
}
