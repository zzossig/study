# STL Algorithm

```cpp
#include <algorithm>
```

## `find` with primitive types

```cpp
std::vector<int> vec {1,2,3};
auto loc = std::find(vec.begin(), vec.end(), 3);

if (loc != vec.end())
  std::cout << *loc << std::endl;
```

## `find` with user-defined types

`operator==` is used and must be provided by your class

```cpp
std::vector<Player> team {/* assume initialized */};
Player p {"Hero", 100, 12};

auto loc = std::find(team.begin(), team.end(), p);
```

## `for_each` - using functor

```cpp
struct Square_Functor {
  void operator() (int x) {
    std::cout << x*x << " ";
  }
};
Square_Functor square;

std::vector<int> vec {1,2,3,4};
std::for_each(vec.begin(), vec.end(), square);
// 1 4 9 16
```

## `for_each` - using a functio pointer

```cpp
void square(int x) {
  std::cout << x*x << " ";
}
std::vector<int> vec {1,2,3,4};
std::for_each(vec.begin(), vec.end(), square);
// 1 4 9 16
```

## `for_each` - using a lambda expression

```cpp
std::vector<int> vec{1,2,3,4};
std::for_each(vec.begin(), vec.end(), [](int x) { std::cout << x*x << " "; });
// 1 4 9 16
```

## `count`

```cpp
std::vector<int> vec {1,2,3,4,5,1,2,1};
int cnt = std::count(vec.begin(), vec.end(), 1);
std::cout << cnt << std::endl;
// 3
```

## `count_if`

```cpp
std::vector<int> vec{1,2,3,4,5,1,2,1,100};
int cnt = std:count_if(vec.begin(), vec.end(), [](int x) { return x%2 == 0; });
std::cout << cnt << std::endl;
// 4
```

## `replace`

```cpp
std::vector <int> vec{1,2,3,4,5,1,2,1};
std::replace(vec.begin(), vec.end(), 1, 100);
for (auto i:vec) {
  std::cout << i << " ";
}
// 100 2 3 4 5 100 2 100
```

## `all-of`, `any-of`

```cpp
std::vector<int> vec1 {1,3,5,7,9,1,3,13,19,5};
std::all_of(vec1.begin(), vec1.end(), [](int x) { return x>10; }); // false
std::any_of(vec1.begin(), vec1.end(), [](int x) { return x>10; }); // true
```

## `string_transform`

```cpp
std::string str1 {"This is test"};
std::transform(str1.begin(), str1.end(), str1.begin(), ::toupper); // ::toupper defined in <cctype>
```