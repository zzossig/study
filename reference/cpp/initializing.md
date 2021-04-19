# Initializing

## string

```cpp
// C-style Strings
char my_name[] {"Frank"}; // F r a n k \0
my_name[5] = 'y' // Problem

char my_name[8] {"Frank"};
my_name[5] = 'y' // ok

char my_name[8];
my_name = "Frank"; // Error
strcpy(my_name, "Frank") // ok

// C++ Strings
#include <string>
using namespace std;

string s1; // empty
s1 = "C++ Rocks!";
string s2 {"Frank"}; // Frank
string s3 {s2}; // Frank
string s4 {"Frank", 3}; // Fra
string s5 {s3, 0, 2}; // Fr
string s6 (3, 'X') // XXX

string sentence;
sentence = "C++" + " is powerful"; // Illegal since the two c-style strings cannot concatenated.
```

## variable

```cpp
int age; // uninitialized
int age = 21; // c-like initialization
int age (21); // constructor initialization
int age {21}; // c++11 list initialization syntax
int age {}; // initialize to zero
```

## array

```cpp
int test_scores[5] {100,95,99,87,88};
int high_score_per_level[10] {3,5}; // init to 3,5 and remaining to 0

const double days_in_year {365};
double hi_temperatures[days_in_year] {0}; // init all to zero

int another_array[] {1,2,3,4,5}; // size automatically calculated

int movie_rating[3][4] // multi-dimensional arrays
{
  {0,4,3,5},
  {2,3,3,5},
  {1,4,4,5},
};
```

## vector

```cpp
#include <vector>

vector <char> vowels;
vector <int> test_scores;

vector <char> vowels (5);
vector <int> test_scores (10);

vector <char> vowels {'a', 'e', 'i', 'o', 'u'};
vector <int> test_scores {100, 98, 89, 85, 93};
vector <double> hi_temperatures (365, 80.0); // vector size: 365, all values initialized to 80.0
```

## enum

```cpp
enum Color {
  red, green, blue
};
Color screen_color {green}; 
```

## switch

```cpp
switch (screen_color) {
  case red:
    cout << "color is red" << endl;
    break;
  case green:
    cout << "color is green" << endl;
    break;
  case blue:
  {
    char anotherVal{}; // variable declaration should be wrapped with curlybraces in switch statement.
    cout << "color is blue" << endl;
    break;
  }
  default:
}
```

## for

```cpp
for (int i{0}; i < 15; ++i)
  cout << i << endl;
  
for (;;)
  cout << "Endless loop" << endl;

int scores [] {100, 90, 97};
for (int score : scores)
  cout << score << endl;

for (auto score : scores)
  cout << score << endl;

for (auto score : {60.2, 80.1, 90.0, 78.2})
  cout << score << endl;

for (auto c : "Frank")
  cout << c << endl;
```

## while

```cpp
int i{0};
while (i <= 10) {
  cout << i << endl;
  i++
}

do {
  cout << i << endl;
  i++
} while(i <= 10)
```

## Dynamic Memory Allocation

```cpp
int *int_ptr {nullptr};
int_ptr = new int;
delete int_ptr;

int size {};
int *int_ptr {nullptr};
array_ptr = new int[size];
delete[] array_ptr;
```
