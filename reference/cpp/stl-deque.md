# STL Deque

```cpp
#include <deque>
```

## features

- Deque stands for `double ended queue`
- Elements ar `not` stored in continuous memory
- Can expand and contract as needed
- Handled automatically
- Direct element access
- Rapid insertion and deletion at the front and back (contant time)
- Insertion or removal of elements (linear time)

## initialization and assignment

```cpp
std::deque<int> d{1,2,3,4,5};
std::deque<int> d(10, 100); // ten 100s

std::deque<std::string> stooges {
  std::string("Larry"),
  "Moe",
  std::string("Curly")
};

d = {2,4,6,8,10};
```

## common methods

```cpp
std::deque<int> d {1,2,3,4,5};

std::cout << d.Size(); // 5
std::cout << d.max_size; // a very large number
std::cout << d.at(0); // 1
std::cout << d[1]; // 2
std::cout << d.front(); // 1
std::cout << d.back(); // 5

d.clear()

Person p1 {"Larry", 18};
std::deque<Person> d;

d.push_back(p1);
d.pop_back();

d.push_front(Person{"Larry", 18});
d.pop_front();

d.emplace_back("Larry", 18);
d.emplace_front("Moe", 24);
```

## deque size can automatically increase or decrease

```cpp
std::vector<int> vec{1,2,3,4,5,6,7,8,9,10};
std::deque<int> d;

for (cont auto &elem:vec) {
  if (elem%2==0)
    d.push_back(elem);
  else
    d.push_front(elem);
}
```

## copy

```cpp
std::vector<int> vec{1,2,3,4,5,6,7,8,9,10};
std::deque<int> d;

std::copy(vec.begin(), vec.end(), std::front_inserter(d)); // 10 9 8 7 6 5 4 3 2 1
d.clear();
std::copy(vec.begin(), vec.end(), std::back_inserter(d)); // 1 2 3 4 5 6 7 8 9 10
```