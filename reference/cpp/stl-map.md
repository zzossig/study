# STL Map

```cpp
#include <map>
```

## features

- Similar to a dictionary
- Elements are stored as Key, Value pairs (std:pair)
- Ordered by key
- No duplicate elements (keys are unique)
- Direct element access using the key
- All iterators available and invalidate when corresponding element is deleted

## initialization and assignment

```cpp
std::map<std::string, int> m {
  {"Larry", 18},
  {"Moe", 25}
};
```

## common methods

```cpp
std::map<std::string, std::string> m {
  {"Bob", "Butcher"},
  {"Anne", "Baker"},
  {"George", "Candlestick maker"}
};

std::cout << m.size(); // 3
std::cout << m.max_size; // a very large number

std::pair<std::string, std::string> p1 {"James", "Mechanic"};
m.insert(p1);
m.insert(std::make_pair("Roger", "Ranger"));

m["Frank"] = "Teacher"; // insert
m["Frank"] = "Instructor"; // update value
m.at("Frank") = "Professor"; // update value

m.erase("Anne"); // erase Anne

if (m.find("Bob") != m.end())
  std::cout << "Found Bob!";

int num = m.count("Bob"); // 0 or 1
m.clear(); // remove all elements
m.empty(); // true or false
```

## std::multi_map

- `#include <map>`
- Ordered by key
- Allows duplicate elements
- All iterators are available

## std::unordered_map

- `#include <unordered_map>`
- Elements are unordered
- No duplicate elements allowed
- No reverse iterators are allowed

## std::unordered_multimap

- `#include <unordered_map>`
- Elements are unordered
- Allows duplicate elements
- No reverse iterators are allowed

## insert

```cpp
std::map<std::string, std::set<int>> grades {
  {"Larry", {100,90}},
  {"Moe", {94}},
  {"Curly", {80,90,100}}
};

grades["Larry"].insert(95);

auto it = grades.find("Moe");
if (it != grades.end())
  it->second.insert(1000);
```