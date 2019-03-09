def main():
    print(urlify("Mr John Smith    ", 13))

def urlify(string, realSize):
    string = list(string)

    i = len(string) - 1
    for j in range(realSize-1, -1, -1):
        if string[j] == " ":
            string[i-2:i+1] = ["%", "2", "0"]
            i -= 3
        else:
            string[i] = string[j]
            i -= 1
    
    return "".join(string)

main()