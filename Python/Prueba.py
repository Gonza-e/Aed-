def posicion_letra(letra: str) -> int:
    """
    Dada una letra, devuelve su posición en el alfabeto.
    A = 1, B = 2, ..., Z = 26
    """
    letra = letra.upper()
    return ord(letra) - ord('A') + 1


# Lista de letras del alfabeto
a = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']


def desc(letra: str, k: int) -> str:
    """
    Desplaza la letra 'k' posiciones hacia atrás en el alfabeto.
    Si llega a 'A', vuelve a 'Z' (cíclico).
    """
    i = posicion_letra(letra) - 1   # índice (0-25)
    i = (i + k) % 26                # desplazamiento circular
    print(i)
    return a[i]


# Programa principal
k = int(input("Ingrese un numero: "))
palabra = input("Ingrese una palabra: ")

palabra_desc = ""

for letra in palabra:
    if letra != " ":
        palabra_desc += desc(letra, k)

print("Palabra original:", palabra)
print("Palabra desplazada:", palabra_desc)
