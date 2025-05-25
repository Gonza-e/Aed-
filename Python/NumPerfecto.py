def numPerfecto(num1,suma,num2: int) -> bool:
    if (num2 == 0):
        if (suma == num1): 
            return True 
        else: 
            return False 
    else: 
        if ((num1 % num2) == 0): 
            return numPerfecto(num1, suma + num2, num2 - 1)
        else: 
            return numPerfecto(num1, suma, num2 - 1)
        

print("Ingrese un numero")
numero = int(input())

if numPerfecto(numero, 0,numero - 1): 
    print("El numero es perfecto")
else: 
    print("El numero no es perfecto")