#Desarrollar una funcion que verifique que un numero es divdivisible

def divDivisible(num,num1,divs: int) -> int:
    if num1 == 0:
        if num % divs == 0:
            return True 
        else: 
            return False 
    else:
        if num % num1 == 0:
            return divDivisible(num,num1-1,divs+1)
        else:
            return divDivisible(num,num1-1,divs)
        

while True:
    numero = int(input("Ingrese un numero: \n"))

    if numero >= 0:
        if divDivisible(numero,numero,0):
            print("El numero es divDivisble")
        else:
            print("El numero no es divDivisible")
    else:
        break 

