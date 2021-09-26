ReadFull

ReadFull은 r에서 buf로 정확히 len(buf) 바이트를 읽습니다. 
복사된 바이트 수와 읽은 바이트 수가 적으면 오류를 반환합니다. 
오류는 읽은 바이트가 없는 경우에만 EOF입니다. 
모든 바이트가 아닌 일부를 읽은 후에 EOF가 발생하면 ReadFull은 ErrUnexpectedEOF를 반환합니다. 
반환 시 err == nil인 경우에만 n == len(buf)입니다. 
r이 적어도 len(buf) 바이트를 읽은 오류를 반환하면 오류가 삭제됩니다.

Read메서드 
Read주어진 바이트 슬라이스를 데이터로 채우고 채워진 바이트 수와 오류 값을 반환합니다. io.EOF스트림이 종료되면 오류를 반환합니다 .