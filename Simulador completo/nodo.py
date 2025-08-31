import tkinter as tk

def borde_por_rayo(x1, y1, xc, yc, x_min, y_min, x_max, y_max):
    """Devuelve el punto donde el rayo (x1,y1)->(xc,yc) intersecta el borde del rectángulo destino."""
    dx = xc - x1
    dy = yc - y1
    candidatos = []

    if dx != 0:
        # Lado izquierdo
        t = (x_min - x1) / dx
        y = y1 + t * dy
        if t > 0 and y_min <= y <= y_max:
            candidatos.append((t, x_min, y))
        # Lado derecho
        t = (x_max - x1) / dx
        y = y1 + t * dy
        if t > 0 and y_min <= y <= y_max:
            candidatos.append((t, x_max, y))

    if dy != 0:
        # Lado superior
        t = (y_min - y1) / dy
        x = x1 + t * dx
        if t > 0 and x_min <= x <= x_max:
            candidatos.append((t, x, y_min))
        # Lado inferior
        t = (y_max - y1) / dy
        x = x1 + t * dx
        if t > 0 and x_min <= x <= x_max:
            candidatos.append((t, x, y_max))

    if not candidatos:
        # Fallback: clamp al borde más cercano
        x = min(max(xc, x_min), x_max)
        y = min(max(yc, y_min), y_max)
        return x, y

    _, x2, y2 = min(candidatos, key=lambda k: k[0])
    return x2, y2


class Nodo:
    def __init__(self, canvas, x, y, valor):
        self.canvas = canvas
        self.x = x
        self.y = y
        self.valor = valor
        self.siguiente = None
        self.flecha = None

        # Dibujar rectángulos y texto
        self.rect_dato = canvas.create_rectangle(x, y, x+60, y+30, fill="lightblue")
        self.rect_ptr = canvas.create_rectangle(x+60, y, x+100, y+30, fill="lightgreen")
        self.texto = canvas.create_text(x+30, y+15, text=str(valor))

        self.items = [self.rect_dato, self.rect_ptr, self.texto]

    def contiene(self, x, y):
        return self.x <= x <= self.x+100 and self.y <= y <= self.y+30

    def mover(self, dx, dy, objetos):
        self.x += dx
        self.y += dy
        for item in self.items:
            self.canvas.move(item, dx, dy)

        if self.flecha and self.siguiente:
            self.canvas.delete(self.flecha)
            self.conectar(self.siguiente)

        for obj in objetos:
            if obj.siguiente == self:
                if obj.flecha:
                    self.canvas.delete(obj.flecha)
                obj.conectar(self)

    def conectar(self, otro):
        if self.flecha:
            self.canvas.delete(self.flecha)
        self.siguiente = otro

        # Origen: centro del lado derecho del nodo origen
        x1, y1 = self.x + 80, self.y + 15  #valor orignal self.x + 100

        # Rect destino (ancho 100 si es Nodo, 80 si es Puntero)
        ancho_dest = 100 if hasattr(otro, "rect_dato") else 80
        x_min, y_min = otro.x, otro.y
        x_max, y_max = otro.x + ancho_dest, otro.y + 30

        # Centro geométrico del destino
        xc, yc = otro.x + ancho_dest / 2, otro.y + 15

        # Punto de entrada correcto en el borde
        x2, y2 = borde_por_rayo(x1, y1, xc, yc, x_min, y_min, x_max, y_max)

        # Curva suave
        xm = (x1 + x2) / 2
        ym = (y1 + y2) / 2 - 40

        self.flecha = self.canvas.create_line(
            x1, y1, xm, ym, x2, y2,
            smooth=True, arrow=tk.LAST, width=2
        )


    def eliminar_flecha(self):
        if self.flecha:
            self.canvas.delete(self.flecha)
            self.flecha = None
            self.siguiente = None
