# Hacer una funcion recursiva que verifique que los digitos de un numero respetan el siguiente patron 
# par impar par impar... (siendo el ultimo numero par)

def patron(num, dig1, dig2: int) -> bool:
    if num == 0:
        return True 
    else:
        if num < 10 and (num % 2 == 0):
            return True 
        else:
            dig1 = num % 10 
            dig2 = (num // 10) % 10 
            if (dig1 % 2 == 0) and (dig2 % 2 != 0):
                return patron(num // 100,dig1,dig2)
            else:
                return False
    

num1 = int(input("Ingrese un numero \n"))

if patron(num1,0,0):
    print("El numero respeta el patron")
else:
    print("El numero no respeta el patron")