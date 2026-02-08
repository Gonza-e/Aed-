def eliminarDig(num,val: int) -> int:
    if num == 0:
        return 0
    else:
        if (num % 10) != val:
            return eliminarDig(num//10,val) * 10 + num % 10
        else: 
            return eliminarDig(num//10,val)
        

#print(eliminarDig(1251215,1))

def divisores(num,num1: int) -> int:
    if num1 == 1: 
        return 1
    else:
        if num % num1 == 0:
            return 1 + divisores(num,num1-1)
        else:
            return divisores(num,num1-1)
    
def esPrimo(num: int) -> bool:
    if divisores(num,num) <= 2:
        return True 
    else:
        return False 

def primoPerfecto(num,numog,suma: int) -> int:
    if (num == 0) and (esPrimo(numog)) and (esPrimo(suma)):
        return True 
    else:
        if esPrimo(num%10):
            return primoPerfecto(num//10,numog,suma+num)
        else:
            return primoPerfecto(num//10,numog,suma)
        
#if primoPerfecto(227,227,0):
#    print("El numero es primo perfecto")


def patronDos(num,dig: int) -> bool:
    if (num == 0) or ((num < 10) and (num % 10) == dig-2):
        return True 
    else:
        dig = num % 10
        if (num//10) % 10 == dig-2:
            return patronDos(num//10,dig)
        else:
            return False 
        
#numero = int(input("Ingrese un numero para comprobar: \n"))

#if patronDos(numero,0):
#    print("El patron si se cumple")
#else: 
#   print("El patron no se cumple")

def factorial2(num: int) -> int:
    if num <= 0:
        return 1
    else:
        return factorial2(num-2) * num 

#print(factorial2(5)) 

def mul(num1,num2: int) -> int:
    if num2 == 0 or num1 == 0:
        return 0
    else:
        if num2 == 1:
            return num1
        else:
            return mul(num1,num2-1) + num1 
        
#print(mul(3,7))

bandera = False 
A = [5,2,7,26,63,23,2,52]
def burbuja(lista: list):
    bandera = False
    while not bandera:
        bandera = True
        for i in range(0,len(A)-1):
            if lista[i] > lista[i+1]:
                x = lista[i]
                lista[i] = lista[i+1]
                lista[i+1] = x 
                bandera = False 

    return lista

def seleccion(lista: list):
    for i in range(0,len(lista)-1):
        x = lista[i]
        for j in range(i+1,len(lista)):
            if x > lista[j]:
                x = lista[j]
                mayor = j 

        lista[mayor] = lista[i]
        lista[i] = x 

    return lista 

#print(seleccion(A))

def busquedaBinaria(A:list,i:str,ind:int):
    izq = 0
    der = len(A)
    medio = (izq+der)//2
    while izq < der and (i != A[medio]):
        if i < A[medio]:
            der = medio 
        else: 
            izq = medio  
    
        medio = (izq+der)//2

    if i == A[medio]:
        return medio
    else:
        return -1 
    
Ar = ["A","B","C","D","E","F","G","H","I","J","K"]

#print(busquedaBinaria(Ar,"B",0))

def invertirArreglo(A:list,ind:int):
    x: str
    if ind < 11:
        x = A[10-ind]
        invertirArreglo(A,ind+1)
        A[ind] = x

    return A 

#print(invertirArreglo(Ar,0))

def insercionDirecta(A:list):
    for i in range(1,len(A)):
        actual = A[i]
        j = i-1
        while (j>=0) and (A[j]>actual):
            A[j+1] = A[j]
            j = j-1
        
        A[j+1] = actual 

    return A

print(insercionDirecta(A))

    