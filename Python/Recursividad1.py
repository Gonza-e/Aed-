class Nodo:
    def __init__(self,numero):
        self.num = numero
        self.prox = None 

class Lista:
    def __init__(self):
        self.prim = None

    def cargarNodo(self,numero):
        nuevo = Nodo(numero) 
        nuevo.prox = self.prim
        self.prim = nuevo 

    def mostrarNodo(self):
        p = self.prim
        while p is not None:
            print(f"Número: {p.num}")
            p = p.prox


def esPar(n: int) -> bool:
    if (n % 2) == 0:
        return True 
    else: 
        return False 

lista = Lista()

arreglo = [
    1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
    11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
    21, 22, 23, 24, 25, 26, 27, 28, 29, 30
]

def verificar(a,indice):
    if indice < 30:
        if esPar(a[indice]):
            lista.cargarNodo(a[indice])

        return verificar(a,indice + 1)


verificar(arreglo,0)

lista.mostrarNodo()