# Wizard

Permite crear la funcionalidad de un modelo, como storage, sqlserver, postgresql, handler, etc.

## Instalación:
```bash
$ go get -u github.com/alexyslozada/wizard
```

## Compilación:
```bash
$ go build
```

## Crear la carpeta dist
```bash
$ mkdir dist
```

## Ejecución
```bash
$ ./wizard -model=nombremodelo -table=nombretabla -fields=campo1,campo2,campo3
```

Los flag son obligatorios:
* model: Es el nombre del modelo.
* table: Es el nombre de la tabla en el motor de BD.
* fields: Es un listado de los nombres de los campos de la tabla separado por comas sin espacios. El sistema utiliza esos nombres para el sql y le hace un CamelCase para los nombres de los campos del modelo.
