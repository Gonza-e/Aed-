class Nodo():
    def __init__(self, numero):
        self.numero = numero 
        self.izq = None
        self.der = None 

class ABB():
    def __init__(self):
        self.prim = None # The root of the tree

    # --- Insertion Method ---
    def cargaArbol(self, numero):
        """Public method to start the insertion process."""
        if self.prim is None:
            self.prim = Nodo(numero)
        else:
            # Calls the private helper method starting from the root
            self._insertar_recursivo(self.prim, numero)

    def _insertar_recursivo(self, actual, numero):
        """
        Private helper method for recursive insertion.
        actual: The current node being examined.
        """
        if numero < actual.numero:
            # Go left
            if actual.izq is None:
                actual.izq = Nodo(numero)
            else:
                self._insertar_recursivo(actual.izq, numero)
        elif numero > actual.numero:
            # Go right
            if actual.der is None:
                actual.der = Nodo(numero)
            else:
                self._insertar_recursivo(actual.der, numero)
        # If numero == actual.numero (duplicate), do nothing in a standard BST

    # --- Traversal Method ---
    def enOrden(self):
        """Public method to start the in-order traversal."""
        print("Recorrido en Orden (In-Order Traversal):")
        self._enOrden_recursivo(self.prim)
        print("-" * 30)

    def _enOrden_recursivo(self, actual):
        """
        Private helper method for recursive in-order traversal.
        Base Case: Stop when the current node is None.
        Order: LEFT -> ROOT -> RIGHT
        """
        if actual is not None:
            self._enOrden_recursivo(actual.izq)  # 1. Traverse Left Subtree
            print(f"{actual.numero}", end=" ")  # 2. Visit Root Node (Print)
            self._enOrden_recursivo(actual.der) # 3. Traverse Right Subtree

# --- Execution ---
arbol = ABB()

arbol.cargaArbol(10)
arbol.cargaArbol(30)
arbol.cargaArbol(50) # Duplicates are generally ignored or handled separately in a BST, here it behaves like a unique value if inserted
arbol.cargaArbol(50)
arbol.cargaArbol(40)
arbol.cargaArbol(70)

# The expected structure of the tree should be:
#         10
#          \
#          30
#           \
#            50
#           /  \
#          40  70

arbol.enOrden()