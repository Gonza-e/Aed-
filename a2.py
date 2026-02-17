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
        self.ult = None 

    def cargaEncolada(self, num):
        nuevo_nodo = Nodo(num)

        if self.prim is None:
            nuevo_nodo.prox = self.prim 
            self.prim = nuevo_nodo
            self.ult = nuevo_nodo
        else:
            nuevo_nodo.prox = self.ult.prox
            self.ult.prox = nuevo_nodo
            self.ult = nuevo_nodo  

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
        
        if self.prim is None:
            nuevo_maestro.prox = self.prim 
            self.prim = nuevo_maestro
        else:
            p = self.prim 
            ant = None 
            while (p is not None) and esMayor(nuevo_maestro.acceso_lista,p.acceso_lista):
                ant = p 
                p = p.prox 
            
            if p == self.prim:
                nuevo_maestro.prox = self.prim 
                self.prim = nuevo_maestro
            else:
                ant.prox = nuevo_maestro
                nuevo_maestro.prox = p 

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

def esMayor(a,n) -> bool:
    if n is None:
        return False
    else:
        if a is None:
            return True 
        else:
            if a.dato == n.dato:
                return esMayor(a.prox,n.prox)
            else:
                if a.dato < n.dato:
                    return True 
                else:
                    return False 

lista_a = ListaSimple()
for n in "ABCDAS":
    lista_a.cargaEncolada(n)

lista_b = ListaSimple()
for n in "ABCDASA":
    lista_b.cargaEncolada(n)

lista_c = ListaSimple()
for n in "ABCDAS":
    lista_c.cargaEncolada(n)

maestra = ListaDeListas()
maestra.guardar_acceso_lista(lista_a.prim)
maestra.guardar_acceso_lista(lista_b.prim)
maestra.guardar_acceso_lista(lista_c.prim)

lista_a.prim = None
lista_b.prim = None

maestra.mostrar_todo()