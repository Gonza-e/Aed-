def traducir(num: int):
    if num > 1:
        traducir(num // 2)
    
    print(num % 2)
    


traducir(11)
