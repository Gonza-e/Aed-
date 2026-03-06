def long(n:int) -> int:
    if n == 0:
        return 0
    else:
        return 1+long(n//10)
    
def binario(n,pos: int) -> int:
    if n % 10 == 1: 
        return long(n)-1
    else:
        return binario(n//10,pos)
    
def conv(x:int) -> int:
    if x < 2:
        return x 
    else:
        return conv(x//2)*10 + (x%2)
    
num = 60

n = conv(num)
print(n)
print(binario(n,0))
    
def sumar(dni:int) -> int:
    if dni == 0:
        return 1
    else:
        return sumar(dni//10) + dni%10
    
#31 

dni = 45817140

print(sumar(dni))
