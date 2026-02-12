class Nodo:
    def __init__(self, numero):
        self.dato = numero 
        self.prox = None 

class NodoMaestro:
    def __init__(self, puntero_a_lista):
        self.acceso_lista = puntero_a_lista 
        self.cantidad_elementos = 0
        self.prox = None 
    
class ListaSimple:
    def __init__(self):
        self.prim = None 

    def cargaApilada(self, num):
        nuevo_nodo = Nodo(num)
        nuevo_nodo.prox = self.prim 
        self.prim = nuevo_nodo

def calcular_longitud(nodo: Nodo) -> int:
    if nodo is None:
        return 0
    return 1 + calcular_longitud(nodo.prox)
    
class ListaDeListas: 
    def __init__(self):
        self.prim = None

    def guardar_acceso_lista(self, prim_nodo: Nodo):
        nuevo_maestro = NodoMaestro(prim_nodo)
        nuevo_maestro.cantidad_elementos = calcular_longitud(prim_nodo)
        
        nuevo_maestro.prox = self.prim 
        self.prim = nuevo_maestro

    def mostrar_todo(self):
        p_maestro = self.prim 
        while p_maestro is not None:
            print(f"\n--- Sublista detectada (Longitud: {p_maestro.cantidad_elementos}) ---")
            
            p_datos = p_maestro.acceso_lista 
            while p_datos is not None:
                print(f"  [Dato: {p_datos.dato}]", end=" -> ")
                p_datos = p_datos.prox
            
            print("None")
            p_maestro = p_maestro.prox 

lista_a = ListaSimple()
for n in [10, 20, 30]:
    lista_a.cargaApilada(n)

lista_b = ListaSimple()
for n in [5, 4, 3, 2, 1]:
    lista_b.cargaApilada(n)

maestra = ListaDeListas()
maestra.guardar_acceso_lista(lista_a.prim)
maestra.guardar_acceso_lista(lista_b.prim)

lista_a.prim = None
lista_b.prim = None

maestra.mostrar_todo()