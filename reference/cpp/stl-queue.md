# STL Queue

```cpp
#include <queue>
```

## features

- First-in First-out data structure
- Implemented as an adapter over other STL container(can be implemented as a list or deque)
- Elements are pushed at the back and popped from the front
- No iterators are supported

## initialization

```cpp
std::queue<int> q; // deque
std::queue<int, std::list<int>> q2; // list
std::queue<int, std::deque<int>> q3; // deque
```

## common methods

```cpp
std::queue<int> q;
q.push(10);
q.push(20);
q.push(30);

std::cout << q.front(); // 10
std::cout << q.back(); // 30

q.pop(); // remove 10
q.pop(); // remove 20

std::cout << q.size(); // 1
```

# STL Priority Queue

```cpp
#include <queue>
```

## features

- Allows insertions and removal of elements in order from the front of the container
- Elements are stored internally as a vector by default
- Elements are inserted in priority order(Largest value will always be at the front)
- No iterators are supported

## common methods

```cpp
std::priority_queue<int> pq; // vector

pq.push(10);
pq.push(20);
pq.push(3);
pq.push(4);

std::cout << pq.top(); // 20 (largest)
pq.pop(); // remove 20
pq.top(); // 10 (largest)
```