You want ParseFastq to accept an argument of type io.Reader, not
*io.Reader. *os.File implements io.Reader, but *os.File cannot be
treated as a pointer to an io.Reader.

When passing around interfaces, you typically pass by value because
the interfaces contain either a pointer or a primitive type.

-> 값 전달이 가능함.