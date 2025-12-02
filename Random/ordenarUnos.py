def ordenar(num: int):
    if num > 0:
        if num % 10 == 1:
            print(num % 10)

        ordenar(num // 10)

        if num % 10 == 0:
            print(num % 10)


ordenar(1010011) 

