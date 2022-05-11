## seldon model infer

run inference on a model

### Synopsis

call a model with a given input and get a prediction

```
seldon model infer [flags]
```

### Options

```
  -f, --file-path string        inference payload file
  -h, --help                    help for infer
      --inference-host string   seldon inference host (default "0.0.0.0")
      --inference-mode string   inference mode rest or grpc (default "rest")
      --inference-port int      seldon scheduler port (default 9000)
  -i, --iterations int          inference iterations (default 1)
  -m, --model-name string       model name for inference
```

### Options inherited from parent commands

```
  -r, --show-request    show request
  -o, --show-response   show response (default true)
```

### SEE ALSO

* [seldon model](seldon_model.md)	 - manage models

###### Auto generated by spf13/cobra on 16-Apr-2022