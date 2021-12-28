
# Getting start

* Download

```bash
wget https://github.com/drummerzzz/brtools/raw/master/bin/brtools

```

* Add execute permission

```bash
chmod +x brtools
```

* Add to path

```bash

# Move to local bin
sudo mv brtools /usr/local/bin
```

* Or you can add another path of your choice

```bash
# Example:
export PATH=$PATH:"$HOME/Downloads/brtools"

```

* And per last close current terminal and open a new and try

```bash
brtools
```

# Quick starter

```bash
brtools --help
```

* Generate a cpf with mask

```bash
brtools cpf
# 865.353.851-83
```

* Generate a cpf without mask

```bash
brtools cpf --clean
# 25285688353
```

Or

```bash
brtools cpf -c
# 25285688353
```

* Validate a cpf

```bash
brtools cpf --validate 25285688353
# true
```

Or

```bash
brtools cpf -v 865.353.851-83
# true
```
