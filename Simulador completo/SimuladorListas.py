import tkinter as tk
import subprocess
import sys
import os

# Rutas de los dos simuladores
SIMULADOR_SIMPLE = os.path.join(os.path.dirname(__file__), r"C:\Users\Gonzalo\Desktop\Algoritmos\Aed-\Simulador completo", "main.py")
SIMULADOR_DOBLE = os.path.join(os.path.dirname(__file__), r"C:\Users\Gonzalo\Desktop\Algoritmos\Aed-\Simulador completo", "main2.py")

class MenuPrincipal:
    def __init__(self, root):
        self.root = root
        self.root.title("Menú Principal - Simuladores de Listas")
        self.root.geometry("400x300")

        tk.Label(root, text="Seleccione un simulador", font=("Times New Roman", 16)).pack(pady=20)

        btn_simple = tk.Button(root, text="Listas Simples", font=("Times New Roman", 16), command=self.abrir_simple)
        btn_simple.pack(pady=10)

        btn_doble = tk.Button(root, text="Listas Dobles", font=("Times New Roman", 16), command=self.abrir_doble)
        btn_doble.pack(pady=10)

        btn_ayuda = tk.Button(root, text="Ayuda", font=("Times New Roman", 12), command=self.mostrar_ayuda)
        btn_ayuda.pack(pady=10)


    def abrir_simple(self):
        self.root.destroy()
        subprocess.run([sys.executable, SIMULADOR_SIMPLE])

    def abrir_doble(self):
        self.root.destroy()
        subprocess.run([sys.executable, SIMULADOR_DOBLE])

    def mostrar_ayuda(self):
        ayuda_win = tk.Toplevel(self.root)
        ayuda_win.title("Ayuda")
        ayuda_win.geometry("500x400")

        texto = (
            "📘 Bienvenido al simulador de listas.\n\n"
            "👉 Puede elegir entre listas simples o dobles.\n"
            "👉 Use click derecho en nodos/punteros para ver opciones.\n"
            "👉 Puede mover nodos con el mouse.\n"
            "👉 Use F8 o el botón ← para volver al menú.\n\n"
            "Este simulador es educativo y le ayudará a comprender cómo\n"
            "funcionan las listas enlazadas de manera interactiva."
        )

        label = tk.Label(ayuda_win, text=texto, justify="left", font=("Arial", 11))
        label.pack(padx=20, pady=20)


if __name__ == "__main__":
    root = tk.Tk()
    app = MenuPrincipal(root)
    root.mainloop()
