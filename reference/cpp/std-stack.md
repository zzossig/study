# STL Stack

```cpp
#include <stack>
```

## features

- Last-in First-out data structure
- Implemented as an adapter over other STL container(Can be implemented as a vector, list, or deque)
- All operations occur on one end of the stack (top)
- No iterators are supported

## initialization

```cpp
std::stack<int> s;                    // deque
std::stack<int, std::vector<int>> s1; // vector
std::stack<int, std::list<int>> s2;   // list
std::stack<int, std::deque<int>> s3;  // deque
```

## common methods

```cpp
std::stack<int> s;
s.push(10);
s.push(20);
s.push(30);

std::cout << s.top(); // 30
s.pop(); // 30 is removed
s.pop(); // 20 is removed
std::cout << s.size(); // 1

s.empty() // true or false
```