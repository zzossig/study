# STL Set

```cpp
#include <set>
```

## initialization and assignment

```cpp
std::set<int> s{1,2,3,4,5};
s = {2,4,6,8,10};
```

## common methods

```cpp
std::set<int> s{4,1,1,3,3,2,5}; // 1 2 3 4 5

std::cout << s.size(); // 5
std::cout << s.max_size; // a very large number

s.insert(7); // No concept of front and back
s.erase(3); // erase the 3

auto it = s.find(5)
if (it != s.end())
  s.erase(it);

s.count(1); // if exist, return true else false
s.clear();
s.empty();
```

# STL Multi Set

```cpp
#include <set>
```

## features

- Sorted by ket
- Allows duplicate elements
- All iterators are available

# STL Unordered Set

```cpp
#include <unordered_set>
```

## features

- Elements are unordered
- No duplicate element allowed
- Elements cannot be modified - must be erased and new element inserted
- No reverse iterators are allowed

# STL Unordered Multiset

```cpp
#include <unordered_set>
```

## features

- Elements are unordered
- Allows duplicate elements
- No reverse iterators are allowed

## set pair

```cpp
std::set<std::string> s{"A","B","C"};
auto result = s.insert("D"); // result is a pair <iterator, bool>

std::cout << "first: " << *(result.first) << std::endl; // D
std::cout << "second: " << result.second << std::endl; // true
```