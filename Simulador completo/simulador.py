import tkinter as tk
from nodo import Nodo
from puntero import Puntero
import SimuladorListas


class SimuladorLista:
    def __init__(self, root):
        self.root = root
        self.root.title("Lista Simple Enlazada Interactiva")
        self.canvas = tk.Canvas(root, width=1000, height=600, bg="white")
        self.canvas.pack()

        # Atajo con F8 para volver al menú
        self.root.bind("<F8>", self.volver_menu)

        # Botón gráfico para volver atrás
        btn_back = tk.Button(self.root, text="← Menú", command=self.volver_menu)
        btn_back.place(x=5, y=5)

        self.objetos = []
        self.nodo_drag = None
        self.nodo_origen = None
        self.nodo_click = None
        self.last_x, self.last_y = 0, 0

        self.canvas.bind("<Button-1>", self.click_izquierdo)
        self.canvas.bind("<B1-Motion>", self.arrastrar)
        self.canvas.bind("<ButtonRelease-1>", self.soltar)
        self.canvas.bind("<Button-3>", self.click_derecho)

        frame = tk.Frame(root)
        frame.pack()

        self.entry_valor = tk.Entry(frame)
        self.entry_valor.pack(side=tk.LEFT)

        btn_add = tk.Button(frame, text="Agregar nodo", command=self.agregar_nodo)
        btn_add.pack(side=tk.LEFT)

        btn_add_ptr = tk.Button(frame, text="Agregar puntero", command=self.agregar_puntero)
        btn_add_ptr.pack(side=tk.LEFT)



    def agregar_nodo(self):
        valor = self.entry_valor.get()
        if not valor:
            valor = f"N{len(self.objetos)+1}"

        x = 100 + (len(self.objetos) % 4) * 150
        y = 100 + (len(self.objetos) // 4) * 100

        nodo = Nodo(self.canvas, x, y, valor)
        self.objetos.append(nodo)

    def agregar_puntero(self):
        nombre = self.entry_valor.get()
        if not nombre:
            nombre = f"P{len(self.objetos)+1}"

        x = 100 + (len(self.objetos) % 4) * 150
        y = 400 + (len(self.objetos) // 4) * 100

        ptr = Puntero(self.canvas, x, y, nombre)
        self.objetos.append(ptr)

    def click_izquierdo(self, event):
        for obj in self.objetos:
            if obj.contiene(event.x, event.y):
                self.nodo_drag = obj
                self.last_x, self.last_y = event.x, event.y
                return

    def arrastrar(self, event):
        if self.nodo_drag:
            dx = event.x - self.last_x
            dy = event.y - self.last_y
            self.nodo_drag.mover(dx, dy, self.objetos)
            self.last_x, self.last_y = event.x, event.y

    def soltar(self, event):
        self.nodo_drag = None

    def click_derecho(self, event):
        for obj in self.objetos:
            if obj.contiene(event.x, event.y):
                # Si ya hay un origen seleccionado → conectar directo
                if self.nodo_origen:
                    if self.nodo_origen != obj:
                        self.nodo_origen.conectar(obj)
                    self.nodo_origen = None
                    return

                # Si no hay origen, abrimos menú
                self.nodo_click = obj
                menu = tk.Menu(self.root, tearoff=0)

                if isinstance(obj, Puntero):
                    # Menú reducido para punteros
                    menu.add_command(label="Crear conexión", command=self.opcion_conectar)
                    menu.add_command(label="Eliminar conexión", command=self.opcion_eliminar)
                    menu.add_command(label="Eliminar objeto", command=self.opcion_borrar)
                else:
                    # Menú completo para nodos
                    menu.add_command(label="Crear conexión", command=self.opcion_conectar)
                    menu.add_command(label="Eliminar conexión", command=self.opcion_eliminar)
                    menu.add_command(label="Eliminar objeto", command=self.opcion_borrar)

                menu.tk_popup(event.x_root, event.y_root)
                return


    def opcion_conectar(self):
        if self.nodo_origen is None:
            self.nodo_origen = self.nodo_click
        else:
            if self.nodo_origen != self.nodo_click:
                self.nodo_origen.conectar(self.nodo_click)
            self.nodo_origen = None

    def opcion_eliminar(self):
        self.nodo_click.eliminar_flecha()
        if self.nodo_origen == self.nodo_click:
            self.nodo_origen = None

    def opcion_borrar(self):
        obj = self.nodo_click
        obj.eliminar_flecha()
        for otro in self.objetos:
            if otro.siguiente == obj:
                otro.eliminar_flecha()
        for item in getattr(obj, "items", []):
            self.canvas.delete(item)
        if obj in self.objetos:
            self.objetos.remove(obj)
        if self.nodo_origen == obj:
            self.nodo_origen = None

    def volver_menu(self, event=None):
        self.root.destroy()
        root = tk.Tk()
        app = SimuladorListas.MenuPrincipal(root)
        root.mainloop()

