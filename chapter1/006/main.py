def main():
    print(compression("aabcccccaaa"))

    print(compression("aaa"))
    print(compression("aa"))
    print(compression("aabb"))
    print(compression(""))
    print(compression("abcd"))
    print(compression("abcdddddd"))
    print(compression("aaAA"))
    print(compression("aaaAAA"))

def compression(string):
    if len(string) == 0:
        return ""

    buffer = []

    count = 0
    lastChar = string[0]
    for c in string:
        if lastChar != c:
            buffer.append(f'{lastChar}{count}')
            count = 1
        else:
            count += 1

        lastChar = c

    if count > 0:
        buffer.append(f'{lastChar}{count}')
    
    new_string = "".join(buffer)

    if len(string) <= len(new_string):
        return string
    else:
        return new_string

main()