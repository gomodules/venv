# venv

```bash
$ mkdir /tmp/virtual-secrets

$ echo "x1" > /tmp/virtual-secrets/var1
$ export VAR1=vs:///tmp/virtual-secrets/var1

$ echo "x2" > /tmp/virtual-secrets/var2
$ export VAR2=vs:///tmp/virtual-secrets/var2

$ go run main.go sh -c "echo VAR1=\$VAR1 VAR2=\$VAR2"
VAR1=x1 VAR2=x2
```
