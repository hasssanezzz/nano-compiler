# Nano compiler

I want to create a compiler/interpreter for a lisp-like lang.
Will call it **nano lang**.

### Khawater

- lists with initial function name `(func p1 p2 p3)`
- non function lists `[a, b, c, d]`
- no commas allowed
- every pair of parenthesis have to return something
- function declaration return `nil`

#### variable declaration
```lisp
(set x 78)
```

#### function definition
```lisp
(func some_function [x] [
    (line1)
    (line2)
    (line3)
])

(func dec [x] 
    (set x
        (- x 1)))
    
(func odd [x] (== 1 (% x 2))) -> nil
```

#### string concatenation
```lisp
(str "hello" "world") -> `helloworld`
(+ 34 -24) -> 10
(str ((+ 34 -24))) -> `10`
```

#### print stuff
```lisp
(print [1 2 3]) -> [1 2 3]
(print "Hello" "world") -> ERROR print takes only one parameter
```

#### conditionals
```lisp
(if (= 5 5) "Hello" "Bye") -> "Hello"
(if (= 5 -5) "Hello" "Bye") -> "Bye"
```

#### loops
```lisp
(while cond [
    (line 1)
    (line 2)
])

(set x 10)
(while (< x 10) [
    (doSomething x)
    (set x (- x 1))
])
```


#### pattern matching
```lisp
(set x 7)
(match x
    [comparison return]
    [1 "one"]
    [2 "two"]
    [-1 _]
    [_ "Unknown"]
)

```

## Test programs

```lisp
(set x 10)

(if (< x 100)
    (print "X is small")
    (print "X is big"))
```