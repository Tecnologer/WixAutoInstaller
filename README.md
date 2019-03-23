# Wix Automation Installer
Actualiza la version en los archivos para crear un instalador usando Wix.

Archivos afectados:
- Version y ProductGUID en el archivo `Product.wxs`
- Version (`AssemblyVersion` & `AssemblyFileVersion`) en `AssemblyInfo.cs`

# Depencies

- `go get -u github.com/tecnologer/uuid`

# Como compilar
- Clonar el repositorio<br>
  > `go get github.com/tecnologer/WixAutoInstaller`
  
  Esto descargara el codigo en la siguiente ruta:
  > `$GOPATH/src/github.com/tecnologer/WixAutoInstaller`
- Accedemos a la carpeta con el codigo
  > `cd $GOPATH/src/github.com/tecnologer/WixAutoInstaller`
- Compilamos usando
  > `go build`
- Este comando generara un executable llamado:
  - `WixAutoinstaller.exe` en windows
  - `WixAutoinstaller` en sistemas operativos unix.

**Nota:** `$GOPATH` es una [variable de entorno][2] que apunta a la ruta donde estan los repositorios de Go. En Windows se usa con `%`; por ejemplo: `cd %GOPATH%\src\github.com\tecnologer\WixAutoInstaller`.

# Como usarlo
 Una vez compilado, se usara el binario `WixAutoinstaller[.exe]`. 

 Recibe tres parametros:
- `-p` ruta del archivo `.wix`, en el cual se actualizara el numero de version y ID unico del instalador (GUID).
- `-a` ruta del archivo `AssemblyInfo.cs`, en el cual se actualiza solo la version en dos variables: `AssemblyVersion` & `AssemblyFileVersion`.
- `-v` el valor de la version. Ejemplo: 1.0.1.001

## Ejemplo
### Windows
```bat
WixAutoinstaller.exe -p C:\Users\Tecnologer\source\repos\FontInstallerInstaller\AppInstaller\Product.wxs -a C:\Users\Tecnologer\source\repos\FontInstaller\FontInstaller\Properties\AssemblyInfo.cs -v 0.0.1.0010
```

[1]: https://golang.org/dl/
[2]: https://medium.com/@01luisrene/como-agregar-variables-de-entorno-s-o-windows-10-e7f38851f11f