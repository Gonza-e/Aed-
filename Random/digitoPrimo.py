def esPrimo(num, num1: int) -> bool:
    if num < 2 :
        return False
    else:
        if num1 == 0:
            return True 
        else:
            if (num % num1) == 0 and (num1 != 1):
                return False 
            else: 
                return esPrimo(num,num1 - 1)
            

k = int(input("Ingrese un numero \n"))
numero = 3456

while k > 1:
    k -= 1
    numero = numero // 10 

digito = numero % 10
print(digito)
print(esPrimo(digito,digito - 1))

    
