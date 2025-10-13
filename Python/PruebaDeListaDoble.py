class Nodo:
    def __init__(self, letra):
        self.letra = letra
        self.prox = None
        self.ant = None

class ListaDoble:
    def __init__(self):
        self.prim = None
        self.ult = None

    # Función recursiva para verificar capicúa 
    def capicua(self, proximo, anterior):
        if proximo == anterior or (proximo.ant == anterior):  # caso base
            return True
        else:
            if proximo.letra == anterior.letra:
                return self.capicua(proximo.prox, anterior.ant)
            else:
                return False

    # Procedimiento de inserción al final
    def insertar(self, letra):
        nuevo = Nodo(letra)

        if self.prim is None:  # lista vacía
            self.prim = nuevo
            self.ult = nuevo
        else:
            nuevo.ant = self.ult
            self.ult.prox = nuevo
            self.ult = nuevo

    # Método para mostrar lista (opcional, ayuda en prueba) 
    def mostrar(self):
        p = self.prim
        while p:
            print(p.letra, end=" <-> ")
            p = p.prox
        print("Nil")

# --- Programa Principal ---
lista = ListaDoble()
print("Presione ENTER entre cada letra que coloque, presione 0 y luego ENTER para terminar")
letra = ' '

while letra != "0":
    letra = input("Ingrese una letra: ")
    if letra != "0":
        lista.insertar(letra)


print("\nLista cargada:")
lista.mostrar()

# Verificar si es capicúa
es_capicua = lista.capicua(lista.prim, lista.ult)
print("\n¿La lista es capicúa?:", "Sí" if es_capicua else "No")
