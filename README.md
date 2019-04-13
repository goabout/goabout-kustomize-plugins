# Kustomize plugins


## sopsdotenv

Set secret key-values from a sops-encrypted dotenv file.


### Build

    go build -buildmode plugin -o ~/.config/kustomize/plugins/kvSources/sopsdotenv.so sopsdotenv/main.go


## Usage

    secretGenerator:
    - name: example
      kvSources:
      - name: sopsdotenv
        pluginType: go
        args:
        - somefile.sops.env


## sopsfiles

Create secret entries from sops-encrypted files.


### Build

    go build -buildmode plugin -o ~/.config/kustomize/plugins/kvSources/sopsfiles.so sopsfiles/main.go


## Usage

    secretGenerator:
    - name: example
      kvSources:
      - name: sopsfiles
        pluginType: go
        args:
        - somefile.txt=somefile.sops.txt
