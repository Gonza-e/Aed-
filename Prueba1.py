def casiPerfectos(n1,n2,s) -> bool:
    if n1 == 1: 
        return True
    else:
        if n2 == 0:
            if s == n1-1:
                return True 
            else:
                return False 
        else:
            if n1 % n2 == 0:
                return casiPerfectos(n1,n2-1,s+n2)
            else:
                return casiPerfectos(n1,n2-1,s)
        

while True:
    cont = 0
    A = int(input("Ingrese un numero positivo mayor a 0: \n"))
    B = int(input("Ingrese un numero positivo mayor a 0: \n"))

    if (A>0) and (A<B):
        for i in range(A,B+1):
            if casiPerfectos(i,i-1,0):
                cont += 1
    else:
        print("Ingrese otro rango")

    print(f"La cantidad de numeros casi perfectos en el rango es de: {cont}")
    


