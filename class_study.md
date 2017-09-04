# Class study

https://docs.microsoft.com/zh-cn/dotnet/csharp/programming-guide/classes-and-structs/classes

## Define / New

- Field

```
static field
initialize field in the constructor
public private protected internal
```

- Constant

```
class Calendar1
{
    public const int months = 12;
}
```

- Attribute

https://docs.microsoft.com/zh-cn/dotnet/csharp/programming-guide/classes-and-structs/properties

```
get/set

public class SaleItem
{
   string name;
   decimal cost;

   public string Name
   {
      get => name;
      set => name = value;
   }

   public decimal Price
   {
      get => cost;
      set => cost = value;
   }
}
```

- Method
- Event

- Operator

```
public struct Complex
{
    public int real;
    public int imaginary;

    public Complex(int real, int imaginary)
    {
        this.real = real;
        this.imaginary = imaginary;
    }

    public static Complex operator +(Complex c1, Complex c2)
    {
        return new Complex(c1.real + c2.real, c1.imaginary + c2.imaginary);
    }
}

class TestComplex
{
    static void Main()
    {
        Complex num1 = new Complex(2, 3);
        Complex num2 = new Complex(3, 4);
        Complex sum = num1 + num2;
    }
}
```

- Indexer

```
class SampleCollection<T>
{
   private T[] arr = new T[100];

   public T this[int i]
   {
      get { return arr[i]; }
      set { arr[i] = value; }
   }
}

class Program
{
   static void Main()
   {
      var stringCollection = new SampleCollection<string>();
      stringCollection[0] = "Hello, World";
      Console.WriteLine(stringCollection[0]);
   }
}
```

- Constructor

```
public class Child : Person
{
   private static int maximumAge;

   public Child(string lastName, string firstName) : base(lastName, firstName)
   { }

   static Child() => maximumAge = 18;
}
```

- Destructor
- Internal Class

```
class Container
{
    class Nested
    {
        Nested() { }
    }
}
```

## Static class

https://docs.microsoft.com/zh-cn/dotnet/csharp/programming-guide/classes-and-structs/static-classes-and-static-class-members

## Singleton class

https://msdn.microsoft.com/zh-TW/Library/ms998558.aspx

```
public class Singleton
{
   private static Singleton instance;

   private Singleton() {}

   public static Singleton Instance
   {
      get
      {
         if (instance == null)
         {
            instance = new Singleton();
         }
         return instance;
      }
   }
}
```

## Inherit

https://docs.microsoft.com/zh-cn/dotnet/csharp/programming-guide/classes-and-structs/inheritance

### Virtual

```
public class Bird
{
  public virtual void Eat()
  {
    Console.WriteLine("I am a bird, I like eat worm.");
  }
}

public  class Magpie : Bird
{
  public override void Eat()
  {
    Console.WriteLine("I am a magpie, I like eat worm.");
  }
}
```

### Abstract

```
abstract class ShapesClass
{
    abstract public int Area();
}

class Square : ShapesClass
{
    int side = 0;

    public override int Area()
    {
        return side * side;
    }
}
```

### Override



### Interface

https://docs.microsoft.com/zh-cn/dotnet/csharp/programming-guide/interfaces/

```
interface IEquatable<T>
{
    bool Equals(T obj);
}

public class Car : IEquatable<Car>
{
    public string Make {get; set;}
    public string Model { get; set; }
    public string Year { get; set; }

    // Implementation of IEquatable<T> interface
    public bool Equals(Car car)
    {
        if (this.Make == car.Make &&
            this.Model == car.Model &&
            this.Year == car.Year)
        {
            return true;
        }
        else
            return false;
    }
}
```

## Delete
