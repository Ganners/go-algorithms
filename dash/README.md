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

Solution
--------

I first experimented with simply shuffling a letter at a time to try and match
the output. I realised quickly that `hot` cannot change to `dot` when the goal
is `cog` using this method.

I then decided that it might be best solved as a graph problem, which is what I
have developed.

The algorithmic complexity in total is O((M * 2N) + N^2) where M is the number
of letters and N is the number of items in the dictionary. There is
reusability, however, on the first set of parentheses as the graph can be
reused to perform another search.
