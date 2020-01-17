1ipackage cmd\n\n// This file is auto-generated!\n// Do not edit.\n\nimport (\n    "time"\n\n    "github.com/spf13/cast"\n)\n
s|\([^,]*\),\([^,]*\),\([^,]*\)|func decode\1(in, out interface{}) {\n    *out.(*\2) = cast.To\1(in)\n}\n\n// Get\1 reads \3 from the configuration file\nfunc Get\1(key string, val *\2) {\n    getConfig(key, val, decode\1)\n}\n|g
