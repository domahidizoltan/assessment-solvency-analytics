# Homework

The following two tasks are designed to give us a comprehensive image about your programming skills in different domains. We expect you to solve these tasks according to your current knowledge and experience. If you like the tasks or you want to further demonstate your involvement, we are happy to see the solution for the optional tasks.

You can use any programming language, but we prefer python3 and Go. In case you choose another programming language, you must provide a dockerized execution environment, and a detailed description of how to run your solution.

Feel free to use the internet, but please do not seek any assistance from a third party. If you have difficulties interpreting the tasks, do not hesitate to contact us.

If you find the solution for any of the tasks on the internet please provide us the source and the way you found it, we will give you extra credit for that.

Our preferences:
- 100% statement coverage
- Baby steps
- Dockerized test and development environment
- Clean code

*Please consider our preferences only if you are familiar with these methodologies, in case not, we are not expecting you to do so.*

## Algorithmic task
A list of base ten integers `haystack`, and a list of digits `needle` are given. Find the sequence of the lowest possible indexes of the `haystack`, containing the digits of the `needle` in order. One index of the `haystack` should only occur for a digit in the `needle` once. The n-th element of the result should be the index of the `haystack`, which contains the n-th digit of the `needle`.

- `haystack[result[n]]` should contain the digit `needle[n]`
- `result[n]` should be smaller than `result[n+1]`
- `result[n]` should be the smallest possible

```python
>>> haystack = [662, 154063, 38, 1, 946773, 7877907760054, 332, 76826670, 7653639346039, 90593, 2567954972664]
>>> needle = [6, 5, 4]

>>> find_first_occurance(haystack, needle)
[0, 1, 4]
```
```python
>>> haystack = [5, 3, 5]
>>> needle = [3, 5]

>>> find_first_occurance(haystack, needle)
[1, 2]
```
### Optional algorithmic task 1
Find the sequence with the lowest possible indexes in the `haystack` where the distance between the first and the last index is lower or equal to the `maximum_distance`.

- `haystack[result[n]]` should contain the digit `needle[n]`
- `result[n]` should be smaller than `result[n+1]`
- `result[n] - result[0]` should be equal to or smaller than `maximum_distance`
- `result[n]` should be the smallest possible

```python
>>> haystack = [662, 154063, 38, 1, 946773, 7877907760054, 332, 76826670, 7653639346039, 90593, 2567954972664]
>>> needle = [6, 5, 4]
>>> maximum_distance = 3

>>> find_first_occurance_with_max_distance_limit(haystack, needle, maximum_distance)
[7, 8, 10]
```

### Optional algorithmic task 2
Find the sequence with the lowest possible indexes in the `haystack` where the distance between the first and the last index is minimal.

- `haystack[result[n]]` should contain the digit `needle[n]`
- `result[n]` should be smaller than `result[n+1]`
- `result[n] - result[0]` should be as small as possible
- `result[n]` should be the smallest possible

```python
>>> haystack = [662, 154063, 38, 1, 946773, 7877907760054, 332, 76826670, 7653639346039, 90593, 2567954972664]
>>> needle = [6, 5, 4]

>>> find_first_occurance_with_minimum_possible_distance(haystack, needle)
[8, 9, 10]
```

## Business task
The task is to validate a json document against a schema. The document and the schema are given as a single json.

```json
{
    "schema": {},
    "document": {}
}
```

Validate that the document contains exactly the keys which are in the schema. Allow keys to be omitted from the document when `required: false` is provided for the corresponding schema element.

***Example valid 1***
```json
{
    "schema": {
        "key1": {
            "required": true
        },
        "key2": {
            "required": false
        }
    },
    "document": {
        "key1": 1,
        "key2": 2
    }
}
```

***Example valid 2 (Optional field is missing)***
```json
{
    "schema": {
        "key1": {
            "required": true
        },
        "key2": {
            "required": false
        }
    },
    "document": {
        "key1": 1
    }
}
```

***Example invalid (`document.key1` is missing)***
```json
{
    "schema": {
        "key1": {
            "required": true
        },
        "key2": {
            "required": false
        }
    },
    "document": {
        "key2": 2
    }
}
```

***Example invalid (`document.key3` is unexpected)***
```json
{
    "schema": {
        "key1": {
            "required": true
        },
        "key2": {
            "required": false
        }
    },
    "document": {
        "key1": 1,
        "key2": 2,
        "key3": 3
    }
}
```

### Optional business task 1
In addition to the above requirements, also validate the type of the fields of the document based on the `type` field in the schema. The possible types are `integer`, `string` and `boolean`.


***Example valid***
```json
{
    "schema": {
        "key1": {
            "type": "string",
            "required": true
        },
        "key2": {
            "type": "boolean",
            "required": true
        },
        "key3": {
            "type": "integer",
            "required": true
        }
    },
    "document": {
        "key1": "value1",
        "key2": true,
        "key3": 1
    }
}
```

***Example invalid (`document.key1` is not a string)***
```json
{
    "schema": {
        "key1": {
            "type": "string",
            "required": true
        },
        "key2": {
            "type": "boolean",
            "required": true
        },
        "key3": {
            "type": "integer",
            "required": true
        }
    },
    "document": {
        "key1": 1,
        "key2": true,
        "key3": 4
    }
}
```

### Optional business task 2
In addition to the above requirements, also validate the structure of the whole input json, including the schema, disallowing any unexpected or malformed keys or fields.

***Example invalid (Unexpected key `something_else`)***
```json
{
    "schema": {},
    "document": {},
    "something_else": {}
}
```

***Example invalid (Incomplete `schema.key1.type` is missing)***
```json
{
    "schema": {
        "key1": {
            "required": false
        }
    },
    "document": {}
}
```

***Example invalid (Invalid schema, unexpected key `schema.key1.something_else`)***
```json
{
    "schema": {
        "key1": {
            "type": "string",
            "required": false,
            "something_else": null
        }
    },
    "document": {}
}
```
