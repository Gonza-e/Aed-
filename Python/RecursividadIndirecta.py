def par(num: int) -> str:
    if num == 0:
        return "Es par"
    else:
        if (num % 10) % 2 != 0:
            return impar(num//10)
        else:
            return par(num//10)
        
def impar(num: int) -> str:
    if num == 0:
        return "Es impar"
    else:
        if (num % 10) % 2 == 0:
            return par(num//10)
        else:
            return impar(num//10)
        

numero = int(input("Ingrese un numero \n"))

print(par(numero))
    
