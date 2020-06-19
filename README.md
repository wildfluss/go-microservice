# microservice

Get

```

brew install skaffold
```

- Get k3d v1.7.0 [here](hhttps://github.com/rancher/k3d/releases/tag/v1.7.0)

  Get v1.7.x, in June 2020 v3.0.0 [does not support](https://github.com/rancher/k3d/blob/master/docs/faq/v1vsv3-comparison.md) `--enable-registry`

- kustomize

  ```
  brew install kustomize
  ```

- skaffold

  TBW

Run

```bash
k3d c --enable-registry --publish 8080:8080@k3d-k3s-default-worker-0 --workers 1
echo '127.0.0.1 registry.local' | sudo tee -a /etc/hosts
export KUBECONFIG="$(k3d get-kubeconfig)"
skaffold dev --default-repo registry.local:5000
kubectl port-forward microservice-deployment-6c79bf66cc-sj8gw 8080:8080
curl http://localhost:8080
You've hit microservice-deployment-6c79bf66cc-sj8gw
```
