# Type and Variable

## Basic type

- integer
- float
- double
- string
- array
- map
- set
- list
- queue
- collection
- stack

## Constant Variable

```C#
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

## Static Type / Dynamic Type

- static type

You must specify the type when define a variable. Such as C, C++, Java, C#.

- dynamic type

The type is associated with run-time. Such as Ruby, Python, Javascript.

## Strong Type / Weak Type

- strong type

The language will not implicit the type of variable. When check the type is not match, it will raise ane exception. Such as Java, C#, Ruby.

- weak type

Fox example, The language tries to implicit conversion from strings to number. Such as C.

## Type Conversion

static language use type conversion to change the type of variable.

```C#
double d = 5673.74;
int i = (int)d;
```

## Type derivation

dynamic language use this way to derivation the type of a variable.

## Value Type / Reference Type

[Value Type and Reference Type](http://www.tutorialsteacher.com/csharp/csharp-value-type-and-reference-type)

- value type

It holds a data value within its own memory space (save in stack).

- reference type

It holds address within its own memory space. The reference value is saved in the heap.  
Such as: String, Array, Class, Delegates.

```C#
// boxing unboxing
int i = 123;      // a value type
object o = i;     // boxing
int j = (int)o;   // unboxing
```

When boxing, change from value type to reference type, save from stack to heap.  
When unboxing, change from reference type to value type, the value save from heap to stack.
