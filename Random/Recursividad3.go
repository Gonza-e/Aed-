"Desarrollar una función recursiva que realice la búsqueda binaria en un arreglo
ordenado. La búsqueda binaria consiste en dividir repetidamente el arreglo en dos
mitades y comparar el valor buscado con el valor en el medio, reduciendo así el
espacio de búsqueda a la mitad en cada paso. Tome como ejemplo el siguiente
vector de 7 posiciones
2 5 6 14 45 67 70"

Accion binaria es 
 Ambiente 
	Funcion busquedaBinaria(izq,der,medio,num: entero; a: arreglo[1...7] de enteros): entero 
		si izq = der entonces 
			busquedaBinaria:= 0 
		sino 
			si num < a[medio] entonces 
				busquedaBinaria:= busquedaBinaria(izq,medio,(izq + der) div 2,num,a)
			sino 
				si num > a[medio] entonces 
					busquedaBinaria:= busquedaBinaria(medio,der,(izq + der) div 2,num,a)
				sino 
					si num = a[medio] entonces 
						busquedaBinaria:= a[medio]
					Fsi
				Fsi 
			Fsi 
		Fsi 
	FFuncion