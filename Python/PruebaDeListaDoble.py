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
        while p is not None:
            print(p.letra, end=" <-> ")
            p = p.prox
        print("Nil")

# --- Programa Principal ---
lista = ListaDoble()
palabra = input("Ingrese una palabra \n")

for letra in palabra:
    lista.insertar(letra)

lista.mostrar()

# Verificar si es capicúa
if lista.capicua(lista.prim,lista.ult):
    print("La palabra es capicua")
else:
    print("La palabra no es capicua")
