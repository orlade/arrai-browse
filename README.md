# Arr.ai Browse

Browse is a simple web server that evaluates arr.ai scripts (a poor man's CGI).

## Usage

To install `browse` as an executable:

```bash
make install
```

`cd` to the directory you want to browse, and run:

```bash
arrai-browse
```

Navigate to http://localhost:4000 in the browser, and browse the `.arrai` files in the file system subtree.

To pass args to the script (i.e. populate `//os.args`), pass one or more `args=` query parameters.
