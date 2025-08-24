def conv_caracter(num: int) -> str:
    if num == 1: return "1"
    if num == 0: return "0"

def conv_binario(num: int) -> str:
    if num == 0:
        return "0"
    elif num == 1:
        return "1"
    else:
        return conv_binario(num // 2) + conv_caracter(num % 2)

    
num = int(input("Ingrese un numero: "))

while num != 99:
    print(conv_binario(num))
    num = int(input("Ingrese nuevo numero: "))