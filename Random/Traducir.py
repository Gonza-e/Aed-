def traducir(num: int) -> int:
    if num < 2:
        return num 
    else: 
        return traducir(num // 2) * 10 + num % 2
    

print(traducir(8))