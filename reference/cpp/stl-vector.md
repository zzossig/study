# STL Vector

```cpp
#include <vector>
```

## std::vector initialization and assignment

```cpp
std::vector<int> vec {1,2,3,4,5};
std::vector<int> vec1 (10, 100); // ten 100s

std::cout << vec.size();     // 5
std::cout << vec.capacity(); // 5
std::cout << vec.max_size(); // a very large number

std::cout << vec.at(0);      // 1
std::cout << vec[1];         // 2

std::cout << vec.front();    // 1
std::cout << vec.back();     // 5
```

```cpp
Person p1 {"Larry", 18};
std::vector<Person> vec;

vec.push_back(p1); // add p1 to the back
vec.pop_back();    // remove p1 from the back

vec.push_back(Person{"Larry", 18});
vec.emplace_back("Larry", 18); // efficient!!
```

## std::vector - common methods

```cpp
std::vector<int> vec1 {1,2,3,4,5};
std::vector<int> vec2 {10,20,30,40,50};

auto it = std::find(vec1.begin(), vec1.end(), 3);
vec1.insert(it, 10); // 1,2,10,3,4,5

it = std::find(vec1.begin(), vec1.end(), 4);
std::insert(it, vec2.begin(), vec2.end()); // 1,2,10,3,10,20,30,40,50,4,5

vec2.swap(vec1);

std::sort(vec1.begin(), vec1.end());

vec1.clear(); // remove all elements
vec2.erase(vec2.begin(), vec2.begin()+2);

auto it = vec2.begin();
while (it != vec2.end()) {
  if (*it%2 == 0)
    vec.erase(it);
  else
    it++;
}
```

```cpp
std::vector<int> vec1 {1,2,3,4,5};
std::vector<int> vec2 {10,20};

std::copy(vec1.begin(), vec1.end(), std::back_inserter(vec2));
// vec1: [1 2 3 4 5]
// vec2: [10 20 1 2 3 4 5]

vec1 = {1,2,3,4,5,6,7,8,9,10};
vec2 = {10,20};

std::copy_if(vec1.begin(), vec1.end(), std::back_inserter(vec2), [](int x) { return x%2 == 0; });
```

```cpp
std::vector<int> vec1 {1,2,3,4,5};
std::vector<int> vec2 {10,20,30,40,50};
std::vector<int> vec3;

std::transform(vec1.begin(), vec1.end(), vec2.begin(),
  std::back_inserter(vec3),
  [](int x, int y) { return x*y; });
// vec3: [10 40 90 160 250]
```

```cpp
```
