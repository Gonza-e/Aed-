class Nodo:
    def __init__(self,numero):
        self.num = numero 
        self.prox = None
    
class ListaSimple:
    def __init__(self):
        self.prim = None
        self.ult = None 
    
    def cargaOrdenada(self,numero):
        nuevo = Nodo(numero)
        
        if self.prim is None:   #si prim = nil entonces 
            nuevo.prox = self.prim 
            self.prim = nuevo 
        else: 
            p = self.prim
            ant = None
            #Si queres comprobar como funciona la carga ordenada ascendente o descendente solo cambia el "<"
            while (p is not None) and (nuevo.num < p.num):  #Mientras p <> nil y *q.num < *p.num hacer
                ant = p 
                p = p.prox 
            
            if p == self.prim:  #Insertar antes del primer elemento (si p = prim entonces)
                nuevo.prox = self.prim 
                self.prim = nuevo 
            else:   #Para los demas casos de insercion (al medio o al final)
                nuevo.prox = ant.prox 
                ant.prox = nuevo 

    def cargaApilada(self,numero):
        nuevo = Nodo(numero)

        nuevo.prox = self.prim  #*q.prox:= prim 
        self.prim = nuevo   #prim:= q
    
    def cargaEncolada(self,numero):
        nuevo = Nodo(numero)

        p = self.prim
        ant = None 
        while p is not None:
            ant = p
            p = p.prox

        if p == self.prim:  #si p = prim entonces
            self.prim = nuevo 
            nuevo.prox = None
        else:   #Carga encolada
            nuevo.prox = ant.prox 
            ant.prox = nuevo 

    def cargaEncolada2(self,numero):
        nuevo = Nodo(numero)
        if self.prim is None:
            nuevo.prox = None
            self.prim = nuevo 
        else:
            nuevo.prox = None
            self.ult.prox = nuevo 

        self.ult = nuevo 

    def mostrarNodo(self):
        p = self.prim 

        while p is not None:    #Mientras p <> nil hacer 
            print(f"{p.num}", end = " -> ")
            p = p.prox
        print("nil")

    def borrarNodo(self,valor):
        if self.prim is None:   #si prim = nil entonces
            print("Error lista vacia")
        else:
            p = self.prim 
            val = valor
            ant = None
            while (p is not None) and (p.num != val):   #Mientras p <> nil y *p.num <> val hacer  
                ant = p 
                p = p.prox 
            
            if p == self.prim:  #si p = prim
                self.prim = self.prim.prox 
            else:   #Para todos los demas casos de eliminacion (ultimo elemento y elemento del medio)
                ant.prox = p.prox 
    
        e = p 
        del e 

def long(A: ListaSimple) -> int: 
    if A.prox is None or A.prox == A:
        return 1 
    else:
        return 1 + long(A.prox)

#De aca en adelante son solamente los bucles para poder probar las cargas y eliminaciones
#Podes reemplazar cargaEncolada por Ordenada o Apilada

lista = ListaSimple()
lista.cargaOrdenada("Gonzalo")
lista.cargaOrdenada("Zahira")
lista.cargaOrdenada("Mario")
lista.cargaOrdenada("Amarilla")
lista.mostrarNodo()

print(f"Longitud de la lista: {long(lista.prim)}")
bandera = True
while bandera:
    numero = int(input("Ingrese un valor para eliminar: \n"))
    
    if numero >= 0: 
        lista.borrarNodo(numero)
    else: 
        bandera = False
    
    if lista.prim == None:
        bandera = False
    lista.mostrarNodo()

