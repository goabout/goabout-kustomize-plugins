# Kustomize plugins

[Kustomize](https://github.com/kubernetes-sigs/kustomize) plugins previously used at GoAbout.

**Note: these plugins only work for some older kustomize commits around kustomize 2. For a plugin
for Kustomize 3+, see [kustomize-sopssecret-plugin](https://github.com/goabout/kustomize-sopssecret-plugin).**


## Running kustomize with plugins

    kustomize build --enable_alpha_goplugins_accept_panic_risk

## sopsdotenv

Set secret key-values from a sops-encrypted dotenv file.


### Build

    OUTDIR=${XDG_CONFIG_HOME:-$HOME/.config}/kustomize/plugins/kvSources
    go build -buildmode plugin -o $OUTDIR/sopsdotenv.so sopsdotenv/main.go


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

    OUTDIR=${XDG_CONFIG_HOME:-$HOME/.config}/kustomize/plugins/kvSources
    go build -buildmode plugin -o $OUTDIR/sopsfiles.so sopsfiles/main.go


## Usage

    secretGenerator:
    - name: example
      kvSources:
      - name: sopsfiles
        pluginType: go
        args:
        - somefile.txt=somefile.sops.txt


## Docker image

### Building

    docker build -t goabout/goabout-kustomize-plugins .

### Extracting binary and plugins

    cid=$(docker create goabout/k8s-infra-deploy)
    docker cp $cid:/usr/local/bin/kustomize .
    docker cp $cid:/root/.config/kustomize/plugin/kvSources ${XDG_CONFIG_HOME:-$HOME/.config}/kustomize/plugin
