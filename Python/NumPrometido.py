def numPrometido(num1,num2,num3,suma: int) -> bool: 
    if (num3 == 0): 
        if ((num2 + 1) == suma):
            return True
        else: 
            return False
    else: 
        if ((num1 % num3) == 0): 
            return numPrometido(num1,num2,num3-1,suma+num3)
        else: 
            return numPrometido(num1,num2,num3-1,suma)
        

numero1 = int(input("Ingrese un valor: "))
numero2 = int(input("Ingrese otro valor: "))


if numPrometido(numero1, numero2, numero1 - 1, 0) and numPrometido(numero2, numero1, numero2 - 1, 0):
    print("Los números son prometidos")
else:
    print("Los números no son prometidos")
