def cerosMult3(num, cont: int) -> bool: 
    if (num == 0): 
        if ((cont % 3) == 0):
            return True 
        else: 
            return False 
    else: 
        if ((num % 10) == 0):
            return cerosMult3(num // 10, cont + 1)
        else: 
            return cerosMult3(num // 10, cont)


print("Ingrese un numero")
numero = int(input())

if cerosMult3(numero, 0): 
    print("La cantidad de ceros es mult de 3")
else: 
    print("La cantidad de ceros no es mult de 3")  
