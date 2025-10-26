import tkinter as tk
from nodo2 import Nodo2
from puntero2 import Puntero2
import SimuladorListas

class SimuladorListaDoble:
    def __init__(self, root):
        self.root = root
        self.root.title("Lista Doble Enlazada Interactiva")
        self.canvas = tk.Canvas(root, width=1000, height=600, bg="white")
        self.canvas.pack()

        
        # Atajo con F8 para volver al menú
        self.root.bind("<F8>", self.volver_menu)

        # Botón gráfico para volver atrás
        btn_back = tk.Button(self.root, text="← Menú", command=self.volver_menu)
        btn_back.place(x=5, y=5)

        self.objetos = []
        self.objeto_drag = None
        self.objeto_click = None
        self.objeto_origen = None
        self.tipo_conexion = None
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
        x = 100 + (len(self.objetos)%4)*150
        y = 100 + (len(self.objetos)//4)*100
        nodo = Nodo2(self.canvas, x, y, valor)
        self.objetos.append(nodo)

    def agregar_puntero(self):
        nombre = self.entry_valor.get()
        if not nombre:
            nombre = f"P{len(self.objetos)+1}"
        x = 100 + (len(self.objetos)%4)*150
        y = 400 + (len(self.objetos)//4)*100
        ptr = Puntero2(self.canvas, x, y, nombre)
        self.objetos.append(ptr)

    def click_izquierdo(self, event):
        for obj in self.objetos:
            if obj.contiene(event.x, event.y):
                self.objeto_drag = obj
                self.last_x, self.last_y = event.x, event.y
                return

    def arrastrar(self, event):
        if self.objeto_drag:
            dx = event.x - self.last_x
            dy = event.y - self.last_y
            self.objeto_drag.mover(dx, dy, self.objetos)
            self.last_x, self.last_y = event.x, event.y

    def soltar(self, event):
        self.objeto_drag = None

    def click_derecho(self, event):
        for obj in self.objetos:
            if obj.contiene(event.x, event.y):
                # Si ya hay un origen seleccionado, conectar directamente
                if self.objeto_origen and self.tipo_conexion:
                    if self.objeto_origen != obj:
                        self.objeto_origen.redibujar_flecha(obj, tipo=self.tipo_conexion)
                    # reset
                    self.objeto_origen = None
                    self.tipo_conexion = None
                    return

                # Si NO hay origen, abrimos el menú normalmente
                self.objeto_click = obj
                menu = tk.Menu(self.root, tearoff=0)

                if isinstance(obj, Puntero2):
                    # Menú reducido para punteros
                    menu.add_command(label="Crear conexión", command=lambda: self.opcion_conectar("siguiente"))
                    menu.add_command(label="Eliminar conexión", command=lambda: self.opcion_eliminar(None))
                    menu.add_command(label="Eliminar objeto", command=self.opcion_borrar)
                else:
                    # Menú completo para nodos
                    menu.add_command(label="Conectar como siguiente", command=lambda: self.opcion_conectar("siguiente"))
                    menu.add_command(label="Conectar como anterior", command=lambda: self.opcion_conectar("anterior"))
                    menu.add_command(label="Eliminar conexión siguiente", command=lambda: self.opcion_eliminar("siguiente"))
                    menu.add_command(label="Eliminar conexión anterior", command=lambda: self.opcion_eliminar("anterior"))
                    menu.add_command(label="Eliminar objeto", command=self.opcion_borrar)

                menu.tk_popup(event.x_root, event.y_root)
                return


    def opcion_conectar(self, tipo):
        if self.objeto_origen is None:
            self.objeto_origen = self.objeto_click
            self.tipo_conexion = tipo
        else:
            if self.objeto_origen != self.objeto_click:
                self.objeto_origen.redibujar_flecha(self.objeto_click, tipo=self.tipo_conexion)
            self.objeto_origen = None
            self.tipo_conexion = None

    def opcion_conectar_puntero(self, tipo):
        if self.objeto_origen is None:
            self.objeto_origen = self.objeto_click
            self.tipo_conexion = tipo
        else:
            if self.objeto_origen != self.objeto_click:
                self.objeto_origen.redibujar_flecha(self.objeto_click, tipo=self.tipo_conexion)
            self.objeto_origen = None
            self.tipo_conexion = None

    def opcion_eliminar(self, tipo):
        self.objeto_click.eliminar_flecha(tipo=tipo)
        if self.objeto_origen == self.objeto_click:
            self.objeto_origen = None
            self.tipo_conexion = None

    def opcion_borrar(self):
        obj = self.objeto_click
        obj.eliminar_flecha(tipo=None)
        for otro in self.objetos:
            if otro.siguiente == obj:
                otro.eliminar_flecha("siguiente")
            if otro.anterior == obj:
                otro.eliminar_flecha("anterior")
        for item in getattr(obj, "items", []):
            self.canvas.delete(item)
        if obj in self.objetos:
            self.objetos.remove(obj)
        if self.objeto_origen == obj:
            self.objeto_origen = None
            self.tipo_conexion = None

    def volver_menu(self, event=None):
        self.root.destroy()
        root = tk.Tk()
        app = SimuladorListas.MenuPrincipal(root)
        root.mainloop()
