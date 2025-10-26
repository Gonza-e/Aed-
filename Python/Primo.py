def esPrimo(num1,num2:int) -> bool:
    if num1 == 1:
        return False
    else:
        if num2 == 0: 
            return True
        else:
            if ((num1 % num2) == 0) and (num2 != 1):
                return False
            else:
                return esPrimo(num1,num2-1)

num = int(input("Ingrese un numero: "))

if esPrimo(num,num//2):
    print("El numero es primo")
else: 
    print("El numero no es primo")

