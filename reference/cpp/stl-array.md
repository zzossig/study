# STL Array

```cpp
#include <array>
```

## std::array - initialization and assignment

```cpp
std::array<int, 5> arr1 { {1,2,3,4,5} };
std::array<std::string, 3> stooges {
  std::string("Larry"),
  "Moe",
  std::string("Curly")
};
arr1 = {2,4,6,8,10};
```

## std::array - common methods

```cpp
std::array<int, 5> arr {1,2,3,4,5};
std::cout << arr.size(); // 5
std::cout << arr.at(0); // 1
std::cout << arr[1]; // 2
std::cout << arr.front(); // 1
std::cout << arr.back(); // 5
std::cout << arr.empty(); // 0 (false)
std::cout << arr.max_size(); // 5
std::cout << arr.fill(10); // fills all to 10
arr.swap(arr1) // swaps the 2 arrays
int *data = arr.data(); // get raw array address
```

## sort

```cpp
std::array<int, 5> arr{2,1,4,5,3};
std::sort(arr.begin(), arr.end());
```

## min_element, max_element

```cpp
std::array<int, 5> arr{2,1,4,5,3};
std::array<int, 5>::iterator min_num = std::min_element(arr.begin(), arr.end());
auto max_num = std::max_element(arr.begin(), arr.end());
```

## adjacent_find

```cpp
std::array<int, 5> arr {2,1,3,3,5};
auto adjacent = std::adjacent_find(arr.begin(), arr.end());
if (adjacent != arr.end())
  std::cout << "Adjacent element found with value: " << *adjacent << std::endl;
  // 3
```

## accumulate

accumulate is from `#include <numeric>`
```cpp
std::array<int, 5> arr{1,2,3,4,5};

int sum = std::accumulate(arr.begin(), arr.end(), 0)
std::cout << "Sum: " << sum << std::endl;
```

## count

```cpp
std::array<int, 10> arr {1,2,3,1,2,3,3,3,3,3};

int count = std::count{arr.begin(), arr.end(), 3};
// count = 6
```

## count_if

```cpp
std::array<int, 10> arr{1,2,3,50,60,70,80,200,300,400};
int count = std::count_if(arr.begin(), arr.end(), [](int x) { return x>10 && x<200; });
// count = 4
```