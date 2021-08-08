# serification

_serializable + specification = serification_

## Problem

* Serializing Specifications to JSON and SQL, unserializing JSON to Specifications

## Example

Converting:

```go
package main

var spec = PostCreatedBetween("2020-01-01", "2020-01-31").And(PostHasTags("tag_A", "tag_B").Or(PostHasTags("tag_A", "tag_C")))
```

into:

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

and vice versa.

## License

MIT
