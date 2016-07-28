Dictionary Dash
===============

Task
----

Given two words (start and end) and the dictionary, find the length of the
shortest transformation

sequence from start to end, such that:

 + Only one letter can be changed at a time
 + Each intermediate word must exist in the given dictionary
 + At each step, exactly one character is replaced with another character

### For example:

```
start = "hit"
end = "cog"
dictionary = ["hit","dot","dog","cog","hot","log"]
```

As one of the shortest transformations is
`"hit" -> "hot" -> "dot" -> "dog" -> cog"`, return its length 4.
