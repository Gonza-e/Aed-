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
        
        # Variables para mover la vista del canvas (pan)
        self.pan_start_x = 0
        self.pan_start_y = 0

        self.canvas.bind("<Alt-ButtonPress-1>", self.start_pan)
        self.canvas.bind("<Alt-B1-Motion>", self.do_pan)
        self.canvas.config(scrollregion=(0, 0, 5000, 3000))

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
        canvas_x = self.canvas.canvasx(event.x)
        canvas_y = self.canvas.canvasy(event.y)

        for obj in self.objetos:
            if obj.contiene(canvas_x, canvas_y):
                self.nodo_drag = obj
                self.last_x, self.last_y = canvas_x, canvas_y
                return

    def arrastrar(self, event):
        if self.nodo_drag:
            canvas_x = self.canvas.canvasx(event.x)
            canvas_y = self.canvas.canvasy(event.y)

            dx = canvas_x - self.last_x
            dy = canvas_y - self.last_y

            self.nodo_drag.mover(dx, dy, self.objetos)
            self.last_x, self.last_y = canvas_x, canvas_y

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

    def start_pan(self, event):
        self.pan_start_x = event.x
        self.pan_start_y = event.y

    def do_pan(self, event):
        dx = event.x - self.pan_start_x
        dy = event.y - self.pan_start_y

        self.canvas.xview_scroll(int(-dx/2), "units")
        self.canvas.yview_scroll(int(-dy/2), "units")

        self.pan_start_x = event.x
        self.pan_start_y = event.y

