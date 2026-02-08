class Nodo():
    def __init__(self, numero):
        self.numero = numero 
        self.izq = None
        self.der = None 

class ABB():
    def __init__(self):
        self.prim = None 

    def cargaArbol(self, numero):
        if self.prim is None:
            self.prim = Nodo(numero)
        else:
            self.insertarRecursivo(self.prim, numero)

    def insertarRecursivo(self, p, numero):
        if numero < p.numero:
            if p.izq is None:
                p.izq = Nodo(numero)
            else:
                self.insertarRecursivo(p.izq, numero)
        elif numero > p.numero:
            if p.der is None:
                p.der = Nodo(numero)
            else:
                self.insertarRecursivo(p.der, numero)

def enOrden(nodo: ABB):
    if nodo is not None:
        enOrden(nodo.izq)
        print(f"{nodo.numero}",end=" -> ")
        enOrden(nodo.der)

def postOrden(nodo: ABB):
    if nodo is not None:
        postOrden(nodo.izq)
        postOrden(nodo.der)
        print(f"{nodo.numero}",end=" -> ")

def preOrden(nodo: ABB):
    if nodo is not None:
        print(f"{nodo.numero}",end=" -> ")
        preOrden(nodo.izq)
        preOrden(nodo.der)

def busqueda(prim: ABB, buscar:int) -> int:
    if prim is None:
        return -1

    if prim.numero == buscar:
        return prim.numero 

    if buscar < prim.numero:
        return busqueda(prim.izq,buscar)
    else:
        return busqueda(prim.der,buscar)

def altura(nodo: ABB) -> int:
    if nodo is None:
        return 0 
    else:
        if altura(nodo.izq) > altura(nodo.der):
            return 1 + altura(nodo.izq)
        else:
            return 1 + altura(nodo.der)

def camino(nodo: ABB,valor: int):    
    if nodo is not None:
        print(f"{nodo.numero}",end=" -> ")

        if valor > nodo.numero:
            camino(nodo.der,valor)
        else:
            if valor < nodo.numero:
                camino(nodo.izq,valor)

def nodos(nodo: ABB) -> int:
    if nodo is None:
        return 0 
    else:
        if nodo is not None:
            return 1 + nodos(nodo.izq) + nodos(nodo.der)
        


arbol = ABB()

arbol.cargaArbol(50)
arbol.cargaArbol(30)
arbol.cargaArbol(10)
arbol.cargaArbol(60)
arbol.cargaArbol(40)
arbol.cargaArbol(70)
arbol.cargaArbol(80)

print(f"La altura del arbol es: {altura(arbol.prim)}")
print("Recorrido en orden:")
enOrden(arbol.prim)
print("\nRecorrido post orden:")
postOrden(arbol.prim)
print("\nRecorrido pre orden:")
preOrden(arbol.prim)
print(f"\nValor: {busqueda(arbol.prim,6)}")

print(f"La cantidad de nodos del arbol es: {nodos(arbol.prim)}")

