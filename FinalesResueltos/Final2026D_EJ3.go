Accion ej3 es 
	Ambiente
		Funcion experimento(num: entero): entero 
			si num = 3 entonces 
				experimento:= 7 
			sino 
				si num = 2 entonces 
					experimento:= 5+experimento(random())
				sino 
					si num = 1 entonces 
						experimento:= 3+experimento(random())
					Fsi 
				Fsi 
			Fsi 
		FFuncion 
		n: entero 
	Proceso 
		n:= random()

		Esc("Tiempo que tardó la rata en salir:",experimento(n),"minutos")
	FProceso 
FAccion 
