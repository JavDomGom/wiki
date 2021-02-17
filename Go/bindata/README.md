## Bindata

Para añadir ficheros estáticos al binario, como `.html`, `.js`, etc, de modo que no haga falta que existan dichos fiechos en local:

1. Instala go-bindata:

```bash
~$ sudo apt install go-bindata
```

2. Prueba a crear esta estructura de archivos y directorios:

```bash
~$ tree
.
├── files
│   └── index.html
└── main.go

1 directory, 2 files
```

Archivo **index.html**:

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <title>Hello world!</title>
    </head>

    <body>
        <h1>Hi, {{.Name}}</h1>
    </body>
</html>
```

Archivo **main.go**:

```go
package main

import (
    "html/template"
    "log"
    "net/http"
)

var tmpl *template.Template

func init() {
    data, err := Asset("files/index.html")
    if err != nil {
        log.Fatal(err)
    }
    tmpl = template.Must(template.New("tmpl").Parse(string(data)))
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        tmpl.Execute(w, map[string]string{"Name": "Javier"})
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

3. Ejecuta estos comando en el siguiente orden:

```bash
~$ go-bindata files
```

Este te generará un archivo llamado bindata.go.

```bash
~$ go build
```

Este te generará el binario.

```bash
~$ ./test
```

Este lanzará el servidor HTTP sin necesidad de que existan en local los archivos que tengas dentro de la carpeta `files`.
