# Iterator

## Iterator begin and end method

```cpp
#include <vector>

int main() {
  std::vector<int> vec {1,2,3};
  std::vector<int>::interator it = vec.begin();
  // or
  auto it = vec.begin();
}
```

## Operations with iterators

| Operation | Description | Type of Iterator|
|---|---|---|
| ++it | Pre-increment | All |
| it++ | Post-increment | All |
| it = it1 | Assignment | All |
| *it | Dereference | Input and Output |
| it-> | Arrow operator | Input and Output |
| it == it1 | Comparison for equality | Input |
| it != it1 | Comparison for inequality | Input |
| --it | Pre-decrement | Bidirectional |
| it-- | Post-decrement | Bidirectional |
| it + i, it += i | Increment | Random Access |
| it - i, it -= i | Decrement | Random Access |
| it < it1, it <= it1 | Comparison | Random Access |
| it > it1, it >= it1 | Comparison | Random Access |

## Using iterator - std::vector

```cpp
std::vector<int> vec {1,2,3};
std::vector<int>::iterator it = vec.begin();

while (it != vec.end()) {
  std::cout << *it << " ";
  ++it;
}
// 1 2 3
```

```cpp
std::vector<int> vec {1,2,3};
for (auto it = vec.begin(); it != vec.end(); it++) {
  std::cout << *it << " ";
}
// 1 2 3
```

## Using iterator - std::set

```cpp
std::set<char> suits {'C','H','S','D'};
auto it = suits.begin();

while (it != suits.end()) {
  std::cout << *it << " " << std::endl;
  ++it;
}
// C H S D
```

## Reverse iterator

```cpp
std::vector<int> vec {1,2,3};
std::vector<int>::reverse_iterator it = vec.begin();
// auto it = vec.rbegin();
while (it != vec.end()) {
  std::cout << *it << " ";
  ++it;
}
// 3 2 1
```

## Other iterators

- begin() and end(): iterator
- cbegin() and cend(): const_iterator
- rbegin() and rend(): reverse_iterator
- crbegin() and crend(): const_reverse_iterator
