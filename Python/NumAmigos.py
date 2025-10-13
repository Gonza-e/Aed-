def numAmigos(num1,res1,num2,res2,suma1,suma2: int) -> bool:
    if (res1 == 0) and (res2 == 0):
        if (num1 == suma2) and (num2 == suma1):
            return True
        else:
            return False 
    else:
        if res1 > 0:
            if (num1 % res1) == 0:
                return numAmigos(num1,res1-1,num2,res2,suma1+res1,suma2)
            else:
                return numAmigos(num1,res1-1,num2,res2,suma1,suma2)
        if res2 > 0:
            if (num2 % res2) == 0:
                return numAmigos(num1,res1,num2,res2-1,suma1,suma2+res2)
            else: 
                return numAmigos(num1,res1,num2,res2-1,suma1,suma2)
            
n1 = 220
n2 = 284 

if numAmigos(n1,n1-1,n2,n2-1,0,0):
    print("Los numeros son amigos")
else: 
    print("Los numeros no son amigos")

