def digitos(n:int) -> bool:
    if n == 0:
        return True
    else:
        if (n % 10) % 2 == 0:
            return digitos(n // 10)
        else:
            return False


numero = int(input("Ingrese un numero: \n"))

if digitos(numero):
    print("Todos los digitos del numero son pares")
else:
    print("No todos los digitos del numero son pares")