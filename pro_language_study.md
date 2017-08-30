# Program language study

## Type and variable

### Constant variable

```
static class Constants
{
    public const double Pi = 3.14159;
    public const int SpeedOfLight = 300000;
}

class Program
{
    static void Main()
    {
        double radius = 5.3;
        double area = Constants.Pi * (radius * radius);
        int secsFromSun = 149476000 / Constants.SpeedOfLight; // in km
    }
}
```

### Static type / Dynamic type

```
static type: You must specify the type when define a variable. Such as C, C++, Java, C#.

dynamic type: The type is associated with run-time. Such as Ruby, Python, Javascript.
```

### Strong type / Weak type

```
strong type: The language will not implicit the type of variable. When check the type is not match, it will raise ane exception. Such as Java, C#, Ruby.

weak type: Fox example, The language tries to implicit conversion from strings to number. Such as C.
```

### Type conversion

```
static language use type conversion to change the type of variable.

C#

double d = 5673.74;
int i;
i = (int)d;

Convert
```

### Type derivation

```
dynamic language use this way to derivation the type of a variable.
```

### Value type / Reference type

http://www.tutorialsteacher.com/csharp/csharp-value-type-and-reference-type

```
It holds a data value within its own memory space (save in stack).

It holds address within its own memory space. The reference value is saved in the heap.
Such as: String, Array, Class, Delegates.

C#: boxing unboxing
int i = 123;      // a value type
object o = i;     // boxing
int j = (int)o;   // unboxing

When boxing, change from value type to reference type, save from stack to heap.
When unboxing, change from reference type to value type, the value save from heap to stack.
```

### Basic type

```
integer, fload, double, string, array, map, set, list, queue, collection
```

### Random number

### Condition

https://www.tutorialspoint.com/ruby/ruby_if_else.htm

```
if/else
case/when
.. ? .. : ..
```

### Loop

```
for
while
each
```

### Recursive

https://zh.wikipedia.org/wiki/%E6%B1%89%E8%AF%BA%E5%A1%94

### Function

```
function define(key, optional argument, out argument)(https://www.dotnetperls.com/out)
function call(pass by value, pass by reference(out, ref, class))
```

- Out / Ref

http://www.cnblogs.com/wang_yb/archive/2011/05/18/2050574.html

- Template function

https://docs.microsoft.com/zh-cn/dotnet/csharp/programming-guide/generics/generic-methods

```
static void Swap<T>(ref T lhs, ref T rhs)
{
    T temp;
    temp = lhs;
    lhs = rhs;
    rhs = temp;
}

public static void TestSwap()
{
    int a = 1;
    int b = 2;

    Swap<int>(ref a, ref b);
    // Swap(ref a, ref b);
    System.Console.WriteLine(a + " " + b);
}
```

- Delegate

https://docs.microsoft.com/zh-cn/dotnet/csharp/programming-guide/delegates/using-delegates

```
public delegate void Del(string message);

// Create a method for a delegate.
public static void DelegateMethod(string message)
{
    System.Console.WriteLine(message);
}

// Instantiate the delegate.
Del handler = DelegateMethod;

// Call the delegate.
handler("Hello World");
```

- Lambda / Anonymous function

https://docs.microsoft.com/zh-cn/dotnet/csharp/programming-guide/statements-expressions-operators/anonymous-functions

https://docs.microsoft.com/zh-cn/dotnet/csharp/programming-guide/statements-expressions-operators/lambda-expressions

http://www.cnblogs.com/jackFloyd/p/4844073.html

- Closures

http://www.cnblogs.com/jackFloyd/p/4844073.html

http://www.cnblogs.com/frankfang/archive/2011/08/03/2125663.html

http://www.cnblogs.com/jujusharp/archive/2011/08/04/C-Sharp-And-Closure.html

### exception

### 存儲管理

```
存儲管理
靜態分配/棧分配/堆分配
對象週期
垃圾回收
```

### using namespace

### 並發

### async/await

### file

### event

### TODO

```
指針和遞歸類型（递归类型，就是可以在其对象中包含一个或多个本类型对象的引用类型）
```