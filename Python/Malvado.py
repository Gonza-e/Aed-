def malvado(num,cont: int) -> bool:
    if num == 0:
        if (cont % 2) == 0:
            return True
        else:
            return False
    else:
        if (num % 10) == 1:
            return malvado(num // 10,cont + 1)
        else: 
            return malvado(num // 10,cont)
        
numero = int(input("Ingrese un numero: "))

if malvado(numero,0):
    print("El numero es malvado")
else: 
    print("El numero no es malvado")