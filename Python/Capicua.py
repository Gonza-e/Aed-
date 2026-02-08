def capicua(num1, num2, num_og: int) -> bool: 
    if (num1 == 0): 
        if (num2 == num_og):
           return True
        else: 
            return False 
    else: 
        return capicua(num1 // 10, (num2 * 10) + (num1 % 10), num_og)


print("Ingrese un numero")
numero = int(input())

if capicua(numero, 0, numero): 
    print("El numero es capicua")
else: 
    print("El numero no es capicua") 



