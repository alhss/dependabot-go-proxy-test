package main

import (
    "k8s.io/client-go/kubernetes"
    "github.com/alhss/private-go-utils"
)

func main() {
    _ = kubernetes.NewForConfig
}
