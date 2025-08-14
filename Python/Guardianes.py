def Guardianes(A,B,sumA,sumB,resA,resB: int) -> bool:
    if (resA == 0) and (resB == 0):
        if (sumA == B) and (sumB == A): 
            return True
        else: 
            return False
    else: 
        if resA > 0: 
            if (resA % 2 != 0) and (A % resA == 0):
                return Guardianes(A,B,sumA+resA,sumB,resA-1,resB)
            else: 
                return Guardianes(A,B,sumA,sumB,resA-1,resB)
        if resB > 0:
            if (resB % 2 == 0) and (B % resB == 0):
                return Guardianes(A,B,sumA,sumB+resB,resA,resB-1)
            else:
                return Guardianes(A,B,sumA,sumB,resA,resB-1)

numA= int(input("Ingrese un numero A para evaluarlo: "))
numB = int(input("Ingrese un numero B para evaluarlo: "))

if Guardianes(numA,numB,0,0,numA-1,numB-1):
    print('Los numeros son guardianes')
else: 
    print('Los numeros no son guardianes')
