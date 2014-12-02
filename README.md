argonauts
=========

JSON command line toolset

### `csvtojson`

Converts from CSV to lines of JSON.

```
cat data.csv | csvtojson
```

### `jsontrans`

Transforms JSON objects.

```
cat source.json | jsontrans {flags}
```

```
Usage of jsontrans:
  -array=0: group objects into arrays of this size
  -lower=false: make fields lowercase
```