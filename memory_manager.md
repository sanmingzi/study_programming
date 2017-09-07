# Memory manager

http://www.helper3.com/detail/id/3634

## Demo

```
void f()
{
  int* p=new int[5];
}
```

The point p allocated a stack memory. <br/>
The value that 'p' pointed is allocated in the heap. <br/>
The stack memory save a point 'p' which point to a block of heap memory.

## Memory locations

- heap
- stack
- free (自由存储区)
- global/static
- constant

## stack VS heap

### Management style

Programmer should control the release of heap memory, otherwise it will have memory leak.

### Size

- Heap: 4G
- Stack: much smaller than heap

### Growth direction

Heap grow with memory address increase. <br/>
Stack grow with memory address decrease.

### Allocate way

Heap memory allocate dynamic. <br/>
Local variable allocate in stack is static, we can use 'allocate method' to allocate memory in stack.

### Allocate efficiency

Stack is more effective than heap.
