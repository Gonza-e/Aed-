def fact(n:int) -> int:
    if n == 0:
        return 1 
    else:
        return fact(n-1) * n 
    

numero = int(input("Ingrese un numero para calcular su factorial \n"))

print("Factorial:",fact(numero))