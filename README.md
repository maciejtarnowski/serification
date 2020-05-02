# serification

_serializable + specification = serification_

## Problem

* Zapis i odczyt struktury przypominającej Specification Pattern do JSON-a

## Rozwiązanie

Kod zamieniający:

```go
package main

var spec = PostCreatedBetween("2020-01-01", "2020-01-31").And(PostHasTags("tag_A", "tag_B").Or(PostHasTags("tag_A", "tag_C")))
```

na:

```json
{
  "type": "and",
  "left": {
    "type": "post_created_between",
    "lower": "2020-01-01",
    "upper": "2020-01-31"
  },
  "right": {
    "type": "or",
    "left": {
      "type": "post_has_tags",
      "tags": ["tag_A", "tag_B"]
    },
    "right": {
      "type": "post_has_tags",
      "tags": ["tag_A", "tag_C"]
    }
  }
}
```

i z powrotem
