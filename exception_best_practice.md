# Exception best practice

[Reference](https://www.codeproject.com/Articles/9538/Exception-Handling-Best-Practices-in-NET#Whenindoubt,don%27tAssert,throwanException23)

## Plan for worst

### Pre-checking

```
Check type, format, range early.
```

### Do not trust external data

```
Data from registry, database, disk, socket, file is not reliable.
We can not trust the data from config file, maybe someone would edit it and corrupt it.
```

### Reliable devices

```
Video, mouse, keyboard
```

### Write operation fail

- Permission
- Device is not there
- Not enough space
- Device has a physical fault

```
Both read and write have the same issues above.
```

## Code safely

### Don't throw new Exception()

```
Derive you own exception class from ApplicationException.
You can set a specialized exception handler for exceptions throw by framework and another for exceptions thrown by yourself.
```

### Don't put important exception information on the message field

```
Create fields to store the exception information, not on the message.
```

### Put a single catch(Exception ex) per thread

```
Generic exception handling should be done in a central point in your application. Each thread needs a separate try/catch block.
```

### Generic exception caught should be published

```
Log generic exception and only one time.
```

### Log stack trace of exception

```
Exception.message
Exception.stack_trace, this is priceless.
```

### Catch the specific exception

```
public class MyClass
{
    public static string ValidateNumber(string userInput)
    {
        try
        {
            int val = GenericLibrary.ConvertToInt(userInput);
            return "Valid number";
        }
        catch (Exception)
        {
            return "Invalid number";
        }
    }
}

public class GenericLibrary
{
    public static int ConvertToInt(string userInput)
    {
        return Convert.ToInt32(userInput);
    }
}
```

```
This code is not good, if the exception is "OutOfMemoryException", it return 'Invalid number', although the userInput is valid.
We should only handle the exception we know how to handle it.
```

### Finally block

```
Cleanup code, such as closing streams, restoring state.
This code in finally block will run after the function return.
```

```
string ReadTempFile(string FileName)
{
    try
    {
        string fileContents;
        using (StreamReader sr = new StreamReader(FileName))
        {
            fileContents = sr.ReadToEnd();
        }
        File.Delete(FileName);
        return fileContents;
    }
    catch (Exception)
    {
        File.Delete(FileName);
        throw;
    }
}
```

Better code:

```
string ReadTempFile(string FileName)
{
    try
    {
        using (StreamReader sr = new StreamReader(FileName))
        {
            return sr.ReadToEnd();
        }
    }
    finally
    {
        File.Delete(FileName);
    }
}
```
