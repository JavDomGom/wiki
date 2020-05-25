# Kubernetes

## Extraer la configuración de un deployment en formato YAML

Primero listar los deployments:
```bash
~$ kubectl get deployment
```

Luego extraer la información en formato YAML:
```bash
~$ kubectl get deployment deployment_name -o yaml | tee -a salida.yaml
```

## Copiar un fichero desde Master a un pod de un nodo

```bash
~$ kubectl cp archivo.txt namespace_name/pod_name:absolute_path
```

## Listar secretos

```bash
~$ kubectl get secret
```
