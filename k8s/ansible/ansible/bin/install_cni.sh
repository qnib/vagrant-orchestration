#!/bin/bash

export KUBECONFIG="/home/vagrant/.kube/config"
kubectl apply -f https://github.com/flannel-io/flannel/releases/latest/download/kube-flannel.yml
