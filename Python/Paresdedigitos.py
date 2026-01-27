def paresConsecutivos(num,dig: int) -> bool:
    if num == 0:
        return False 
    else:
        if num % 10 == dig: 
            return True 
        else:
            return paresConsecutivos(num // 10, num % 10)
        

if paresConsecutivos(1223,1067):
    print("El numero contiene un par de digitos consecutivos")
else:
    print("El numero no contiene pares de digitos consecutivos")