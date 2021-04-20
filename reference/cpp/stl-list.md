# STL List

```cpp
#include <list>
```

## features

- Dynamic size
- Direct element access is `not` provided
- Rapid insertion and deletion of elements anywhere in the container (contant time)
- All iterators available and invalidate when corresponding element is deleted

## initialization and assignment

```cpp
std::list<int> l{1,2,3,4,5};
std::list<int> ll(10, 100); // ten 100s

std::list<std::string> stooges{
  std::string("Larry"),
  "Moe",
  std::string("Curly")
};
l = {2,4,6,8,10};
```

## common methods

```cpp
std::list<int> l{1,2,3,4,5};

std::cout << l.size(); // 5
std::cout << l.max_size; // a very large number
std::cout << l.front(); // 1
std::cout << l.back(); // 5

Person p1 {"Larry", 18};
std::list<Person> l;

l.push_back(p1);
l.pop_back();

l.push_front(Person{"Larry", 18});
l.pop_front();

l.emplace_back("Larry", 18);
l.emplace_front("Moe", 24);
```

## methods that use iterators

```cpp
std::list<int> l{1,2,3,4,5};
auto it = std::find(l.begin(), l.end(), 3);

l.insert(it, 10); // 1 2 10 3 4 5
l.erase(it); // erases the 3, 1 2 10 4 5
l.resize(2); // 1 2
l.resize(5); // 1 2 0 0 0
```

## insert and advance

```cpp
std::list<int> l{1,2,3,4,5,6,7,8,9,10};
auto it = std::find(l.begin(), l.end(), 5);
if (it != l.end())
  l.insert(it, 100);

std::list<int> l2{1000,2000,3000};
l.insert(it, l2.begin(), l2.end());

std::advance(it, -4); // point to the 100
l.erase(it); // remove the 100 and the iterator becomes invalid
```

# STL Forward List

```cpp
#include <forward_list>
```

## features
- Dynamic size
- Direct element access is `not` provided
- Rapid insertion and deletion of elements anywhere in the container
- Reverse iterators not available. Iterators invalidate when corresponding element is deleted

## common methods

```cpp
std::forward_list<int> l {1,2,3,4,5};

std::cout << l.size(); // Not available
std::cout << l.max_size; // a very large number

std::cout << l.front(); // 1
std::cout << l.back(); // Not available

Person p1 {"Larry", 18};
std::forward_list<Person> l;

l.push_front(p1);
l.pop_front();
l.emplace_front("Moe", 24);
```