def numerosFelices(num1, num2) -> bool:
    if num2 == 0:
        if num1 == 1:
            return True
        else: 
            return False 
    else:
        nueva_suma = 0
        while num1 > 0:
            digito = num1 % 10
            nueva_suma += digito ** 2
            num1 //= 10 # division entera para ir reduciendo el numero
        return numerosFelices(nueva_suma, num2 - 1)


numero = input("Ingrese un numero para evaluarlo: ")

if numerosFelices(int(numero),15):
    print("El numero es feliz")
else: 
    print("El numero no es feliz")
